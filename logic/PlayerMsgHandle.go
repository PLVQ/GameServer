package logic

import (
	"gameServer/log"
	"gameServer/mgr"
	"gameServer/msg"
	"gameServer/msg/pbproto"
	"gameServer/sessionmgr"

	"google.golang.org/protobuf/proto"
)

func HandlePlayerActLogin(msgBuf []byte, msgHead *msg.SSMsgHead) {
	req := &pbproto.CSPlayerActLogin{}
	if err := proto.Unmarshal(msgBuf, req); err != nil {
		log.Log.WithField("Error", err.Error()).Error()
	}

	log.Log.WithField("data", req).Debug()

	player := mgr.GetPlayer(req.Uin)
	if player == nil {
		player = mgr.CreatePlayer(req.Uin)
	}

	sessionmgr.CreateSession(req.Uin, msgHead.AgentID)

	msg.Send2Client(req.Uin, 1, req)

	// mgr.AddPlayerAgent(req.RoleID, req.)
}

func init() {
	RegShakeHandle(1, HandlePlayerActLogin)
}
