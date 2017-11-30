package managers

import (
	"time"

	"github.com/laidingqing/feichong/helpers"
	"github.com/laidingqing/feichong/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// InsertBusinessData 新增订单业务数据
func InsertBusinessData(biz models.Business) string {
	biz.ID = bson.NewObjectId()
	biz.CreatedAt = time.Now()
	biz.CapitalInfo = models.CapitalInfo{
		BusinessID: biz.ID.Hex(),
	}
	biz.TaxInfo = models.TaxInfo{
		BusinessID: biz.ID.Hex(),
	}
	biz.ProfitInfo = models.ProfitInfo{
		BusinessID: biz.ID.Hex(),
	}
	query := func(c *mgo.Collection) error {
		return c.Insert(biz)
	}

	err := executeQuery(businessCollectionName, query)
	if err != nil {
		return ""
	}

	return biz.ID.Hex()
}

// FindOrderBusiness ...
func FindOrderBusiness(orderID string) ([]models.Business, error) {
	var data []models.Business

	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"orderID": orderID}).Sort("sorter").All(&data)
	}

	err := executeQuery(businessCollectionName, query)
	if err != nil {
		return data, err
	}
	return data, nil
}

// FindBusinessByID ...
func FindBusinessByID(businessID string) (models.Business, error) {
	var data models.Business

	query := func(c *mgo.Collection) error {
		return c.FindId(bson.ObjectIdHex(businessID)).One(&data)
	}

	err := executeQuery(businessCollectionName, query)

	log := helpers.NewLogger()
	log.Log("data", data.ID)
	if err != nil {
		return data, err
	}
	return data, nil
}

// UpdateCapitalByBusiness ..
func UpdateCapitalByBusiness(capital models.CapitalInfo) (models.CapitalInfo, error) {
	business, err := FindBusinessByID(capital.BusinessID)
	log := helpers.NewLogger()
	log.Log("data", business.ID)
	if err != nil {
		return models.CapitalInfo{}, err
	}

	business.CapitalInfo = capital

	query := func(c *mgo.Collection) error {
		return c.UpdateId(bson.ObjectIdHex(capital.BusinessID), business)
	}

	executeQuery(businessCollectionName, query)

	return business.CapitalInfo, nil
}

// UpdateProfitByBusiness ..
func UpdateProfitByBusiness(profit models.ProfitInfo) (models.ProfitInfo, error) {

	business, err := FindBusinessByID(profit.BusinessID)

	if err != nil {
		return models.ProfitInfo{}, err
	}

	business.ProfitInfo = profit

	query := func(c *mgo.Collection) error {
		return c.UpdateId(bson.ObjectIdHex(profit.BusinessID), business)
	}

	executeQuery(businessCollectionName, query)

	return business.ProfitInfo, nil
}

// UpdateTaxByBusiness ..
func UpdateTaxByBusiness(tax models.TaxInfo) (models.TaxInfo, error) {

	business, err := FindBusinessByID(tax.BusinessID)

	if err != nil {
		return models.TaxInfo{}, err
	}

	business.TaxInfo = tax

	query := func(c *mgo.Collection) error {
		return c.UpdateId(bson.ObjectIdHex(tax.BusinessID), business)
	}

	executeQuery(businessCollectionName, query)

	return business.TaxInfo, nil
}
