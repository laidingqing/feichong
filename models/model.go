package models

import "time"
import "errors"
import (
	"gopkg.in/mgo.v2/bson"
)
// Pagination 分页

type Pagination struct {
	Data interface{} `bson:"-" json:"data"`
	TotalCount int `bson:"-" json:"totalCount"`
}

// User 用户信息
type User struct {
	ID          bson.ObjectId    `bson:"_id" json:"id"`
	UserName    string    `bson:"username" json:"username"`
	Password    string    `bson:"password" json:"password"`
	Salt				string    `bson:"salt" json:"-"`
	Nick        string    `bson:"nick" json:"nick"`
	Email       string    `bson:"email" json:"email"`
	Name        string    `bson:"name" json:"name"`
	Phone       string    `bson:"phone" json:"phone"`
	CompanyName string    `bson:"companyName" json:"companyName"`
	Admin       bool      `bson:"admin" json:"admin"`
	OpenID      string    `bson:"openID" json:"openID"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	Avatar	    string    `bson:"avatar" json:"avatar"`
}

// Order 订单信息
type Order struct {
	ID        bson.ObjectId    `bson:"_id" json:"id"`
	UserID    string    `bson:"userID" json:"userID"`
	OrderNO   string    `bson:"orderNO" json:"orderNO"`
	Saler			string    `bson:"saler" json:"saler"`
	Type      int       `bson:"type" json:"type"` //订单业务类型
	Teams     []string  `bson:"teams" json:"teams"`
	Views     []string  `bson:"views" json:"views"`     //都谁可查看订单
	Editors   []string  `bson:"editors" json:"editors"` //都谁可编辑订单
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	Status    int       `bson:"status" json:"status"`
	Company	  string    `bson:"companyName" json:"companyName"`
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

// 咨询表单记录
type Consult struct {
		ID string `bson:"_id" json:"id"`
		From string `bson:"from" json:"from"`
		Invite string `bson:"invite" json:"invite"`
		Name string `bson:"name" json:"name"`
		Phone string `bson:"phone" json:"phone"`
		Description string `bson:"description" json:"description"`
		CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
		Catalog string `bson:"catalog" json:"catalog"`
		Biz string `bson:"biz" json:"biz"`
}

var (
	ErrUserNotFound = errors.New("用户不存在")
)
