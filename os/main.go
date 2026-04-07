package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("--- 1. Informasi Sistem & Lingkungan ---")
	// Mendapatkan variabel environment
	path := os.Getenv("PATH")
	fmt.Println("Variabel env PATH:", path)

	// Mendapatkan working directory saat ini (tempat eksekusi program)
	pwd, err := os.Getwd()
	if err != nil {
		log.Println("Gagal mendapatkan working directory:", err)
	} else {
		fmt.Println("Working Directory (Getwd):", pwd)
	}

	// Mendapatkan nama host
	hostname, _ := os.Hostname()
	fmt.Println("Hostname sistem:", hostname)

	fmt.Println("\n--- 2. Manipulasi File Dasar ---")
	fileName := "contoh_file.txt"
	// Membuat file baru dengan os.Create
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// Pastikan file dihapus dan ditutup setelah selesai
	defer func() {
		file.Close()
		os.Remove(fileName)
		fmt.Printf("File %s telah dihapus.\n", fileName)
	}()

	// Menulis data ke file
	file.WriteString("Halo dari package os!\nBaris kedua di sini.\n")
	fmt.Printf("Berhasil membuat dan menulis ke file: %s\n", fileName)

	// Membaca file menggunakan os.ReadFile
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("Gagal membaca file:", err)
	} else {
		fmt.Printf("Isi dari %s:\n%s", fileName, string(data))
	}

	fmt.Println("\n--- 3. Informasi File (os.Stat) dan Pengecekan Keberadaan File ---")
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File tidak ditemukan!")
		} else {
			log.Println("Error saat os.Stat:", err)
		}
	} else {
		fmt.Printf("Info File -> Nama: %s, Ukuran: %d bytes, Mode (Permission): %v, Direktori?: %t\n",
			fileInfo.Name(), fileInfo.Size(), fileInfo.Mode(), fileInfo.IsDir())
	}

	// Mengecek file yang tidak ada
	_, errNotFound := os.Stat("file_palsu_ini_tidak_ada.abc")
	if os.IsNotExist(errNotFound) {
		fmt.Println("Pengecekan IsNotExist: Benar, file_palsu_ini_tidak_ada.abc tidak ada.")
	}

	fmt.Println("\n--- 4. Manipulasi Direktori ---")
	dirName := "contoh_direktori/sub_direktori"

	// Membuat direktori beserta parent-nya dengan permission 0755
	err = os.MkdirAll(dirName, 0755)
	if err != nil {
		log.Println("Gagal membuat direktori:", err)
	} else {
		fmt.Println("Berhasil membuat direktori:", dirName)
	}

	// Membersihkan direktori
	defer func() {
		os.RemoveAll("contoh_direktori")
		fmt.Println("Direktori contoh_direktori dan isinya telah dihapus.")
	}()

	// Membaca isi direktori menggunakan os.ReadDir (menggantikan ioutil.ReadDir)
	entries, err := os.ReadDir(".") // Membaca current directory
	if err != nil {
		log.Println("Gagal membaca direktori saat ini:", err)
	} else {
		fmt.Println("Daftar 3 file/folder pertama di current directory:")
		count := 0
		for _, entry := range entries {
			fmt.Printf("- %s (Direktori: %t)\n", entry.Name(), entry.IsDir())
			count++
			if count >= 3 {
				break
			}
		}
	}
}
