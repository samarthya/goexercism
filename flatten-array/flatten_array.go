package flatten

// Flatten It returns a flat array
func Flatten(i interface{}) (result []interface{}) {
	result = make([]interface{}, 0)
	// Check for the type.
	if list, ok := i.([]interface{}); ok {
		// It is a list, range iver element by element.
		for _, element := range list {
			if element == nil {
				continue
			}
			result = append(result, Flatten(element)...)
		}
		return result
	}
	return append(result, i)

}
