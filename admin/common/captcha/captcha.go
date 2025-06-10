package captcha

import (
	"errors"
	"image/color"

	"github.com/golang/glog"
	"github.com/mojocn/base64Captcha"
	cap "github.com/mojocn/base64Captcha"
)

var fonts = []string{"3Dumb.ttf", "Comismsh.ttf", "DENNEthree-dee.ttf", "DeborahFancyDress.ttf",
	"Flim-Flam.ttf", "RitaSmith.ttf", "actionj.ttf", "chromohv.ttf", "wqy-microhei.ttc"}

const (
	CAP_HEIGHT      = 80
	CAP_WIDTH       = 240
	CAP_NOISE_COUNT = 0
	CAP_LINE_NUMBER = 4
	CAP_LENGTH      = 5
	CAP_SOURCE      = "1234567890qwertyuioplkjhgfdsazxcvbnm"
	CAP_COLOUR_R    = 3
	CAP_COLOUR_G    = 102
	CAP_COLOUR_B    = 214
	CAP_COLOUR_A    = 125
	CAP_FONT        = "wqy-microhei.ttc"
)

func SetUpCaptcha(driverString *cap.DriverString) (captcha *cap.Captcha, err error) {
	if driverString == nil {
		driverString = &cap.DriverString{
			Height:          CAP_HEIGHT,
			Width:           CAP_WIDTH,
			NoiseCount:      CAP_NOISE_COUNT,
			ShowLineOptions: CAP_LINE_NUMBER,
			Length:          CAP_LENGTH,
			Source:          CAP_SOURCE,
			BgColor: &color.RGBA{
				R: CAP_COLOUR_R,
				G: CAP_COLOUR_G,
				B: CAP_COLOUR_B,
				A: CAP_COLOUR_A,
			},
			Fonts: []string{CAP_FONT},
		}
	}

	err = ValidateCaptchaDriver(driverString)
	if err != nil {
		glog.Error(err)
		return
	}

	captcha = cap.NewCaptcha(driverString.ConvertFonts(), cap.DefaultMemStore)

	return
}

func ValidateCaptchaDriver(driverString *cap.DriverString) (err error) {
	fontsMap := make(map[string]struct{})

	for _, font := range fonts {
		fontsMap[font] = struct{}{}
	}

	for _, font := range driverString.Fonts {
		if _, ok := fontsMap[font]; !ok {
			err := ErrInvalidCaptchaFont
			glog.Error(err)
			return err
		}
	}

	if driverString.Height == 0 {
		err := ErrInvalidCaptchaHeight
		glog.Error(err)
		return err
	}

	if driverString.Width == 0 {
		err := ErrInvalidCaptchaWidth
		glog.Error(err)
		return err
	}

	if driverString.Length == 0 {
		err := ErrInvalidCaptchaLength
		glog.Error(err)
		return err
	}

	if driverString.Height > driverString.Width {
		err := ErrInvalidCaptcha
		glog.Error(err)
		return err
	}

	return
}

func GenerateCaptcha(captcha *base64Captcha.Captcha) (captchaID string, captchaImage string, err error) {
	if captcha == nil {
		err = errors.New("验证码为空")
		glog.Error(err)
		return
	}

	captchaID, captchaImage, _, err = captcha.Generate()
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func VerifyCaptcha(captcha *base64Captcha.Captcha, captchaID string, captchaValue string) (successed bool, err error) {
	if captcha == nil {
		err = errors.New("验证码为空")
		glog.Error(err)
		return
	}

	if captchaID == "" || captchaValue == "" {
		err = errors.New("验证码为空")
		glog.Error(err)
		return
	}

	successed = captcha.Store.Verify(captchaID, captchaValue, true)

	if !successed {
		err = errors.New("验证码错误，请重试")
		glog.Error(err)
		return
	}

	return
}
