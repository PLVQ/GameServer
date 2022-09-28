package sessionmgr

import "gameServer/comm"

type Session struct {
	Uin    uint64
	Agent  int64
	Echo   uint32
	Status int32
}

var sessionMap map[uint64]*Session

func init() {
	sessionMap = map[uint64]*Session{}
}

func CreateSession(iUin uint64, iAgent int64) *Session {
	session := &Session{
		Uin:    iUin,
		Agent:  iAgent,
		Status: comm.STATUS_HANDSHAKE,
	}

	sessionMap[iUin] = session

	return session
}

func GetSession(iUin uint64) *Session {
	session, ok := sessionMap[iUin]
	if !ok {
		return nil
	}

	return session
}

func SetSessionWorking(iUin uint64) {
	session, ok := sessionMap[iUin]
	if ok {
		session.Status = comm.STATUS_WORKING
	}
}

// func updateSession(iUin uint64, iAgent int64, iEcho uint32) {
// 	session, ok := sessionMap[iUin]
// 	if ok {
// 		session.agent = iAgent
// 		session.echo = iEcho
// 	}
// }
