package gate
import (
	"server/msg"
	"server/login"
)

func init() {
	msg.Processor.SetRouter(&msg.CS_PlayerLogin{},login.ChanRPC)
}
