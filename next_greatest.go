package main

import (
	"fmt"
	"sort"
	"strconv"
)

func sortedNumber(number int) int{
	var arry []int
	for number>0{
		arry=append(arry,number%10)
		number/=10
	}
	sort.Ints(arry)
	num:=0
	for _,e := range arry{
		num=(num*10)+e
	}
	return num
}

func reverseSortedNumber(number int) int{
	var arry []int
	for number>0{
		arry=append(arry,number%10)
		number/=10
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arry)))
	num:=0
	for _,e := range arry{
		num=(num*10)+e
	}
	return num
}

func getNumber(num string) int{
	i,err:=strconv.Atoi(num)
	if err!=nil{
		fmt.Println(err)
	}
	return i
}

func getNumberString(num int) string{
	i:=strconv.Itoa(num)
	return i
}

func genNumber(nums string, pos int) int{
	nums=nums[:pos-1]+nums[len(nums)-1:len(nums)]+nums[pos:len(nums)-1]+nums[pos-1:pos]
	nums=nums[:pos]+getNumberString(sortedNumber(getNumber(nums[pos:])))
	return getNumber(nums)
}

func nextGreatest(nums string) int{
	i:=getNumber(nums)
	answer:=0
	if i<reverseSortedNumber(i){
		if i==sortedNumber(i){
			i:=getNumber(nums[:len(nums)-2]+nums[len(nums)-1:]+nums[len(nums)-2:len(nums)-1])
			return i
		}else{
			for i:=1;i<len(nums);i++{
				num:=getNumber(nums[i:])
				if num==reverseSortedNumber(num){
					answer=genNumber(nums,i)
					break
				}
			}
			return answer
		}
	}else{
		return 0
	}
}

func main(){
	input:=[]string{"218765","4321","1234", "534976"}
	for _,e:=range input{
		fmt.Println("The number is - ", e)
		temp:=nextGreatest(e)
		if temp!=0{
			fmt.Println("Answer is - ", temp)
		}else{
			fmt.Println("Not possible")
		}
	}
}