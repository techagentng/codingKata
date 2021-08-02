package main

import "fmt"

func FindOutlier(integer []int){
	//var name []int
	var even []int
	for _, v := range integer {
		if v %2 == 0 {
			even = append(even, v)
			fmt.Printf("even numbers: %v\n", even)
		}
	}
}

//func anotherLoop(){
//	var arr [10]int
//	for i=0; i < 10; i++{
//		fmt.Print(arr[i] + 10)
//	}
//}


func main(){
	FindOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36})
}