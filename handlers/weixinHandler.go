package handlers

import (
	"net/http"

	"github.com/laidingqing/feichong/helpers"
	"github.com/laidingqing/feichong/models"
	"github.com/laidingqing/feichong/managers"
)

const codeParam = "code"

// LoginSession ...
func LoginSession(w http.ResponseWriter, r *http.Request) {
	var weixin models.Weixin
	helpers.GetWeixinBody(w, r, &weixin)
	log := helpers.NewLogger()
	log.Log("code", weixin.Code, "encrypted", weixin.EncryptedData, "iv", weixin.Iv)
	session, errCode, err := helpers.GetSessionID(weixin.Code)

	if err != nil{
		helpers.SetResponse(w, http.StatusBadRequest, err)
	}
  // 创建用户
	user := managers.GetUserByOpenID(session.OpenID)
	if( user.ID.Hex() == ""){
		//New User
		var wxUser = models.User{
			Nick: weixin.UserInfo.NickName,
			Avatar: weixin.UserInfo.AvatarUrl,
			OpenID: session.OpenID,
		}
		managers.InsertUser(wxUser)
		errCode = 0
	}else{
		//Update User
	}

	if errCode == 0 {
		helpers.SetResponse(w, http.StatusOK, session)
	} else {
		helpers.SetResponse(w, http.StatusNotFound, err)
	}
}
