package main

import (
	"fmt"
	"sort"
)

func doBinarySearch(array []int, target int) int {
	left := 0
	right := len(array) - 1
	pos := 0
	if target >= array[right] {
		return right
	}
	for right >= left {
		mid := right + (left-right)/2
		if target <= array[mid] {
			pos = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return pos
}

func checkDuplicates(array []int, index int, value int) int {
	temp := index
	for i := index + 1; i < len(array); i++ {
		if value != array[i] {
			return temp
		} else {
			temp = i
		}
	}
	return index
}

func main() {
	var N int          //total number of soldiers
	var Q int          //total number of rounds
	var BishuPower int // bishu's power (varies each round)
	fmt.Scan(&N)

	SoldierPowers := make([]int, N) // power of each soldier
	for i := 0; i < N; i++ {
		fmt.Scan(&SoldierPowers[i])
	}
	sort.Ints(SoldierPowers)
	SumPower := make([]int, N)
	SumPower[0] = SoldierPowers[0]
	for i := 1; i < N; i++ {
		SumPower[i] = SumPower[i-1] + SoldierPowers[i]

	}

	fmt.Scan(&Q)

	for ; 0 < Q; Q-- {
		fmt.Scan(&BishuPower)
		result := doBinarySearch(SoldierPowers, BishuPower)
		correctIndex := checkDuplicates(SoldierPowers, result, SoldierPowers[result])
		fmt.Println(correctIndex+1, SumPower[correctIndex])
	}
}
