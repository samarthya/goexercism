package sublist

export type Relation string

const (
	EQUAL     Relation = "equal"
	SUBLIST   Relation = "sublist"
	SUPERLIST Relation = "superlist"
	UNEQUAL   Relation = "unequal"
)

// Find the sublist
func Sublist(a []int, b []int) Relation {

	lenA := len(a)
	lenB := len(b)

	if lenA == 0 && lenB > 0 {
		return SUBLIST
	} else if lenB == 0 && lenA > 0 {
		return SUPERLIST
	} else if lenA == 0 && lenB == 0 {
		return EQUAL
	} else {

		if lenA <= lenB {
			for i, val := range b {
				if a[0] == val {
					bIndex := i + 1
					aIndex := 1

					for aIndex < lenA && bIndex < lenB {
						if b[bIndex] == a[aIndex] {
							bIndex++
							aIndex++
							continue
						} else {
							break
						}
					}
					if aIndex == lenA {
						if lenA == lenB {
							return EQUAL
						} else {
							return SUBLIST
						}
					}
				}
			}
		} else {
			for i, val := range a {
				if b[0] == val {
					bIndex := 1
					aIndex := i + 1

					for bIndex < lenB {
						if b[bIndex] == a[aIndex] {
							bIndex++
							aIndex++
							continue
						} else {
							break
						}
					}
					if bIndex == lenB {
						if lenA == lenB {
							return EQUAL
						} else {
							return SUPERLIST
						}
					}
				}
			}
		}
	}
	return UNEQUAL
}

/*
for j, val := range a {
			for i, ival := range b {
				if val == ival {
					// Next elements
					bIndex := i + 1
					aIndex := j + 1

					if lenA <= lenB {
						// While the aIndex is less than the length of the list.
						for aIndex < lenA {
							if b[bIndex] == a[aIndex] {
								bIndex++
								aIndex++
								continue
							} else {
								return UNEQUAL
							}
						}
						if aIndex == lenA && bIndex < lenB {
							return SUBLIST
						} else if aIndex == lenA && bIndex == lenB {
							return EQUAL
						}
					} else if lenB < lenA {
						for bIndex < len(b) {
							if b[bIndex] == a[aIndex] {
								bIndex++
								aIndex++
								continue
							} else {
								return UNEQUAL
							}
						}
						if bIndex == lenB && aIndex < lenA {
							return SUPERLIST
						} else if aIndex == lenA && bIndex == lenB {
							return EQUAL
						}
					}
				}
			}
		}
*/
