package playerdata

import (
	"github.com/name5566/leaf"
)

type PlayerInfo struct {
	Id int64
	Account string `xorm:"unique"`
	Pwd string
}

func (info *PlayerInfo)GetPlayerInfo() (playerInfo PlayerInfo)  {
	return *info;
}


func (info *PlayerInfo) PlayerLogin(playerAccount,playerPwd string) *PlayerInfo {
	has,_:= leaf.DataEngine.Where("Account=? and Pwd=?",playerAccount,playerPwd).Get(&info)
	if !has{
		return nil
	}
	return info
}

func (info *PlayerInfo) IsContainAccount(playerAccount string) bool{
	has,_:=leaf.DataEngine.Where("Account=?",playerAccount).Get(&info)
	return has
}

func (info *PlayerInfo) PlayerRegister(playerAccount,playerPwd string) *PlayerInfo{
	has:= info.IsContainAccount(playerAccount)
	if has{
		return nil
	}
	info.Account = playerAccount
	info.Pwd = playerPwd
	leaf.DataEngine.Insert(info)
	return info
}