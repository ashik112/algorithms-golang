package main

import (
	"fmt"
)

func doBinarySearch(array  []int, target int) int {
	left := 0
	right := len(array) - 1
	for left<=right{
		mid := int((left+right)/2)
		if array[mid]==target{
			return mid
		} else if target>array[mid] {
			left = mid+1
		} else {
			right = mid -1
		}
	}
	return -1
}

func main()  {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	result := doBinarySearch(primes,17)
	if result!= -1{
		fmt.Println("Found ",primes[result]," at index:", result)
	} else {
		fmt.Println("Not found!")
	}
}
