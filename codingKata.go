/*You are given an array (which will have a length of at least 3, but could be very large) containing integers.
The array is either entirely comprised of odd integers or entirely comprised of
even integers except for a single integer N. Write a method that takes the array as
an argument and returns this "outlier" N.*/
package  main

import "fmt"

func  FindOutlier(integers []int) int{
	//Declare an empty array to store odd numbers from loop
	var odd []int
	var even []int
	for _, val := range integers {
		//check for even number
		if val %2 == 0 {
			even = append(even, val)
		}else {
			odd = append(odd, val)
		}
	}
		if len(even) == 1{
			return even[0]
	}
	return odd[0]
}
	func main(){
		fmt.Println(FindOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))
	}