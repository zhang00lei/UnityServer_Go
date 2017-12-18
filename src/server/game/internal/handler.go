package internal

import (
	"reflect"
	"server/msg"
	"github.com/name5566/leaf/gate"
	"time"
)

func init()  {
	handler(&msg.CS_Heartbeat{},cs_Heartbeat)
}

func handler(m interface{},h interface{})  {
	skeleton.RegisterChanRPC(reflect.TypeOf(m),h)
}

func cs_Heartbeat(args[] interface{}){
	result := msg.SC_Heartbeat{}
	client := args[1].(gate.Agent)
	result.ServerTime = time.Time{}.Unix()
	client.WriteMsg(&result)
}