package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

//var Processor network.Processor
//var Processor = json.NewProcessor()
var Processor = protobuf.NewProcessor()
func init() {
	//Processor.Register( &Hello{})
	Processor.Register(&CS_Login{})
	Processor.Register(&SC_Login{})
}
