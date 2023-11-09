package wutils

type WUtilSlice struct{}

// Intersect 取两个列表的交集
func (w *WUtilSlice) Intersect(s1, s2 []string) []string {
	m := map[string]bool{}
	for _, v := range s1 {
		m[v] = true
	}
	var res []string
	for _, v := range s2 {
		if m[v] {
			res = append(res, v)
		}
	}
	return res
}

// DifferentSet 取两个列表的差集
func (w *WUtilSlice) DifferentSet(s1, s2 []string) []string {
	m := map[string]string{}
	var res []string

	for _, v := range s1 {
		if _, ok := m[v]; !ok {
			m[v] = v
		}
	}

	for _, v := range s2 {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}

	return res
}

// Merge 合集
func (w *WUtilSlice) Merge(s1, s2 []string) []string {
	var res []string
	for _, v := range s1 {
		res = append(res, v)
	}

	for _, v := range s2 {
		res = append(res, v)
	}
	return res
}

// MergeRepeatedElement 合集去重
func (w *WUtilSlice) MergeRepeatedElement(s1, s2 []string) []string {
	var res []string
	for _, v := range s1 {
		res = append(res, v)
	}

	for _, v := range s2 {
		res = append(res, v)
	}
	return w.RemoveRepeatedElement(res)
}

// RemoveRepeatedElement 数组切片去重
func (w *WUtilSlice) RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
