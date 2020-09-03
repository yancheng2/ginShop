package util

import (
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
)

var globalSession *session.Manager

func sessionStart(){
	globalSession, _ = session.NewManager("ginshop", "gosessionid", 3600)
	go globalSession.GC()
}
