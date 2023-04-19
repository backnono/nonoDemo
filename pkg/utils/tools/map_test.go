// @Author nono.he 2023/4/19 16:07:00
package tools

import (
	"fmt"
	"testing"
)

func Test_Keys(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}

	fmt.Println(Keys(m))
}

func Test_Values(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}

	fmt.Println(Values(m))
}

func Test_FiltrateBy(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	fi := func(key int, value int) bool {
		if value == 3 {
			return false
		}
		return true
	}
	fmt.Println(FiltrateBy(m, fi))
}

func Test_FiltrateByKeys(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}

	fmt.Println(FiltrateByKeys(m, []int{1, 2}))
}

func Test_FiltrateByValues(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}

	fmt.Println(FiltrateByValues(m, []int{1, 2}))
}

func Test_MapToEntries(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}

	fmt.Println(MapToEntries(m))
}

func Test_EntriesToMap(t *testing.T) {
	e := []Entry[int, int]{{Key: 1, Value: 1}, {Key: 2, Value: 2}, {Key: 3, Value: 3}}

	fmt.Println(EntriesToMap(e))
}

func Test_Invert(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}

	fmt.Println(Invert(m))
}

func Test_Assign(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	n := map[int]string{4: "d", 5: "e"}
	fmt.Println(Assign(m, n))
}

// 创建一个map,其中key是通过iteratee函数从原map派生得到的。value保持不变。
func Test_MapUpdateKeys(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	fi := func(key int, value string) string {
		return value
	}
	fmt.Println(MapUpdateKeys(m, fi))
}

// 创建一个map,其中value是通过iteratee函数从原map派生得到的。key保持不变。
func Test_MapUpdateValues(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	fi := func(key int, value string) int {
		return key
	}
	fmt.Println(MapUpdateValues(m, fi))
}
