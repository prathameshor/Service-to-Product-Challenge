package day_three

func ContainerWithMostWater(height []int) int {
	volume := 0
	i, j := 0, len(height)-1
	for i < j {
		val := (j - i) * min(height[i], height[j])
		if val > volume {
			volume = val
		}

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return volume
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
