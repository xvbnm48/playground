package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Count("hira dazzle", "ra"))
	// hasilnya adalah 1 karena "ra" hanya muncul 1 kali
	fmt.Println(strings.ContainsAny("hira dazzle", "ale"))
	fmt.Println(strings.HasPrefix("hira dazzle", "hira"))
	// hasilnya adalah true karena "hira" adalah awalan dari "hira dazzle"
}
