package casuserRequest

type RegisterCasUser struct {
	Username   string `json:"username"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	RePassword string `json:"rePassword"`
	Captcha    string `json:"captcha"`
	CaptchaId  string `json:"captchaId"`
}
