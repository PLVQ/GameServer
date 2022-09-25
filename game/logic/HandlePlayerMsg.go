package logic

import (
	"gameServer/log"
	"gameServer/msg"

	"google.golang.org/protobuf/proto"
)

func HandlePlayerLogin(msgBuf []byte) {
	req := &msg.CSPlayerActLogin{}
	if err := proto.Unmarshal(msgBuf, req); err != nil {
		log.Log.WithField("Error", err.Error()).Error()
	}

	log.Log.WithField("data", req).Debug()
	// mgr.AddPlayerAgent(req.RoleID, req.)
}
