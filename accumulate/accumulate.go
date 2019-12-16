package accumulate

type funcToCall func(string) string

func Accumulate(inputStrings []string, op funcToCall) (outStrings []string) {

	for _, val := range inputStrings {
		outStrings = append(outStrings, op(val))
	}

	return outStrings
}
