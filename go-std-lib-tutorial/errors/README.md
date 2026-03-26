# Modul: `errors`

## Ringkasan
Package `errors` mengimplementasikan fungsi-fungsi untuk membuat dan memanipulasi nilai error.

## Penjelasan Lengkap (Fungsi & Tujuan)
Di dalam bahasa Go, error diperlakukan layaknya nilai biasa (value) yang dikembalikan oleh fungsi. Package `errors` memberikan fungsi-fungsi standar untuk membangun dan mengelola *error* ini.

**Tujuan dan Fungsi Utama:**
*   **Membuat Error Dasar:** Membuat instance error baru yang sederhana berisi pesan teks menggunakan `errors.New()`.
*   **Error Wrapping (Membungkus):** Saat sebuah error berpindah dari fungsi terdalam ke atas, Anda bisa menambahkan konteks ("gagal membaca file: <error asli>") tanpa menghilangkan jejak error aslinya. (Walaupun *wrapping* biasanya dilakukan lewat `fmt.Errorf("%w", err)`).
*   **Mengecek Jenis Error (Is):** Fungsi `errors.Is(err, TargetErr)` sangat berguna untuk mengecek apakah sebuah rantai error *mengandung* error tertentu yang Anda cari. Ini menggantikan cara lama menggunakan `==`.
*   **Mengekstrak Tipe Error (As):** Fungsi `errors.As` memungkinkan Anda mengambil instance error kustom (beserta field tambahan di dalamnya) dari sebuah *wrapped error*.

**Mengapa menggunakan `errors`?**
Penanganan error yang baik adalah inti dari kestabilan sistem Go. Penggunaan fungsi `errors.Is` dan `errors.As` (diperkenalkan sejak Go 1.13) adalah standar wajib di Go modern agar aplikasi bisa merespons error dengan spesifik (misal: "Jika error adalah NotFound, return 404, jika timeout, return 500").

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/errors](https://pkg.go.dev/errors)
