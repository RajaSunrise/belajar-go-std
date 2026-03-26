# Modul: `strings`

## Ringkasan
Package `strings` mengimplementasikan fungsi-fungsi sederhana untuk memanipulasi string UTF-8.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `strings` adalah utilitas utama di Go untuk melakukan manipulasi atau operasi pada tipe data teks (`string`).

**Tujuan dan Fungsi Utama:**
*   **Pencarian Teks:** Mengecek apakah string mengandung teks tertentu (`strings.Contains`), mencari indeks kemunculan suatu kata (`strings.Index`), atau mengecek awalan/akhiran (`strings.HasPrefix`, `strings.HasSuffix`).
*   **Manipulasi Huruf:** Mengubah teks menjadi huruf kapital semua (`strings.ToUpper`) atau huruf kecil semua (`strings.ToLower`).
*   **Memecah dan Menggabungkan:** Memisahkan satu kalimat menjadi kumpulan kata atau baris (`strings.Split`) dan sebaliknya, menggabungkan beberapa kata menjadi satu kalimat (`strings.Join`).
*   **Pembersihan:** Menghapus spasi atau karakter tertentu di awal/akhir string (`strings.Trim`, `strings.TrimSpace`).
*   **Penggantian:** Mengganti sebagian kata di dalam string dengan kata lain (`strings.Replace`, `strings.ReplaceAll`).

**Mengapa menggunakan `strings`?**
Sangat krusial untuk memproses input teks dari pengguna. Misalnya, saat melakukan pencarian di database (seringkali teks diubah ke *lowercase* terlebih dahulu), memvalidasi format (menghapus spasi yang tidak disengaja), atau memparsing data yang dipisahkan oleh koma (CSV).

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/strings](https://pkg.go.dev/strings)
