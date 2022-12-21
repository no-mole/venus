package codec

import (
	"github.com/no-mole/venus/agent/venus/structs"
	"google.golang.org/protobuf/proto"
)

var encoder = proto.MarshalOptions{}
var decoder = proto.UnmarshalOptions{}

func SetEncoder(newEncoder proto.MarshalOptions) {
	encoder = newEncoder
}

func SetDecoder(newDecoder proto.UnmarshalOptions) {
	decoder = newDecoder
}

func Decode(buf []byte, msg proto.Message) error {
	return decoder.Unmarshal(buf, msg)
}

func Encode(msgType structs.MessageType, msg proto.Message) (buf []byte, err error) {
	buf = []byte{byte(msgType)}
	buf, err = encoder.MarshalAppend(buf, msg)
	return
}
