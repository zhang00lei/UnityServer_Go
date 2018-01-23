package msg
import (
	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()
func init() {
	Processor.Register(&CS_PlayerLogin{})
	Processor.Register(&SC_PlayerLogin{})

	Processor.Register(&CS_PlayerRegister{})
	Processor.Register(&SC_PlayerRegister{})

	Processor.Register(&CS_Heartbeat{})
	Processor.Register(&SC_Heartbeat{})

	Processor.Register(&CS_PlayerInfo{})
	Processor.Register(&SC_PlayerInfo{})

	Processor.Register(&CS_PlayerReady{})
	Processor.Register(&SC_PlayerReady{})

	Processor.Register(&CS_PlayerCancelReady{})
	Processor.Register(&SC_PlayerCancelReady{})

	Processor.Register(&SC_PlayerCard{})
}
