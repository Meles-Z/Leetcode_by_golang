package main

import "fmt"

func linearSearch(num []int, target int) int{
	for i:=0; i<len(num)-1; i++{
		if num[i]==target{
			return i
		}
	}
	return -1
}
func main(){
	dd:=[]int{1,2,3,4}
	fmt.Println(linearSearch(dd, 3))
}