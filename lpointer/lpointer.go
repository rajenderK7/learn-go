package lpointer

func SwapByRef(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func SwapByRefWOTemp(a *int, b *int) {
	// a = 1, b = 2
	// a = a + b => a = 1 + 2 = 3
	// b = a - b => b = 3 - 2 = 1
	// a = a - b => a = 3 - 1 = 2
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func UnderstandMapModify(m map[string]string) {
	m["virat"] = "kohli"
}

// Slice has a Slice header, which contains the
// following metadata -> length, capacity and
// pointer to underlying array.
// When a slice is passed as a parameter to a function
// only the slice header is copied and passed which will contain
// also contain the pointer to the original array.
// Therefore any modification done within the capacity
// of the passed slice, will reflect in the original slice.
// However, when we use functions such as append, delete etc.
// in the function new memory is allocated and pointer to the
// original array of the slice is lost.
func UnderstandSliceModify(slc []int) {
	n := len(slc)
	slc[0] = 99
	slc[n-1] = 200
}
