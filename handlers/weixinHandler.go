package handlers

import (
	"net/http"

	"github.com/laidingqing/feichong/helpers"
	"github.com/laidingqing/feichong/models"
)

const codeParam = "code"

// LoginSession ...
func LoginSession(w http.ResponseWriter, r *http.Request) {
	var weixin models.Weixin
	helpers.GetWeixinBody(w, r, &weixin)

	session, errCode, err := helpers.GetSessionID(weixin.Code)
	if errCode == 0 {
		helpers.SetResponse(w, http.StatusOK, session)
	} else {
		helpers.SetResponse(w, http.StatusNotFound, err)
	}
}
