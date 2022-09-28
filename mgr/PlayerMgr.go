package mgr

import "gameServer/model"

var playerPool map[uint64]*model.Player

func init() {
	playerPool = make(map[uint64]*model.Player)
}

func GetPlayer(uin uint64) *model.Player {
	player, ok := playerPool[uin]
	if !ok {
		return nil
	}

	return player
}

func CreatePlayer(uin uint64) *model.Player {
	player := &model.Player{}
	playerPool[uin] = player

	return player
}
