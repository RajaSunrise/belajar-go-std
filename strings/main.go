package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "belajar golang standard library"

	// Mengubah ke huruf kapital
	fmt.Println("ToTitle:", strings.ToTitle(text))
	fmt.Println("ToUpper:", strings.ToUpper(text))

	// Mengecek apakah string mengandung substring tertentu
	fmt.Println("Contains 'golang':", strings.Contains(text, "golang"))

	// Menghitung jumlah kemunculan
	fmt.Println("Count 'a':", strings.Count(text, "a"))

	// Memisahkan string menjadi slice
	splitted := strings.Split(text, " ")
	fmt.Printf("Split: %#v\n", splitted)

	// Menggabungkan slice menjadi string
	joined := strings.Join(splitted, "-")
	fmt.Println("Join:", joined)
}
