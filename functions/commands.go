package functions

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// only for up, low, cap
func CreateCommands(inp []string) []string {
	var res []string
	temp := ""
	for i := 0; i < len(inp); i++ {
		if (inp[i] == "(low" || inp[i] == "(up" || inp[i] == "(cap") && len(temp) == 0 {
			temp += inp[i]
		} else if (inp[i] == ")" || strings.Contains(inp[i], ")")) && len(temp) > 0 {
			if inp[i] == ")" {
				temp += inp[i]
			} else {
				temp += " " + inp[i]
			}
			res = append(res, temp)
			temp = ""
		} else if inp[i] == "," && len(temp) > 0 {
			temp += inp[i]
		} else if len(temp) == 0 {
			res = append(res, inp[i])
		} else if len(temp) > 0 {
			if temp[len(temp)-1] == '(' || temp[len(temp)-1] == '[' || temp[len(temp)-1] == '{' {
				temp += inp[i]
			} else {
				temp += " " + inp[i]
			}
		}
	}
	return res
}

func isBinStr(s string) bool {
	for _, l := range s {
		if l != '0' && l != '1' {
			return false
		}
	}
	return true
}

func isNotCommand(s string) bool {
	if s == "(bin)" || s == "(hex)" || s == "(low)" || s == "(up)" || s == "(cap)" || strings.HasPrefix(s, "(low") || strings.HasPrefix(s, "(up") || strings.HasPrefix(s, "(cap") {
		return false
	}
	return true
}

func isNotPunctuation(s string) bool {
	if s == "," || s == "." || s == "!" || s == "?" || s == "'" || s == ":" || s == ";" || s == "(" || s == ")" {
		return false
	}
	return true
}

func findIndexBin(inp []string, index int) int {
	for i := index - 1; i >= 0; i-- {
		if inp[i] != "" && isBinStr(inp[i]) && isNotCommand(inp[i]) && isNotPunctuation(inp[i]) {
			return i
		} else if !isBinStr(inp[i]) && isNotCommand(inp[i]) && isNotPunctuation(inp[i]) {
			return -2
		}
	}
	return -1
}

func ApplyAllCommands(inp []string) []string {
	for i := 0; i < len(inp); i++ {
		switch {
		case inp[i] == "(bin)":
			// fmt.Println(i)
			ApplyBin(inp, i)
		case inp[i] == "(hex)":
			ApplyHex(inp, i)
		case inp[i] == "(low)" || inp[i] == "(up)" || inp[i] == "(cap)":
			ind := findIndexStr(inp, i)
			if ind < 0 {
				inp[i] = ""
			} else {
				ApplyToStr(inp[i], inp, i, ind)
			}
		case inp[i] == "(low":
			ApplyToSevralStr("(low", inp, i)

		case inp[i] == "(up":
			ApplyToSevralStr("(up", inp, i)

		case inp[i] == "(cap":
			ApplyToSevralStr("(cap", inp, i)
		}
	}
	return inp
}

func checkValidityOfCommand(cmd string, inp []string, index int) int {
	if index+1 >= len(inp) { // this check didn't happend
		fmt.Println("ERROR: not valid command started with (cap OR (low OR (up [stopped right here]")
		return -1
	} else if inp[index+1] != "," {
		fmt.Println("ERROR: missing comma after (cap OR (low OR (up")
		return -1
	} else {
		// do the rest of the check
		if index+2 >= len(inp) {
			fmt.Println("ERROR: not valid command started with (cap OR (low OR (up [stopped after comma]")
			return -1
		} else if isNumeric(inp[index+2][:len(inp[index+2])-1]) && strings.Contains(inp[index+2], ")") {
			num, err := strconv.Atoi(inp[index+2][:len(inp[index+2])-1])
			if err != nil {
				fmt.Println(err.Error())
			}
			return num
		} else {
			fmt.Println("ERROR: the number is not valid or negative")
			return -1
		}
	}
}

func isNumeric(s string) bool {
	for _, l := range s {
		if l < '0' || l > '9' {
			return false
		}
	}
	return true
}

