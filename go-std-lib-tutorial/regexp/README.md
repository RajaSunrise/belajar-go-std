# Modul: `regexp`

## Ringkasan
Package `regexp` mengimplementasikan pencarian dan manipulasi teks berbasis ekspresi reguler (Regular Expression).

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `regexp` (Regular Expression) memberikan bahasa *pattern-matching* (pencocokan pola) tingkat lanjut yang sangat cepat dalam mencari atau memvalidasi format string tertentu. Go menggunakan sintaks RE2, yang dirancang untuk mencegah eksekusi berlebihan akibat pencocokan yang tak berujung (linear time execution).

**Tujuan dan Fungsi Utama:**
*   **Kompilasi Pola:** Fungsi `regexp.MustCompile` mengonversi string pola pencarian (`^[0-9]+$`) menjadi objek Regexp yang dapat dieksekusi efisien. Digunakan fungsi `MustCompile` (alih-alih `Compile`) agar program langsung berhenti (panic) jika regex-nya ditulis salah.
*   **Validasi:** Menggunakan `.MatchString()` untuk sekadar mengonfirmasi apakah format sebuah teks sudah sesuai kriteria (contoh validasi Email, No. Telepon, KTP, dll).
*   **Pencarian Lanjut:** Mencari seluruh kata yang sesuai pola di dalam dokumen yang panjang menggunakan `.FindAllString()`.
*   **Sensor atau Penggantian:** Mengganti bagian dari kalimat yang *match* dengan pola tertentu menggunakan nilai baru melalui `.ReplaceAllString()` (misalnya menyembunyikan/sensor password atau data pribadi di dalam *log* file).

**Mengapa menggunakan `regexp`?**
Penyortiran teks dengan package `strings` saja tidak selalu cukup karena terkadang format data yang dicari dinamis (contoh: "Carikan saya kalimat apapun yang diawali Rp. dan diikuti 4 hingga 8 angka"). `regexp` adalah solusi baku untuk masalah tersebut.

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/regexp](https://pkg.go.dev/regexp)
