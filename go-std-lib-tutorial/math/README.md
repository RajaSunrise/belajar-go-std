# Modul: `math`

## Ringkasan
Package `math` menyediakan konstanta dasar dan fungsi matematika.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `math` berisi banyak fungsi dan konstanta yang berhubungan dengan matematika dasar dan lanjutan. Hampir semua fungsi di package `math` bekerja pada nilai tipe `float64` untuk memberikan presisi tertinggi yang didukung CPU secara *native*.

**Tujuan dan Fungsi Utama:**
*   **Konstanta Penting:** Menyediakan nilai pasti untuk konstanta matematika seperti Pi (`math.Pi`) dan e (`math.E`).
*   **Fungsi Dasar:** Menghitung nilai absolut (`math.Abs`), akar kuadrat (`math.Sqrt`), atau pemangkatan (`math.Pow`).
*   **Pembulatan:** Membulatkan angka pecahan ke atas (`math.Ceil`), ke bawah (`math.Floor`), atau ke bilangan bulat terdekat (`math.Round`).
*   **Trigonometri:** Fungsi sinus, kosinus, tangen, dll (`math.Sin`, `math.Cos`, `math.Tan`).

**Mengapa menggunakan `math`?**
Anda membutuhkannya ketika mengembangkan aplikasi yang membutuhkan kalkulasi presisi tinggi, seperti permainan (untuk menghitung jarak antar titik 2D/3D), simulasi fisika, algoritma statistik, atau sekadar membulatkan angka desimal dalam sebuah perhitungan harga pada aplikasi keuangan (e-commerce).

## Daftar Fungsi Umum dan Cara Penggunaannya

### 1. Pembulatan Angka (Rounding)
*   **`math.Round(x)`**: Membulatkan ke integer terdekat. (0.5 ke atas, 0.4 ke bawah).
*   **`math.Ceil(x)`**: Membulatkan selalu ke atas (Ceiling/Plafon).
*   **`math.Floor(x)`**: Membulatkan selalu ke bawah (Floor/Lantai).
```go
angka := 3.14159

fmt.Println("Round:", math.Round(angka)) // Output: 3
fmt.Println("Ceil:", math.Ceil(angka))   // Output: 4
fmt.Println("Floor:", math.Floor(angka)) // Output: 3

angka2 := 3.8
fmt.Println("Round:", math.Round(angka2)) // Output: 4
```

### 2. Akar dan Pangkat (Roots & Exponents)
*   **`math.Sqrt(x)`**: Menghitung akar kuadrat dari `x`. (Hanya menerima angka positif).
*   **`math.Pow(x, y)`**: Menghitung `x` pangkat `y` ($x^y$).
```go
// Akar kuadrat dari 25
akar := math.Sqrt(25) // Output: 5

// 2 pangkat 10
pangkat := math.Pow(2, 10) // Output: 1024
```

### 3. Nilai Maksimum, Minimum, dan Absolut
Sangat berguna dalam perulangan untuk mencari nilai terbesar atau terkecil.
*   **`math.Max(x, y)`**: Mengembalikan nilai yang lebih besar antara dua angka.
*   **`math.Min(x, y)`**: Mengembalikan nilai yang lebih kecil antara dua angka.
*   **`math.Abs(x)`**: Mengembalikan nilai absolut (menghilangkan tanda minus pada angka negatif).
```go
terbesar := math.Max(10.5, 50.2) // Output: 50.2
terkecil := math.Min(-5, 2)      // Output: -5

jarak := math.Abs(-150) // Output: 150 (jarak tidak pernah negatif)
```

### 4. Konstanta Matematika Dasar
Nilai presisi tinggi dari konstanta yang sering digunakan di matematika dan fisika.
```go
luasLingkaran := math.Pi * math.Pow(jariJari, 2)
```


## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
