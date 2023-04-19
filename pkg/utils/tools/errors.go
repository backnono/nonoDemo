package tools

//Try calls the function and return false in case of error.
//橱柜函数,可以让我们不用担心直接调用可能发生panic的函数。这给很多算法的实现带来了简单性。
func Try(callback func() error) (ok bool) {
	ok = true

	defer func() {
		if err := recover(); err != nil {
			ok = false
			return
		}
	}()

	err := callback()
	if err != nil {
		ok = false
	}

	return
}
