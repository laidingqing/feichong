package handlers

import (
	"net/http"

	"github.com/laidingqing/feichong/helpers"
	"github.com/laidingqing/feichong/managers"
	"github.com/laidingqing/feichong/models"
)

const userIDParam = "userId"

// GetUsers ...
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := managers.GetUsers()
	if len(users) > 0 {
		helpers.SetResponse(w, http.StatusOK, users)
	} else {
		helpers.SetResponse(w, http.StatusNotFound, nil)
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

	isCreated := managers.InsertUser(user)

	if isCreated != "" {
		helpers.SetResponse(w, http.StatusCreated, nil)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, nil)
	}
}
