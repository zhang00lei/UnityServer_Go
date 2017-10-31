package internal

import (
	"reflect"
	"server/msg"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/gate"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
func init() {
	handleMsg(&msg.CS_PlayerLogin{},CS_Login)
}

func CS_Login(args[] interface{}){
	m:=args[0].(*msg.CS_PlayerLogin)
	log.Debug(m.Account)
	client:=args[1].(gate.Agent)
	result:=msg.SC_PlayerLogin{}
	result.LoginResult= msg.SC_PlayerLogin_ACCOUNT_ERROR;
	client.WriteMsg(&result)
}