// @Author nono.he 2023/4/19 20:22:00
package tools

import (
	"fmt"
	"testing"
)

func Test_Ternary(t *testing.T) {
	s := Ternary(false, 1, 2)
	fmt.Println(s)
}

func Test_If(t *testing.T) {
	s := If(true, 2)
	fmt.Println(s)
}
