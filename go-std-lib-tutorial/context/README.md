# Modul: `context`

## Ringkasan
Package `context` membawa deadline, sinyal pembatalan, dan data rentas layanan.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `context` adalah salah satu standar emas di Go yang jarang ditemukan di bahasa pemrograman lain. Sederhananya, `context` adalah "tas ransel" informasi yang dibawa oleh sebuah HTTP *Request* (atau tugas tertentu) dan selalu dioper / diteruskan dari satu fungsi ke fungsi turunannya sampai proses *request* itu selesai atau dibatalkan.

**Tujuan dan Fungsi Utama:**
*   **Pembatalan Terkoordinasi (Cancellation):** Jika seorang klien (user via browser) meminta halaman laporan berat yang butuh waktu 5 detik untuk dirender, namun di detik ke-2 ia kesal dan menutup browsernya (koneksi terputus). Web server akan tahu klien telah pergi. `context` yang mendeteksinya lalu meneriakkan sinyal batal ke semua *goroutine*, database, dan service lainnya untuk menghentikan kalkulasi, sehingga menghemat daya pemrosesan server CPU! (`context.WithCancel`).
*   **Batas Waktu (Timeout):** Membatasi berapa lama kita mau menunggu sistem luar. Misal memanggil API ke Payment Gateway eksternal dibatasi maksimal 10 detik. Lebih dari itu, anggap gagal. (`context.WithTimeout`).
*   **Kontekstual Value:** Mengangkut sedikit data (misalnya ID Tracing, ID Request, atau User ID hasil *Middleware Auth*) ke kedalaman tumpukan fungsi sistem (dari routing -> layer layanan -> layer repositori database) secara aman tanpa perlu menambahkan argumen satu-persatu pada fungsi-fungsi tersebut (`context.WithValue`).

**Mengapa menggunakan `context`?**
Jika Anda merancang Microservice (aplikasi berbasis API-API kecil yang saling berbicara), package ini **sangat absolut dan wajib**. Tanpanya, jika ada salah satu sistem tetangga yang mati dan lambat merespons, Goroutine di server Go Anda akan menumpuk tidak terbatasi (karena menunggu merespons), kehabisan RAM, dan membuat seluruh server Anda ikutan hancur (*Cascading Failure*).

## Daftar Fungsi Umum dan Cara Penggunaannya

### 1. Context Dasar
Setiap pembuatan fungsi *root* dimulai dari sini. Selalu gunakan ini pada tingkat paling awal program (misal fungsi `main`). Semua turunan context akan berasal darinya.
```go
ctx := context.Background() // Context kosong tanpa batas waktu.
```
*(Catatan: jika di dalam server `net/http`, context secara otomatis sudah disediakan di dalam request: `r.Context()`)*.

### 2. Membatasi Waktu (Timeout)
Digunakan saat querying database atau memanggil HTTP luar. *Timeout* ini akan secara mandiri mengirim sinyal `Done()` jika waktu telah habis.
```go
// Memberikan batas waktu 3 detik pada context awal (Background)
ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
defer cancel() // Selalu biasakan memanggil cancel di akhir blok agar memory tak bocor

// Saat Anda mengoper context ini ke database (misal: db.QueryContext(ctx, "SELECT..."))
// Jika query melebihi 3 detik, fungsi QueryContext otomatis akan return error "context deadline exceeded"
```

### 3. Pembatalan Manual (Cancel)
Digunakan jika Anda memiliki sekumpulan proses Goroutine, dan jika salah satunya menemukan *Error*, Anda ingin menyuruh goroutine sisanya berhenti semua.
```go
ctx, cancel := context.WithCancel(context.Background())

go func() {
    time.Sleep(2 * time.Second)
    fmt.Println("Pekerja 1 selesai tugasnya, menyuruh semua stop!")
    cancel() // Memanggil pembatalan
}()

// Worker yang lain akan terus menunggu sambil mengecek apakah disuruh berhenti
select {
case <-time.After(5 * time.Second):
    fmt.Println("Menunggu terlalu lama.")
case <-ctx.Done():
    // Saluran (channel) Done tertutup akibat pemanggilan cancel() di goroutine pertama
    fmt.Println("Operasi dibatalkan, Worker lain segera berhenti beroperasi:", ctx.Err())
}
```

### 4. Mengirimkan Nilai Tersisip (Value)
Biasanya digunakan oleh *Middleware* HTTP untuk menyimpan data siapa user yang login dan memberikannya pada controller.
```go
ctx := context.WithValue(context.Background(), "user_id", 99)

// Jauh di dalam fungsi lain, Anda bisa menerimanya kembali:
id := ctx.Value("user_id").(int)
fmt.Println("ID Anda adalah:", id)
```


## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
