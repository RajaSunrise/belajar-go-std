# Modul: `strconv`

## Ringkasan
Package `strconv` memfasilitasi konversi data dari String ke tipe data primitif lain dan sebaliknya.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `strconv` (String Conversion) sangat vital. Seringkali, data yang masuk ke sistem aplikasi Go selalu berupa tipe string text (misalnya, data form HTML, isi *query parameter* URL HTTP, variabel environment OS). Jika kita perlu melakukan operasi matematika atau logika, data string itu harus diubah.

**Tujuan dan Fungsi Utama:**
*   **Konversi Dasar Cepat:** `strconv.Atoi` (ASCII to Integer) mengonversi tipe teks seperti `"100"` menjadi integer `100`. Sedangkan fungsi sebaliknya adalah `strconv.Itoa` (Integer to ASCII).
*   **Parsing Mendalam (Parse):** Fungsi *parse* digunakan saat kita butuh fleksibilitas tambahan pada tipe tertentu: `strconv.ParseInt`, `strconv.ParseFloat`, dan `strconv.ParseBool` (`"true"`, `"1"`, `"t"` akan menjadi true Boolean).
*   **Pemformatan (Format):** `strconv.FormatInt`, `strconv.FormatFloat` berguna untuk mengubah nilai angka dengan kontrol yang spesifik (misal, menentukan *base* biner/heksadesimal atau berapa angka desimal di belakang koma) kembali menjadi teks String.

**Mengapa menggunakan `strconv`?**
Go adalah tipe data kaku (Strongly Typed). Anda tidak bisa sekadar menambahkan string `"5"` dengan int `5` tanpa mengubah salah satu tipe tersebut terlebih dahulu. Package `strconv` menangani hal ini dengan cara paling aman tanpa "magic" tersembunyi, sehingga potensi bug konversi dapat terdeteksi melalui *return value Error* nya.

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/strconv](https://pkg.go.dev/strconv)
