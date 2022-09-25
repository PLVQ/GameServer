package model

type PlayerDirty struct {
	m_bDirty bool
}

func (pthis *PlayerDirty) IsDirty() bool {
	return pthis.m_bDirty
}

func (pthis *PlayerDirty) MarkDirty() {
	pthis.m_bDirty = true
}

func (pthis *PlayerDirty) ClearDirty() {
	pthis.m_bDirty = false
}

func (pthis *PlayerDirty) OnDirty() bool {
	return true
}
