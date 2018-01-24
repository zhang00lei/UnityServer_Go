package internal

import (
	"reflect"
	"server/msg"
	"github.com/name5566/leaf/gate"
	"time"
	"server/login/playerdata"
)

func init()  {
	handler(&msg.CS_Heartbeat{},cs_Heartbeat)
	handler(&msg.CS_PlayerInfo{},cs_PlayerInfo)
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

func cs_PlayerInfo(args[] interface{}) {
	info:=args[0].(*msg.CS_PlayerInfo)
	result:=msg.SC_PlayerInfo{}
	playerInfo, has := playerdata.GetInfoById(info.PlayerId)
	if has {
		userInfo := msg.PlayerInfo{}
		userInfo.PlayerAccount = playerInfo.Account
		userInfo.PlayerId = playerInfo.Id
		userInfo.PlayerCoin = playerInfo.CoinCount
		userInfo.PlayerName = playerInfo.PlayerNickName
		result.Result = 0
		result.PlayerInfo = &userInfo
	}else {
		result.Result = -1
		result.PlayerInfo = nil
	}
	client:=args[1].(gate.Agent)
	client.WriteMsg(&result)
}

func cs_GetCard(args[] interface{}){

}