# Modul: `time`

## Ringkasan
Package `time` menyediakan fungsionalitas untuk mengukur dan menampilkan waktu.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `time` adalah pustaka bawaan untuk segala sesuatu yang berhubungan dengan tanggal, jam, durasi, dan jeda.

**Tujuan dan Fungsi Utama:**
*   **Mendapatkan Waktu Saat Ini:** Mengetahui jam berapa sistem berjalan (`time.Now()`).
*   **Durasi (Duration):** Menghitung rentang waktu antar dua kejadian (`time.Since`, `time.Until`) atau mendefinisikan durasi seperti "2 jam 30 menit" (`time.Hour`, `time.Minute`).
*   **Formatting Waktu:** Mengubah objek waktu menjadi string yang bisa dibaca ("01 Januari 2025") atau sebaliknya (parsing string menjadi waktu). Uniknya di Go, layout format waktu menggunakan tanggal referensi standar: `Mon Jan 2 15:04:05 MST 2006`.
*   **Penundaan (Sleep) dan Timer:** Membuat program berhenti sementara (`time.Sleep`) atau membuat eksekusi tertunda (`time.After`, `time.Ticker`).

**Mengapa menggunakan `time`?**
Hampir setiap aplikasi memerlukan pencatatan waktu: *timestamp* untuk log error, waktu pembuatan data di database (*created_at*, *updated_at*), mengatur masa kedaluwarsa sesi login (token JWT), atau menjalankan fungsi secara berkala (cron job).

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/time](https://pkg.go.dev/time)
