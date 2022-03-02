package main

import (
	"sync/atomic"
	"testing"
)

func TestAtomicValue(t *testing.T) {
	vInt := atomic.Value{}
	vInt.Store(1)
	vInt.Store(2)
	t.Log(vInt.Load().(int))
	t.Log(vInt.Load().(int))

	v1 := atomic.Value{}
	v1.Store(struct{}{})
	v1.Store(struct{}{})
	t.Log(v1.Load())
	t.Log(v1.Load())

	v2 := atomic.Value{}
	ch2, err := v2.Load().(chan struct{})
	t.Log(ch2, err)
	v2.Store(make(chan struct{}))
	ch2, err = v2.Load().(chan struct{})
	t.Log(ch2, err)
	if ch2 != nil {
		t.Log("ch2 not nil")
		close(ch2)
	}
	v2.Store((chan struct{})(nil))
	ch2, err = v2.Load().(chan struct{})
	t.Log(ch2, err)
	if ch2 != nil {
		t.Log("ch2 not nil--")
		close(ch2)
	}
}
