package  main

import "fmt"

func  FindOutlier(integers []int) int{
	//Declare an empty array to store odd numbers from loop
	var odd []int
	var even []int
	//var result []int
	//var result []int
	for _, val := range integers {
		//check for even number
		if val %2 == 0 {
			even = append(even, val)
			//fmt.Println(even)
		}else {
			odd = append(odd, val)
		   }
		}
		//fmt.Println(even)
		//fmt.Println(odd)
	//fmt.Println(len(odd))
		if len(even) == 1{
			return even[0]
	}
	return odd[0]

}
	func main(){
		fmt.Println(FindOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))
	}