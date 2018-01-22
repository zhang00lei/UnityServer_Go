package gate
import (
	"server/msg"
	"server/login"
	"server/game"
)

func init() {
	msg.Processor.SetRouter(&msg.CS_PlayerLogin{},login.ChanRPC)
	msg.Processor.SetRouter(&msg.CS_PlayerRegister{},login.ChanRPC)
	msg.Processor.SetRouter(&msg.CS_Heartbeat{},game.ChanRPC)
	msg.Processor.SetRouter(&msg.CS_PlayerInfo{},game.ChanRPC)
}
