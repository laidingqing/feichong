package managers

import (
	"time"

	"github.com/laidingqing/feichong/models"
	"github.com/laidingqing/feichong/helpers"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetUsers 获取所有用户
func GetUsers(page int, size int) (models.Pagination, error) {
	bson := bson.M{} // 查询条件
	logger := helpers.NewLogger()
	var users []models.User
	session := getSession()
	defer session.Close()
	c := session.DB(databaseName).C(userCollectionName)
	q := c.Find(bson)
	total, err := q.Count()
	logger.Log("err", err, "total", total)

	q.All(&users)

	logger.Log("data", len(users))
	return models.Pagination{
		Data: users,
		TotalCount: total,
	}, err
}

// GetUserByID 根据用户编号获取用户信息
func GetUserByID(userID string) models.User {

	var user models.User

	if !bson.IsObjectIdHex(userID) {
		return user
	}

	query := func(c *mgo.Collection) error {
		return c.FindId(userID).One(&user)
	}

	executeQuery(userCollectionName, query)

	return user
}

// GetUserByUserName 根据用户获取用户信息
func GetUserByUserName(username string) models.User {

	var user models.User

	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"username": username}).One(&user)
	}

	executeQuery(userCollectionName, query)

	return user
}

// InsertUser 新增用户
func InsertUser(user models.User) string {
	user.ID = bson.NewObjectId()
	user.CreatedAt = time.Now()
	query := func(c *mgo.Collection) error {
		return c.Insert(user)
	}

	err := executeQuery(userCollectionName, query)
	if err != nil {
		return ""
	}

	return user.ID.Hex()
}
