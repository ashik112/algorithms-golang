package main

import "fmt"

func main() {
	var T,N,K,P,a int
	fmt.Scanf("%d",&T)
	for ;T>0;T--{
		fmt.Scanf("%d",&N)
		fmt.Scanf("%d",&K)
		fmt.Scanf("%d",&P)
		for ;K>0;K-- {
			if fmt.Scanf("%d",&a);a <= P {
				P++
			}
		}
		if P > N {
			fmt.Println(-1)
			continue
		}
		fmt.Println(P)
	}
}
