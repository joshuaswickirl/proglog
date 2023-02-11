package server

import (
	"fmt"
	"sync"
)

type log struct {
	mu      sync.Mutex
	records []record
}

type record struct {
	value  []byte
	offset uint64
}

func newLog() *log {
	return &log{}
}

func (c *log) append(record record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.offset, nil
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")

func (c *log) read(offset uint64) (record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}
