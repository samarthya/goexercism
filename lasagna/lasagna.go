package lasagna

// OvenTime time in the oven
const (
	OvenTime = 40
)

// RemainingOvenTime() time to remain in the oven
func RemainingOvenTime(t int) int {
	return OvenTime - t
}

// PreparationTime() time to prepare
func PreparationTime(n int) int {
	return 2 * n
}

// ElapsedTime Preparation time + time in oven
func ElapsedTime(l, t int) int {
	return PreparationTime(l) + t
}
