package problem0461

func hammingDistance1(x int, y int) int {
	c:=x^y
	count:=0
	for c!=0{
		d:=c>>1
		if (c-(d<<1))==1{
			count++
		}
		c=c>>1
	}
	return count
}

func hammingDistance(x int, y int) int {
	c:=x^y
	count:=0
	for c!=0{
		if (c&1)==1{
			count++
		}
		c=c>>1
	}
	return count
}
