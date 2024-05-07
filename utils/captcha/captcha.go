package captcha

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	Key    string
	Base64 string
	Answer string
	Width  int
	Height int
}

var store = base64Captcha.DefaultMemStore

func NewCaptcha() *Captcha {
	return &Captcha{
		Width:  144,
		Height: 48,
	}
}

// 创建验证码
func (cha *Captcha) Create() error {
	//config struct for digits
	//数字验证码配置
	driver := base64Captcha.DriverMath{
		Width:           cha.Width,
		Height:          cha.Height,
		NoiseCount:      1,
		ShowLineOptions: base64Captcha.OptionShowHollowLine,
		BgColor:         &color.RGBA{R: 255, G: 255, B: 255, A: 255},
	}

	clint := base64Captcha.NewCaptcha(driver.ConvertFonts(), store)
	id, b64, answer, err := clint.Generate()
	if err != nil {
		return err
	}
	cha.Key = id
	cha.Base64 = b64
	cha.Answer = answer
	return nil
}

// 验证验证码
func (cha *Captcha) Verify(key, answer string) bool {
	return store.Verify(key, answer, true)
}
