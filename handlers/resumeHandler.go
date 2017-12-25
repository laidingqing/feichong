package handlers

import (
	"net/http"

	"github.com/laidingqing/feichong/helpers"
	"github.com/laidingqing/feichong/managers"
	"github.com/laidingqing/feichong/models"
)

const jobUserIDParam = "userId"

// GetUserResume ..
func GetUserResume(w http.ResponseWriter, r *http.Request) {
	userID := helpers.GetParam(r, jobUserIDParam)
	res, err := managers.GetResumeByUser(userID)

	if err == nil {
		helpers.SetResponse(w, http.StatusOK, res)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	}

}

// UpdateUserResume ..
func UpdateUserResume(w http.ResponseWriter, r *http.Request) {
	userID := helpers.GetParam(r, jobUserIDParam)
	var resume models.Resume
	helpers.GetResumeBody(w, r, &resume)

	res, err := managers.UpdateResumeByUser(userID, resume)

	if err == nil {
		helpers.SetResponse(w, http.StatusOK, res)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	}

}
