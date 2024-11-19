package main

import "fmt"

func occurenceOfString(char string) int{
	count:=0
	for i:=0; i<len(char); i++{
		for j:=i+1; j<len(char); j++{
			if char[i]==char[j]{
				count++
				break
			}
		}
	}
	return count
}
func main(){
	ch:="melesmm"
	fmt.Println(occurenceOfString(ch))
}