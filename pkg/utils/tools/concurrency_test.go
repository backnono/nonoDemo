// @Author nono.he 2023/4/19 17:28:00
package tools

import (
	"fmt"
	"sync"
	"testing"
)

func Test_Do(t *testing.T) {
	var locker sync.Mutex
	s := synchronize{locker: &locker}

	fi := func() error {
		panic("oops")
	}

	go s.Do(fi)
	go s.Do(fi)
}

func Test_Async(t *testing.T) {
	squares := Async(func() []int {
		return []int{1, 2, 3}
	})

	s := <-squares // s is []int{1, 2, 3}
	fmt.Println(s)
}

func Test_Synchronize(t *testing.T) {
	s := Synchronize()
	// 等价于
	var mutex sync.Mutex
	s1 := synchronize{locker: &mutex}
	fmt.Println(s)
	fmt.Println(s1)
}
