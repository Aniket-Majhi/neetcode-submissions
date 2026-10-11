func findMaxConsecutiveOnes(nums []int) int {
	var count, tmp int

	for i := 0; i < len(nums); i++ {
		val := nums[i]
		
		// If val == 1:(tmp + 1) * 1 -> increments
		// If val == 0:(tmp + 1) * 0 -> resets to 0
		//This is so sneeky. I wuw it
		tmp = (tmp + 1) * val
		if tmp > count {
			count = tmp
		}
	}
	return count
}
