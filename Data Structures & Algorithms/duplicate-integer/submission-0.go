func hasDuplicate(nums []int) bool {
	deem:=make(map[int]struct{})
   	for _, val := range nums {
		if _,exist:=deem[val];exist{
			return true
		}
		deem[val]=struct{}{}
	}
	return false
}
