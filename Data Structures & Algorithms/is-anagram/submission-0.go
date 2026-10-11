func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}
	var cha [26]int
	for i:=0;i< len(s);i++{
		cha[s[i]-'a']++
		cha[t[i]-'a']--
	}
	for _,count:=range cha{
		if count>0{
			return false
		}
	}
	return true
}
