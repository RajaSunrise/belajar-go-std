package main

import (
	"errors"
	"fmt"
)

// Membuat error custom sebagai variabel global
var ErrNotFound = errors.New("data tidak ditemukan")

func findData(id int) (string, error) {
	if id == 0 {
		// Menggunakan fmt.Errorf dengan %w untuk membungkus (wrap) error
		return "", fmt.Errorf("pencarian id %d gagal: %w", id, ErrNotFound)
	}
	return "Data Rahasia", nil
}

func main() {
	// Error biasa
	err1 := errors.New("ini adalah error sederhana")
	fmt.Println("Error:", err1)

	// Membungkus dan mengecek error
	data, err := findData(0)
	if err != nil {
		fmt.Println("Terjadi kesalahan:", err)

		// errors.Is digunakan untuk mengecek apakah suatu error merupakan/membungkus error tertentu
		if errors.Is(err, ErrNotFound) {
			fmt.Println("=> Penanganan khusus: Menampilkan halaman 404 Not Found")
		}
	} else {
		fmt.Println("Data:", data)
	}

	// errors.As digunakan untuk mengekstrak tipe error custom tertentu (jika ada)
}
