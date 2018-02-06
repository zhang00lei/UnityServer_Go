package internal

import (
	"reflect"
	"server/msg"
	"github.com/name5566/leaf/gate"
	"server/login/playerdata"
	"math/rand"
	"time"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

var readyPlayerMap map[uint64]gate.Agent

func init() {
	handleMsg(&msg.CS_PlayerReady{},cs_PlayerReady)
	handleMsg(&msg.CS_PlayerCancelReady{},cs_PlayerCancelReady)
	handleMsg(&msg.CS_PlayerGrabHost{},cs_PlayerGrabHost)
	handleMsg(&msg.CS_PutCard{},cs_PutCard)
	readyPlayerMap=make(map[uint64]gate.Agent)
}

func cs_PlayerReady(args[] interface{})  {
	var resultTemp int32
	playerInfo:=args[1].(gate.Agent).UserData().(playerdata.PlayerInfo)
	result:=msg.SC_PlayerReady{}
	//加入匹配列表
	if readyPlayerMap[playerInfo.Id] ==nil {
		isMatchSuccess,matchResult:= getPlayerMatchResult()
		if !isMatchSuccess {
			readyPlayerMap[playerInfo.Id] = args[1].(gate.Agent)
		}else {
			resultTemp=0
			result.Result = &resultTemp
			args[1].(gate.Agent).WriteMsg(&result)
			matchResult = append(matchResult, args[1].(gate.Agent))
			dispatchCard(matchResult)
			return
		}
		resultTemp=0
	}else{
		resultTemp=-1
	}
	result.Result = &resultTemp
	args[1].(gate.Agent).WriteMsg(&result)
}

func getPlayerMatchResult() (bool,[]gate.Agent){
	//满足匹配条件
	if len(readyPlayerMap)>=0{
		count:=0
		matchPlayer := make([]gate.Agent, 0)
		for _,v:=range readyPlayerMap{
			count++
			matchPlayer= append(matchPlayer, v)
			if count==2{
				break
			}
		}
		return true,matchPlayer
	}
	return false,nil
}

//分发卡牌
/**
 	cardId/4==卡牌值(0:A 1:1 ------>11:Q 12:K,52小王 53大王)
	cardId%4==卡牌类型(0:红 1:方 2:黑 3:梅)
 */
func  dispatchCard(matchPlayer []gate.Agent) {
	cardSourceInfo := make([]int32, 54)
	cardInfo:=make([]int32,54)
	for i := 0; i < 54; i++ {
		cardSourceInfo[i] = int32(i);
	}
	for i:=0;i<len(cardInfo) ;i++  {
		rand.Seed(time.Now().Unix())
		rnd:=rand.Intn(len(cardSourceInfo))
		cardInfo[i] = cardSourceInfo[rnd]
		cardSourceInfo=append(cardSourceInfo[:rnd],cardSourceInfo[rnd+1:]...)
	}
	for i:=0;i<len(matchPlayer) ;i++  {
		playerCard:=msg.SC_PlayerCard{}
		playerPosition:=int32(i)
		playerCard.Position = &playerPosition
		playerIdTemp:=matchPlayer[i].(gate.Agent).UserData().(playerdata.PlayerInfo).Id
		playerCard.PlayerId = &playerIdTemp
		playerCard.CardId = cardInfo[i*13:i*12+13]
		matchPlayer[i].WriteMsg(&playerCard)
	}
}

func cs_PlayerCancelReady(args[] interface{}){
	var resultTemp int32
	playerId:=args[1].(gate.Agent).UserData().(playerdata.PlayerInfo).Id
	result:=msg.SC_PlayerCancelReady{}
	if readyPlayerMap[playerId] != nil{
		delete(readyPlayerMap,playerId)
		resultTemp=0
	}else {
		resultTemp=-1
	}
	result.Result = &resultTemp
	args[1].(gate.Agent).WriteMsg(&result)
}

//抢地主逻辑
func cs_PlayerGrabHost(args[] interface{})  {

}

//玩家出牌
func cs_PutCard(args[] interface{}){

}