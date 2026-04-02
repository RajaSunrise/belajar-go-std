package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Mengambil variabel environment
	path := os.Getenv("PATH")
	fmt.Println("Isi dari variabel PATH:", path)

	// Membuat file baru
	file, err := os.Create("contoh.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer os.Remove("contoh.txt") // Hapus file setelah selesai

	// Menulis ke file
	file.WriteString("Halo dari package os!\n")
	fmt.Println("Berhasil menulis ke file contoh.txt")
}
