# Modul: `context`

## Ringkasan
Package `context` mengelola siklus hidup, pembatalan, dan batas waktu untuk proses-proses konkuren.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `context` membawa (carry) deadline, sinyal pembatalan (cancellation signals), dan variabel khusus dari satu fungsi ke fungsi lainnya dalam sebuah *chain* (rantai) panggilan API.

**Tujuan dan Fungsi Utama:**
*   **Pembatalan (Cancellation):** Jika klien memutus koneksi HTTP (menutup browser sebelum selesai loading), server perlu tahu agar bisa menghentikan query database di baliknya untuk menghemat *resource*. `context.WithCancel` digunakan untuk ini.
*   **Batas Waktu (Timeout/Deadline):** Anda membuat API call ke layanan luar. Jika layanan itu lambat merespons dalam 5 detik, Anda ingin membatalkannya dan mengembalikan error. `context.WithTimeout` menangani kasus ini.
*   **Membawa Data (Value):** Menyisipkan data kecil yang berkaitan dengan *request* (seperti ID pengguna dari token autentikasi atau Request ID untuk log) menembus banyak lapisan fungsi secara aman.

**Mengapa menggunakan `context`?**
Di arsitektur *microservice* modern, sebuah *request* biasanya melewati banyak sistem (API -> Database -> Cache -> API Eksternal). Context adalah "tali" yang mengikat semua proses tersebut sehingga jika sesuatu terjadi (timeout atau klien *disconnect*), seluruh proses turunannya bisa segera dihentikan dengan rapi.

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/context](https://pkg.go.dev/context)
