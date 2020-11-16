package model

import "github.com/mojocn/base64Captcha"

type ConfigJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

type Captcha struct {
	Id   string `json:"id"`
	B64s string `json:"b64s"`
}
