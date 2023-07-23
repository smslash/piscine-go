package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	items := make([]string, 0)
	currentItem := ""
	ans := 0
	for _, c := range str {
		if c == ' ' {
			ans++
			if currentItem != "" {
				items = append(items, currentItem)
				currentItem = ""
			}
		} else {
			currentItem += string(c)
		}
	}

	if currentItem != "" {
		items = append(items, currentItem)
	}
	summary := make(map[string]int)
	if ans-len(items)+1 != 0 {
		summary[""] = ans - len(items) + 1
	}
	for _, item := range items {
		if item != "" {
			summary[item]++
		}
	}
	return summary
}
