# Modul: `io`

## Ringkasan
Package `io` menyediakan antarmuka dasar untuk Input/Output.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `io` menyediakan abstraksi tingkat rendah yang sangat kuat untuk menangani data yang mengalir (streaming data). Daripada memuat seluruh file ke memori sekaligus, `io` memungkinkan kita memproses data sepotong demi sepotong.

**Tujuan dan Fungsi Utama:**
*   **Abstraksi Reader dan Writer:** Mendefinisikan interface `io.Reader` (sumber data yang bisa dibaca) dan `io.Writer` (tujuan data yang bisa ditulis). Semua hal di Go yang "bisa dibaca" (file, network connection, request body HTTP) biasanya mengimplementasikan `io.Reader`.
*   **Menyalin Data:** Fungsi `io.Copy` memungkinkan penyalinan data secara efisien dari sebuah Reader ke sebuah Writer tanpa menghabiskan memori (contoh: menyalin file ke HTTP response).
*   **Utilitas I/O:** Membaca semua data hingga akhir (`io.ReadAll`), atau membatasi ukuran data yang dibaca (`io.LimitReader`).

**Mengapa menggunakan `io`?**
Kekuatan Go terletak pada interface. Dengan menggunakan `io.Reader` dan `io.Writer`, kode Anda menjadi sangat fleksibel. Sebuah fungsi kompresi data (*gzip*) misalnya, bisa mengambil input dari File, HTTP Request, ataupun Memory Buffer, selama semuanya mematuhi kontrak `io.Reader`.

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/io](https://pkg.go.dev/io)
