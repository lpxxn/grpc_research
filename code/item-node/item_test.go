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

/*

原子操作由底层硬件支持，而锁则由操作系统的调度器实现。锁应当用来保护一段逻辑，对于一个变量更新的保护，原子操作通常会更有效率，并且更能利用计算机多核的优势，如果要更新的是一个复合对象，则应当使用atomic.Value封装好的实现

在 Go（甚至是大部分语言）中，一条普通的赋值语句其实不是一个原子操作。例如，在32位机器上写int64类型的变量就会有中间状态，因为它会被拆成两次写操作（MOV）——写低 32 位和写高 32 位，如下图所示：
64位变量的赋值操作
如果一个线程刚写完低32位，还没来得及写高32位时，另一个线程读取了这个变量，那它得到的就是一个毫无逻辑的中间变量，这很有可能使我们的程序出现诡异的 Bug。
这还只是一个基础类型，如果我们对一个结构体进行赋值，那它出现并发问题的概率就更高了。很可能写线程刚写完一小半的字段，读线程就来读取这个变量，那么就只能读到仅修改了一部分的值。这显然破坏了变量的完整性，读出来的值也是完全错误的。
面对这种多线程下变量的读写问题，我们的主角——atomic.Value登场了，它使得我们可以不依赖于不保证兼容性的unsafe.Pointer类型，同时又能将任意数据类型的读写操作封装成原子性操作（让中间状态对外不可见）。
*/
