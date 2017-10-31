package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"server/conf"
	"server/game"
	"server/gate"
	"server/login"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath
	lconf.MySqlUserName = conf.Server.MySqlUserName
	lconf.MySqlUserPwd = conf.Server.MySqlUserPwd
	lconf.MySqlAddr = conf.Server.MySqlAddr
	lconf.DbName = conf.Server.DbName
	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
