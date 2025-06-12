package llm

import (
	"encoding/json"
	"fmt"
	"regexp"

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
