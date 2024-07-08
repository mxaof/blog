package login

type Captcha struct {
	KeyLong            int `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`                                    // 验证码长度
	ImgWidth           int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`                                 // 验证码宽度
	ImgHeight          int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"`                              // 验证码高度
	OpenCaptcha        int `mapstructure:"open-captcha" json:"openCaptcha" yaml:"open-captcha"`                        // 防爆破验证码开启次数，0代表每次登录都需要验证码，其他数字代表错误密码次数，如输错三次后登录时需要填写验证码
	OpenCaptchaTimeOut int `mapstructure:"open-captcha-timeout" json:"openCaptchaTimeOut" yaml:"open-captcha-timeout"` // 防爆破验证码超时时间，单位：s(秒)
}
type SysCaptchaResponse struct {
	CaptchaId     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
	OpenCaptcha   bool   `json:"openCaptcha"`
}
type VerifyRequest struct {
	CaptchaId     string `json:"captchaId"`
	CaptchaAnswer string `json:"captchaAnswer"`
}
