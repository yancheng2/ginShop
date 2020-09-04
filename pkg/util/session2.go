package util

//
//import (
//	"github.com/astaxie/session"
//	_ "github.com/astaxie/session/providers/memory"
//	"net/http"
//)
//
//var globalSession *session.Manager
//var sess session.Session
//
//// 初始化session
//func SessionsInit(){
//	globalSession, _ = session.NewManager("memory", "gosessionid", 3600)
//	go globalSession.GC()
//}
//func SessionSetUp(w http.ResponseWriter, r *http.Request) session.Session {
//	//运行打开session
//	sess = globalSession.SessionStart(w, r)
//	return sess
//
//}
//func SetSession(key string, value interface{}) error {
//	err := sess.Set(key, value)
//	return err
//}
//
//func GetSession(key string)interface{}{
//	val := sess.Get(key)
//	return val
//}
//
//func DelSession(key string) error {
//	err := sess.Delete(key)
//	return err
//}
//
//func GetSessionID() string {
//	return sess.SessionID()
//}
