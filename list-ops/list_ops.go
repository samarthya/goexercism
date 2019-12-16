package listops

// Two parameters
type binFunc func(int, int) int

// One parameter
type unaryFunc func(int) int

// predFunc
type predFunc func(int) bool

//IntList is a slice of ints
type IntList []int

// Foldr Method with a reciever argument. Reduces from R to L
func (list IntList) Foldr(reducer binFunc, initial int) int {

	if list.Length() == 0 {
		return initial
	}

	for i := list.Length() - 1; i >= 0; i-- {
		initial = reducer(list[i], initial)
	}

	return initial
}

//Foldl from left to right
func (list IntList) Foldl(reducer binFunc, initial int) int {

	if list.Length() == 0 {
		return initial
	}
	for i := range list {
		initial = reducer(initial, list[i])
	}
	return initial
}

//Map returns the list
func (list IntList) Map(mapfunc unaryFunc) IntList {
	rangeList := make(IntList, list.Length())
	for i, v := range list {
		rangeList[i] = mapfunc(v)
	}
	return rangeList
}

//Concat concatenates
func (list IntList) Concat(listsOfLists []IntList) IntList {
	totallen, counter := 0, 0
	for _, v := range listsOfLists {
		totallen += v.Length()
	}
	flattened := make(IntList, 0, totallen)

	for _, v := range listsOfLists {
		for _, c := range v {
			flattened = flattened.Append([]int{c})
			counter++
		}
	}
	return list.Append(flattened)
}

//Reverse reverses the list
func (list IntList) Reverse() IntList {
	reversedList := make(IntList, list.Length())
	for i, v := range list {
		reversedList[list.Length()-i-1] = v
	}
	return reversedList
}

//Filter returns a list of each element of the list that satisfies the filter function
func (list IntList) Filter(filterFunc predFunc) IntList {
	filteredList := make(IntList, 0, list.Length())
	c := 0
	for _, v := range list {
		if filterFunc(v) {
			filteredList = filteredList.Append([]int{v})
			c++
		}
	}
	return filteredList
}

//Append appends the argument
func (list IntList) Append(appendList IntList) IntList {
	appended := make(IntList, list.Length()+appendList.Length())
	initialLenght := list.Length()
	for i, v := range list {
		appended[i] = v
	}
	for i, v := range appendList {
		appended[initialLenght+i] = v
	}
	return appended
}

//Length length of the list
func (list IntList) Length() int {
	listlen := 0
	for range list {
		listlen++
	}
	return listlen
}
