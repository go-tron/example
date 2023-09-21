package model

type LoginReq struct {
	ChannelId string `json:"channelId"`
	Phone     string `json:"phone" validate:"required"`
	Code      string `json:"code" validate:"required,len=6"`
}

type LoginRes struct {
	Token  string `json:"token,omitempty"`
	UserId int64  `json:"-"`
	Phone  string `json:"phone"`
}
