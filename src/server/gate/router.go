package gate
import (
	"server/msg"
	"server/login"
	"server/game"
	"server/gamelogic"
)

func init() {
	msg.Processor.SetRouter(&msg.CS_PlayerLogin{},login.ChanRPC)
	msg.Processor.SetRouter(&msg.CS_PlayerRegister{},login.ChanRPC)

	msg.Processor.SetRouter(&msg.CS_Heartbeat{},game.ChanRPC)
	msg.Processor.SetRouter(&msg.CS_PlayerInfo{},game.ChanRPC)

	msg.Processor.SetRouter(&msg.CS_PlayerReady{},gamelogic.ChanRPC)
	msg.Processor.SetRouter(&msg.CS_PlayerCancelReady{},gamelogic.ChanRPC)
}
