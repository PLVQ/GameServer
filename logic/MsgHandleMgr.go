package logic

import (
	"gameServer/log"
	"gameServer/msg"

	"google.golang.org/protobuf/proto"
)

var shakeHanleMap map[uint32]func([]byte, *msg.SSMsgHead)
var msgHanleMap map[uint32]func([]byte)

func init() {
	shakeHanleMap = make(map[uint32]func([]byte, *msg.SSMsgHead))
	msgHanleMap = make(map[uint32]func([]byte))
}

func RegMsgHandle(iCmd uint32, fHandle func([]byte)) {
	msgHanleMap[iCmd] = fHandle
}

func RegShakeHandle(iCmd uint32, fHandle func([]byte, *msg.SSMsgHead)) {
	shakeHanleMap[iCmd] = fHandle
}

func GetMsgHand(iCmd uint32) func([]byte) {
	fHandle, ok := msgHanleMap[iCmd]
	if !ok {
		return nil
	}

	return fHandle
}

func GetShakeHand(iCmd uint32) func([]byte, *msg.SSMsgHead) {
	fHandle, ok := shakeHanleMap[iCmd]
	if !ok {
		return nil
	}

	return fHandle
}

func HandleTestMsg(msgBuf []byte) {
	data := &msg.Hello{}
	if err := proto.Unmarshal(msgBuf, data); err != nil {
		log.Log.WithField("Error", err.Error()).Error()
	}

	log.Log.WithField("data", data).Debug()
	// cnt++
	// log.Log.WithField("cnt", cnt).Debug()
}
