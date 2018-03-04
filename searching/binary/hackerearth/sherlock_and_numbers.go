package main

import (
	"fmt"
	)


func binarySearch(array  []int, target int) int {
	left := 0
	right := len(array) - 1
	for left<=right{
		mid := right+(left-right)/2
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

func main(){
	var T int //No of Test cases
	var N int //Nth number
	var K int //K number of array elements
	var P int //Pth smallest element to find
	var A []int
	fmt.Scan(&T)
	fmt.Println(	)
	//fmt.Println("Total Test Case: ",T)
	for i:=0;i<T;i++{
		var B []int
		length:=0
		fmt.Scan(&N)
		fmt.Scan(&K)
		fmt.Scan(&P)
		//fmt.Println("N, K, P: ",N,K,P)
		A=make([]int,K)
		for j:=0;j<K;j++{
			fmt.Scan(&A[j])
		}
		for k:=1;k<=N;k++{
			//fmt.Println( binarySearch(A,k))
			if binarySearch(A,k)==-1{
				B=append(B,k)
				length++
			}
		}
		if P<=length{
			fmt.Println(B[P-1])
		} else {
			fmt.Println(-1)
		}
		//fmt.Println()
		//fmt.Println(" A: ",A," length: ",length, " B: ",B)
		//fmt.Println("B[p]: ", B[P]," B len: ", len(B))
	}
}