package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasSuffix("hira dazzle", "dazzle"))       // Output: true
	fmt.Println(strings.HasSuffix("vini emerald green", "green")) // Output: false
	fmt.Println(strings.HasSuffix("meia hime", "hime"))           // Output: true
	fmt.Println(strings.HasSuffix("idol local", "lo"))            // Output: false
}
