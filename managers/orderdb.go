package managers

import (
	"github.com/laidingqing/feichong/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetOrders 获取所有订单
func GetOrders() []models.Order {

	var orders []models.Order
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&orders)
	}

	executeQuery(orderCollectionName, query)
	return orders
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
	order.ID = bson.NewObjectId().Hex()
	query := func(c *mgo.Collection) error {
		return c.Insert(order)
	}

	err := executeQuery(orderCollectionName, query)
	if err != nil {
		return ""
	}

	return order.ID
}
