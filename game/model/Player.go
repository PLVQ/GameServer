package model

import (
	"gameServer/game/db"
)

type Player struct {
	RoleData db.RoleData

	cmptList []IPlayerCmpt

	ActivityOpen PlayerActivityOpen
}

func (pthis *Player) OnInit() {
}

func (pthis *Player) OnAfterRoleLogin() {
	for _, cmpt := range pthis.cmptList {
		cmpt.OnAfterRoleLogin()
	}
}

func (pthis *Player) OnTick() {
	for _, cmpt := range pthis.cmptList {
		cmpt.OnTick()
	}
}

func (pthis *Player) OnZeroTime() {
	for _, cmpt := range pthis.cmptList {
		cmpt.OnZeroTime()
	}
}

func (pthis *Player) InitPlayerCmpt() {
	pthis.cmptList = make([]IPlayerCmpt, 0)

	pthis.AddCmpt(&pthis.ActivityOpen)
}

func (pthis *Player) Init() {
	for _, cmpt := range pthis.cmptList {
		cmpt.Init(pthis)
	}
}

func (pthis *Player) AddCmpt(cmpt IPlayerCmpt) {
	pthis.cmptList = append(pthis.cmptList, cmpt)
}
