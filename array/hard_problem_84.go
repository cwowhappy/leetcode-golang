package array

type Node struct {
	Val int
	Index int
}

func Answer1(heights []int) int {
	max := 0
	stack := make([]Node, len(heights))
	j := -1
	for i := 0; i < len(heights); i ++ {
		end := i
		for ;j >= 0 && heights[i] < stack[j].Val; j-- {
			begin := -1
			if j > 0 {
				begin = stack[j-1].Index
			}
			if tmp := stack[j].Val * (end - begin - 1); tmp > max {
				max = tmp
			}
		}
		if j < 0 || stack[j].Val < heights[i] {
			j ++
			stack[j] = Node{
				Val: heights[i],
				Index: i,
			}
		} else if stack[j].Val == heights[i] {
			stack[j].Index = i
		}
	}
	end := len(heights)
	for ; j >= 0; j-- {
		begin := -1
		if j > 0 {
			begin = stack[j-1].Index
		}
		if tmp := stack[j].Val * (end - begin - 1); tmp > max {
			max = tmp
		}
	}

	return max
}