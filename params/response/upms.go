package response

type UserInfoResponse struct {
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	RealName string `json:"realName"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}
