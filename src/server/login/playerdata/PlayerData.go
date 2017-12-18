package playerdata

import (
	"github.com/name5566/leaf"
)

type PlayerInfo struct {
	Id int64
	Account string `xorm:"unique"`
	Pwd string
	CoinCount int32
	CreatedAt int64 `xorm:"created"`
}

func (info *PlayerInfo)GetPlayerInfo() (playerInfo PlayerInfo)  {
	return *info;
}

func (info *PlayerInfo) PlayerLogin(playerAccount,playerPwd string) *PlayerInfo {
	has,_:= leaf.DataEngine.Where("Account=? and Pwd=?",playerAccount,playerPwd).Get(info)
	if !has{
		return nil
	}
	return info
}

func (info *PlayerInfo) IsContainAccount(playerAccount string) bool{
	has,_:=leaf.DataEngine.Where("Account=?",playerAccount).Get(info)
	return has
}

func (info *PlayerInfo) PlayerRegister(playerAccount,playerPwd string) bool{
	has:= info.IsContainAccount(playerAccount)
	if has{
		return false
	}
	info.Account = playerAccount
	info.Pwd = playerPwd
	info.CoinCount = 0
	leaf.DataEngine.Insert(info)
	return true
}