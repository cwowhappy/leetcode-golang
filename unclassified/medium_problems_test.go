package unclassified

import "testing"

func TestCanJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	result := true
	if CanJump(nums) != result {
		t.Error()
	}
}

func TestJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	result := 2
	if JumpV2(nums) != result {
		t.Error()
	}

}
