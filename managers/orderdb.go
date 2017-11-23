package managers

import (
	"github.com/laidingqing/feichong/models"
	"github.com/laidingqing/feichong/helpers"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetOrders 获取所有订单
func GetOrders(page int, size int, catalog int) (models.Pagination, error) {

	bson := bson.M{"type": catalog} // 查询条件
	logger := helpers.NewLogger()
	var orders []models.Order
	session := getSession()
	defer session.Close()
	c := session.DB(databaseName).C(orderCollectionName)
	q := c.Find(bson)
	total, err := q.Count()
	logger.Log("err", err, "total", total)

	q.All(&orders)

	logger.Log("data", len(orders))
	return models.Pagination{
		Data: orders,
		TotalCount: total,
	}, err
}

// GetOrderByID 根据用户编号获取订单信息
func GetOrderByID(orderID string) (models.Order, error) {

	var order models.Order

	if !bson.IsObjectIdHex(orderID) {
		return order, nil
	}

	query := func(c *mgo.Collection) error {
		return c.FindId(orderID).One(&order)
	}

	err := executeQuery(orderCollectionName, query)

	return order, err
}

// PutOrder 修改订单谁可见
func PutOrder(orderID string, order models.Order) (models.Order, error) {

	if !bson.IsObjectIdHex(orderID) {
		return order, nil
	}

	query := func(c *mgo.Collection) error {
		return c.UpdateId(orderID, order)
	}

	err := executeQuery(orderCollectionName, query)

	return order, err
}

// InsertOrder 新增订单
func InsertOrder(order models.Order) string {
	order.ID = bson.NewObjectId()
	query := func(c *mgo.Collection) error {
		return c.Insert(order)
	}

	err := executeQuery(orderCollectionName, query)
	if err != nil {
		return ""
	}

	return order.ID.Hex()
}
