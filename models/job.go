package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	//DistrictConditionAll 福州区域
	DistrictConditionAll = iota //不限
	//DistrictConditionCS  仓山区
	DistrictConditionCS //
	//DistrictConditionTJ 台江区
	DistrictConditionTJ
	//DistrictConditionGL 鼓楼区
	DistrictConditionGL
	//DistrictConditionJA 晋安区
	DistrictConditionJA
	//DistrictConditionMH 闽侯县
	DistrictConditionMH
	//DistrictConditionCL 长乐市
	DistrictConditionCL
	//DistrictConditionMW 马尾区
	DistrictConditionMW
	//DistrictConditionFQ 福清市
	DistrictConditionFQ
	// DistrictConditionLJ 连江县
	DistrictConditionLJ
	//DistrictConditionPT 平潭县
	DistrictConditionPT
	//DistrictConditionLY 罗源县
	DistrictConditionLY
)
const (
	//JobTitleAll 职位，不限
	JobTitleAll = iota + 1
	//JobTitleKJ 会计
	JobTitleKJ
	//JobTitleCN 出纳
	JobTitleCN
	//JobTitleCWGW 财务顾问
	JobTitleCWGW
	//JobTitleJS 结算
	JobTitleJS
	//JobTitleSW 税务
	JobTitleSW
	//JobTitleSJ 审计
	JobTitleSJ
	//JobTitleFK 风控
	JobTitleFK
	//JobTitleCWJL 财务经理
	JobTitleCWJL
	//JobTitleCFO CFO
	JobTitleCFO
	//JobTitleCWZJ 财务总监
	JobTitleCWZJ
	//JobTitleCWZG 财务主管
	JobTitleCWZG
)

const (
	//WorkExpAll 工作经验，不限
	WorkExpAll = iota + 1
	//WorkExp102 应届生
	WorkExp102
	//WorkExp103 1年以内
	WorkExp103
	//WorkExp104 1-3年
	WorkExp104
	//WorkExp105 3-5年
	WorkExp105
	//WorkExp106 5-10年
	WorkExp106
	//WorkExp107 10年以上
	WorkExp107
)

const (
	//DegreeAll 学历 不限
	DegreeAll = iota + 1
	//Degree207 中专及以下
	Degree207
	//Degree206 高中
	Degree206
	//Degree202 大专
	Degree202
	//Degree203 本科
	Degree203
	//Degree204 硕士
	Degree204
	//Degree205 博士
	Degree205
)

const (
	//SalaryAll 薪资要求，不限
	SalaryAll = iota + 1
	//Salary2 3k以下
	Salary2
	//Salary3 3-5k
	Salary3
	//Salary4 5-10k
	Salary4
	//Salary5 10-15k
	Salary5
	//Salary6 15-20k
	Salary6
	//Salary7 20-30k
	Salary7
	//Salary8 30-50k
	Salary8
	//Salary9 50k以上
	Salary9
)

//EnterpriseInfo 企业雇主信息
type EnterpriseInfo struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	UserID      string        `bson:"userID" json:"userID"`
	BossID      string        `bson:"bossID" json:"bossID"`   //boss 直聘中的ID
	Name        string        `bson:"name" json:"name"`       //名称
	Intro       string        `bson:"intro" json:"intro"`     //介绍
	Address     string        `bson:"address" json:"address"` //地址
	Code        string        `bson:"code" json:"code"`       //代码证
	Website     string        `bson:"website" json:"website"`
	Tel         string        `bson:"tel" json:"tel"`
	CompanyType string        `bson:"companyType" json:"companyType"` //企业类型
	ResTime     string        `bson:"resTime" json:"resTime"`         //注册时间
	Capital     string        `bson:"capital" json:"capital"`         //注册金
	Industry    string        `bson:"industry" json:"industry"`       //行业
	LOGO        string        `bson:"logoURL" json:"logoURL"`
	Province    string        `bson:"province" json:"province"`
	City        string        `bson:"city" json:"city"`
	District    string        `bson:"district" json:"district"`
	CreatedAt   time.Time     `bson:"createdAt" json:"createdAt"`
	UpdateAt    time.Time     `bson:"updateAt" json:"updateAt,omitempty"`
	IsAuth      bool          `bson:"isAuth" json:"isAuth,omitempty"`
}

