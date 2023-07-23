package piscine

func CollatzCountdown(start int) int {
	step := 0
	for {
		if start <= 0 {
			return -1
		} else if start == 1 {
			return step
		}

		if start%2 == 0 {
			step++
			start /= 2
		} else {
			step++
			start = start*3 + 1
		}
	}
}
