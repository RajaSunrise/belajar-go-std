# Modul: `os/exec`

## Ringkasan
Package `os/exec` digunakan untuk menjalankan perintah (commands) eksternal di sistem operasi.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `os/exec` membungkus (wrap) fungsi-fungsi dari *low-level operating system* sehingga program Go Anda dapat memicu (execute), mengontrol, dan membaca *output* dari aplikasi atau *shell script* lain di komputer/server tersebut.

**Tujuan dan Fungsi Utama:**
*   **Memanggil Perintah/Aplikasi Eksternal:** Menyiapkan perintah sistem (seperti `ls`, `ping`, `ffmpeg`, atau skrip python) menggunakan `exec.Command`.
*   **Membaca Output:** Mengambil semua *output text* (stdout) dari aplikasi yang dijalankan menggunakan fungsi `.Output()` atau `.CombinedOutput()`.
*   **Mengontrol Input/Output Berjalan (Streaming):** Menyambungkan *pipe* sehingga program Go bisa "menyuapi" input (*stdin*) perlahan-lahan ke aplikasi eksternal, atau membaca responnya secara *realtime* (*stdoutpipe*).
*   **Mengontrol Siklus Hidup Proses:** Memulai proses di belakang layar (`.Start()`) tanpa memblokir Go, lalu nanti ditunggu (`.Wait()`) atau dibunuh jika memakan waktu terlalu lama.

**Mengapa menggunakan `os/exec`?**
Sangat berguna untuk otomatisasi *deployment*, membuat *CLI tools*, atau menggunakan pustaka pihak ketiga yang bukan ditulis di Go (misal: memanggil aplikasi *FFmpeg* dari dalam Go untuk mengonversi video).

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/os/exec](https://pkg.go.dev/os/exec)
