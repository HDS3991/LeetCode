package longestconsecutive

func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
		return 0
	}
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	maxLen := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentLen := 1
			for numSet[currentNum+1] {
				currentNum++
				currentLen++
			}
			if currentLen > maxLen {
				maxLen = currentLen
			}
		}
	}
	return maxLen
}