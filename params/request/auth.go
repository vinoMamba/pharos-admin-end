package request

type DingtalkLoginRequest struct {
	AuthCode string `json:"authCode" `
}

type PwdLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
