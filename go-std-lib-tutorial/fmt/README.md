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

## Daftar Fungsi Umum dan Cara Penggunaannya

### 1. `fmt.Println()` dan `fmt.Print()`
Digunakan untuk mencetak output ke layar. `Println` menambahkan baris baru (`\n`) di akhir kalimat, sedangkan `Print` tidak.
```go
fmt.Println("Halo, dunia!") // Output: Halo, dunia! (dengan enter)
fmt.Print("Halo, ")
fmt.Print("dunia!\n")       // Output: Halo, dunia! (dengan enter manual)
```

### 2. `fmt.Printf()` (Print Format)
Mencetak dengan format tertentu menggunakan *verbs*. Sangat berguna untuk menggabungkan variabel ke dalam string tanpa menggunakan operator `+`.
*   `%v` : Nilai default dari variabel (tipe apapun).
*   `%+v`: Menampilkan nama field jika variabel adalah sebuah struct.
*   `%T` : Menampilkan tipe data dari variabel tersebut.
*   `%s` : Format sebagai string biasa.
*   `%d` : Format sebagai integer (basis 10).
*   `%f` : Format sebagai float (desimal).

```go
nama := "Budi"
umur := 25
tinggi := 170.5
fmt.Printf("Nama saya %s, umur %d tahun, dan tinggi %f cm.\n", nama, umur, tinggi)
```

### 3. `fmt.Sprintf()` (String Print Format)
Fungsi ini memiliki kegunaan yang persis sama dengan `Printf`, namun alih-alih mencetak hasilnya ke layar, fungsi ini **mengembalikan nilainya sebagai sebuah string baru**. Ini sangat sering digunakan untuk menyusun pesan log, query database dinamis, atau menyusun HTML *raw*.
```go
pesanError := fmt.Sprintf("Gagal memuat pengguna dengan ID %d", userID)
// pesanError sekarang berisi string yang bisa disimpan atau dikembalikan
```


## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
