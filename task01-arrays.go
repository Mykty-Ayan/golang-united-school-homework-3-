package homework

func average(input [15]float32) (result float32) {
	sum := float32(0)
	lastNonZeroIndex := 0
	for i := 0; i < len(input); i++ {
		if input[i] != 0 {
			lastNonZeroIndex = i
		}
	}
	for i := 0; i <= lastNonZeroIndex; i++ {
		sum += input[i]
	}
	result = sum / float32(lastNonZeroIndex+1)
	return result
}
