package io_util

type GetFileContentReq struct {
	FileID string `json:"file_id"`
}

type GetFileContentResp struct {
	ReturnCode int    `json:"returnCode"`
	ReturnDesc string `json:"returnDesc"`
	Success    bool   `json:"success"`
	Data       struct {
		Content Content `json:"content"`
	} `json:"data"`
}

type Content struct {
	Content string  `json:"content"`
	Images  []Image `json:"images"`
}

type Image struct {
	ID          string `json:"id"`
	Data        string `json:"data"`
	Description string `json:"description"`
}

type UploadFileExtractResp struct {
	ReturnCode int    `json:"returnCode"`
	ReturnDesc string `json:"returnDesc"`
	Success    bool   `json:"success"`
	Data       struct {
		FileID string `json:"file_id"`
	} `json:"data"`
}
