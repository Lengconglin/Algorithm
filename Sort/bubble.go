package main

import (
	"fmt"
)

func bubbleSort(items []int){
	var length = len(items)
	for j:=0;j<length-1;j++ {
		for i:=length-1;i>j;i-- {
			if items[i]<items[i-1] {
				items[i-1],items[i]=items[i],items[i-1]
			}
		}
	}
}

func main() {
	items:=[]int{4,200,3,9,5,6,143,1,300,500,30,1000}
	bubbleSort(items)
	fmt.Println(items)
}

