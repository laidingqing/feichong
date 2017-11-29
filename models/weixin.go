package models

// Weixin Login会话请求体
type Weixin struct {
	Code          string   `json:"code"`
	EncryptedData string   `json:"encryptedData"`
	Iv            string   `json:"iv"`
	UserInfo      UserInfo `json:"userInfo"`
}

// UserInfo ..
type UserInfo struct {
	AvatarURL string `json:"avatarUrl"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Gender    int    `json:"gender"`
	NickName  string `json:"nickName"`
	Province  string `json:"province"`
	UserID    string `json:"userId"`
}
