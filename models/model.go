package models

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// BusinessStatus 记账业务状态 取单、做账、报税、回访，完成
type BusinessStatus int

const (
	// BusinessStatusUnknown 未知
	BusinessStatusUnknown BusinessStatus = iota
	// BusinessStatusGet 取单
	BusinessStatusGet
	// BusinessStatusDo 做账
	BusinessStatusDo
	// BusinessStatusTax 报税
	BusinessStatusTax
	// BusinessStatusRes 回访
	BusinessStatusRes
	// BusinessStatusFinish 完成
	BusinessStatusFinish
)

// OrderStatus 订单状态
type OrderStatus int

const (
	// OrderStatusUnknown 未知
	OrderStatusUnknown OrderStatus = iota
	// OrderStatusDoing 正常
	OrderStatusDoing
	// BusinessStatusDeleted 删除
	OrderStatusDeleted
)

// Pagination 分页
type Pagination struct {
	Data       interface{} `bson:"-" json:"data"`
	TotalCount int         `bson:"-" json:"totalCount"`
}

// User 用户信息
type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	UserName    string        `bson:"username" json:"username,omitempty"`
	Password    string        `bson:"password" json:"password,omitempty"`
	Salt        string        `bson:"salt" json:"-"`
	Nick        string        `bson:"nick" json:"nick"`
	Email       string        `bson:"email" json:"email,omitempty"`
	Name        string        `bson:"name" json:"name,omitempty"`
	Phone       string        `bson:"phone" json:"phone,omitempty"`
	CompanyName string        `bson:"companyName" json:"companyName,omitempty"`
	Admin       bool          `bson:"admin" json:"admin,omitempty"`
	OpenID      string        `bson:"openID" json:"openID,omitempty"`
	CreatedAt   time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
	Avatar      string        `bson:"avatar" json:"avatar,omitempty"`
}

// Order 订单信息
type Order struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	OrderNO     string        `bson:"orderNO" json:"orderNO"`
	Catalog     int           `bson:"catalog" json:"catalog"`       //订单业务类型, 1: 账务记账，2：企业注册
	OrderMonth  int           `bson:"orderMonth" json:"orderMonth"` //合同月份
	StartAt     time.Time     `bson:"startDate" json:"startDate"`
	CreatedAt   time.Time     `bson:"createdAt" json:"createdAt"`
	ExpiredAt   time.Time     `bson:"expiredAt" json:"expiredAt"`
	Status      OrderStatus   `bson:"status" json:"status"`
	Company     string        `bson:"companyName" json:"companyName"`
	SalerID     *mgo.DBRef    `bson:"saler,omitempty" json:"salerId,omitempty"` //业务
	UserID      *mgo.DBRef    `bson:"userid,omitempty" json:"userId,omitempty"`
	ServiceID   *mgo.DBRef    `bson:"serviceid,omitempty" json:"serviceId,omitempty"` //客服
	AdviserID   *mgo.DBRef    `bson:"adviserid,omitempty" json:"adviserId,omitempty"` //财务顾问
	SalerInfo   User          `bson:"-" json:"salerInfo"`                             //业务员
	UserInfo    User          `bson:"-" json:"userInfo"`                              //所属用户
	ServiceInfo User          `bson:"-" json:"serviceInfo"`                           //客服
	AdviserInfo User          `bson:"-" json:"adviserInfo"`                           //顾问
}

