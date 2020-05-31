package problem0020

func isValid1(s string) bool {
	a:=make([]byte,len(s))
	last:=0
	for i:=0;i<len(s);i++{
		current:=s[i]
		switch current{
		case ')':
			if last<=0{
				return false
			}
			if a[last-1]=='('{
				last--
			}else{
				a[last]=current
				last++
			}
		case '}':
			if last<=0{
				return false
			}
			if a[last-1]=='{'{
				last--
			}else{
				a[last]=current
				last++
			}
		case ']':
			if last<=0{
				return false
			}
			if a[last-1]=='['{
				last--
			}else{
				a[last]=current
				last++
			}
		default:
			a[last]=current
			last++
		}
	}
	if last==0{
		return true
	}
	return false
}

func isValid(s string) bool {
	a:=make([]byte,0)
	for i:=0;i<len(s);i++{
		current:=s[i]
		switch current{
		case ')':
			if len(a)==0{
				return false
			}
			if a[len(a)-1]=='('{
				a=a[:len(a)-1]
			}else{
				a=append(a,current)
			}
		case '}':
			if len(a)==0{
				return false
			}
			if a[len(a)-1]=='{'{
				a=a[:len(a)-1]
			}else{
				a=append(a,current)
			}
		case ']':
			if len(a)==0{
				return false
			}
			if a[len(a)-1]=='['{
				a=a[:len(a)-1]
			}else{
				a=append(a,current)
			}
		default:
			a=append(a,current)
		}
	}
	if len(a)==0{
		return true
	}
	return false
}
