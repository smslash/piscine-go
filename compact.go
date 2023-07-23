package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	newSlice := make([]string, 0, len(slice))
	count := 0

	for _, str := range slice {
		if str != "" {
			newSlice = append(newSlice, str)
			count++
		}
	}

	*ptr = newSlice

	return count
}
