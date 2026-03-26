# Modul: `sort`

## Ringkasan
Package `sort` menyediakan utilitas untuk mengurutkan (sorting) data *slice* dan *collections*.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `sort` mengimplementasikan logika pengurutan dan pencarian biner (*binary search*) untuk berbagai jenis tipe data. Go modern sudah memberikan banyak *helper* fungsi yang membuat pengurutan menjadi sangat ringkas.

**Tujuan dan Fungsi Utama:**
*   **Tipe Bawaan:** Mengurutkan tipe primitif seperti slice of integer (`sort.Ints`), slice of float (`sort.Float64s`), atau slice of string (`sort.Strings`).
*   **Pengurutan Tipe Kustom (Slice of Struct):** Di Go 1.8 ke atas, `sort.Slice` memungkinkan Anda mengurutkan *slice* tipe apa pun (termasuk *struct*) hanya dengan menyediakan fungsi *callback* yang memberi tahu mana dari dua elemen yang "lebih kecil" atau harus muncul duluan.
*   **Pengecekan Urutan:** Fungsi seperti `sort.IsSorted` atau `sort.SliceIsSorted` dapat memeriksa apakah sebuah *slice* sudah terurut dengan benar atau belum.
*   **Pencarian Biner:** `sort.Search` menyediakan cara super efisien ($O(\log n)$) untuk mencari index suatu nilai di dalam data array/slice yang *sudah terurut*.

**Mengapa menggunakan `sort`?**
Jika Anda harus menampilkan daftar *leaderboard* game, mengurutkan hasil query yang didapat dari API berdasarkan tanggal atau harga, package `sort` menyediakan algoritma *quicksort/timsort* yang sangat teroptimasi dan *out of the box* tanpa Anda perlu menulis algoritmanya dari awal.

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/sort](https://pkg.go.dev/sort)
