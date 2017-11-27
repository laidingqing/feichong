package handlers

import (
	"net/http"
	"strconv"
	"github.com/laidingqing/feichong/helpers"
	"github.com/laidingqing/feichong/managers"
	"github.com/laidingqing/feichong/models"
)

const userIDParam = "userId"
const userPageParam = "page"
const userSizeParam = "size"


const orderPageParam = "page"
const orderSizeParam = "size"
const orderCatalogParam = "catalog"

// GetUsers ...
func GetUsers(w http.ResponseWriter, r *http.Request) {
	page := helpers.GetInParam(r, userPageParam)
	size := helpers.GetInParam(r, userSizeParam)

	pageIndex, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(size)

	users, err := managers.GetUsers(pageIndex, pageSize)
	if err == nil {
		helpers.SetResponse(w, http.StatusOK, users)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	}
}

// GetUserByID ..
func GetUserByID(w http.ResponseWriter, r *http.Request) {

	userID := helpers.GetParam(r, userIDParam)

	model := managers.GetUserByID(userID)

	if (models.User{}) == model {
		helpers.SetResponse(w, http.StatusNotFound, nil)
	} else {
		helpers.SetResponse(w, http.StatusOK, model)
	}
}

// PostUser create user
func PostUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	helpers.GetUserBody(w, r, &user)
	user.Salt = user.UserName
	user.Password = helpers.CalculatePassHash(user.Password, user.Salt)
	isCreated := managers.InsertUser(user)

	if isCreated != "" {
		helpers.SetResponse(w, http.StatusCreated, nil)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, nil)
	}
}

// SelfUsers 获取管理用户
func SelfUsers(w http.ResponseWriter, r *http.Request) {
	users := managers.GetUsersBySelf()
	helpers.SetResponse(w, http.StatusOK, users)
}

// LoginUser 用户登录
func LoginUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	helpers.GetUserBody(w, r, &user)

	isUser := managers.GetUserByUserName(user.UserName)

	if isUser.Password != helpers.CalculatePassHash(isUser.Password, isUser.Salt){
		helpers.SetResponse(w, http.StatusBadRequest, models.ErrUserNotFound)
	}

	helpers.SetResponse(w, http.StatusCreated, isUser)

}
