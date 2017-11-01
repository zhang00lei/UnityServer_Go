package gate
import (
	"server/msg"
	"server/login"
)

func init() {
	msg.Processor.SetRouter(&msg.CS_PlayerLogin{},login.ChanRPC)
	msg.Processor.SetRouter(&msg.CS_PlayerRegister{},login.ChanRPC)
}
