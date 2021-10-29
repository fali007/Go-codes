package main

import(
	"fmt"
)

func solution(inp []int){
	res:=make(map[int]int)
	for _,val:=range inp{
		if res[val]!=0{
			res[val]+=1
		}else{
			res[val]=1
		}
		if res[val-1]!=0{
			res[val-1]+=1
		}else{
			res[val-1]=1
		}
	}
	fmt.Printf("\n %+v \n", res)
	var m int=0
    for _,val:= range res{
        if val>m{
            m=val
        }
    }
    fmt.Println("solution :",m)
}

func main(){
	a:=[]int{4,6,5,3,3,1}
	solution(a)
	a=[]int{1,2,2,3,1,2}
	solution(a)
}