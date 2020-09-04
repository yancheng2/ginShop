package util

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var session = sessions.NewCookieStore([]byte("my_sessions_secret"))

func GetSession(key_name string, c *gin.Context) (*sessions.Session, error) {
	val, err := session.Get(c.Request, key_name)
	if err != nil {
		return nil, err
	}
	return val, err
}

func SetSession(key_name string, val interface{}, c *gin.Context) error {
	if key_name == "" || val == nil {
		return nil
	}
	// 先获取
	getSession, err := GetSession(key_name, c)
	if err != nil {
		return err
	}
	// 再保存
	getSession.Values[key_name] = val
	err1 := getSession.Save(c.Request, c.Writer)
	if err1 != nil {
		return err1
	}
	return nil
}

func GetSessID() {
}
