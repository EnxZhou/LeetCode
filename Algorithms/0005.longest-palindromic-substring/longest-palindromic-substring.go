package problem0005

func longestPalindrome(s string) string {
	if len(s)==0{
		return ""
	}
	var minI,maxJ int

	dp:=make([][]int,len(s))
	for i:=0;i<len(s)-1;i++{
		dp[i]=make([]int,len(s))
		dp[i][i]=1
		if s[i]==s[i+1]{
			dp[i][i+1]=1
			minI=i
			maxJ=i+1
		}else{
			dp[i][i+1]=0
		}
	}
	dp[len(s)-1]=make([]int,len(s))
	dp[len(s)-1][len(s)-1]=1
	for sl:=2;sl<len(s);sl++{
		for i:=0;i<=len(s)-sl-1;i++{
			if dp[i+1][i+sl-1]==1&&(s[i]==s[i+sl]){
				dp[i][i+sl]=1
				if sl>(maxJ-minI){
					minI=i
					maxJ=i+sl
				}
			}else{
				dp[i][i+sl]=0
			}
		}
	}
	return s[minI:maxJ+1]
}
