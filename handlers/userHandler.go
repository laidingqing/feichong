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
const userNameParam = "username"

const orderPageParam = "page"
const orderSizeParam = "size"
const orderCatalogParam = "catalog"
const orderNoParam = "orderNo"

// GetUsers ...
func GetUsers(w http.ResponseWriter, r *http.Request) {

	helpers.ValidateTokenMiddleware(w, r, func(w http.ResponseWriter, r *http.Request) {
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
	})

}

// GetUserByID ..
func GetUserByID(w http.ResponseWriter, r *http.Request) {

	userID := helpers.GetParam(r, userIDParam)
	model := managers.GetUserByID(userID)
	helpers.SetResponse(w, http.StatusOK, model)
}

// PutUserByID Update User Profile ...
func PutUserByID(w http.ResponseWriter, r *http.Request) {
	var user models.User
	helpers.GetUserBody(w, r, &user)
	model := managers.UpdateUserByID(user)
	helpers.SetResponse(w, http.StatusOK, model)
}

// PutUserSecurity Update User Profile ...
func PutUserSecurity(w http.ResponseWriter, r *http.Request) {
	userID := helpers.GetParam(r, userIDParam)
	var user models.User
	helpers.GetUserBody(w, r, &user)
	user.Salt = user.UserName
	user.Password = helpers.CalculatePassHash(user.Password, user.Salt)
	model, err := managers.UpdateUserPasswordAndName(userID, user.UserName, user.Password)
	if err != nil {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	} else {
		helpers.SetResponse(w, http.StatusOK, models.User{
			ID: model.ID,
		})
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

// EnterPriseUsers ...
func EnterPriseUsers(w http.ResponseWriter, r *http.Request) {
	users := managers.GetUsersByEnterPrise()
	helpers.SetResponse(w, http.StatusOK, users)
}

// LoginUser 用户登录
func LoginUser(w http.ResponseWriter, r *http.Request) {
	log := helpers.NewLogger()
	log.Log("username", "dfasdfasdf")

	var user models.User
	helpers.GetUserBody(w, r, &user)

	isUser := managers.GetUserByUserName(user.UserName)

	if isUser.Password != helpers.CalculatePassHash(user.Password, isUser.Salt) {
		helpers.SetResponse(w, http.StatusBadRequest, models.ErrUserNotFound)
	} else {
		t, err := helpers.CreateJWT()
		if err != nil {
			helpers.SetResponse(w, http.StatusForbidden, err)
		} else {
			var session = &helpers.Jscode2Session{
				SessionKey: t,
				UserID:     isUser.ID.Hex(),
				Name:       isUser.Name,
			}
			helpers.SetResponse(w, http.StatusOK, session)
		}
	}
}

// CheckUserName ...
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	var user models.User
	username := helpers.GetInParam(r, userNameParam)
	log := helpers.NewLogger()
	log.Log("username", username)
	user = managers.GetUserByUserName(username)
	helpers.SetResponse(w, http.StatusOK, user)
}
