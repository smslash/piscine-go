package piscine

func ShoppingListSort(slice []string) []string {
	sortedSlice := slice[:]
	for i := 0; i < len(sortedSlice); i++ {
		for j := i + 1; j < len(sortedSlice); j++ {
			if len(sortedSlice[i]) > len(sortedSlice[j]) {
				sortedSlice[i], sortedSlice[j] = sortedSlice[j], sortedSlice[i]
			}
		}
	}
	return sortedSlice
}
