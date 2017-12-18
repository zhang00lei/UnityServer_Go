package internal

import (
	"reflect"
	"server/login/playerdata"
	"github.com/name5566/leaf/gate"
	"server/msg"
	"github.com/name5566/leaf"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
func init() {
	handleMsg(&msg.CS_PlayerLogin{},cs_Login)
	handleMsg(&msg.CS_PlayerRegister{},cs_Register)
}

func cs_Login(args[] interface{}) {
	loginInfo := args[0].(*msg.CS_PlayerLogin)
	playerInfo := new(playerdata.PlayerInfo).PlayerLogin(loginInfo.Account, loginInfo.Pwd)
	result := msg.SC_PlayerLogin{}
	if playerInfo != nil {
		args[1].(gate.Agent).SetUserData(playerInfo)
		result.LoginResult = msg.SC_PlayerLogin_SUCCESS
	} else {
		has := new(playerdata.PlayerInfo).IsContainAccount(loginInfo.Account)
		if !has {
			result.LoginResult = msg.SC_PlayerLogin_ACCOUNT_ERROR
		} else {
			result.LoginResult = msg.SC_PlayerLogin_PWD_ERROR
		}
	}
	client := args[1].(gate.Agent)
	client.WriteMsg(&result)
}

func cs_Register(args[] interface{}){

	leaf.DataEngine.Sync2(new(playerdata.PlayerInfo))
	regsterInfo:=args[0].(*msg.CS_PlayerRegister)
	isInsert:=new(playerdata.PlayerInfo).PlayerRegister(regsterInfo.Account,regsterInfo.Pwd)
	result:=msg.SC_PlayerRegister{}

	if isInsert{
		result.RegisterResult = msg.SC_PlayerRegister_SUCCESS
	}else {
		result.RegisterResult = msg.SC_PlayerRegister_ACCOUNT_ERROR
	}
	client:=args[1].(gate.Agent)
	client.WriteMsg(&result)
}