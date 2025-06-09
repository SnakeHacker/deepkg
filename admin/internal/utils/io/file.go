package io

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/golang/glog"
)

// DownloadFile 根据文件地址下载文件并返回 byte 数组
func DownloadFile(fileURL string) ([]byte, error) {
	// 发送 HTTP GET 请求
	resp, err := http.Get(fileURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download file, status code: %d", resp.StatusCode)
	}

	// 读取响应体并转换为 byte 数组
	fileBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

// ParseFileName 解析文件下载地址，获取文件名和文件类型
func ParseFileName(fileURL string) (fileName string, contentType string, err error) {
	// 解析文件地址
	u, err := url.Parse(fileURL)
	if err != nil {
		return "", "", err
	}

	// 获取文件名
	fileName = path.Base(u.Path)

	// 发送 HEAD 请求获取文件类型
	resp, err := http.Head(fileURL)
	if err != nil {
		return fileName, "", err
	}
	defer resp.Body.Close()

	// 获取 Content-Type
	contentType = resp.Header.Get("Content-Type")

	// 如果 Content-Type 为空，尝试根据文件扩展名猜测
	if contentType == "" {
		contentType = getContentTypeByExtension(fileName)
	}

	return fileName, contentType, nil
}

// getContentTypeByExtension 根据文件扩展名判断 Content-Type
func getContentTypeByExtension(fileName string) string {
	ext := strings.ToLower(path.Ext(fileName))
	switch ext {
	case ".txt":
		return "text/plain"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".pdf":
		return "application/pdf"
	case ".zip":
		return "application/zip"
	case ".json":
		return "application/json"
	case ".xml":
		return "application/xml"
	case ".html", ".htm":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".doc":
		return "application/msword"
	case ".docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case ".xls":
		return "application/vnd.ms-excel"
	case ".xlsx":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	case ".ppt":
		return "application/vnd.ms-powerpoint"
	case ".pptx":
		return "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	case ".key":
		return "application/vnd.apple.keynote"
	default:
		return "application/octet-stream"
	}
}

func CreateDirIfNotExist(dirPath string) (err error) {
	if DoesDirExist(dirPath) {
		return
	}

	err = os.MkdirAll(dirPath, os.ModePerm)
	return
}

func DoesDirExist(dirPath string) bool {
	stat, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}
	if !stat.IsDir() {
		return false
	}
	return true
}

func GetFileContent(client *resty.Client, fileServer string, fileID string) (content string, err error) {
	fileReq := GetFileContentReq{
		FileID: fileID,
	}

	time.Sleep(1 * time.Second)
	payload := []byte{}
	payload, err = json.Marshal(fileReq)
	if err != nil {
		glog.Error(err)
		return
	}

	// 发送POST请求
	result := &resty.Response{}
	result, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(fmt.Sprintf("%s/api/file/content", fileServer))
	if err != nil {
		glog.Error(err)
		return
	}

	fileResp := GetFileContentResp{}

	err = json.Unmarshal(result.Body(), &fileResp)
	if err != nil {
		glog.Error(err)
		return
	}

	// 若fileResp.success为false，最多重试3次，成功后立即退出
	for i := 0; i < 100 && !fileResp.Success; i++ {
		glog.Infof("获取文件内容，重试第%d次", i+1)
		time.Sleep(3 * time.Second)
		result, err = client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(payload).
			Post(fmt.Sprintf("%s/api/file/content", fileServer))
		if err != nil {
			glog.Error(err)
			continue
		}

		err = json.Unmarshal(result.Body(), &fileResp)
		if err != nil {
			glog.Error(err)
			continue
		}

		if fileResp.Success {
			break // 成功获取内容后立即退出循环
		}
	}

	content = fileResp.Data.Content.Content
	return
}

func UploadFileExtract(client *resty.Client, fileServer string, header *multipart.FileHeader, file multipart.File, enableMultimodal string) (fileID string, err error) {
	result := &resty.Response{}
	result, err = client.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetFileReader("file", header.Filename, file).
		SetFormData(map[string]string{
			"enable_multimodal": enableMultimodal,
		}).
		Post(fmt.Sprintf("%s/api/file/upload", fileServer))
	if err != nil {
		glog.Error(err)
		return
	}

	fileResp := UploadFileExtractResp{}

	err = json.Unmarshal(result.Body(), &fileResp)
	if err != nil {
		glog.Error(err)
		return
	}

	fileID = fileResp.Data.FileID
	return
}
