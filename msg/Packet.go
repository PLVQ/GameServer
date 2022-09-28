package msg

import (
	"unsafe"

	"google.golang.org/protobuf/proto"
)

func UnPackSSMsgHead(headBuf []byte) *SSMsgHead {
	head := *(**SSMsgHead)(unsafe.Pointer(&headBuf))
	return head
}

func PackSSMsg(iCmd uint32, iAgent int64, iEcho uint32, pMsg proto.Message) []byte {
	msgBodyBuf, err := proto.Marshal(pMsg)
	if err != nil {
		return nil
	}
	msgLen := SS_MSG_HEAD_LEN + len(msgBodyBuf)
	msgBuf := make([]byte, msgLen)

	SSMsgHead := &SSMsgHead{
		MsgLen:  uint32(len(msgBodyBuf)),
		Echo:    iEcho,
		Cmd:     iCmd,
		AgentID: iAgent,
	}

	msgHeadByte := &MsgHeadByte{
		addr: uintptr(unsafe.Pointer(SSMsgHead)),
		cap:  SS_MSG_HEAD_LEN,
		len:  SS_MSG_HEAD_LEN,
	}

	msgHeadBuf := *(*[]byte)(unsafe.Pointer(msgHeadByte))

	copy(msgBuf[:SS_MSG_HEAD_LEN], msgHeadBuf)
	copy(msgBuf[SS_MSG_HEAD_LEN:], msgBodyBuf)

	return msgBuf
}
