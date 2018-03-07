package binarysearch

func BinarySearch(nums []int, target int) int {
	mid := len(nums) / 2

	if nums[mid] == target {
		return mid
	} else if target > nums[mid] {
		return BinarySearch(nums[mid+1:], target)
	} else {
		return BinarySearch(nums[:mid], target)
	}
}
