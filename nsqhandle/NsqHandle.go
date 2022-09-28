package nsqhandle

import (
	"gameServer/comm"
	"gameServer/logic"
	"gameServer/msg"

	"github.com/nsqio/go-nsq"
)

type ClientMsgConsumer struct {
	status int32
}

func NewClientMsgConsumer() *ClientMsgConsumer {
	return &ClientMsgConsumer{
		status: 1,
	}
}

func (pThis *ClientMsgConsumer) HandleMessage(message *nsq.Message) error {
	// 解析消息头
	if len(message.Body) > msg.SS_MSG_HEAD_LEN {
		msgHeadBuf := message.Body[:msg.SS_MSG_HEAD_LEN]
		msgHead := msg.UnPackSSMsgHead(msgHeadBuf)
		if pThis.status == comm.STATUS_START {
			handle := logic.GetShakeHand(msgHead.Cmd)
			if handle != nil {
				handle(message.Body[msg.SS_MSG_HEAD_LEN:], msgHead)
				pThis.status = comm.STATUS_HANDSHAKE
			}
		} else if pThis.status == comm.STATUS_HANDSHAKE {
			handle := logic.GetShakeHand(msgHead.Cmd)
			if handle != nil {
				handle(message.Body[msg.SS_MSG_HEAD_LEN:], msgHead)
				pThis.status = comm.STATUS_WORKING
			}
		} else if pThis.status == comm.STATUS_WORKING {
			handle := logic.GetMsgHand(msgHead.Cmd)
			if handle != nil {
				handle(message.Body[msg.SS_MSG_HEAD_LEN:])
			}
		}
	}

	message.Finish()
	return nil
}
