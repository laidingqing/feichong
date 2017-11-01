package models

import "time"

// User 用户信息
type User struct {
	ID          string    `bson:"_id" json:"id"`
	Nick        string    `bson:"nick" json:"nick"`
	Email       string    `bson:"email" json:"email"`
	Name        string    `bson:"name" json:"name"`
	Phone       string    `bson:"phone" json:"phone"`
	CompanyName bool      `bson:"companyName" json:"companyName"`
	Admin       bool      `bson:"admin" json:"admin"`
	OpenID      string    `bson:"openID" json:"openID"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
}

// Order 订单信息
type Order struct {
	ID        string    `bson:"_id" json:"id"`
	UserID    string    `bson:"userID" json:"userID"`
	Type      int       `bson:"type" json:"type"` //订单业务类型
	Teams     []string  `bson:"teams" json:"teams"`
	Views     []string  `bson:"views" json:"views"`     //都谁可查看订单
	Editors   []string  `bson:"editors" json:"editors"` //都谁可编辑订单
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	Status    int       `bson:"status" json:"status"`
}

// Business 业务信息
type Business struct {
	ID          string    `bson:"_id" json:"id"`
	OrderID     string    `bson:"orderID" json:"orderID"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	Description string    `bson:"description" json:"description"`
	Year        int       `bson:"year" json:"year"`
	Month       int       `bson:"month" json:"month"`
	Type        int       `bson:"type" json:"type"` //业务类型: 取单、做账、报税、回访
}

// CapitalInfo 资金情况
type CapitalInfo struct {
	ID         string  `bson:"_id" json:"id"`
	OrderID    string  `bson:"orderID" json:"orderID"`
	Year       int     `bson:"year" json:"year"`
	Month      int     `bson:"month" json:"month"`
	CashAmt    float64 `bson:"cashAmt" json:"cashAmt"`       //现金
	DepositAmt float64 `bson:"depositAmt" json:"depositAmt"` //存款
	ReceiveAmt float64 `bson:"receiveAmt" json:"receiveAmt"` //应收
	PayAmt     float64 `bson:"payAmt" json:"payAmt"`         //应付
}

// ProfitInfo 利润情况
type ProfitInfo struct {
	ID       string  `bson:"_id" json:"id"`
	OrderID  string  `bson:"orderID" json:"orderID"`
	Year     int     `bson:"year" json:"year"`
	Month    int     `bson:"month" json:"month"`
	InAmt    float64 `bson:"inAmt" json:"inAmt"`
	OutAmt   float64 `bson:"outAmt" json:"outAmt"`
	TotalAmt float64 `bson:"totalAmt" json:"totalAmt"`
}

// TaxInfo 纳税情况
type TaxInfo struct {
	ID        string    `bson:"_id" json:"id"`
	OrderID   string    `bson:"orderID" json:"orderID"`
	Year      int       `bson:"year" json:"year"`
	Month     int       `bson:"month" json:"month"`
	VatAmt    float64   `bson:"vatAmt" json:"vatAmt"` //增值税
	AbbAmt    float64   `bson:"abbAmt" json:"abbAmt"` //城建税
	Other     float64   `bson:"other" json:"other"`   //其它
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}
