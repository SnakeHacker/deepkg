package captcha

import (
	"fmt"
)

type errScope string

const (
	errUnknown errScope = ""
	errCaptcha          = "captcha_error"
)

var (

	// Captcha error
	ErrInvalidCaptchaFont   = makeError(errCaptcha, "invalid captcha font")
	ErrInvalidCaptchaHeight = makeError(errCaptcha, "invalid captcha height")
	ErrInvalidCaptchaWidth  = makeError(errCaptcha, "invalid captcha width")
	ErrInvalidCaptchaLength = makeError(errCaptcha, "invalid captcha length")
	ErrInvalidCaptcha       = makeError(errCaptcha, "captcha height can not larger than width")
)

func makeError(scope errScope, msg ...string) error {
	return fmt.Errorf("[%s]: %s", scope, msg)
}
