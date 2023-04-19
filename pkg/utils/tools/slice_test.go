package tools

import (
	"fmt"
	"testing"
)

func Test_SliceContainsInt(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	if !SliceContains[int](input, 2) {
		t.Fail()
	}
	if SliceContains[int](input, -1) {
		t.Fail()
	}
}

func Test_SliceContainsFloat(t *testing.T) {
	input := []float32{1, 2, 3, 4, 5, 6}
	if !SliceContains[float32](input, 2) {
		t.Fail()
	}
	if SliceContains[float32](input, -1) {
		t.Fail()
	}
}

func Test_SliceContainsStrings(t *testing.T) {
	input := []string{"a", "b", "c"}
	if !SliceContains[string](input, "a") {
		t.Fail()
	}
	if SliceContains[string](input, "x") {
		t.Fail()
	}
}

func Test_SliceContainsStruct(t *testing.T) {
	input := []struct{ Name string }{
		{Name: "a"},
		{Name: "b"},
	}
	if !SliceContains[struct{ Name string }](input, struct{ Name string }{Name: "a"}) {
		t.Fail()
	}
	if SliceContains[struct{ Name string }](input, struct{ Name string }{Name: GetLocalIP()}) {
		t.Fail()
	}
}

func Test_SliceIntersection(t *testing.T) {
	result := SliceIntersection([]int{1, 2, 3, 4}, []int{2, 3, 4}, []int{3})
	fmt.Printf("%v\n", result)
	if len(result) != 1 || result[0] != 3 {
		t.Fail()
	}
}

func Test_SliceUnion(t *testing.T) {
	result := SliceUnion([]int{1}, []int{2}, []int{3})
	fmt.Printf("%v\n", result)
	if len(result) != 3 {
		t.Fail()
	}
}

func Test_SliceDifference(t *testing.T) {
	result := SliceDifference([]int{1, 2, 3}, []int{2, 3}, nil)
	fmt.Printf("%v\n", result)
	if len(result) != 1 {
		t.Fail()
	}

	type x struct {
		Name string
	}
	result1 := SliceDifference([]x{
		{Name: "a"},
		{Name: "b"},
	}, []x{
		{Name: "a"},
	}, func(a, b x) bool {
		return a.Name == b.Name
	})
	fmt.Printf("%v\n", result1)
	if len(result1) != 1 {
		t.Fail()
	}
}

func Test_IndexOf(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	fi := func(value int) bool {
		return value > 3
	}
	fmt.Println(IndexOf(input, fi))
}

func Test_LastIndexOf(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	fi := func(value int) bool {
		return value > 3
	}
	fmt.Println(LastIndexOf(input, fi))
}
