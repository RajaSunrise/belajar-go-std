package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--- 1. Waktu Saat Ini dan Format ---")
	now := time.Now()
	fmt.Println("Waktu saat ini:", now)

	// Format waktu dalam Go menggunakan waktu patokan spesifik:
	// Mon Jan 2 15:04:05 MST 2006
	fmt.Println("Format RFC3339:", now.Format(time.RFC3339))
	fmt.Println("Format Custom (YYYY-MM-DD HH:MM:SS):", now.Format("2006-01-02 15:04:05"))

	fmt.Println("\n--- 2. Parsing Waktu ---")
	dateStr := "2024-08-17 10:00:00"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		fmt.Println("Gagal parsing:", err)
	} else {
		fmt.Println("Hasil Parsing:", parsedTime)
	}

	fmt.Println("\n--- 3. Operasi Waktu dan Durasi (Duration) ---")
	// Menambahkan durasi
	tomorrow := now.Add(24 * time.Hour)
	fmt.Println("Besok pada waktu yang sama:", tomorrow)

	// Mengurangi waktu (menggunakan nilai negatif pada Add atau Sub untuk antar time.Time)
	yesterday := now.Add(-24 * time.Hour)
	fmt.Println("Kemarin:", yesterday)

	// Durasi antara dua waktu
	durationDif := tomorrow.Sub(now)
	fmt.Println("Selisih Besok dan Sekarang:", durationDif.Hours(), "jam")

	// time.Since dan time.Until
	fmt.Println("Waktu berlalu sejak kemarin (Since):", time.Since(yesterday))
	fmt.Println("Waktu hingga besok (Until):", time.Until(tomorrow))

	fmt.Println("\n--- 4. Zona Waktu (Location) ---")
	// Load zona waktu WIB (Asia/Jakarta)
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Gagal meload timezone:", err)
	} else {
		timeWIB := now.In(loc)
		fmt.Println("Waktu di Asia/Jakarta (WIB):", timeWIB.Format("2006-01-02 15:04:05 MST"))
	}

	utcTime := now.UTC()
	fmt.Println("Waktu dalam UTC:", utcTime.Format("2006-01-02 15:04:05 MST"))

	fmt.Println("\n--- 5. Timer dan Ticker ---")
	// Timer mengeksekusi sesuatu SEKALI setelah jeda waktu tertentu
	fmt.Println("Menunggu 1 detik menggunakan time.Timer...")
	timer := time.NewTimer(1 * time.Second)
	<-timer.C // block sampai channel mengirimkan sinyal
	fmt.Println("Timer selesai!")

	// time.After adalah shorthand untuk membuat timer dan menunggu channelnya
	fmt.Println("Menunggu 500 milidetik menggunakan time.After...")
	<-time.After(500 * time.Millisecond)
	fmt.Println("Selesai menunggu.")

	// Ticker mengeksekusi sesuatu BERKALI-KALI dengan interval tertentu
	fmt.Println("Memulai ticker setiap 300 milidetik (akan berdetak 3 kali)...")
	ticker := time.NewTicker(300 * time.Millisecond)

	// Gunakan loop dan batasi detakan
	count := 0
	for t := range ticker.C {
		fmt.Println("Tick pada:", t.Format("15:04:05.000"))
		count++
		if count >= 3 {
			ticker.Stop() // Penting untuk menghentikan ticker guna menghindari memory leak
			break
		}
	}
	fmt.Println("Ticker dihentikan.")
}
