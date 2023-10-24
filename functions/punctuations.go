package functions

func cleanPunctuations(arr []string) []string {
	res := []string{}
	return res
}

func ReplaceUnknownRunes(inp []string) []string {
	for i, w := range inp {
		if w == "‘" {
			inp[i] = "'"
		}
	}
	return inp
}

func findIndexNonEmptyStr(inp []string, index int) int {
	for i := index - 1; i >= 0; i-- {
		if inp[i] != "\n" && inp[i] != "" {
			return i
		}
	}
	return -1
}

func MergeQuotes(inp []string) ([]string, bool) {
	var opened bool = false
	for i := 0; i < len(inp); i++ {
		if inp[i] == "'" && !opened {
			opened = true
		} else if inp[i] != "'" && inp[i] != "\n" && opened {
			inp[i-1] += inp[i]
			inp[i] = ""
		} else if inp[i] == "'" && opened {
			ind := findIndexNonEmptyStr(inp, i)
			inp[ind] += inp[i]
			inp[i] = ""
			opened = false
		} else if inp[i] == "\n" && i-1 >= 0 && inp[i-1] == "'" && opened {
			inp[i-1], inp[i] = inp[i], inp[i-1]
		}
	}
	return inp, opened
}

func addOneSign(inp []string, currIndex int) []string {
	for i := currIndex; i >= 0; i-- {
		if currIndex > 0 && inp[i-1] != "" {
			inp[i-1] += inp[currIndex]
			inp[currIndex] = ""
			break
		}
	}
	return inp
}

func isSinglePunctuation(s string) bool {
	if s == "." || s == "," || s == "!" || s == "?" || s == ";" || s == ":" {
		return true
	}
	return false
}

func AddAllSigns(inp []string) []string {
	for i := 0; i < len(inp); i++ {
		if isSinglePunctuation(inp[i]) || (!(isValidString(inp[i])) && inp[i] != "\n" && inp[i] != "") {
			inp = addOneSign(inp, i)
		}
	}
	return inp
}
