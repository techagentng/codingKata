package  main

import (
	//"bufio"
	//"fmt"
	//"io"
	//"os"
	//"strconv"
	//"strings"
	"fmt"
	"math"
)

/*
Given a square matrix, calculate the absolute difference between the sums of its diagonals.
For example, the square matrix  is shown below:

1 2 3
4 5 6
9 8 9

123 456 989
*/
//func diagonalDifference(arr [][]int32) int32{
//	  leftDiagonal :=int32(0); rightDiagonal :=int32(0)
//	for i := range integers {
//		leftDiagonal += arr[i][i]
//		fmt.Println(arr[i][i])
//		rightDiagonal += arr[i][len(arr)-1-i]
//	}
//	return int32(math.Abs(float64(leftDiagonal - rightDiagonal)))
//}
//
//func main(){
//
//	diagonalDifference([][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}})
//}

func diagonalDifference(arr [][]int32) int32 {
	leftDiagonal :=int32(0); rightDiagonal :=int32(0)
	for i:=  range arr {
		leftDiagonal +=arr[i][i]
		rightDiagonal +=arr[i][len(arr)-1-i]
	}
	return int32(math.Abs(float64(leftDiagonal - rightDiagonal)))
}
func main(){
	fmt.Println(diagonalDifference([][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}))
}