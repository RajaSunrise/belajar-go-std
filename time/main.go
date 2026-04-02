package main

import (
	"fmt"
	"time"
)

func main() {
	// Waktu saat ini
	now := time.Now()
	fmt.Println("Waktu saat ini:", now)

	// Format waktu (Go menggunakan waktu referensi khusus: Mon Jan 2 15:04:05 MST 2006)
	fmt.Println("Format (YYYY-MM-DD):", now.Format("2006-01-02"))

	// Menambahkan waktu
	tomorrow := now.Add(24 * time.Hour)
	fmt.Println("Besok:", tomorrow)

	// Parsing waktu
	parsedTime, _ := time.Parse("2006-01-02", "2024-08-17")
	fmt.Println("Hari Kemerdekaan 2024:", parsedTime)

	// Sleep / Jeda eksekusi
	fmt.Println("Tunggu 1 detik...")
	time.Sleep(1 * time.Second)
	fmt.Println("Selesai!")
}
