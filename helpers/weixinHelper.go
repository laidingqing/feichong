package helpers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"

	"log"
)

// Session struct
type Session struct {
	SessionID string
	UserInfo  map[string]interface{}
	WxSession Jscode2Session
	Data      map[string]interface{}
}

// Jscode2Session success code
type Jscode2Session struct {
	ExpiresIn  int64  `json:"expiresIn"`
	OpenID     string `json:"openID"`
	SessionKey string `json:"sessionKey"`
}

// WxUserInfoWater ..
type WxUserInfoWater struct {
	Appid     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

// SessionStorage interface
type SessionStorage interface {
	Get(string) (Session, error)
	GetByOpenID(string) (Session, error)
	Set(string, Session) error
	Destroy(string) error
}

// TestStorage session test //demo, faeture support : redis memcached mysql
type TestStorage struct {
	Data       map[string]Session
	DataOpenID map[string]Session
	Lk         sync.Mutex
	SessionStorage
}

//get session by sessionId
func (this *TestStorage) Get(sessionId string) (Session, error) {
	this.Lk.Lock()
	defer this.Lk.Unlock()
	var sess Session
	var ok bool
	if sess, ok = this.Data[sessionId]; ok {
		return sess, nil
	} else {
		return sess, fmt.Errorf("Don't Exists")
	}
}

// GetByOpenID get session by openid
func (that *TestStorage) GetByOpenID(openID string) (Session, error) {
	that.Lk.Lock()
	defer that.Lk.Unlock()
	var sess Session
	var ok bool
	if sess, ok = that.DataOpenID[openID]; ok {
		return sess, nil
	} else {
		return sess, fmt.Errorf("Don't Exists")
	}
}

// Set set session by sessionid
func (that *TestStorage) Set(sessionID string, sess Session) error {
	that.Lk.Lock()
	defer that.Lk.Unlock()
	that.Data[sessionID] = sess
	that.DataOpenID[sess.WxSession.OpenID] = sess
	return nil
}

// Destroy destroy session by sessionId
func (that *TestStorage) Destroy(sessionID string) error {
	that.Lk.Lock()
	defer that.Lk.Unlock()
	var sess Session
	var ok bool
	if sess, ok = that.Data[sessionID]; ok {
		delete(that.Data, sessionID)
		delete(that.DataOpenID, sess.WxSession.OpenID)
	}
	return nil
}

//global sesion storage
var ss = NewTestStorage()

// NewTestStorage demo storage
func NewTestStorage() *TestStorage {
	ts := &TestStorage{}
	ts.Data = make(map[string]Session, 0)
	ts.DataOpenID = make(map[string]Session, 0)
	return ts
}

//为nil时则表示为系统内部错误，小程序前端获取接口errMsg为null则提示自定义错误，同示可显示logId，方便调试
var errMsgMap = map[int64]interface{}{
	400001: "无效sessionid",
	500001: nil, //code为空
	500002: nil,
	500003: nil,
	500004: nil,
}

// RandStr ..
func RandStr(len int64) string {
	b := make([]byte, int(math.Ceil(float64(len)/2.0)))
	rand.Seed(time.Now().UnixNano())
	rand.Read(b)
	return hex.EncodeToString(b)[0:len]
}

//fetch wx session by oauth code
func getWxSession(code string) (Jscode2Session, int64, error) {
	var ret Jscode2Session
	if code == "" {
		return ret, 500001, fmt.Errorf("Code为空")
	}
	urls := url.Values{}
	urls.Add("appid", "wxf5a6ca5f27f3d5cc")
	urls.Add("secret", "ef76ee8e04486cd4c4ba81188b5b9ccd")
	urls.Add("js_code", code)
	urls.Add("grant_type", "authorization_code")
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?" + urls.Encode())
	res, err := http.Get(url)
	if err != nil {
		return ret, 500002, err
	}
	info, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ret, 500003, err
	}
	err = json.Unmarshal(info, &ret)
	if err != nil || ret.OpenID == "" {
		return ret, 500004, fmt.Errorf(string(info))
	}
	return ret, 0, nil
}

//获取session
func getSession(code string) (Session, int64, error) {
	var sess Session
	res, errCode, err := getWxSession(code)
	log.Println(res, errCode, err)
	if err != nil {
		return sess, errCode, err
	}
	sess, err = ss.GetByOpenID(res.OpenID)
	if err != nil {
		//如果不存在
		sess.SessionID = RandStr(168)
		sess.WxSession = res
		sess.UserInfo = make(map[string]interface{}, 0)
		sess.Data = make(map[string]interface{}, 0)
		sess.Data["createTime"] = time.Now().String() //todo
		ss.Set(sess.SessionID, sess)
	} else if sess.WxSession.SessionKey != res.SessionKey {
		sess.WxSession = res
		ss.Set(sess.SessionID, sess)
	} else {

	}
	return sess, 0, nil
}

// GetSessionID  常规则情况下请使用GetSessionId
func GetSessionID(code string) (Jscode2Session, int64, error) {
	sess, errCode, err := getSession(code)
	return sess.WxSession, errCode, err
}

// GetUserInfo ..
func GetUserInfo(sessionID string) (Session, error) {
	var sess Session
	var err error
	sess, err = ss.Get(sessionID)
	return sess, err
}
