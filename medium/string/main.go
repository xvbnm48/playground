package main

import "fmt"

func main() {
	// deklarasi variabel
	var idol = "hira dazzle"
	var nama = `Vini`
	var nama2 = `vini adalah idol dari hira dazzle,
	dengan julukan emerald green`
	fmt.Println(idol)
	fmt.Println(nama)
	fmt.Println(nama2)

	// unicode
	// result is 日本
	fmt.Println("\u65E5\u672C")

	// concatenation
	fmt.Println(nama, "idol dari", idol)

	// Substring
	nama4 := "Vini"
	// print substring from nama
	fmt.Println("hasil dari substring dari start index 1 dan end index 3 adalah ", string(nama4[1:3]))
}
