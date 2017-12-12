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
}