// Business 业务信息-交接单
type Business struct {
	ID          bson.ObjectId      `bson:"_id" json:"id"`
	Seq         int                `bson:"sorter" json:"-"`
	OrderID     string             `bson:"orderID" json:"orderID"`
	OrderNO     string             `bson:"orderNO" json:"orderNO"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	Description string             `bson:"description" json:"description"`
	Year        int                `bson:"year" json:"year"`
	Month       int                `bson:"month" json:"month"`
	Star        int                `bson:"star" json:"star"`       //客户评星
	Comment     string             `bson:"comment" json:"comment"` //客户评价
	ProfitInfo  ProfitInfo         `bson:"profitInfo" json:"profitInfo"`
	TaxInfo     TaxInfo            `bson:"taxInfo" json:"taxInfo"`
	Progress    []BusinessProgress `bson:"progress" json:"progress"`
}

// BusinessProgress ...
type BusinessProgress struct {
	Status    BusinessStatus `bson:"catalog" json:"catalog"`
	CreatedAt time.Time      `bson:"createdAt" json:"createdAt"`
}

// CapitalInfo 资金情况, 不用
type CapitalInfo struct {
	OrderID    string  `bson:"orderID" json:"orderID"`
	BusinessID string  `bson:"businessID" json:"businessID"`
	Year       int     `bson:"year" json:"year"`
	Month      int     `bson:"month" json:"month"`
	CashAmt    float64 `bson:"cashAmt" json:"cashAmt"`       //现金
	DepositAmt float64 `bson:"depositAmt" json:"depositAmt"` //存款
	ReceiveAmt float64 `bson:"receiveAmt" json:"receiveAmt"` //应收
	PayAmt     float64 `bson:"payAmt" json:"payAmt"`         //应付
}

// ProfitInfo 利润情况， 待完善
type ProfitInfo struct {
	OrderID    string  `bson:"orderID" json:"orderID"`
	BusinessID string  `bson:"businessID" json:"businessID"`
	Year       int     `bson:"year" json:"year"`
	Month      int     `bson:"month" json:"month"`
	InAmt      float64 `bson:"inAmt" json:"inAmt"`
	OutAmt     float64 `bson:"outAmt" json:"outAmt"`
	TotalAmt   float64 `bson:"totalAmt" json:"totalAmt"`
}

// TaxInfo 纳税情况
type TaxInfo struct {
	OrderID      string    `bson:"orderID" json:"orderID"`
	BusinessID   string    `bson:"businessID" json:"businessID"`
	Year         int       `bson:"year" json:"year"`
	Month        int       `bson:"month" json:"month"`
	VatAmt       float64   `bson:"vatAmt" json:"vatAmt"`             //增值税
	PersonAmt    float64   `bson:"personAmt" json:"personAmt"`       //个税
	AbbAmt       float64   `bson:"abbAmt" json:"abbAmt"`             //城建税
	EducationAmt float64   `bson:"educationAmt" json:"educationAmt"` //教育附加
	LocalAmt     float64   `bson:"localAmt" json:"localAmt"`         //地方
	WaterAmt     float64   `bson:"waterAmt" json:"waterAmt"`         //水利
	TaxAmt       float64   `bson:"taxAmt" json:"taxAmt"`             //税金
	Reported     bool      `bson:"reported" json:"reported"`
	CreatedAt    time.Time `bson:"createdAt" json:"createdAt"`
	ReportedAt   time.Time `bson:"reportedAt" json:"reportedAt"`
}

// Consult 咨询表单记录
type Consult struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	From        string        `bson:"from" json:"from"`
	Invite      string        `bson:"invite" json:"invite"`
	CompanyName string        `bson:"companyName" json:"companyName"`
	Name        string        `bson:"name" json:"name"`
	Phone       string        `bson:"phone" json:"phone"`
	Description string        `bson:"description" json:"description"`
	CreatedAt   time.Time     `bson:"createdAt" json:"createdAt"`
	Catalog     string        `bson:"catalog" json:"catalog"` //业务分类
	Biz         string        `bson:"biz" json:"biz"`         //业务描述
}

// FeedBack 评价
type FeedBack struct {
	Star    int    `json:"star"`    //客户评星
	Comment string `json:"comment"` //客户评价
}

var (
	// ErrUserNotFound 用户不存在
	ErrUserNotFound = errors.New("用户不存在")
)
