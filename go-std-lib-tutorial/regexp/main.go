package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Compile regex (Gunakan MustCompile untuk inisialisasi global agar panic jika salah)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	emails := []string{
		"user@example.com",
		"invalid-email",
		"admin@domain.co.id",
	}

	for _, email := range emails {
		// MatchString mengembalikan boolean
		match := re.MatchString(email)
		fmt.Printf("%s -> Valid? %v\n", email, match)
	}

	// Mencari dan mengganti text
	text := "Nomor telepon saya adalah 0812-3456-7890 dan 0899-8888-7777"
	rePhone := regexp.MustCompile(`\d{4}-\d{4}-\d{4}`)

	// FindAllString mencari semua kecocokan
	phones := rePhone.FindAllString(text, -1)
	fmt.Println("\nNomor telepon yang ditemukan:", phones)

	// ReplaceAllString mengganti kecocokan
	anonymized := rePhone.ReplaceAllString(text, "[SENSOR]")
	fmt.Println("Teks disensor:", anonymized)
}
