# Modul: `os`

## Ringkasan
Package `os` menyediakan antarmuka untuk fungsionalitas sistem operasi.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `os` menyediakan antarmuka (interface) independen-platform untuk berinteraksi dengan sistem operasi di bawahnya (Windows, Linux, macOS, dll). Desainnya mirip dengan Unix, namun fungsi-fungsinya disamarkan sehingga bisa berjalan mulus di berbagai OS.

**Tujuan dan Fungsi Utama:**
*   **Manipulasi File dan Direktori:** Membuka, membuat, menghapus, atau membaca file (`os.Open`, `os.Create`, `os.Remove`, `os.Mkdir`).
*   **Environment Variables:** Mengambil atau mengatur variabel environment sistem (`os.Getenv`, `os.Setenv`).
*   **Informasi Proses dan Sistem:** Mendapatkan argumen command-line (`os.Args`), nama host (`os.Hostname`), dan direktori kerja saat ini (`os.Getwd`).
*   **Keluar dari Program:** Menghentikan program secara paksa dengan kode status tertentu (`os.Exit`).

**Mengapa menggunakan `os`?**
Jika aplikasi Anda perlu berinteraksi langsung dengan sistem tempat ia berjalan—seperti membaca konfigurasi dari *environment variables*, menulis log ke dalam sebuah file, atau membaca file yang diunggah pengguna—package `os` adalah pintu masuk utamanya.

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/os](https://pkg.go.dev/os)
