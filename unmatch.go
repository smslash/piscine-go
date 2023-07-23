package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		ans := 0
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] && i != j {
				ans++
				if ans == 2 {
					ans = 0
				}
			}
		}
		if ans < 1 {
			return a[i]
		}
	}
	return -1
}
