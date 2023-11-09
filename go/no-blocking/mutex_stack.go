package noblocking

import "sync"

type MutexStack struct {
	v  []interface{}
	mu sync.Mutex
}

func NewMutexStack() *MutexStack {
	return &MutexStack{v: make([]interface{}, 0)}
}

func (s *MutexStack) Push(v interface{}) {
	s.mu.Lock()
	s.v = append(s.v, v)
	s.mu.Unlock()
}

func (s *MutexStack) Pop() interface{} {
	s.mu.Lock()
	var v interface{}
	if len(s.v) > 0 {
		v = s.v[len(s.v)-1]
		s.v = s.v[:len(s.v)-1]
	}
	s.mu.Unlock()
	return v
}
