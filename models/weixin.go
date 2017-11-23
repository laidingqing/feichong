package models

// Weixin Login会话请求体
type Weixin struct {
	Code          string `json:"code"`
	EncryptedData string `json:"encryptedData"`
	Iv            string `json:"iv"`
	UserInfo      UserInfo `json:"userInfo"`
}

type UserInfo struct {
	AvatarUrl string `json:"avatarUrl"`
	City string `json:"city"`
	Country string `json:"country"`
	Gender string `json:"gender"`
	NickName string `json:"nickName"`
	Province string `json:"province"`
}
