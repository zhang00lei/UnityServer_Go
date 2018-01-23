package internal

import (
	"reflect"
	"server/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(msg.CS_PlayerReady{},cs_PlayerReady)
	handleMsg(msg.CS_PlayerCancelReady{},cs_PlayerCancelReady)
}

func cs_PlayerReady(args[] interface{})  {
	log.Debug("cs_playerReady")
	result:=msg.SC_PlayerReady{}
	result.Result=0
	client := args[1].(gate.Agent)
	client.WriteMsg(&result)
}

func cs_PlayerCancelReady(args[] interface{})  {
	log.Debug("cancel ready")
	result:=msg.SC_PlayerCancelReady{}
	result.Result=0
	client:=args[1].(gate.Agent)
	client.WriteMsg(&result)
}