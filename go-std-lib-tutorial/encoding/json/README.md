# Modul: `encoding/json`

## Ringkasan
Package `encoding/json` mengimplementasikan encoding dan decoding data format JSON.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `encoding/json` menyediakan sarana untuk mengubah (serialize) objek data Go (seperti *struct* atau *map*) menjadi teks berformat JSON, dan sebaliknya (deserialize).

**Tujuan dan Fungsi Utama:**
*   **Marshal:** Proses konversi dari struktur data di Golang (misal `struct`) menjadi JSON byte array. Fungsi utamanya adalah `json.Marshal`.
*   **Unmarshal:** Proses konversi kebalikan, dari teks/byte JSON ke dalam bentuk *struct* atau *map* di Go menggunakan `json.Unmarshal`.
*   **Struct Tags:** Go menggunakan *struct tags* (contoh: `` `json:"id"` ``) untuk memetakan nama field antara Go dan JSON, termasuk opsi spesifik seperti `omitempty` untuk mengabaikan field yang kosong saat di-*marshal*, atau tanda strip `-` untuk menyembunyikan field tersebut.
*   **Streaming Encoder/Decoder:** Untuk file JSON yang sangat besar atau membaca JSON langsung dari HTTP Body tanpa menghabiskan memori besar, digunakan tipe `json.Encoder` dan `json.Decoder`.

**Mengapa menggunakan `encoding/json`?**
JSON (JavaScript Object Notation) adalah standar *de facto* untuk pertukaran data di web saat ini (terutama RESTful API). Hampir setiap aplikasi Go yang berbicara dengan aplikasi *frontend* (Web/Mobile) atau layanan *backend* lain akan mengandalkan package ini.

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/encoding/json](https://pkg.go.dev/encoding/json)
