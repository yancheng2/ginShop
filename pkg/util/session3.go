package util

//import (
//	"github.com/gin-contrib/sessions"
//	"github.com/gin-contrib/sessions/cookie"
//	"github.com/gin-gonic/gin"
//)
//var store sessions.Store
//func SessionSetUp(){
//	r := gin.Default()
//	store = cookie.NewStore([]byte("secret"))
//	r.Use(sessions.Sessions("mysession",store))
//
//}
//
//func GetSession(c *gin.Context,key string) interface{} {
//	sess := sessions.Default(c)
//	option := sessions.Options{MaxAge: 3600}
//	sess.Options(option)
//	ret := sess.Get(key)
//	return ret
//}
//
//func SetSession(c *gin.Context,key string, val interface{}) error {
//	sess := sessions.Default(c)
//	option := sessions.Options{MaxAge: 3600}
//	sess.Options(option)
//	sess.Set(key,val)
//	err := sess.Save()
//	return err
//}
//
//func SessionID(c *gin.Context) {
//	//c.Cookie("")
//}
