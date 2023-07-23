package piscine

func SortIntegerTable(table []int) {
	for j := 0; j < len(table)-1; j++ {
		for i := 0; i < len(table)-1; i++ {
			if table[i] > table[i+1] {
				tmp := table[i]
				table[i] = table[i+1]
				table[i+1] = tmp
			}
		}
	}
}
