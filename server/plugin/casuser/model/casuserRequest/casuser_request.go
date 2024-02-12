package casuserRequest

type RegisterCasUser struct {
	BaseCasUser
	RePassword string `json:"rePassword"`
}

type LoginCasUser struct {
	BaseCasUser
}

type BaseCasUser struct {
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}
