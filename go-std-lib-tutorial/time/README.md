# Modul: `time`

## Ringkasan
Package `time` menyediakan fungsionalitas untuk mengukur dan menampilkan waktu.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `time` adalah pustaka bawaan untuk segala sesuatu yang berhubungan dengan tanggal, jam, durasi, dan eksekusi tertunda. Sistem Go menggunakan standar *timezone-aware* yang sangat presisi hingga level *nanosecond*.

**Tujuan dan Fungsi Utama:**
*   **Mendapatkan Waktu Saat Ini:** Mengetahui jam dan tanggal sistem beroperasi (`time.Now()`).
*   **Durasi (Duration):** Menghitung rentang waktu antar dua kejadian (`time.Since`, `time.Until`) atau mendefinisikan konstanta durasi baku seperti "2 jam 30 menit" menggunakan manipulasi `time.Duration` (misal: `2 * time.Hour`).
*   **Formatting dan Parsing Waktu:** Mengubah objek internal waktu menjadi string yang bisa dibaca ("01 Januari 2025") atau sebaliknya (parsing teks dari *API request* menjadi objek tipe Waktu). Uniknya di Go, format layout waktu sangat berbeda dari bahasa lain yang menggunakan %Y atau %m. Go menggunakan tanggal spesifik sebagai cetakan dasar (Template): **`Mon Jan 2 15:04:05 MST 2006`**. Angka 1, 2, 3(Jam: 03/15), 4, 5, 6, 7(MST).
*   **Penundaan (Sleep) dan Timer Periodik:** Membuat program/goroutine berhenti sementara (`time.Sleep`) atau membuat eksekusi tertunda berbasis *Channel* (`time.After`, `time.Ticker`).

**Mengapa menggunakan `time`?**
Tidak ada aplikasi *backend* yang tidak memerlukan data waktu. Anda mencatat *timestamp* untuk log aktivitas error, mengisi *field* `created_at` dan `updated_at` ke database, mengatur batas waktu (timeout) pada sebuah koneksi HTTP agar aplikasi tidak *hang*, hingga mengatur cron-job sederhana (fungsi yang dijalankan otomatis setiap jam 12 malam).

## Daftar Fungsi Umum dan Cara Penggunaannya

### 1. Waktu Saat Ini dan Operasi Penambahan
```go
// Mengambil waktu server saat ini
sekarang := time.Now()

// Mengambil waktu besok (saat ini + 24 Jam)
besok := sekarang.Add(24 * time.Hour)

// Mengambil waktu 1 jam 30 menit yang lalu
lalu := sekarang.Add(- (1 * time.Hour + 30 * time.Minute))

// Mendapatkan komponen spesifik
tahun, bulan, hari := sekarang.Date()
jam, menit, detik := sekarang.Clock()
```

### 2. Format dan Parsing Waktu
Ingat baik-baik format kunci di Go: `2006` (Tahun), `01` (Bulan), `02` (Hari), `15` (Jam 24h), `04` (Menit), `05` (Detik).
```go
sekarang := time.Now()

// Menjadikan string
formatStandar := sekarang.Format("2006-01-02 15:04:05")
formatIndo := sekarang.Format("02-01-2006")
fmt.Println("Hari ini:", formatIndo)

// Parsing string menjadi time.Time
teksWaktu := "2024-08-17 10:00:00"
parsedTime, err := time.Parse("2006-01-02 15:04:05", teksWaktu)
if err != nil {
    panic(err)
}
fmt.Println("Waktu yang diparsing:", parsedTime)
```

### 3. Mengukur Durasi (Duration)
Digunakan untuk *benchmarking* atau mengukur seberapa lama sebuah fungsi API dijalankan.
```go
mulai := time.Now()

// ... (Proses yang lama, misal query database) ...
time.Sleep(2 * time.Second) // Simulasi penundaan 2 detik

lamaEksekusi := time.Since(mulai)
fmt.Printf("Query memakan waktu: %s\n", lamaEksekusi)
```

### 4. Ticker (Jadwal Berulang)
Sangat bermanfaat jika Anda memiliki sebuah Goroutine di belakang layar yang tugasnya adalah membersihkan *cache* aplikasi setiap 5 menit.
```go
// Menjalankan sesuatu setiap 1 detik
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()

counter := 0
for t := range ticker.C {
    fmt.Println("Tick pada:", t)
    counter++
    if counter >= 5 {
        break // Berhenti setelah 5 detik
    }
}
```


## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
