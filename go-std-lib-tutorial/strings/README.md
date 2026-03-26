# Modul: `strings`

## Ringkasan
Package `strings` mengimplementasikan fungsi-fungsi untuk memanipulasi string UTF-8.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `strings` adalah utilitas utama di Go untuk melakukan manipulasi, inspeksi, dan modifikasi pada tipe data teks (`string`). Semua string di Go pada dasarnya adalah *read-only slice of bytes*, dan package `strings` mengelola pengubahan tipe data tersebut secara efisien tanpa harus berhubungan langsung dengan indeks byte.

**Tujuan dan Fungsi Utama:**
*   **Pencarian Teks:** Mengecek apakah string mengandung teks tertentu (`strings.Contains`), mencari indeks awal kemunculan suatu kata (`strings.Index`), atau mengecek awalan/akhiran (`strings.HasPrefix`, `strings.HasSuffix`).
*   **Manipulasi Huruf:** Mengubah teks menjadi huruf kapital semua (`strings.ToUpper`) atau huruf kecil semua (`strings.ToLower`).
*   **Memecah dan Menggabungkan:** Memisahkan satu kalimat teks menjadi kumpulan kata atau baris array (`strings.Split`) dan sebaliknya, menggabungkan beberapa array teks menjadi satu kalimat panjang (`strings.Join`).
*   **Pembersihan:** Menghapus spasi (whitespace) tak terlihat atau karakter tertentu di awal/akhir string (`strings.Trim`, `strings.TrimSpace`).
*   **Penggantian:** Mengganti sebagian kata di dalam string dengan kata lain (`strings.Replace`, `strings.ReplaceAll`).

**Mengapa menggunakan `strings`?**
Manipulasi string adalah operasi yang terjadi di mana-mana. Saat Anda memvalidasi input dari pengguna form web, Anda seringkali perlu menghapus spasi tak sengaja (Trim) atau menormalkan nama mereka ke format huruf kapital tertentu. Package `strings` memiliki performa yang sangat dioptimalkan untuk berbagai macam skenario manipulasi teks harian tersebut.

## Daftar Fungsi Umum dan Cara Penggunaannya

### 1. Pencarian dan Pengecekan (Search & Check)
Berguna untuk logika validasi seperti memfilter daftar berdasarkan *search query*.
```go
teks := "golang is awesome"

// Apakah mengandung kata "go"?
ada := strings.Contains(teks, "go") // Output: true

// Apakah diawali dengan "golang"?
prefix := strings.HasPrefix(teks, "golang") // Output: true

// Apakah diakhiri dengan "bad"?
suffix := strings.HasSuffix(teks, "bad") // Output: false

// Berapa kali huruf 'a' muncul?
jumlah := strings.Count(teks, "a") // Output: 2
```

### 2. Memecah dan Menggabungkan (Split & Join)
Sering digunakan untuk membaca format data spesifik, misalnya URL Path atau CSV.
```go
data := "apel,jeruk,mangga,pisang"

// Memecah berdasarkan koma menjadi []string (Slice of String)
buah := strings.Split(data, ",")
// buah sekarang berisi: ["apel", "jeruk", "mangga", "pisang"]

// Menggabungkan kembali dengan pemisah ' | '
gabung := strings.Join(buah, " | ")
// Hasil: "apel | jeruk | mangga | pisang"
```

### 3. Merubah Ukuran Huruf (Casing)
Sangat berguna ketika ingin membandingkan dua input dari pengguna tanpa mempedulikan huruf besar kecil (Case Insensitive comparison).
```go
input := "   bElAjAr gOlang   "

// Bersihkan spasi di awal dan akhir
bersih := strings.TrimSpace(input) // "bElAjAr gOlang"

// Ubah ke huruf kecil semua
kecil := strings.ToLower(bersih) // "belajar golang"

// Ubah ke huruf besar semua
besar := strings.ToUpper(bersih) // "BELAJAR GOLANG"
```

### 4. Mencari dan Mengganti Teks (Replace)
```go
kalimat := "saya suka kopi, karena kopi itu enak."

// Mengganti semua kata "kopi" menjadi "teh" (-1 artinya ganti semua yang ditemukan)
baru := strings.Replace(kalimat, "kopi", "teh", -1)
// Atau bisa menggunakan ReplaceAll
baru = strings.ReplaceAll(kalimat, "kopi", "teh")
// Hasil: "saya suka teh, karena teh itu enak."
```


## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
