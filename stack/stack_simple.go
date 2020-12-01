package stack

import "errors"

type Stack []interface{}

func (stack *Stack) Pop() (interface{}, error) {
	var value interface{}
	var err error
	if value, err = stack.Top(); err != nil {
		return nil, err
	}
	tmpStack := *stack
	*stack = tmpStack[:stack.Length()-1]
	return value, nil
}

func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

func (stack *Stack) Top() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	tmpStack := *stack
	return tmpStack[stack.Length()-1], nil
}

func (stack *Stack) IsEmpty() bool {
	return stack.Length() == 0
}

func (stack *Stack) Length() int {
	return len(*stack)
}

//Problem 20
func IsValid(s string) bool {
	stack := make(Stack, 0)
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