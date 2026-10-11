func findMaxConsecutiveOnesBranchless(nums []int) int {
	var count, tmp int

	for i := 0; i < len(nums); i++ {
		val := nums[i]
		
		// If val == 1: tmp becomes (tmp + 1) * 1 -> increments
		// val == 0: tmp becomes (tmp + 1) * 0 -> resets to 0
		//this is so sneaky. I wuw it
		tmp = (tmp + 1) * val
		if tmp > count {
			count = tmp
		}
	}
	return count
}
