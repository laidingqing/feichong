package managers

import (
	"time"

	"github.com/laidingqing/feichong/helpers"
	"github.com/laidingqing/feichong/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// InsertBusinessData 新增订单业务数据
func InsertBusinessData(biz models.Business) (models.Business, error) {
	biz.ID = bson.NewObjectId()
	biz.CreatedAt = time.Now()
	// biz.CapitalInfo = models.CapitalInfo{
	// 	BusinessID: biz.ID.Hex(),
	// }
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
		return models.Business{}, err
	}

	return biz, nil
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

// FindOrderBusinessByDate ...
func FindOrderBusinessByDate(orderID string, year int, month int) (models.Business, error) {
	var data models.Business

	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"orderID": orderID, "year": year, "month": month}).One(&data)
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

	query := func(c *mgo.Collection) error {
		return c.UpdateId(bson.ObjectIdHex(capital.BusinessID), business)
	}

	executeQuery(businessCollectionName, query)

	return models.CapitalInfo{}, nil
}

// UpdateFeedbackBusiness ..
func UpdateFeedbackBusiness(businessID string, feedback models.FeedBack) (models.Business, error) {
	business, err := FindBusinessByID(businessID)
	if err != nil {
		return models.Business{}, err
	}
	business.Star = feedback.Star
	business.Comment = feedback.Comment

	query := func(c *mgo.Collection) error {
		return c.UpdateId(bson.ObjectIdHex(businessID), business)
	}

	executeQuery(businessCollectionName, query)

	return business, nil
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

// GetFeedbacks 获取所有咨询列表
func GetFeedbacks(page int, size int) (models.Pagination, error) {
	bsonQuery := bson.M{} // 查询条件
	var consults []models.Consult
	session := getSession()
	defer session.Close()
	c := session.DB(databaseName).C(feedbackCollectionName)
	q := c.Find(bsonQuery)
	total, err := q.Count()
	q.All(&consults)
	return models.Pagination{
		Data:       consults,
		TotalCount: total,
	}, err
}

// PostFeedbacks 获取所有咨询列表
func PostFeedbacks(model models.Consult) (models.Consult, error) {
	model.ID = bson.NewObjectId()
	model.CreatedAt = time.Now()
	query := func(c *mgo.Collection) error {
		return c.Insert(model)
	}
	err := executeQuery(feedbackCollectionName, query)
	if err != nil {
		return models.Consult{
			ID: "",
		}, err
	}

	return model, nil
}
