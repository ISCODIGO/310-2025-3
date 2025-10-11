package main

import "fmt"

func main() {
	// range
	fmt.Println("Range")
	for h := range 5 {
		fmt.Println(h)
	}

	fmt.Println("For (c)")
	// for
	for i := 5; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("While (c)")
	// while (C)
	j := 11
	for j <= 20 {
		fmt.Println(j)
		j++
	}

	fmt.Println("While")
	k := 21
	for true {
		fmt.Println(k)
		k++

		if k > 25 {
			break
		}
	}
}
