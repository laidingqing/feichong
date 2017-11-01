package managers

import (
	"time"

	"github.com/laidingqing/feichong/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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
