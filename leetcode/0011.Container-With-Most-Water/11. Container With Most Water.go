package leetcode

func myMaxArea(height []int) int {
	max, left, right := 0, 0, len(height)-1
	for left < right {
		width := right - left
		high := 0
		if height[left] < height[right] {
			high = height[left]
			left++
		} else {
			high = height[right]
			right--
		}
		tmp := width * high
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		
		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}
