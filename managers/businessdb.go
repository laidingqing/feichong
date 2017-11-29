package managers

import (
	"time"

	"github.com/laidingqing/feichong/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// InsertBusinessData 新增订单业务数据
func InsertBusinessData(biz models.Business) string {
	biz.ID = bson.NewObjectId().Hex()
	biz.CreatedAt = time.Now()
	query := func(c *mgo.Collection) error {
		return c.Insert(biz)
	}

	err := executeQuery(businessCollectionName, query)
	if err != nil {
		return ""
	}

	return biz.ID
}

// FindOrderBusiness ...
func FindOrderBusiness(orderID string, year int, month int) ([]models.Business, error) {
	var data []models.Business

	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"orderID": orderID, "month": month, "year": year}).All(&data)
	}

	err := executeQuery(businessCollectionName, query)
	if err != nil {
		return data, err
	}
	return data, nil
}

// InsertOrderTax 新增用户
func InsertOrderTax(tax models.TaxInfo) string {
	tax.ID = bson.NewObjectId().Hex()
	tax.CreatedAt = time.Now()
	query := func(c *mgo.Collection) error {
		return c.Insert(tax)
	}

	err := executeQuery(orderTaxsCollectionName, query)
	if err != nil {
		return ""
	}

	return tax.ID
}

// GetOrderTaxs 新增用户
func GetOrderTaxs(orderID string, month int) (models.TaxInfo, error) {

	var tax models.TaxInfo
	year, _, _ := time.Now().Date()
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"orderID": orderID, "month": month, "year": year}).One(&tax)
	}

	err := executeQuery(orderTaxsCollectionName, query)
	if err != nil {
		return tax, nil
	}

	return tax, nil
}

// GetOrderCapitals ...
func GetOrderCapitals(orderID string, month int) (models.CapitalInfo, error) {

	var capital models.CapitalInfo
	year, _, _ := time.Now().Date()
	query := func(c *mgo.Collection) error {
		bs := bson.M{"orderID": orderID, "year": year}
		return c.Find(bs).One(&capital)
	}

	err := executeQuery(orderTaxsCollectionName, query)
	if err != nil {
		return capital, nil
	}

	return capital, nil
}
