package model

import (
	"gameServer/comm"
	"gameServer/res"
	"gameServer/res/respb"
	"math"
	"time"

	"gameServer/log"

	"github.com/sirupsen/logrus"
)

const (
	ACTIVITY_OPEN_TYPE_FIXED           = 1 /* 固定开启，类似七日签到 */
	ACTIVITY_OPEN_TYPE_TIME            = 2 /* 指定时间开启 */
	ACTIVITY_OPEN_TYPE_SVR_OPENTIME    = 3 /* 开服时间 */
	ACTIVITY_OPEN_TYPE_SVR_MERGETIME   = 4 /* 合服时间 */
	ACTIVITY_OPEN_TYPE_FIX_DAILY_TIME  = 5 /* 每日固定时间 */
	ACTIVITY_OPEN_TYPE_SVR_WEEK_DAY    = 6 /* 每周固定时间 */
	ACTIVITY_OPEN_TYPE_CREATE_ROLE_DAY = 7 /* 创角天数 */
)

type PlayerActivityOpen struct {
	PlayerCmpt
	LastTick int64
}

func (pthis *PlayerActivityOpen) OnInit() {
	pthis.LastTick = 0
}

func (pthis *PlayerActivityOpen) OnAfterRoleLogin() {
	pthis.NotifyActivityOpenConfig()
}

func (pthis *PlayerActivityOpen) OnTick() {
	if pthis.IsDirty() {
		pthis.NotifyActivityOpenConfig()
	}
	if math.Abs(float64(time.Now().Unix()-pthis.LastTick)) > 5 {
		pthis.NotifyActivityOpenConfig()
		pthis.LastTick = time.Now().Unix()
	}
}

func (pthis *PlayerActivityOpen) IsActivityOpenOrDelay(pCfg *respb.ActivityOpenConfig) bool {
	if pCfg.OpenSvrTimeLimit > 0 {
		if time.Now().Unix() < pCfg.OpenSvrTimeLimit {
			return false
		}
	}

	roleCreateTime := pthis.pPlayer.RoleData.CreateTime
	creatRoleDays := comm.GetDiffDay(time.Now().Unix(), int64(roleCreateTime))
	if creatRoleDays > 0 && creatRoleDays <= uint(pCfg.CreateRoleDays) {
		return false
	}

	tNow := time.Now().Unix()
	bIsOpen := false
	switch pCfg.OpenType {
	case ACTIVITY_OPEN_TYPE_FIXED:
		if pCfg.OpenParam[0] > 0 {
			bIsOpen = true
		}
	case ACTIVITY_OPEN_TYPE_TIME:
		if tNow >= pCfg.OpenParam[0] && pCfg.OpenParam[1] >= tNow {
			bIsOpen = true
		}
	case ACTIVITY_OPEN_TYPE_SVR_OPENTIME:
		if pCfg.OpenParam[0] > 0 {
			bIsOpen = true
		}
	case ACTIVITY_OPEN_TYPE_SVR_MERGETIME:
		if pCfg.OpenParam[0] > 0 {
			bIsOpen = true
		}
	case ACTIVITY_OPEN_TYPE_SVR_WEEK_DAY:
		if time.Unix(tNow, 0).Weekday() == time.Weekday(pCfg.OpenParam[0]) {
			bIsOpen = true
		}
	case ACTIVITY_OPEN_TYPE_CREATE_ROLE_DAY:
		if creatRoleDays >= uint(pCfg.OpenParam[0]) && creatRoleDays <= uint(pCfg.OpenParam[1]) {
			bIsOpen = true
		}
	default:
	}
	return bIsOpen
}

func (pthis *PlayerActivityOpen) NotifyActivityOpenConfig() {
	for _, pCfg := range res.ActivityOpenConfigListData.Data {
		if pthis.IsActivityOpenOrDelay(pCfg) {
			log.Log.WithFields(logrus.Fields{"ActivityID": pCfg.ActivityID, "ActivityName": pCfg.ActivityName}).Info("Open Activitys")
		}
	}
}
