package homework

type PairInput struct {
	Nums   []int
	Target int
}

func longestPrefix(words []string) string {
	if len(words) == 0 {
		return ""
	}
	prefix := []rune(words[0])
	for i := 1; i < len(words); i++ {
		word := []rune(words[i])
		newLen := 0
		minLen := len(prefix)
		if len(word) < minLen {
			minLen = len(word)
		}
		for j := 0; j < minLen; j++ {
			if prefix[j] == word[j] {
				newLen++
			} else {
				break
			}
		}
		prefix = prefix[:newLen]
		if len(prefix) == 0 {
			break
		}
	}
	return string(prefix)
}

func pairIndices(input PairInput) [2]int {
	seen := make(map[int]int)
	for i, num := range input.Nums {
		complement := input.Target - num
		if j, ok := seen[complement]; ok {
			if j < i {
				return [2]int{j, i}
			}
			return [2]int{i, j}
		}
		seen[num] = i
	}
	return [2]int{}
}

func hasDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

func compactSorted(nums []int) []int {
	seen := make(map[int]bool)
	result := make([]int, 0, len(nums))
	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}

func balancedBrackets(text string) bool {
	stack := make([]rune, 0)
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, ch := range text {
		if _, ok := pairs[ch]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}
