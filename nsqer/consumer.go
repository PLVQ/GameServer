package nsqer

import (
	"fmt"
	"gameServer/game/logic"
	"gameServer/log"
	"gameServer/msg"
	"unsafe"

	"github.com/nsqio/go-nsq"
)

type NSQConsumer struct {
	Consumer *nsq.Consumer
}

func (pThis *NSQConsumer) HandleMessage(message *nsq.Message) error {
	// 解析消息头
	SSMsgHeadLen := int(unsafe.Sizeof(msg.SSMsgHead{}))
	if len(message.Body) > SSMsgHeadLen {
		msgHeadBuf := message.Body[:SSMsgHeadLen]
		SSMsgHead := *(**msg.SSMsgHead)(unsafe.Pointer(&msgHeadBuf))

		msgHanle, ok := logic.MsgHanleMap[SSMsgHead.Cmd]
		if ok {
			msgHanle(message.Body[SSMsgHeadLen:])
		}
	}

	message.Finish()
	return nil
}

func InitConsuemr(topic string, channel string) {
	var err error
	nsqConsumer := &NSQConsumer{}
	config := nsq.NewConfig()
	if nsqConsumer.Consumer, err = nsq.NewConsumer(topic, channel, config); err != nil {
		log.Log.WithField("Error", err.Error()).Fatal()
	}

	nsqConsumer.Consumer.AddHandler(nsqConsumer)
	if err = nsqConsumer.Consumer.ConnectToNSQD("127.0.0.1:4150"); err != nil {
		fmt.Println(err)
	}

	<-nsqConsumer.Consumer.StopChan
	log.Log.Error("stop nsq consume")
}
