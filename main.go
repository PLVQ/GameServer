package main

import (
	"gameServer/game/logic"
	"gameServer/nsqer"
)

func main() {
	logic.Init()
	nsqer.InitConsuemr("client", "test")
}
