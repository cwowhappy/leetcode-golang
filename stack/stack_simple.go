package stack

import "github.com/cwowhappy/leetcode-golang/util"

//Problem 20
func IsValid(s string) bool {
	stack := make(util.Stack, 0)
	bracketMap := make(map[string]string)
	bracketMap["]"] = "["
	bracketMap[")"] = "("
	bracketMap["}"] = "{"

	for i := 0; i < len(s); i++ {
		char := s[i:i+1]
		if bracket, ok := bracketMap[char]; ok {
			if value, err := stack.Pop(); err == nil {
				if value != bracket {
					return false
				}
			} else {
				return false
			}
		} else {
			stack.Push(char)
		}
	}
	return stack.IsEmpty()
}