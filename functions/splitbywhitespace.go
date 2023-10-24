package functions

import (
	"regexp"
)

func SplitByDelimiters(input string) []string {
	pattern := "([\\{\\}\\[\\]().,:;!?\\n\\s'])"
	regex := regexp.MustCompile(pattern)

	indices := regex.FindAllStringIndex(input, -1)
	last := 0
	result := make([]string, 0, len(indices)*2)
	for _, idx := range indices {
		result = append(result, input[last:idx[0]])
		result = append(result, input[idx[0]:idx[1]])
		last = idx[1]
	}
	result = append(result, input[last:])

	return result
}

func RemoveSpaces(arr []string) []string {
	var res []string
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" && arr[i] != " " {
			res = append(res, arr[i])
		}
	}
	return res
}

func MergeBracket(inp []string) []string {
	var res []string
	var res2 []string
	temp := ""
	for i := 0; i < len(inp); i++ {
		if (inp[i] == "(" || inp[i] == "[" || inp[i] == "{") && len(temp) == 0 {
			temp += inp[i]
		} else if len(temp) > 0 && inp[i] != "(" {
			temp += inp[i]
			res = append(res, temp)
			temp = ""
		} else {
			if len(temp) > 0 {
				res = append(res, temp)
			} else {
				res = append(res, inp[i])
			}
		}
	}

	res = RemoveSpaces(res)
	temp = ""
	for i := len(res) - 1; i >= 0; i-- {
		if (res[i] == ")" || res[i] == "]" || res[i] == "}") && len(temp) == 0 {
			temp += res[i]
		} else if len(temp) > 0 && res[i] != ")" {
			temp = res[i] + temp
			res2 = append(res2, temp)
			temp = ""
		} else {
			if len(temp) > 0 {
				res2 = append(res2, temp)
			} else {
				res2 = append(res2, res[i])
			}
		}
	}
	res2 = ReverseSlice(res2)

	RemoveSpaces(res2)
	return res2
}

func HandleRestOfBrackets(inp []string) []string {
	var res []string
	var res2 []string
	temp := ""
	for i := 0; i < len(inp); i++ {
		if (inp[i] == "(" || inp[i] == "[" || inp[i] == "{") && len(temp) >= 0 {
			temp += inp[i]
		} else if len(temp) > 0 {
			temp += inp[i]
			res = append(res, temp)
			temp = ""
		} else {
			res = append(res, inp[i])
		}
	}
	res = RemoveSpaces(res)

	for i := len(res) - 1; i >= 0; i-- {
		if (res[i] == ")" || res[i] == "]" || res[i] == "}") && len(temp) >= 0 {
			temp += res[i]
		} else if len(temp) > 0 {
			temp = res[i] + temp
			res2 = append(res2, temp)
			temp = ""
		} else {
			res2 = append(res2, res[i])
		}
	}
	ReverseSlice(res2)
	res2 = RemoveSpaces(res2)
	return res2
}
