package main

import(
	"fmt"
)
func reverse(a string)string{
	temp:=""
	l:=len(a)
	for i:=0;i<l;i++{
		temp=string(a[i])+temp
	}
	return temp
}

func checkPalindrome(a string)bool{
	if a==reverse(a){
		return true
	}
	return false
}

func longestPalindrome(a string)string{
	if len(a)<1{
		return ""
	}
	if checkPalindrome(a){
		return a
	}
	temp1:=longestPalindrome(a[1:])
	temp2:=longestPalindrome(a[:len(a)-1])
	if len(temp1)>len(temp2){
		return temp1
	}
	return temp2
}

func Get2DArry(l,b int, val interface{})[][]interface{}{
	dyn:=[][]interface{}{}
	for i:=0;i<l;i++{
		temp:=[]interface{}{}
		for j:=0;j<b;j++{
			temp=append(temp,val)
		}
		dyn=append(dyn,temp)
	}
	return dyn
}

func print2Darray(a [][]interface{}){
	l:=len(a)
	for i:=0;i<l;i++{
		fmt.Println(a[i])
	}
}

func longestPalindromeDynamic(a string)string{
	l:=len(a)
	dyn:=Get2DArry(l,l, false)
	maxLength:=1
	for i:=0;i<l;i++{
		dyn[i][i]=true
	}
	start:=0
	for i:=0;i<l-1;i++{
		if a[i]==a[i+1]{
			dyn[i][i+1]=true
			start=i
			maxLength=2
		}
	}
	fmt.Println(start, maxLength)
	print2Darray(dyn)
	return " "
}

func main(){
	input:="forgeeksskeegfor"
	fmt.Println("Output - ",longestPalindrome(input))
	// fmt.Println(longestPalindromeDynamic(input))
}