package mgr

var PlayerPool map[uint64]struct{}

func init() {
	PlayerPool = make(map[uint64]struct{})
}

func CreatePlayer(RoleID uint64) {
	PlayerPool[RoleID] = struct{}{}
}
