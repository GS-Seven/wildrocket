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

// InSlice 检查字符串是否存在字符串切片中
func (w *WUtilSlice) InSlice(s string, ts []string) bool {
	for _, v := range ts {
		if v == s {
			return true
		}
	}
	return false
}

// InSliceInterface 检查Interface是否存在Interface切片中
func (w *WUtilSlice) InSliceInterface(s interface{}, ts []interface{}) bool {
	for _, v := range ts {
		if v == s {
			return true
		}
	}
	return false
}

func (w *WUtilSlice) RandSlice(min, max int) []int {
	if max < min {
		min, max = max, min
	}
	l := max - min + 1
	rand.Seed(int64(time.Now().Nanosecond()))
	ls := rand.Perm(l)
	for i, _ := range ls {
		ls[i] += min
	}
	return ls
}

// MergeSlice merge two interface slices to one slice
func (w *WUtilSlice) MergeSlice(s1, s2 []interface{}) (s3 []interface{}) {
	s3 = append(s1, s2...)
	return
}

type reduceType func(interface{}) interface{}
type filterType func(interface{}) bool

func (w *WUtilSlice) SliceReduce(s []interface{}, a reduceType) (s2 []interface{}) {
	for _, v := range s {
		s2 = append(s2, a(v))
	}
	return
}

func (w *WUtilSlice) SliceRand(a []interface{}) (b interface{}) {
	randNum := rand.Intn(len(a))
	b = a[randNum]
	return
}

func (w *WUtilSlice) SliceSum(s []int64) (sum int64) {
	for _, v := range s {
		sum += v
	}
	return
}

func (w *WUtilSlice) SliceFilter(s []interface{}, a filterType) (ftSlice []interface{}) {
	for _, v := range s {
		if a(v) {
			ftSlice = append(ftSlice, v)
		}
	}
	return
}

func (w *WUtilSlice) SliceDiff(s1, s2 []interface{}) (s3 []interface{}) {
	for _, v := range s1 {
		if !w.InSliceInterface(v, s2) {
			s3 = append(s3, v)
		}
	}
	return
}

func (w *WUtilSlice) SliceIntersect(s1, s2 []interface{}) (s3 []interface{}) {
	for _, v := range s1 {
		if !w.InSliceInterface(v, s2) {
			s3 = append(s3, v)
		}
	}
	return
}

func (w *WUtilSlice) SliceChunk(s []interface{}, size int) (chunkSlice [][]interface{}) {
	if size >= len(s) {
		chunkSlice = append(chunkSlice, s)
		return
	}
	end := size
	for i := 0; i < (len(s) - size); i += size {
		chunkSlice = append(chunkSlice, s[i:end])
		end += size
	}
	return
}

func (w *WUtilSlice) SliceRange(start, end, step int64) (intSlice []int64) {
	for i := start; i <= end; i += step {
		intSlice = append(intSlice, i)
	}
	return
}

func (w *WUtilSlice) SlicePad(slice []interface{}, size int, val interface{}) []interface{} {
	if size <= len(slice) {
		return slice
	}
	for i := 0; i < (size - len(slice)); i++ {
		slice = append(slice, val)
	}
	return slice
}

func (w *WUtilSlice) SliceUnique(slice []interface{}) (uniqueSlice []interface{}) {
	for _, v := range slice {
		if !w.InSliceInterface(v, uniqueSlice) {
			uniqueSlice = append(uniqueSlice, v)
		}
	}
	return
}

func (w *WUtilSlice) SliceShuffle(slice []interface{}) []interface{} {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}

func (w *WUtilSlice) MoveStr2Slice(str string, buf []byte) {
	size := binary.Size(buf)
	c := len(str)
	for i := 0; i < size; i++ {
		if i < c {
			buf[i] = byte(str[i])
		} else {
			buf[i] = 0
		}
	}
}

func (w *WUtilSlice) CompareSliceStr(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func (w *WUtilSlice) CompareSliceByte(s1, s2 []byte) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func (w *WUtilSlice) CompareSliceInt(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func (w *WUtilSlice) SliceStr2SliceInt(s []string) (ret []int) {
	for i := 0; i < len(s); i++ {
		v, err := strconv.Atoi(s[i])
		if err == nil {
			ret = append(ret, v)
		}
	}
	return
}

func (w *WUtilSlice) SliceHex2SliceUInt16(s []string) (ret []uint16) {
	for i := 0; i < len(s); i++ {
		v, err := strconv.ParseUint(s[i], 16, 16)
		if err == nil {
			ret = append(ret, uint16(v))
		}
	}
	return
}

func (w *WUtilSlice) ExpandStringSlice(s []string, split string) []string {
	ret := make([]string, 0, len(s))
	for _, v := range s {
		ls := strings.Split(v, split)
		for _, d := range ls {
			if len(d) > 0 {
				ret = append(ret, d)
			}
		}
	}
	return ret
}

