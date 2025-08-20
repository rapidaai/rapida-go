package utils

func AverageFloat32(numbers []float32) float32 {
	if len(numbers) == 0 {
		return 0
	}

	sum := float32(0)
	for _, num := range numbers {
		sum += num
	}

	return sum / float32(len(numbers))
}
