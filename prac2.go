package main

import "fmt"

func FindOutlier(interger []int){
	//var odd []int
	var even []int
	for _, v := range interger {
		if v %2 == 0 {
			even = append(even, v)
			fmt.Printf("even numbers: %v\n", even)
		}
	}
}

func main(){
	FindOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36})
}