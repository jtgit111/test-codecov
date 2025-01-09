package testCodecov

import (
	"testing"
)

func TestSum(t *testing.T) {
	nums := []uint{1, 2, 3, 4}
	res := Sum1(nums)
	if res != 10 {
		t.Error("error")
	} 
}

func TestStats(t *testing.T) {
	tests := []struct {
		input    []int
		expectedMax int
		expectedMin int
		expectedAvg float64
		expectErr bool
	}{
		{[]int{1, 2, 3, 4, 5}, 5, 1, 3.0, false},
		{[]int{-1, -2, -3, -4, -5}, -1, -5, -3.0, false},
		{[]int{0, 0, 0, 0}, 0, 0, 0.0, false},
		{[]int{10}, 10, 10, 10.0, false},
	}

	for _, test := range tests {
		max, min, avg, err := Stats(test.input)

		if test.expectErr {
			if err == nil {
				t.Errorf("expected an error for input %v, got none", test.input)
			}
			continue
		} else {
			if err != nil {
				t.Errorf("unexpected error for input %v: %v", test.input, err)
			}
		}

		if max != test.expectedMax {
			t.Errorf("for input %v, expected max %d, got %d", test.input, test.expectedMax, max)
		}
		if min != test.expectedMin {
			t.Errorf("for input %v, expected min %d, got %d", test.input, test.expectedMin, min)
		}
		if avg != test.expectedAvg {
			t.Errorf("for input %v, expected avg %.2f, got %.2f", test.input, test.expectedAvg, avg)
		}
	}
}

func TestFindMaxFrequencyElement(t *testing.T) {
	tests := []struct {
		arr           []int
		expectedElem  int
		expectedFreq  int
	}{
		{[]int{1, 3, 2, 3, 4, 3, 2, 1}, 3, 3},
		{[]int{1, 1, 2, 2, 3, 3, 4}, 1, 2}, // 1 和 2 的频率相同，返回第一个
		{[]int{5, 5, 5, 1, 1, 2}, 5, 3},
		{[]int{}, 0, 0}, // 空数组
		{[]int{7}, 7, 1}, // 只有一个元素
	}

	for _, test := range tests {
		elem, freq := FindMaxFrequencyElement(test.arr)
		if elem != test.expectedElem || freq != test.expectedFreq {
			t.Errorf("对于输入 %v，期望 (%d, %d)，但得到 (%d, %d)", test.arr, test.expectedElem, test.expectedFreq, elem, freq)
		}
	}
}