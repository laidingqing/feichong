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

	bsonQuery := bson.M{"status": int(models.OrderStatusDoing)} // 查询条件
	logger := helpers.NewLogger()
	var orders []models.Order
	session := getSession()
	defer session.Close()
	c := session.DB(databaseName).C(orderCollectionName)
	q := c.Find(bsonQuery).Skip(page).Limit(size)
	total, err := q.Count()

	q.All(&orders)
	logger.Log("size", len(orders))
	for i := 0; i < len(orders); i++ {
		var o = &orders[i]
		if o.UserID != "" {
			user := GetUserByID(o.UserID)
			o.UserInfo = user
		}
		if o.AdviserID != "" {
			user := GetUserByID(o.AdviserID)
			o.AdviserInfo = user
		}
		if o.SalerID != "" {
			user := GetUserByID(o.SalerID)
			o.SalerInfo = user
		}
		if o.ServiceID != "" {
			user := GetUserByID(o.ServiceID)
			o.ServiceInfo = user
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

// GetOrderByNo 根据合同号获取订单信息
func GetOrderByNo(orderNo string) (models.Order, error) {
	var order models.Order

	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"orderNO": orderNo}).One(&order)
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
		data := bson.M{"$set": bson.M{"status": int(models.OrderStatusDeleted)}}
		return c.Update(bson.M{"_id": bson.ObjectIdHex(orderID)}, data)
	}

	err := executeQuery(orderCollectionName, query)
	if err != nil {
		return err
	}
	return nil
}

// GetOrdersByUser ...
func GetOrdersByUser(userID string) ([]models.Order, error) {

	var orders []models.Order

	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"userid": userID, "status": int(models.OrderStatusDoing)}).All(&orders)
	}

	err := executeQuery(orderCollectionName, query)

	log := helpers.NewLogger()
	log.Log("orders", len(orders))

	orderList := []models.Order{}
	for _, order := range orders {

		order.SalerInfo = GetUserByID(order.SalerID)
		order.AdviserInfo = GetUserByID(order.AdviserID)
		order.ServiceInfo = GetUserByID(order.ServiceID)
		order.UserInfo = GetUserByID(order.UserID)
		orderList = append(orderList, order)
	}

	return orderList, err
}

// GetOrdersByNO 根据合同号查询合同
func GetOrdersByNO(orderNO string) (models.Pagination, error) {
	session := getSession()
	defer session.Close()
	var orders []models.Order
	bsonQuery := bson.M{"orderNO": orderNO} // 查询条件

	c := session.DB(databaseName).C(orderCollectionName)
	q := c.Find(bsonQuery)
	total, err := q.Count()
	q.All(&orders)

	return models.Pagination{
		Data:       orders,
		TotalCount: total,
	}, err
}

// PutOrder 修改订单谁可见
func PutOrder(orderID string, order models.Order) (models.Order, error) {

	if !bson.IsObjectIdHex(orderID) {
		return order, nil
	}

	query := func(c *mgo.Collection) error {
		return c.UpdateId(bson.ObjectIdHex(orderID), order)
	}
	log := helpers.NewLogger()
	err := executeQuery(orderCollectionName, query)
	if err != nil {
		log.Log("put order err", err)
		return models.Order{
			ID: "",
		}, err
	}

	return order, nil
}

// InsertOrder 新增订单
func InsertOrder(order models.Order) (models.Order, error) {
	order.ID = bson.NewObjectId()
	order.CreatedAt = time.Now()
	order.Status = models.OrderStatusDoing

	query := func(c *mgo.Collection) error {
		return c.Insert(order)
	}

	err := executeQuery(orderCollectionName, query)

	if err != nil {
		return models.Order{
			ID: "",
		}, err
	}

	return order, nil
}
