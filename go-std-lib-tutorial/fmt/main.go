package main

import "fmt"

func main() {
	name := "Go Developer"
	// Menggunakan verb %s untuk string dan %d untuk integer
	fmt.Printf("Halo, %s! Selamat datang di Go %d.\n", name, 1)

	// Menggabungkan string
	msg := fmt.Sprintf("Ini adalah string hasil format: %s", name)
	fmt.Println(msg)
}
