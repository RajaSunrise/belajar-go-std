# Modul: `fmt`

## Ringkasan
Package `fmt` mengimplementasikan I/O berformat.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `fmt` adalah salah satu package yang paling sering digunakan di Go. Package ini menyediakan fungsi-fungsi untuk memformat string dan melakukan Input/Output (I/O) dasar yang terinspirasi oleh fungsi `printf` dan `scanf` dalam bahasa C.

**Tujuan dan Fungsi Utama:**
*   **Mencetak ke Console:** Fungsi seperti `fmt.Println`, `fmt.Printf`, dan `fmt.Print` digunakan untuk mencetak teks atau nilai variabel ke standar output (layar).
*   **Memformat String:** Fungsi seperti `fmt.Sprintf` digunakan untuk membuat string baru berdasarkan template dan argumen yang diberikan (tanpa mencetaknya langsung).
*   **Membaca Input:** Fungsi seperti `fmt.Scanf` dan `fmt.Scanln` dapat digunakan untuk membaca input dari pengguna.

**Mengapa menggunakan `fmt`?**
Karena hampir setiap program perlu memberikan *output* informasi kepada pengguna atau *developer* untuk keperluan *debugging*. `fmt` menyediakan cara yang sangat mudah dan fleksibel menggunakan *verbs* (seperti `%s` untuk string, `%d` untuk integer, `%v` untuk nilai default, dan `%+v` untuk struct dengan nama field).

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/fmt](https://pkg.go.dev/fmt)
