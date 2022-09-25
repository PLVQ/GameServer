package model

type IPlayerCmpt interface {
	OnInit()
	OnAfterRoleLogin()
	OnTick()
	OnZeroTime()
	Init(pPlayer *Player)
}

type PlayerCmpt struct {
	PlayerDirty
	pPlayer *Player
}

func (pthis *PlayerCmpt) OnInit() {

}

func (pthis *PlayerCmpt) OnResume() {

}

func (pthis *PlayerCmpt) OnTick() {

}

func (pthis *PlayerCmpt) OnInitCreateRole() {

}

func (pthis *PlayerCmpt) OnPreRoleLogin() {

}

func (pthis *PlayerCmpt) OnAfterRoleLogin() {

}

func (pthis *PlayerCmpt) OnZeroTime() {

}

func (pthis *PlayerCmpt) OnDailyReset() {

}

func (pthis *PlayerCmpt) OnWeekReset() {

}

func (pthis *PlayerCmpt) OnRoleLogout() {

}

func (pthis *PlayerCmpt) OnRoleReconnect() {

}

func (pthis *PlayerCmpt) OnRoleLevelUp() {

}

func (pthis *PlayerCmpt) Init(pPlayer *Player) {
	pthis.pPlayer = pPlayer
	pthis.OnInit()
}

func (pthis *PlayerCmpt) Resume(pPlaye *Player) {
	pthis.pPlayer = pPlaye
	pthis.OnResume()
}

// func (pthis *PlayerCmpt) Tick() {
// 	if (pthis.IsDirty()) {
// 		boo
// 	}
// }
