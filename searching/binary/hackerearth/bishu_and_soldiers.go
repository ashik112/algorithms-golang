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
		if target <= array[mid]{
			pos = mid
			right = mid - 1
		}else  {
			left = mid + 1
		}
	}
	return pos
}

func dup(array []int,index int,value int)int {
	temp:=index
	for i:=index+1;i<len(array);i++{
		if value!=array[i]{
			return temp
		}else{
			temp=i
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
	SumPower:= make([]int, N)
	SumPower[0]=SoldierPowers[0]
	for i:=1;i<N;i++ {
		SumPower[i]=SumPower[i-1]+SoldierPowers[i]

	}

	fmt.Scan(&Q)


	for ; 0 < Q; Q-- {
		fmt.Scan(&BishuPower)
		result := doBinarySearch(SoldierPowers, BishuPower)
		correctIndex:=dup(SoldierPowers,result,SoldierPowers[result])
		fmt.Println(correctIndex+1,SumPower[correctIndex])
	}
}
/*package main

import (
	"fmt"
	"sort"
	"os"
	"bufio"
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
		if target <= array[mid]{
			pos = mid
			right = mid - 1
		}else  {
			left = mid + 1
		}
	}
	return pos
}

func dup(array []int,index int,value int)int {
	temp:=index
	for i:=index+1;i<len(array);i++{
		if value!=array[i]{
			return temp
		}else{
			temp=i
		}
	}
	return index
}

func doSum(array []int) int {
	total:=0
	for i:=0;i<len(array);i++{
		total=total+array[i]
	}
	return total
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
	SumPower:= make([]int, N)
	SumPower[0]=SoldierPowers[0]
	for i:=1;i<N;i++ {
		SumPower[i]=SumPower[i-1]+SoldierPowers[i]

	}

	fmt.Scan(&Q)


	f, err := os.Create("output.txt")
	//check(err)
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	//fmt.Fprintln(w, SoldierPowers)
	//fmt.Fprintln(w, SumPower)
	for ; 0 < Q; Q-- {
		fmt.Scan(&BishuPower)
		result := doBinarySearch(SoldierPowers, BishuPower)
		correctIndex:=dup(SoldierPowers,result,SoldierPowers[result])
		fmt.Fprint(w, "Bisu Power: ",BishuPower," Index: ",result," Value: ",SoldierPowers[result],"  ","length: Soldier->",len(SoldierPowers)," Correct Index: ",dup(SoldierPowers,result,SoldierPowers[result])," :: ")
		//fmt.Fprintln(w,result+1,SumPower[dup(SoldierPowers,result,SoldierPowers[result])])
		fmt.Fprintln(w,correctIndex+1,SumPower[correctIndex])
	}
	fmt.Println(SumPower[9899])
	w.Flush()
}*/
/*
package main

import (
	"fmt"
	"sort"
	"os"
	"bufio"
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
		if target <= array[mid]{
			pos = mid
			right = mid - 1
		}else  {
			left = mid + 1
		}
	}
	return pos
}

func dup(array []int,index int,value int)int {
	temp:=index
	for i:=index+1;i<len(array);i++{
		if value!=array[i]{
			return temp
		}else{
			temp=i
		}
	}
	return index
}

func doSum(array []int) int {
	total := 0
	for _, val := range array {
		total += val
	}
	return total
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
	SumPower:= make([]int, N)
	SumPower[0]=SoldierPowers[0]
	for i:=1;i<N;i++ {
		SumPower[i]=SumPower[i-1]+SoldierPowers[i]
	}

	fmt.Scan(&Q)


	f, err := os.Create("output.txt")
	//check(err)
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	fmt.Fprintln(w, SoldierPowers)
	fmt.Fprintln(w, SumPower)
	for ; 0 < Q; Q-- {
		fmt.Scan(&BishuPower)
		result := doBinarySearch(SoldierPowers, BishuPower)
		fmt.Fprint(w, "Bisu Power: ",BishuPower," Index: ",result," Value: ",SoldierPowers[result],"  ","length: Soldier->",len(SoldierPowers)," Correct Index: ",dup(SoldierPowers,result,SoldierPowers[result])," :: ")
		fmt.Fprintln(w,result+1,SumPower[dup(SoldierPowers,result,SoldierPowers[result])])
	}
	w.Flush()
}*/
