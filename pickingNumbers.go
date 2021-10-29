package main

import(
	"fmt"
)

func max(a, b, c int)int{
	if a>b{
		if a>c{
			return a
		}
		return c
	}else{
		if b>c{
			return b
		}
		return c
	}
}

func solution(inp []int){
	res:=make(map[int]int)
	m:=0
	for _,val:=range inp{
		res[val]+=1
		res[val-1]+=1
		m=max(res[val], res[val-1],m)
	}
	fmt.Println("solution :",m)
}

func main(){
	a:=[]int{4,6,5,3,3,1}
	solution(a)
	a=[]int{1,2,2,3,1,2}
	solution(a)
}