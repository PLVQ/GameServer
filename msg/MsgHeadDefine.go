package msg

type SSMsgHead struct {
	MsgLen  uint32 // 消息长度
	Echo    uint32 // 用于校验包的对应关系，递增
	Cmd     uint32 // 消息ID
	AgentID uint32 // 客户端连接的代理ID
}

type MsgHeadByte struct {
	addr uintptr
	len  int
	cap  int
}
