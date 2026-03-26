# Modul: `io`

## Ringkasan
Package `io` menyediakan antarmuka dasar untuk Input/Output berbasis stream.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `io` tidak mengerjakan hal fungsional yang terlihat di layar, melainkan menyediakan abstraksi antarmuka (interface) tingkat rendah (`io.Reader` dan `io.Writer`) yang menyatukan seluruh komponen sistem Go. Daripada memuat seluruh file berukuran 10 GB ke memori sekaligus (yang akan memicu Out-of-Memory / Server Crash), arsitektur `io` memaksa kita memproses data secara mengalir (streaming) sepotong demi sepotong.

**Tujuan dan Fungsi Utama:**
*   **Kontrak Baku Data Mengalir:** Karena `os.File`, `net.Conn` (koneksi jaringan tcp), HTTP Response Body, hingga *zip archive* semuanya mengimplementasikan interface `io.Reader` (artinya: data bisa dibaca dari sumber tersebut byte-demi-byte) dan `io.Writer` (artinya: bisa ditulis), mereka semua kompatibel satu sama lain.
*   **Penyalinan Aman Memori (Copy):** Fungsi `io.Copy` memungkinkan Anda menyalin aliran file langsung ke aliran Web Response, memindahkan file bergigabyte dengan konsumsi RAM yang stabil di angka beberapa megabyte saja.

**Mengapa menggunakan `io`?**
Kekuatan sejati Go terletak pada *Interface* nya, dan `io` adalah puncak dari desain tersebut. Jika Anda membuat fungsi untuk menghitung jumlah baris, buatlah agar ia menerima argumen bertipe `io.Reader`. Dengan begitu fungsi Anda otomatis bisa menghitung jumlah baris di dalam file di hardisk, jumlah baris pada data yang di-*upload* lewat HTTP, atau jumlah baris di memori internal buffer tanpa perlu diubah sedikitpun!

## Daftar Fungsi Umum dan Cara Penggunaannya

### 1. Antarmuka `io.Reader` dan `io.Writer`
Ini adalah pemahaman konseptual. Segala fungsi yang menuntut `io.Writer` berarti ia ingin "menuliskan sesuatu kepada Anda". Segala yang menuntut `io.Reader` berarti "ia ingin membaca sesuatu dari Anda".
```go
// Definisi Interface di dalam inti Go:
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

### 2. Membaca Semuanya dengan `io.ReadAll`
Jika Anda benar-benar yakin ukuran datanya kecil (misal: respons API berupa JSON pendek), Anda bisa membaca semua aliran itu menjadi array Byte sekaligus ke dalam RAM agar mudah diproses string-nya.
```go
// strings.NewReader mengubah string teks menjadi sebuah io.Reader
sumber := strings.NewReader("Ini adalah data simulasi dari internet.")

dataBytes, err := io.ReadAll(sumber)
if err != nil {
    panic(err)
}
fmt.Println(string(dataBytes))
```

### 3. Menyalin Tanpa Memori Berlebih dengan `io.Copy`
Ini adalah pola yang sangat sering dijumpai saat men-*download* file besar atau melayani file statis dari disk.
```go
fileSumber, _ := os.Open("video_besar.mp4")
defer fileSumber.Close()

fileTujuan, _ := os.Create("copy_video.mp4")
defer fileTujuan.Close()

// io.Copy memindahkan aliran data secara efisien melalui buffer memori kecil (chunking)
bytesDisalin, err := io.Copy(fileTujuan, fileSumber)
fmt.Printf("Berhasil menyalin %d bytes data!\n", bytesDisalin)
```


## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
