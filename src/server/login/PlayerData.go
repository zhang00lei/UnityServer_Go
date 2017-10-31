package login

type PlayerInfo struct {
	PlayerId string
	PlayerName string
}

func (info *PlayerInfo)GetPlayerInfo() (playerInfo PlayerInfo)  {
	return *info;
}

func(info *PlayerInfo)SetPlayerInfo(playerId,playerName string)(playerinfo PlayerInfo){
	info.PlayerId = playerId
	info.PlayerName = playerName
	return info.GetPlayerInfo()
}