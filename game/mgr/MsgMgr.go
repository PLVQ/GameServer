package mgr

var PlayerAgentPool map[uint64]int

func init() {
	PlayerAgentPool = make(map[uint64]int)
}

func AddPlayerAgent(RoleID uint64, AgentID int) {
	PlayerAgentPool[RoleID] = AgentID
}
