package gtools

import "sync"

// IDGen 自增id
type IDGen struct {
	id uint64
	sync.Mutex
}

// NewIDGen 返回一个新的idGen
func NewIDGen() *IDGen {
	return &IDGen{
		id: 0,
	}
}

// Next 获取下一个id
func (id *IDGen) Next() uint64 {
	id.Lock()
	defer id.Unlock()
	id.id = id.id + 1
	return id.id
}
