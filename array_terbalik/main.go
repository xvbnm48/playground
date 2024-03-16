package main

import "fmt"

func main() {
	var N int
	var tabel [100]int
	fmt.Println("masukkan N: ")
	fmt.Scanf("%d", &N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &tabel[i])
	}

	for i := N - 1; i >= 0; i-- {
		fmt.Printf("%d ", tabel[i])
	}
}
