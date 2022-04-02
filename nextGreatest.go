package main

import(
	"fmt"
)

func nextGreatestRaw(a []int)[]int{
	res:=make([]int,0)
	l:=len(a)
	for i:=0;i<l-1;i++{
		min:=1<<31
		for j:=i+1;j<l;j++{
			if a[i]<a[j] && a[j]<min {
				min=a[j]
			}
		}
		if min!=1<<31{
			res=append(res,min)
		}else{
			res=append(res,-1)
		}
	}
	res=append(res,-1)
	return res
}

func main(){
	inp:=[]int{4, 5, 2, 25}
	fmt.Println("Result :",nextGreatestRaw(inp))
	inp=[]int{5,4,3,2,1}
	fmt.Println("Result :",nextGreatestRaw(inp))
	inp=[]int{13, 7, 6, 12}
	fmt.Println("Result :",nextGreatestRaw(inp))
}