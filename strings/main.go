package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("--- 1. Perubahan Case ---")
	text := "belajar golang standard library"
	fmt.Println("Original:", text)
	fmt.Println("ToTitle:", strings.ToTitle(text))
	fmt.Println("ToUpper:", strings.ToUpper(text))
	fmt.Println("ToLower:", strings.ToLower("HURUF KECIL SEMUA"))

	fmt.Println("\n--- 2. Pencarian dan Pengecekan ---")
	// Mengecek apakah string mengandung substring tertentu
	fmt.Println("Contains 'golang':", strings.Contains(text, "golang"))
	fmt.Println("ContainsAny 'z x c':", strings.ContainsAny(text, "z x c"))

	// HasPrefix dan HasSuffix
	fmt.Println("Dimulai dengan 'belajar' (HasPrefix)?", strings.HasPrefix(text, "belajar"))
	fmt.Println("Diakhiri dengan 'library' (HasSuffix)?", strings.HasSuffix(text, "library"))

	// Index dan LastIndex (mencari posisi kemunculan)
	fmt.Println("Posisi pertama 'a' (Index):", strings.Index(text, "a"))
	fmt.Println("Posisi terakhir 'a' (LastIndex):", strings.LastIndex(text, "a"))
	fmt.Println("Posisi 'xyz' (Index tidak ditemukan):", strings.Index(text, "xyz")) // Output -1 jika tidak ketemu

	fmt.Println("\n--- 3. Modifikasi dan Replace ---")
	// Menghitung jumlah kemunculan
	fmt.Println("Jumlah karakter 'a' (Count):", strings.Count(text, "a"))

	// Mengganti substring
	fmt.Println("Replace 'golang' dengan 'Go' (Replace 1x):", strings.Replace(text, "golang", "Go", 1))
	fmt.Println("ReplaceAll 'a' dengan '@':", strings.ReplaceAll(text, "a", "@"))

	fmt.Println("\n--- 4. Split dan Join ---")
	// Memisahkan string menjadi slice
	splitted := strings.Split(text, " ")
	fmt.Printf("Split: %#v\n", splitted)

	// Menggabungkan slice menjadi string
	joined := strings.Join(splitted, "-")
	fmt.Println("Join dengan '-':", joined)

	fmt.Println("\n--- 5. Trim (Membersihkan Spasi / Karakter) ---")
	dirtyText := "   halo dunia!   "
	fmt.Println("Original kotor:", dirtyText)
	fmt.Println("TrimSpace (membersihkan spasi ujung):", strings.TrimSpace(dirtyText))

	customDirty := "!!!peringatan!!!"
	fmt.Println("Trim '!' dari ujung:", strings.Trim(customDirty, "!"))
	fmt.Println("TrimLeft '!':", strings.TrimLeft(customDirty, "!"))
	fmt.Println("TrimPrefix '!!!':", strings.TrimPrefix(customDirty, "!!!"))

	fmt.Println("\n--- 6. Menggunakan strings.Builder ---")
	// strings.Builder jauh lebih efisien dalam alokasi memori dibandingkan konkatenasi dengan operator '+' di dalam loop
	var builder strings.Builder
	builder.WriteString("Ini ")
	builder.WriteString("adalah ")
	builder.WriteString("kalimat ")
	builder.WriteString("dari Builder.")

	// Mengambil hasilnya
	fmt.Println("Hasil Builder:", builder.String())
	fmt.Println("Panjang (Len) Builder:", builder.Len())
	fmt.Println("Kapasitas (Cap) Builder:", builder.Cap())
}
