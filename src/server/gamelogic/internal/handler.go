package internal

import (
	"reflect"
	"server/msg"
	"github.com/name5566/leaf/gate"
	"server/login/playerdata"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

var readyPlayerMap map[uint64]gate.Agent

func init() {
	handleMsg(&msg.CS_PlayerReady{},cs_PlayerReady)
	handleMsg(&msg.CS_PlayerCancelReady{},cs_PlayerCancelReady)
	readyPlayerMap=make(map[uint64]gate.Agent)
}

func cs_PlayerReady(args[] interface{})  {
	playerInfo:=args[1].(gate.Agent).UserData().(playerdata.PlayerInfo)
	result:=msg.SC_PlayerReady{}
	//加入匹配列表
	if readyPlayerMap[playerInfo.Id] ==nil {
		readyPlayerMap[playerInfo.Id] = args[1].(gate.Agent)
		result.Result=0
	}else{
		result.Result=-1
	}
	args[1].(gate.Agent).WriteMsg(&result)
	playerMatch()
}

func playerMatch(){
	if len(readyPlayerMap)>=3{

	}
}

func cs_PlayerCancelReady(args[] interface{}){
	playerId:=args[1].(gate.Agent).UserData().(playerdata.PlayerInfo).Id
	result:=msg.SC_PlayerCancelReady{}
	if readyPlayerMap[playerId] != nil{
		delete(readyPlayerMap,playerId)
		result.Result = 0
	}else {
		result.Result = -1
	}
	args[1].(gate.Agent).WriteMsg(&result)
}