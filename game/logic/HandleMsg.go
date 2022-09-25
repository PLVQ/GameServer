package logic

import (
	"gameServer/log"
	"gameServer/msg"

	"google.golang.org/protobuf/proto"
)

var MsgHanleMap map[uint32]func([]byte)

func Init() {
	MsgHanleMap = make(map[uint32]func([]byte))
	MsgHanleMap[1] = HandleTestMsg
}

func HandleTestMsg(msgBuf []byte) {
	data := &msg.Hello{}
	if err := proto.Unmarshal(msgBuf, data); err != nil {
		log.Log.WithField("Error", err.Error()).Error()
	}

	log.Log.WithField("data", data).Debug()
}
