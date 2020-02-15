package strain

// Ints collection of int
type Ints []int

// Lists array of array of ints
type Lists [][]int

// Strings array of strings
type Strings []string

type boolFn func(int) bool
type boolStringFn func(string) bool
type boolListFn func([]int) bool

// Keep the predicate validated Ints
func (lists Lists) Keep(fnBool boolListFn) (out Lists) {
	if lists == nil {
		return nil
	}
	out = make(Lists, 0)
	for _, val := range lists {
		if fnBool(val) {
			out = append(out, val)
		}
	}
	return
}

// Keep the predicate validated Ints
func (ints Ints) Keep(fnBool boolFn) (out Ints) {

	if ints == nil {
		return nil
	}
	out = make(Ints, 0)

	for _, val := range ints {
		if fnBool(val) {
			out = append(out, val)
		}
	}
	return
}

// Keep the predicate validated Ints
func (s Strings) Keep(boolFn boolStringFn) (out Strings) {
	if s == nil {
		return nil
	}
	out = make(Strings, 0)
	for _, val := range s {
		if boolFn(val) {
			out = append(out, val)
		}
	}
	return
}

// Discard the predicate validated Ints
func (ints Ints) Discard(fnBool boolFn) (out Ints) {
	if ints == nil {
		return nil
	}
	out = make(Ints, 0)
	for _, val := range ints {
		if !fnBool(val) {
			out = append(out, val)
		}
	}
	return
}
