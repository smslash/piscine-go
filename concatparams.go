package piscine

func ConcatParams(args []string) string {
	var s string
	for i := 0; i < len(args); i++ {
		s += args[i]
		if i+1 != len(args) {
			s += string('\n')
		}
	}

	return s
}
