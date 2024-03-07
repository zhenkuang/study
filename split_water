package main

import (
	"fmt"
	"sort"
)

func main() {
	for i := 1; i < 10; i++ {
		high := []int{1, 3, 3, 3, 5}
		printCase(high, i)
	}
}

func printCase(high []int, v int) {
	fmt.Println("Before:", high, "v:", v)
	fmt.Println("After:", splitWater(high, v))
}

func splitWater(high []int, v int) []float32 {
	sort.Ints(high)

	resultWater := make([]float32, len(high))
	// 当要加的水小于第一个水桶里的水，直接加到第一个水桶里
	if v <= high[0] {
		resultWater[0] = float32(high[0] + v)
		for i := 1; i < len(high); i++ {
			resultWater[i] = float32(high[i])
		}
		return resultWater
	}

	// 获取离结果最近的水位线
	result := getWaterLine(high, v)
	// 获取加水后剩余的水
	residue := calculate(result, high)

	// 重新分配剩余的水
	if residue != 0 {
		residueWater := float32(v-residue) / float32(result+1)
		//fmt.Println("cal", float32(v-residue)/float32(result+1))
		for i := 0; i <= result; i++ {
			resultWater[i] = float32(high[result]) + residueWater
		}
	} else {
		for i := 0; i <= result; i++ {
			resultWater[i] = float32(high[result])
		}
	}
	for i := result + 1; i < len(high); i++ {
		resultWater[i] = float32(high[i])
	}
	return resultWater
}

// 二分查找最近的水位线
func getWaterLine(high []int, v int) int {
	left, right := 0, len(high)-1
	var middle int
	var middleWater int
	for left <= right {
		middle = (left + right) / 2
		middleWater = calculate(middle, high)
		if middleWater < v {
			left = middle + 1
		} else if middleWater > v {
			right = middle - 1
		} else {
			return middle
		}
	}
	return left - 1
}

// 计算需要添加的水
func calculate(index int, high []int) int {
	result := 0
	for i := 0; i <= index; i++ {
		result = result + (high[index] - high[i])
	}
	return result
}
