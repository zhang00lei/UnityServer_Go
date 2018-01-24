package internal

import (
	"reflect"
	"server/msg"
	"github.com/name5566/leaf/log"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.CS_PlayerReady{},cs_PlayerReady)
}

func cs_PlayerReady(args[] interface{})  {
	log.Debug("playerReady")
}
