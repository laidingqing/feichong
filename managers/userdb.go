package managers

import (
	"time"

	"github.com/laidingqing/feichong/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetUsers 获取所有用户
func GetUsers() []models.User {

	var users []models.User
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&users)
	}

	executeQuery(userCollectionName, query)
	return users
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

// InsertUser 新增用户
func InsertUser(user models.User) string {
	user.ID = bson.NewObjectId().Hex()
	user.CreatedAt = time.Now()
	query := func(c *mgo.Collection) error {
		return c.Insert(user)
	}

	err := executeQuery(userCollectionName, query)
	if err != nil {
		return ""
	}

	return user.ID
}
