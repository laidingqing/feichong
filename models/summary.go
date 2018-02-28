package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	//SummaryJobType 会计在线汇总信息
	SummaryJobType = iota
)

//Summary 汇总表
type Summary struct {
	ID    bson.ObjectId          `bson:"_id" json:"id"`
	Type  int32                  `bson:"type" json:"type,omitempty"`
	Props map[string]interface{} `bson:"props" json:"props,omitempty"`
}

//ResumeSummary 人才汇总
type ResumeSummary struct {
	Total    int32 `bson:"total" json:"total,omitempty"`       //总共
	Recently int32 `bson:"recently" json:"recently,omitempty"` //新注册的

}

//CorporationSummary 企业汇总
type CorporationSummary struct {
}

//CertifySummary 认证汇总
type CertifySummary struct {
}
