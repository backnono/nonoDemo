package tools

// SliceContains 判断给定的Item是否存在于给定input切片中
func SliceContains[T comparable](input []T, item T) bool {
	for _, v := range input {
		if v == item {
			return true
		}
	}
	return false
}

// SliceContainsFunc 判断给定的Item是否存在于给定input切片中, 将会使用给定的function比较元素
func SliceContainsFunc[T comparable](input []T, item T, compareFunc func(a, b T) bool) bool {
	for _, v := range input {
		if compareFunc(v, item) {
			return true
		}
	}
	return false
}

// SliceIntersection 返回给定的
func SliceIntersection[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return []T{}
	}
	if len(slices) == 1 {
		return slices[0]
	}
	var result []T
	set := map[T]int{}
	for _, slice := range slices {
		for _, v := range slice {
			set[v]++
		}
	}
	for k, v := range set {
		if v == len(slices) {
			result = append(result, k)
		}
	}
	return result
}

func SliceUnion[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return []T{}
	}
	if len(slices) == 1 {
		return slices[0]
	}
	var result []T
	set := map[T]interface{}{}
	for _, slice := range slices {
		for _, v := range slice {
			set[v] = nil
		}
	}
	for k := range set {
		result = append(result, k)
	}
	return result
}

func SliceDifference[T comparable](a []T, b []T, compareFunc func(x, y T) bool) []T {
	var result []T
	if compareFunc == nil {
		for _, item := range a {
			if !SliceContains(b, item) {
				result = append(result, item)
			}
		}
	} else {
		for _, item := range a {
			if !SliceContainsFunc(b, item, compareFunc) {
				result = append(result, item)
			}
		}
	}
	return result
}

// IndexOf 在集合collection中查找满足predicate条件的第一个元素,并返回其索引
func IndexOf[T comparable](collection []T, predicate func(T) bool) int {
	for i, t := range collection {
		if predicate(t) {
			return i
		}
	}

	return -1
}

// LastIndexOf 在集合collection中从后向前查找满足predicate条件的第一个元素,并返回其索引
func LastIndexOf[T comparable](collection []T, predicate func(T) bool) int {
	l := len(collection)

	for i := l - 1; i >= 0; i-- {
		if predicate(collection[i]) {
			return i
		}
	}

	return -1
}
