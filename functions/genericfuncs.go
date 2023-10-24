package functions

func ReverseSlice(inp []string) []string {
	for i := 0; i < len(inp)/2; i++ {
		inp[i], inp[len(inp)-1-i] = inp[len(inp)-1-i], inp[i]
	}
	return inp
}
