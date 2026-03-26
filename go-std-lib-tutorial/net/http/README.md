# Modul: `net/http`

## Ringkasan
Package `net/http` menyediakan implementasi klien dan server HTTP.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `net/http` adalah salah satu "superpower" bahasa Go. Package ini memungkinkan Go untuk membuat web server berkinerja tinggi atau melakukan panggilan API ke server lain tanpa memerlukan framework atau pustaka pihak ketiga.

**Tujuan dan Fungsi Utama:**
*   **Membuat Web Server:** Dengan menggunakan `http.ListenAndServe`, Anda dapat membuat server HTTP untuk melayani *request* dari browser atau aplikasi klien.
*   **Routing (ServeMux):** Mendefinisikan endpoint URL mana yang akan ditangani oleh fungsi yang mana (contoh: `/api/users` ditangani oleh fungsi A). Mulai Go 1.22, *ServeMux* secara *native* mendukung *method-based routing* (misal `GET /items/{id}`).
*   **HTTP Client:** Melakukan HTTP Request (GET, POST, PUT, DELETE) ke sistem eksternal menggunakan `http.Get`, `http.Post`, atau `http.Client`.
*   **Memanipulasi Request & Response:** Membaca header, cookie, body dari *request* masuk, serta mengirim balasan JSON atau HTML.

**Mengapa menggunakan `net/http`?**
Bahasa Go terkenal kuat untuk pembuatan backend / microservice. Package inilah mesin utamanya. Anda bisa membuat API (Application Programming Interface) yang cepat dan tangguh hanya dengan pustaka standar ini.

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/net/http](https://pkg.go.dev/net/http)
