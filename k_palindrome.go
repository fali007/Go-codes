package main

import(
	"fmt"
)
func min(a,b int)int{
	if a>b{
		return b
	}
	return a
}
func reverse(s string)string{
	temp:=""
	l:=len(s)
	for i:=0;i<l;i++{
		temp=string(s[i])+temp
	}
	return temp
}
func checkPalindrome(inp1,inp2 string, m,n int)int{
	if m==0{
		return n
	}
	if n==0{
		return m
	}
	if inp1[m-1]==inp2[n-1]{
		return checkPalindrome(inp1,inp2,m-1,n-1)
	}
	return 1+min(checkPalindrome(inp1,inp2,m-1,n),checkPalindrome(inp1,inp2,m,n-1));
}

func isKpal(inp string, k int)bool{
	rev:=reverse(inp)
	l:=len(inp)
	return (checkPalindrome(inp,rev,l,l)<=k*2)
}

func main(){
	input:="abcdecba"
	fmt.Println(isKpal(input,1))
}