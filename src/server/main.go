package main

import ( 
	
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"server/conf"
	"server/game"
	"server/gate"
	"server/login"
	"server/gamelogic"
	"server/login/playerdata"
)
func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath
	lconf.DB_DriverName = conf.Server.DB_DriverName
	lconf.DB_DataSourceName = conf.Server.DB_DataSourceName

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
		gamelogic.Module,
	)
	leaf.DataEngine.Sync2(new(playerdata.PlayerInfo))
}