func ApplyToSevralStr(cmd string, inp []string, cmdIndex int) {
	num := checkValidityOfCommand(cmd, inp, cmdIndex)

	if num > cmdIndex { // if param bigger than the number of strings (len(inp) - (len(inp) - cmdIndex))
		num = cmdIndex
	} else if num == 0 {
		inp[cmdIndex] = ""
		inp[cmdIndex+1] = "" // clean comma
		inp[cmdIndex+2] = "" // clean second part with num
		return
	}
	currInd := cmdIndex
	for num > 0 && currInd > 0 {
		if currInd-1 >= 0 && isValidString(inp[currInd-1]) {
			ApplyToStr(cmd, inp, cmdIndex, currInd-1)
			num--
		}
		currInd--
	}
}

func ApplyBin(inp []string, index int) {
	ind := findIndexBin(inp, index)
	if ind < 0 {
		fmt.Println("ERROR: invalid bin")
		inp[index] = ""
	} else {
		num := AtoiBase(inp[ind], "01")
		str := strconv.Itoa(num)
		// add any strings between current and the binary value
		inp[ind] = str
		inp[index] = ""
	}
}

func isHexStr(s string) bool {
	for _, l := range s {
		if (!(l >= '0' && l <= '9') && (!(l >= 'A' && l <= 'F'))) && (!(l >= 'a' && l <= 'f')) {
			return false
		}
	}
	return true
}

func findIndexHex(inp []string, index int) int {
	for i := index - 1; i >= 0; i-- {
		if inp[i] != "" && isHexStr(inp[i]) && isNotCommand(inp[i]) && isNotPunctuation(inp[i]) {
			return i
		} else if !isHexStr(inp[i]) && isNotCommand(inp[i]) && isNotPunctuation(inp[i]) {
			return -2
		}
	}
	return -1
}

func IsUpper(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

// Not using
func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func ApplyHex(inp []string, index int) {
	ind := findIndexHex(inp, index)
	if ind < 0 {
		fmt.Println("ERROR: invalid hex")
		inp[index] = ""
	} else {
		num := 0
		if IsUpper(inp[ind]) {
			num = AtoiBase(inp[ind], "0123456789ABCDEF") // only works for capital cases
		} else {
			num = AtoiBase(inp[ind], "0123456789abcdef")
		}
		str := strconv.Itoa(num)
		inp[ind] = str
		inp[index] = ""
	}
}

func isValidString(s string) bool {
	for _, let := range s {
		if (let >= 'a' && let <= 'z') || (let >= 'A' && let <= 'Z') || (let >= '0' && let <= '9') || let == '\'' {
			return true
		}
	}
	return false
}

func findIndexStrAfter(inp []string, index int) int {
	for i := index; i < len(inp); i++ {
		if inp[i] != "\n" && inp[i] != "" && isValidString(inp[i]) && isNotCommand(inp[i]) && isNotPunctuation(inp[i]) {
			return i
		} else if inp[i] != "\n" && inp[i] != "" && !isValidString(inp[i]) && isNotCommand(inp[i]) && isNotPunctuation(inp[i]) {
			return -1
		}
	}
	return -1
}

func findIndexStr(inp []string, index int) int {
	for i := index - 1; i >= 0; i-- {
		if inp[i] != "\n" && inp[i] != "" && isValidString(inp[i]) && isNotCommand(inp[i]) && isNotPunctuation(inp[i]) {
			return i
		}
	}
	return -1
}

func ApplyToStr(cmd string, inp []string, cmdIndex int, strIndex int) {
	var several bool = false
	str := ""
	switch {
	case cmd == "(low)":
		str = strings.ToLower(inp[strIndex])
	case cmd == "(up)":
		str = strings.ToUpper(inp[strIndex])
	case cmd == "(cap)":
		str = strings.ToLower(inp[strIndex])
		str = strings.Title(str)
	case cmd == "(cap":
		str = strings.ToLower(inp[strIndex])
		str = strings.Title(str)
		several = true
	case cmd == "(low":
		str = strings.ToLower(inp[strIndex])
		several = true
	case cmd == "(up":
		str = strings.ToUpper(inp[strIndex])
		several = true

	}
	inp[strIndex] = str
	inp[cmdIndex] = ""
	if several {
		inp[cmdIndex+1] = ""
		inp[cmdIndex+2] = ""
	}
}
