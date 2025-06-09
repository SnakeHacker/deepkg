package errorx

const SUCCESS_CODE = 17000
const DEFAULT_CODE = 1001
const BADREQUEST_CODE = 14000
const UNAUTHORIZED_CODE = 1401

type CodeError struct {
	Code    int    `json:"returnCode"`
	Msg     string `json:"returnDesc"`
	Success bool   `json:"success"`
}

type CodeErrorResponse struct {
	Code    int    `json:"returnCode"`
	Msg     string `json:"returnDesc"`
	Success bool   `json:"success"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(DEFAULT_CODE, msg)
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
