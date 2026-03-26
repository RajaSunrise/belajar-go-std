package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// Menggabungkan path dengan aman (menghindari masalah slash/backslash antar OS)
	path := filepath.Join("users", "admin", "documents", "file.txt")
	fmt.Println("Joined Path:", path)

	// Mendapatkan direktori dan nama file
	dir := filepath.Dir(path)
	file := filepath.Base(path)
	fmt.Println("Directory:", dir)
	fmt.Println("Filename:", file)

	// Mendapatkan ekstensi file
	ext := filepath.Ext(path)
	fmt.Println("Extension:", ext)

	// Mengecek apakah path absolute
	fmt.Println("Is absolute?", filepath.IsAbs(path))

	absPath, _ := filepath.Abs("main.go")
	fmt.Println("Absolute path of main.go:", absPath)
}
