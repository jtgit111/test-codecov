package testCodecov

import (
	"errors"
)

// Stats 计算数组的最大值、最小值和平均值
func Stats(arr []int) (max int, min int, avg float64, err error) {
	if len(arr) == 0 {
		return 0, 0, 0, errors.New("array cannot be empty")
	}

	max = arr[0]
	min = arr[0]
	sum := 0

	for _, value := range arr {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
		sum += value
	}

	avg = float64(sum) / float64(len(arr))
	return max, min, avg, nil
}

func Sum1(nums []uint) (res uint) {
	for _, num := range nums {
		res += num
	}
	return
}

func FindMaxFrequencyElement(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0 // 返回0和0表示数组为空
	}

	frequencyMap := make(map[int]int)
	maxElement := arr[0]
	maxFrequency := 1

	// 统计每个元素的频率
	for _, num := range arr {
		frequencyMap[num]++
		if frequencyMap[num] > maxFrequency {
			maxFrequency = frequencyMap[num]
			maxElement = num
		}
	}

	return maxElement, maxFrequency
}