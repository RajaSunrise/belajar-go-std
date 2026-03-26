package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Menggunakan io.Reader (membaca dari string)
	reader := strings.NewReader("Ini adalah contoh penggunaan io.Reader\n")

	// Menggunakan io.Writer (menulis ke os.Stdout)
	_, err := io.Copy(os.Stdout, reader)
	if err != nil {
		log.Fatal(err)
	}

	// io.ReadAll untuk membaca seluruh isi dari Reader
	reader2 := strings.NewReader("Membaca semua data sekaligus.")
	data, err := io.ReadAll(reader2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hasil ReadAll:", string(data))
}
