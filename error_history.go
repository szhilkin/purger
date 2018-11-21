package main

import (
	"sync"
	"time"
)

type errorHistoryItem struct {
	Time time.Time
	Msg  string
}

type errorHistoryStore struct {
	sync.RWMutex
	data     []errorHistoryItem
	MaxCount int
}

func (e *errorHistoryStore) Add(s string) {
	e.Lock()
	defer e.Unlock()
	e.data = append(e.data, errorHistoryItem{time.Now(), s})
	if len(e.data) > e.MaxCount {
		e.data = e.data[1:]
	}
}

func (e *errorHistoryStore) Get() []errorHistoryItem {
	e.RLock()
	e.RUnlock()
	return e.data
}

func newErrorHistoryStore(MaxCount int) *errorHistoryStore {
	return &errorHistoryStore{MaxCount: MaxCount}
}
