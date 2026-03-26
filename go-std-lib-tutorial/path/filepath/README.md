# Modul: `path/filepath`

## Ringkasan
Package `path/filepath` memanipulasi path sistem file dengan cara yang kompatibel lintas sistem operasi.

## Penjelasan Lengkap (Fungsi & Tujuan)
Jika Anda berurusan dengan direktori dan file, `path/filepath` adalah alat bantu utama. Perbedaan besarnya dengan package `path` biasa adalah `filepath` sadar akan Sistem Operasi tempat dia berjalan (menggunakan backslash `\` di Windows, dan slash `/` di Linux/Mac).

**Tujuan dan Fungsi Utama:**
*   **Penggabungan yang Aman:** Fungsi `filepath.Join` menggabungkan nama-nama folder dan file menggunakan pemisah yang benar sesuai OS. Jangan pernah menggabungkan string path secara manual (`dir + "/" + file`) karena akan rusak di Windows!
*   **Ekstraksi Komponen Path:** Mengambil nama file ekstensi saja (`filepath.Ext`), mendapatkan path folder *parent* nya (`filepath.Dir`), atau mendapatkan nama file murni (`filepath.Base`).
*   **Normalisasi & Absolut:** Membersihkan penulisan *path* yang membingungkan (`filepath.Clean` untuk mengubah `a/../b/c` menjadi `b/c`) atau mendapatkan lokasi lengkap/absolute dari sebuah file (`filepath.Abs`).
*   **Traversing / Mencari:** Menelusuri seluruh direktori dan subdirektori untuk mencari file dengan pola tertentu menggunakan `filepath.Walk` atau `filepath.Glob`.

**Mengapa menggunakan `path/filepath`?**
Kode Go diharapkan bisa di*compile* silang dan berjalan di mana saja. Menggunakan package ini menjamin aplikasi Anda tidak memiliki bug atau gagal mencari file konfigurasi hanya karena program dijalankan dari sistem operasi yang berbeda.

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/path/filepath](https://pkg.go.dev/path/filepath)
