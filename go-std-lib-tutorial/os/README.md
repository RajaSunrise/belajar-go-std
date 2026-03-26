# Modul: `os`

## Ringkasan
Package `os` menyediakan antarmuka untuk fungsionalitas sistem operasi.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `os` menyediakan antarmuka (interface) independen-platform untuk berinteraksi dengan sistem operasi di bawahnya (Windows, Linux, macOS, dll). Desainnya mirip dengan Unix, namun fungsi-fungsinya disamarkan sehingga bisa berjalan mulus di berbagai OS tanpa perlu mengubah kode Anda.

**Tujuan dan Fungsi Utama:**
*   **Manipulasi File dan Direktori:** Membuka, membuat, menghapus, atau membaca file (`os.Open`, `os.Create`, `os.Remove`, `os.Mkdir`).
*   **Environment Variables:** Mengambil atau mengatur variabel environment sistem (`os.Getenv`, `os.Setenv`).
*   **Informasi Proses dan Sistem:** Mendapatkan argumen command-line (`os.Args`), nama host (`os.Hostname`), dan direktori kerja saat ini (`os.Getwd`).
*   **Keluar dari Program:** Menghentikan program secara paksa dengan kode status tertentu (`os.Exit`).

**Mengapa menggunakan `os`?**
Jika aplikasi Anda perlu berinteraksi langsung dengan sistem tempat ia berjalan—seperti membaca konfigurasi dari *environment variables* (untuk memuat rahasia *API key*), menulis log ke dalam sebuah file teks, atau mengambil argumen saat aplikasi dijalankan melalui command line (CLI), package `os` adalah alat wajib.

## Daftar Fungsi Umum dan Cara Penggunaannya

### 1. Variabel Lingkungan (Environment Variables)
Digunakan secara ekstensif pada sistem *backend modern* (seperti arsitektur 12-factor) agar konfigurasi sistem tidak di-hardcode ke dalam kode sumber.
```go
// Mengambil nilai environment variable
dbUser := os.Getenv("DB_USER")
if dbUser == "" {
    fmt.Println("DB_USER belum di-set!")
}

// Menetapkan environment variable (hanya untuk proses ini)
os.Setenv("APP_MODE", "development")
```

### 2. Membaca dan Menulis File
Package `os` menyediakan `os.File` yang mengimplementasikan antarmuka `io.Reader` dan `io.Writer`.
```go
// Membuat file (jika sudah ada, akan ditimpa/truncate)
file, err := os.Create("data.txt")
if err != nil {
    panic(err)
}
defer file.Close()

// Menulis ke file
file.WriteString("Ini baris pertama di dalam file.\n")

// Membaca file
data, err := os.ReadFile("data.txt") // Cara cepat di Go 1.16+
if err != nil {
    panic(err)
}
fmt.Println(string(data))
```

### 3. Mengambil Argumen CLI (`os.Args`)
`os.Args` adalah sebuah slice of string yang berisi argumen command-line yang diberikan saat program dijalankan. Indeks `0` selalu berisi nama atau *path* executable program itu sendiri.
```go
// Menjalankan: go run main.go hello world
argumen := os.Args
// argumen[0] = path/main
// argumen[1] = hello
// argumen[2] = world
fmt.Println("Argumen yang diberikan:", argumen[1:])
```

### 4. Mengakhiri Program (`os.Exit`)
Kadang Anda menemui *fatal error* (seperti gagal menyambung ke database utama saat *startup*) dan perlu mematikan aplikasi dengan segera dan memberi tahu OS bahwa terjadi *error*.
```go
fmt.Println("Aplikasi mengalami error kritis.")
os.Exit(1) // Keluar dengan status 1 (menandakan ada error). Kode 0 artinya sukses.
// Kode di bawah ini TIDAK akan pernah dieksekusi
fmt.Println("Selesai")
```


## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
