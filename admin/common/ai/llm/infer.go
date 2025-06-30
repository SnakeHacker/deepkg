package llm

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/golang/glog"
)

type Client struct {
	Config     Config
	HTTPClient *resty.Client
}

func (dc *Client) Infer(query string, history History) (result string, err error) {

	config := dc.Config

	messages := make([]Message, 0)

	// 添加历史消息
	for _, msg := range history.Messages {
		messages = append(messages, Message{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	// 添加当前查询
	if query != "" {
		// glog.Infof("开始推理，输入query: %s", query)
		messages = append(messages, Message{
			Role:    "user",
			Content: query,
		})
	}

	chatReq := ChatCompletionRequest{
		Model:       config.Model,
		Messages:    messages,
		Stream:      false,
		MaxTokens:   1024 * 16,
		Temperature: 0.1,
		// EnableThinking: false,
	}

	payload, err := json.Marshal(chatReq)
	if err != nil {
		glog.Error(err)
		return
	}

	resp, err := dc.HTTPClient.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", dc.Config.ApiSecret)).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(dc.Config.BaseUrl)
	if err != nil {
		glog.Error(err)
		return
	}

	chatCompletion := ChatCompletion{}
	err = json.Unmarshal(resp.Body(), &chatCompletion)
	if err != nil {
		glog.Infof("%v", string(resp.Body()))
		glog.Error(err)
		return
	}

	// 截断think标签
	result = chatCompletion.Choices[0].Message.Content
	re := regexp.MustCompile(`(?s)<think>.*?</think>`)
	result = re.ReplaceAllString(result, "")

	return
}

func (dc *Client) InferStream(ctx context.Context, query string, history History, ch chan<- string) error {

	reasoning_flag := false
	reasoning_content_start := true
	answer_content_start := true
	config := dc.Config

	if strings.Contains(strings.ToLower(config.BaseUrl), "volces") {
		reasoning_flag = true
	}

	glog.Infof("使用模型配置: model=%s, baseUrl=%s", config.Model, config.BaseUrl)
	glog.Infof("开始流式推理，输入query: %s", query)

	messages := make([]Message, 0)

	// 添加历史消息
	for _, msg := range history.Messages {
		messages = append(messages, Message{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	// 添加当前查询
	messages = append(messages, Message{
		Role:    "user",
		Content: query,
	})

	chatReq := ChatCompletionRequest{
		Model:     config.Model,
		Messages:  messages,
		Stream:    true,
		MaxTokens: 1024 * 8,
	}

	payload, err := json.Marshal(chatReq)
	if err != nil {
		glog.Error(err)
		return err
	}
	// glog.Infof("请求参数: %s", string(payload))

	method := "POST"
	client := &http.Client{
		Timeout: 3 * time.Minute,
	}

	// 创建带有上下文的请求
	req, err := http.NewRequestWithContext(ctx, method, config.BaseUrl, strings.NewReader(string(payload)))
	if err != nil {
		glog.Error(err)
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.ApiSecret))
	req.Header.Add("Content-Type", "application/json")
	glog.Info("已设置请求头部")

	resp, err := client.Do(req)
	if err != nil {
		glog.Error(err)
		return err
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	glog.Info("开始读取流式响应")

	for {
		// 检查上下文是否已取消
		select {
		case <-ctx.Done():
			glog.Info("检测到客户端断开，停止流式推理")
			return ctx.Err()
		default:
		}

		line, err := reader.ReadBytes('\n')
		if !reasoning_flag {
			if err != nil {
				if err == io.EOF {
					glog.Info("流式响应结束")
					break
				}
				// 检查是否是因为连接被关闭导致的错误
				if err == io.ErrUnexpectedEOF || strings.Contains(err.Error(), "connection reset by peer") {
					glog.Info("检测到连接断开")
					return nil
				}
				glog.Error(err)
				return err
			}
		}

		// 跳过空行
		if len(line) <= 1 {
			continue
		}

		// 移除 "data: " 前缀
		data := bytes.TrimPrefix(line, []byte("data: "))
		trimmedData := strings.TrimSpace(string(data))

		// 检查是否是 [DONE] 标记
		if trimmedData == "[DONE]" {
			glog.Info("收到[DONE]标记")
			ch <- "[DONE]"
			break
		}

		// 跳过空数据
		if len(trimmedData) == 0 {
			continue
		}

		var streamResp ChatCompletionStream
		err = json.Unmarshal([]byte(trimmedData), &streamResp)
		if err != nil {
			glog.Errorf("解析响应数据失败: %v, 原始数据: %s", err, trimmedData)
			continue
		}

		if len(streamResp.Choices) > 0 {
			if reasoning_flag {
				delta := streamResp.Choices[0].Delta
				if delta.ReasoningContent != "" {
					if reasoning_content_start {
						ch <- "<think>"
						reasoning_content_start = false
					}
					ch <- delta.ReasoningContent
				} else {
					if answer_content_start {
						ch <- "</think>"
						answer_content_start = false
					}
					ch <- delta.Content
				}
			} else {
				content := streamResp.Choices[0].Delta.Content
				if content != "" {
					ch <- content
				}
			}
		}
	}

	glog.Info("流式推理完成")
	return nil
}
