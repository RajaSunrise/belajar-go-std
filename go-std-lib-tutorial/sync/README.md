# Modul: `sync`

## Ringkasan
Package `sync` menyediakan primitive sinkronisasi dasar seperti mutex dan wait group.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `sync` menyediakan utilitas tingkat rendah yang sangat penting untuk menulis kode *concurrent* (berjalan bersamaan/paralel) secara aman (thread-safe). Go terkenal dengan kemudahan konkurennya melalui *Goroutine*, namun seringkali goroutine-goroutine ini perlu saling berkoordinasi.

**Tujuan dan Fungsi Utama:**
*   **Menunggu Pekerjaan Selesai (WaitGroup):** Jika Anda memecah pekerjaan ke 10 goroutine yang berjalan berbarengan, program utama harus menunggu semuanya selesai sebelum exit. `sync.WaitGroup` adalah cara standar melakukannya.
*   **Mencegah Kondisi Balapan Data (Mutex):** Jika dua atau lebih goroutine mencoba mengubah data pada variabel (memori) yang sama secara bersamaan, akan terjadi kerusakan data (*Race Condition*). `sync.Mutex` mengunci (lock) data tersebut, sehingga hanya 1 goroutine yang bisa mengubahnya dalam satu waktu.
*   **Eksekusi Hanya Sekali (Once):** Menjamin bahwa sebuah fungsi inisialisasi (misalnya koneksi ke *database*) hanya dijalankan tepat *satu kali*, meskipun diakses oleh 1000 goroutine bersamaan, dengan menggunakan `sync.Once`.
*   **Pool Objek:** Menyimpan dan mendaur ulang objek sementara yang sering dipakai untuk mengurangi tekanan pada *Garbage Collector* (`sync.Pool`).

**Mengapa menggunakan `sync`?**
Kapan pun Anda menggunakan keyword `go func()` dan ada memori/variabel bersama (seperti Map, Slice, atau field pada struct) yang dibaca/ditulis bersamaan, Anda diwajibkan memahami dan menggunakan primitive sinkronisasi dari package ini agar aplikasi tidak panik (crash) atau menghasilkan data invalid.

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/sync](https://pkg.go.dev/sync)
