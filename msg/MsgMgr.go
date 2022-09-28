package msg

import (
	"fmt"
	"gameServer/nsqer"
	"gameServer/sessionmgr"

	"google.golang.org/protobuf/proto"
)

func Send2Client(iUin uint64, iCmd uint32, pMsg proto.Message) {
	session := sessionmgr.GetSession(iUin)
	if session == nil {
		fmt.Println("sssssss")
		return
	}
	msgBuf := PackSSMsg(iCmd, session.Agent, session.Echo, pMsg)
	fmt.Println(len(msgBuf))
	nsqer.NsqProducer.Publish("game", msgBuf)
}
