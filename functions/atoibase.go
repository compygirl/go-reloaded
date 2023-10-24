package functions

func AtoiBase(s string, base string) int {
	res := 0
	index := 0
	lent := len(s)
	if len(base) < 2 {
		return 0
	}
	for i := 0; i < len(base)-1; i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return 0
			}
		}
	}
	for _, symbol := range base {
		if symbol == '-' || symbol == '+' {
			return 0
		}
	}

	for _, letter := range s {
		for i := 0; i < len(base); i++ {
			if rune(base[i]) == letter {
				index = i
				break
			}
		}
		res += index * IterativePower(len(base), lent-1)
		lent--
	}
	return res
}

func IterativePower(nb int, power int) int {
	res := 1
	if power >= 0 {
		for i := 1; i <= power; i++ {
			res *= nb
		}
	} else {
		res = 0
	}

	return res
}
