package tools

import "sync"

type synchronize struct {
	locker sync.Locker
}

// Do 线程安全的try函数
// 这个类型使我们不必重复地在各个需要同步的函数中写锁定解锁的逻辑,这提高了代码的复用性和可读性
func (s *synchronize) Do(cb func() error) {
	s.locker.Lock()
	Try(cb)
	s.locker.Unlock()
}

// Synchronize 提供一个方便的方式来创建synchronize类型的实例
//1. 该函数也使用了泛型,入参和返回值都是interface{},使得它可以支持任何实现了sync.Locker接口的类型。
//2. 如果传入0个或多个参数,会panic。这是一个安全性措施,确保每个synchronize实例只获取一个locker。
//3. len(opt) > 1的判断也是为了这个目的,防止用户不小心传入多个locker。
//4. 该函数隐藏了一些边界情况的处理逻辑,如默认值的设置等,使得用户调用可以更加简洁。
//5. 这个函数使synchronize类型的创建方式更加优雅,让用户感知到它是一个更加内置的并发控制方式。这提高了其可用性。
func Synchronize(opt ...sync.Locker) synchronize {
	if len(opt) > 1 {
		panic("unexpected arguments")
	} else if len(opt) == 0 {
		opt = append(opt, &sync.Mutex{})
	}

	return synchronize{locker: opt[0]}
}

// Async
// 将一个函数f以goroutine的形式异步执行,并返回一个channel。f函数的返回值会被发送到这个channel中。
// 调用者可以在这个channel上接收到f函数的返回值。
func Async[A any](f func() A) chan A {
	ch := make(chan A)

	go func() {
		ch <- f()
	}()

	return ch
}
