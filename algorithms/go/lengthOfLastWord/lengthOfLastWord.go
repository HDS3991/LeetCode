package lengthoflastword

func lengthOfLastWord(s string) int {
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}
	ans := 0
	for index >= 0 && s[index] != ' ' {
		ans++
		index--
	}
	return ans
}