//Job 招聘职位
type Job struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	EnterpriseID string        `bson:"enterPriseID" json:"enterPriseID,omitempty"`
	JobTitle     int32         `bson:"jobTitle" json:"-,omitempty"`
	JobTitleView Dictionary    `bson:"-" json:"jobTitleView,omitempty"`
	WorkCity     int32         `bson:"workCity" json:"-,omitempty"`
	WorkCityView Dictionary    `bson:"-" json:"workCityView,omitempty"`
	WorkExp      int32         `bson:"workExp" json:"-,omitempty"`
	WorkExpView  Dictionary    `bson:"-" json:"orkExpView,omitempty"`
	Degree       int32         `bson:"degree" json:"-,omitempty"`
	DegreeView   Dictionary    `bson:"-" json:"degreeView,omitempty"`
	Salary       int32         `bson:"salary" json:"-,omitempty"`
	SalaryView   Dictionary    `bson:"-" json:"salaryView,omitempty"`
	JobDesc      string        `bson:"description" json:"description,omitempty"`
	PublishAt    time.Time     `bson:"publishAt" json:"publishAt,omitempty"`
	Status       int           `bson:"status" json:"status,omitempty"`
}

// Resume 简历
type Resume struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	UserID     string        `bson:"userId" json:"userId,omitempty"`
	CreatedBy  string        `bson:"createBy" json:"createBy,omitempty"`
	Name       string        `bson:"name" json:"name,omitempty"`
	Phone      string        `bson:"phone" json:"phone,omitempty"`
	Email      string        `bson:"email" json:"email,omitempty"`
	Gender     string        `bson:"gender" json:"gender,omitempty"`
	Birth      time.Time     `bson:"birth" json:"birth,omitempty"`
	Bio        string        `bson:"bio" json:"bio,omitempty"`
	Status     int32         `bson:"status" json:"status,omitempty"` //工作状态: 0，离职随时到岗，1,在职，考虑换工作
	Projects   []Project     `bson:"projects" json:"projects,omitempty"`
	Educations []Education   `bson:"educations" json:"educations,omitempty"`
	CreatedAt  time.Time     `bson:"createdAt" json:"createdAt"`
	UpdateAt   time.Time     `bson:"updateAt" json:"updateAt,omitempty"`
	IsAuth     bool          `bson:"isAuth" json:"isAuth,omitempty"`
	Recommand  bool          `bson:"recommand" json:"recommand"`
}

//Education 教育经历
type Education struct {
	Name        string `bson:"name" json:"name"`               //学校
	Major       string `bson:"major" json:"major"`             //专业
	Degree      string `bson:"degree" json:"degree"`           //学历
	StartDate   string `bson:"startAt" json:"startAt"`         //开始日期
	EndDate     string `bson:"endAt" json:"endAt"`             //结束日期
	Description string `bson:"description" json:"description"` //在校经历，描述
}

//Project 工作经历
type Project struct {
	StartDate string   `bson:"startAt" json:"startAt,omitempty"`   //开始日期
	EndDate   string   `bson:"endAt" json:"endAt,omitempty"`       //结束日期
	Position  string   `bson:"position" json:"position,omitempty"` //职位
	Content   string   `bson:"content" json:"content,omitempty"`   //工作内容
	Name      string   `bson:"name" json:"name,omitempty"`         //公司名称
	Tags      []string `bson:"tags" json:"tags,omitempty"`         //标签
}

//Entrust 委托认证
type Entrust struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	UserID         string        `bson:"userID" json:"userID"`
	ResumeID       string        `bson:"resumeID" json:"resumeID"`
	EnterpriseInfo string        `bson:"enterpriseID" json:"enterpriseID"`
	CreatedAt      time.Time     `bson:"createdAt" json:"createdAt"`
}
