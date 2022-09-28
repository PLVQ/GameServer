package main

import (
	"gameServer/nsqer"
	"gameServer/nsqhandle"
)

func main() {
	nsqHandle := nsqhandle.NewClientMsgConsumer()
	nsqer.InitConsuemr("client", "test", nsqHandle)
}
