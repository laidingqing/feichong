package models

// Weixin Login会话请求体
type Weixin struct {
	Code          string `json:"code"`
	EncryptedData string `json:"encryptedData"`
	Iv            string `json:"iv"`
}
