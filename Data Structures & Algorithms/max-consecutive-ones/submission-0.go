func findMaxConsecutiveOnes(nums []int) int {
	var count int
	tmp:=0
	for i:=0;i<len(nums);i++{
		if nums[i]==1{
			tmp++
			if tmp>count{
				count=tmp
			}
		}else{
			tmp=0
		}
	}
	return count
}
