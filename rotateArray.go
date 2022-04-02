package main

import (
	"fmt"
)

func rotateArray(ary []int, rot int) []int{
	l:=len(ary)
	return append(ary[l-rot:],ary[:l-rot]...)
}

func rotate(ary []int, k int){
	l:=len(ary)
	for i:=0;i<k;i++{
		temp:=ary[l-1]
		for j:=l-1;j>0;j--{
			ary[j]=ary[j-1]
		}
		ary[0]=temp
	}
}

func main(){
	ary:=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,15,17,19,20,34}
	rotate(ary,5)
	fmt.Println(ary)
	bst(ary,4)
}

func bstOnRotated(ary []int,key int){
	if len(ary)==0{
		return
	}
	m:=len(ary)/2
	if ary[m]==key{
		fmt.Println(ary[m])
		return
	}else{
		if ary[len(ary)-1]>ary[0]{
			if ary[m]>key{
				bst(ary[:m],key)
			}else{
				bst(ary[m+1:],key)
			}
		}else{
			bst(ary[m+1:],key)
			bst(ary[:m],key)
		}
	}
}