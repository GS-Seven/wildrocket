package wutils

type WUtilSlice struct{}

// StringsInter 取两个字符串列表的交集
func (w *WUtilSlice) StringsInter(s1, s2 []string) []string {
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
