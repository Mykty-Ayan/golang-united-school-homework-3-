package homework

func reverse(input []int64) (result []int64) {
	left := 0
	right := len(input) - 1
	for left < right {
		temp := input[left]
		input[left] = input[right]
		input[right] = temp
		left += 1
		right -= 1
	}
	return input
}
