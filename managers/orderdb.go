package managers

import (
	"time"

	"github.com/laidingqing/feichong/helpers"
	"github.com/laidingqing/feichong/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetOrders 获取所有订单
func GetOrders(page int, size int, catalog int) (models.Pagination, error) {

	bsonQuery := bson.M{"catalog": catalog} // 查询条件
	logger := helpers.NewLogger()
	var orders []models.Order
	session := getSession()
	defer session.Close()
	c := session.DB(databaseName).C(orderCollectionName)
	q := c.Find(bsonQuery)
	total, err := q.Count()

	q.All(&orders)

	for i := 0; i < len(orders); i++ {
		var o = &orders[i]
		var user *models.User
		dbref := &mgo.DBRef{
			Collection: "users",
			Id:         o.SalerID.Id,
		}
		if err = session.DB(databaseName).FindRef(dbref).One(&user); err != nil {
			logger.Log("err", err)
		}
		o.SalerInfo = models.User{
			ID:   user.ID,
			Nick: user.Nick,
		}

		var cusUser *models.User
		userRef := &mgo.DBRef{
			Collection: "users",
			Id:         o.UserID.Id,
		}
		if err = session.DB(databaseName).FindRef(userRef).One(&cusUser); err != nil {
			logger.Log("err", err)
		}

		o.UserInfo = models.User{
			ID:          user.ID,
			Nick:        user.Nick,
			CompanyName: user.CompanyName,
		}

	}

	return models.Pagination{
		Data:       orders,
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

// DeleteOrderByID remove
func DeleteOrderByID(orderID string) error {

	if !bson.IsObjectIdHex(orderID) {
		return nil
	}

	query := func(c *mgo.Collection) error {
		return c.Update(bson.M{"_id": bson.ObjectIdHex(orderID)}, bson.M{"status": int(models.OrderStatusDeleted)})
	}

	err := executeQuery(orderCollectionName, query)

	return err
}

// GetOrdersByUser ...
func GetOrdersByUser(userID string) ([]models.Order, error) {

	var orders []models.Order

	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"userid.$id": bson.ObjectIdHex(userID), "status": int(models.OrderStatusDoing)}).All(&orders)
	}

	err := executeQuery(orderCollectionName, query)

	log := helpers.NewLogger()
	log.Log("orders", len(orders))

	return orders, err
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
func InsertOrder(order models.Order) (models.Order, error) {
	order.ID = bson.NewObjectId()
	order.CreatedAt = time.Now()
	order.Status = models.OrderStatusDoing
	query := func(c *mgo.Collection) error {
		return c.Insert(order)
	}
	order.SalerID = &mgo.DBRef{
		Collection: "users",
		Id:         order.SalerInfo.ID,
	}

	order.UserID = &mgo.DBRef{
		Collection: "users",
		Id:         order.UserInfo.ID,
	}

	err := executeQuery(orderCollectionName, query)
	if err != nil {
		return models.Order{
			ID: "",
		}, err
	}

	return order, nil
}
