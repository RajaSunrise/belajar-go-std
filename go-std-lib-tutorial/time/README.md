# Modul: `time`

## Ringkasan
Package `time` menyediakan kapabilitas lengkap, mutakhir, dan berkinerja tinggi untuk mengukur waktu, menjadwalkan eksekusi (timer/ticker), serta mem-parsing (menerjemahkan) dan memformat (*formatting*) representasi penanggalan (*Date & Time*) secara akurat dengan dukungan penuh *Timezone* (zona waktu) global.

## Penjelasan Lengkap (Fungsi & Tujuan)
Di dunia rekayasa perangkat lunak modern, "waktu" bukanlah konsep yang sederhana. Karena perbedaan zona geografis, perubahan waktu musiman (*Daylight Saving Time* / DST), hingga detik kabisat (*leap seconds*), representasi waktu dalam string teks "14 Januari 2025 jam 10 Pagi" memiliki ribuan makna jika Anda tidak menyertakan konteks koordinat lokasinya (GMT/UTC). Package `time` dari Golang mengatasi kerumitan tersebut dengan membungkus seluruh data jam atomik ke dalam satu buah struktur data pusat bernama **`time.Time`**.

Struktur `time.Time` di Go sangat presisi; ia tidak sekadar menyimpan titik kalender Gregorian, melainkan jam atom beresolusi tingkat **Nanosecond** (10^-9 detik). Lebih uniknya lagi, `time.Time` di Go memisahkan (*encapsulate*) titik absolut *Wall Clock* (jam kalender dinding yang dilihat manusia dan bisa diubah secara manual di OS) dengan titik mutlak *Monotonic Clock* (jam internal *hardware* mesin yang selalu maju dan tak terpengaruh loncatan waktu OS). Desain arsitektur hibrida ini mencegah bug mematikan ketika aplikasi web Anda sedang berjalan, dan tiba-tiba Administrator Server memajukan paksa waktu mesin Linux sejauh 1 jam, fungsi jeda durasi (*Timer/Ticker*) Anda tidak akan panik melainkan terus berjalan secara rasional berdasarkan jam monotonik yang aman tersebut!

**Tujuan dan Fungsi Utama:**
1.  **Pencatatan Instan (Now):** Mendapatkan potret wujud waktu lokal presisi nanodetik persis saat fungsi tersebut dipanggil (`time.Now()`), yang siap ditanamkan (*embedded*) ke dalam baris rekaman log aplikasi atau transaksi database Anda.
2.  **Operasi Aritmetika Waktu:** Melakukan kalkulasi majemuk durasi kalender di masa depan atau masa lampau secara absolut tanpa repot menghitung jumlah lompatan hari dalam 1 bulan (contoh: 30 hari ke depan + 5 jam -> `sekarang.AddDate(0,0,30).Add(5 * time.Hour)`).
3.  **Pengukuran Rentang Waktu (Duration):** Mengalkulasi perbedaan interval detik (*Benchmark* & Profiling). Berapa millisecond *Database Query SQL* ini menyelesaikan permohonannya?
4.  **Format Sintaks Unik:** Go tidak pernah menggunakan template karakter format usang seperti persentase `%Y` atau `%H` dalam mem-parsing kalender ke teks, yang harus dihafalkan sintaks maknanya. Golang memperkenalkan layout basis referensi visual (sebuah momen fiktif **Jan 2 15:04:05 2006 MST**), yang mana Go compiler akan mencocokkan pola urutan tersebut untuk diubah menjadi wujud tanggal.
5.  **Penjadwalan (Scheduling):** Menunda (Sleep) sementara jalannya goroutine yang sedang dieksekusi, atau menggunakan *Ticker Channel* yang menembakkan pemicu kejadian berulang-ulang, cocok untuk siklus latar belakang pembersihan tembolok memori (*cache purge daemon*).

**Mengapa menggunakan `time`?**
Jika aplikasi Anda menyimpan riwayat rekam medis pasien, mencatat rentang penagihan pembayaran lisensi *software* (*recurring billing*), mengecek apakah sesi *token JWT Auth* sudah melampaui batas kedaluwarsa, atau memanajemen eksekusi *Timeout HTTP* terhadap layanan hulu (Upstream Gateway API), tak sedetik pun server *backend* Go Anda berhenti memanggil paket komprehensif nan sakral ini.

---

## Daftar Fungsi Umum dan Cara Penggunaannya Secara Mendalam

### 1. Titik Waktu dan Operasi Matematis (Add & Sub)

Setiap detik di dalam aplikasi Go Anda disimpan dalam bentuk entitas absolut `time.Time`. Berbeda dengan durasi (`time.Duration`) yang menyimpan jeda waktu abstrak tak berwujud seperti "2 jam".

*   **`time.Now()`**: Memotret detik ini secara instan di zona lokal di mana server berdiri.
*   **`.Add(durasi)`**: Melakukan teleportasi masa. Menambahkan atau mengurangi durasi non-tanggal (hanya menambahkan detik/jam/menit) pada titik absolut waktu tersebut.
*   **`.AddDate(tahun, bulan, hari)`**: Melakukan teleportasi penanggalan kalender (perhatikan, jika hari ini 31 Januari dan Anda tambah 1 Bulan, Go cukup pintar melompat bukan ke 31 Februari melainkan ke tanggal matematis terdekat yang valid secara Gregorian, contohnya 2/3 Maret).
*   **`.Sub(waktu_lain)`**: (Subtract / Kurangi). Menghitung "Berapa jarak yang terbentuk (Duration) apabila waktu saya ini dikurangi waktu titik B yang lalu".
*   **`.Before(t)`, `.After(t)`**: Menguji dua objek *time* secara konseptual. Apakah waktu saya ini *Sebelum* waktu dia? (True/False).

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    sekarang := time.Now()

    // Teleportasi masa depan durasi
    duaJamLagi := sekarang.Add(2 * time.Hour)

    // Teleportasi mundur 3 hari yang lalu (Menggunakan AddDate)
    lalu := sekarang.AddDate(0, 0, -3)

    fmt.Printf("Waktu Saat Ini: %v\n", sekarang)
    fmt.Printf("Tiga Hari yang lalu: %v\n", lalu)

    // Perbandingan Titik Waktu (Logika Kondisional)
    if lalu.Before(sekarang) {
        fmt.Println("Ya, tiga hari lalu terjadi sebelum hari ini.")
    }

    // Menghitung selisih waktu secara murni
    // Misal: Kapan tiket diskon ini akan kadaluarsa?
    kadaluarsa := sekarang.Add(48 * time.Hour)
    sisaWaktu := kadaluarsa.Sub(sekarang)

    fmt.Printf("Batas promo tiket habis dalam %v jam.\n", sisaWaktu.Hours())
}
```

---

### 2. Format Referensi Unik Go (Formatting & Parsing)

Perlu dipahami: Jika di bahasa Python atau C untuk menulis format waktu Anda menggunakan token rahasia `"%Y-%m-%d %H:%M:%S"`, dalam bahasa Go sistem template itu dipandang sangat tidak masuk akal (siapa yang hafal perbedaan huruf besar kecil `%M` menit vs `%m` bulan?).

Sebaliknya, sang pencipta Go menetapkan sebuah tanggal tetap di Amerika Serikat sebagai template *referensi* yang sangat berurutan angkanya: `1` (bulan Januari), `2` (Tanggal dua), `3` (Jam Tiga sore/15:00), `4` (Menit 04), `5` (Detik 05), `6` (Tahun 2006), dan `7` (Zona Waktu MST yang beda angkanya 7 dari UTC). Jadi:
*Ajaib!* Format ajaib itu adalah: **`Mon Jan 2 15:04:05 MST 2006`**

*   **`.Format(layout)`**: Memuntahkan `time.Time` kembali menjadi bentuk kalimat teks `string`.
*   **`time.Parse(layout, teks)`**: Menyerap kalimat teks kalender eksternal dari API asing menjadi `time.Time` murni yang terkalibrasi ke standar UTC default.
*   **`time.ParseInLocation(...)`**: Sama seperti di atas namun menyerap teks asing dan memaksanya terikat kepada zona waktu (*Timezone*) lokasi tertentu yang Anda deklarasikan (bukan jatuh ke default UTC).

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Anda bisa mendeklarasikan template format Anda sendiri
    // ASALKAN harus selalu merujuk angka-angka spesifik 2006, 01, 02 dsb
    layoutDatabaseSQL := "2006-01-02 15:04:05"
    layoutOrangIndo   := "02/01/2006 Jam 15:04"
    layoutSistemLog   := time.RFC3339 // Konstanta default ISO8601 standar internet

    waktuSkrg := time.Now()

    // Merubah Waktu menjadi String!
    fmt.Println("Format Standar SQL DB  :", waktuSkrg.Format(layoutDatabaseSQL))
    fmt.Println("Format Human Indo      :", waktuSkrg.Format(layoutOrangIndo))
    fmt.Println("Format Web API Intern  :", waktuSkrg.Format(layoutSistemLog))

    // ---------
    // Kasus Parsing: Mengubah Teks JSON dari Frontend menjadi Waktu
    inputBirtday := "1999-08-17"

    // Kita beritahu Go bahwa "1999" itu mewakili komponen 2006 dst.
    lahirDiri, err := time.Parse("2006-01-02", inputBirtday)
    if err != nil {
        panic(fmt.Sprintf("Format teks kelahiran ngawur! Error: %v", err))
    }

    fmt.Println("\nData Waktu Lahir Di-Parse:", lahirDiri.Format("Hari 02 Bulan 01 Tahun 2006"))
}
```

---

### 3. Eksekusi Jeda Tertunda (Sleep) dan Timer Channel

Meskipun `fmt.Println` mencetak ke layar secara sangat singkat secepat kedipan, seringkali proses Goroutine latar belakang kita diinstruksikan untuk *beristirahat* demi membatasi tingkat lalu lintas (*Rate Limiting*) panggilan ke Server Pihak Ketiga agar IP kita tak diblokir.

*   **`time.Sleep(d)`**: Membekukan Goroutine yang tengah memanggilnya selama hitungan durasi. Fungsi ini hanya melumpuhkan 1 Goroutine yang sedang mengeksekusinya; goroutine *Thread* Go paralel lainnya takkan terganggu sama sekali (Ini adalah kekuatan *non-blocking* Go scheduler!)
*   **`time.After(d)`**: Cara elegan berbasis salurah pipa (`Channel`). Fungsi ini tidak membekukan program. Melainkan ia langsung mengembalikan sebuah Saluran Penerima (`<-chan Time`), dan `d` durasi kemudian, Go di belakang layar akan menembakkan waktu saat itu ke dalam saluran tersebut. Cocok dikombinasikan dengan perintah `select`.
*   **`time.Tick(d)` atau `time.NewTicker(d)`**: Mirip dengan `.After()`, bedanya `.After()` hanya menembak 1 kali lalu mati. Ticker menembakkan jam ke dalam Channel secara periodik abadi (misal: setiap 5 detik tanpa putus) hingga Anda memanggil perintah penutup `.Stop()`.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Memulai Simulasi Proses Sinkronus (Sleep)")
    // Eksekusi akan hangup disini selama 1.5 detik
    time.Sleep(1500 * time.Millisecond)
    fmt.Println("Selesai tertidur pulas!\n")

    // ----------------------
    // Pola Penggunaan Modern (Timeout Channel dgn select)
    fmt.Println("Memulai Simulasi API Request Asinkron")

    // Kita buat channel kosong untuk menandakan proses berat
    saluranProses := make(chan bool)

    // Jalankan goroutine di background (pura-pura proses 3 detik)
    go func() {
        time.Sleep(3 * time.Second)
        saluranProses <- true // Laporkan selesai ke channel
    }()

    // Tunggu balasan menggunakan Blokir Select
    select {
    case <-saluranProses:
        // Jika goroutine background berhasil merespons sebelum batas waktu
        fmt.Println("BERHASIL! Proses selesai lebih cepat.")

    case <-time.After(2 * time.Second):
        // Jika 2 detik berlalu lebih cepat dibanding 3 detik, blok ini yg tereksekusi duluan
        fmt.Println("GAGAL! Timeout: Server backend terlalu lama merespons.")
        // Di aplikasi nyata, Anda langsung return error HTTP 504 Gateway Timeout di sini
    }
}
```

---

## Bagian Lanjutan: Waktu Temporal Abstrak dan Pemrograman Multithreading (Timezone)

Di balik kelenturan luar biasa format "Mon Jan 2 15:04:05 2006 MST" milik package `time` Golang, terdapat ancaman yang mengintai bagi developer *Backend* junior saat sistem berinteraksi di lingkungan *Server Global Database Distribusi* berskala besar. Di bagian ini, kita akan membahas lebih dalam arsitektur waktu berbasis lokal (Timezone/Location).

### 1. Titik Geografis Lokasi (Timezone / Location)

Bayangkan server Go Anda berlokasi di *Data Center* Amazon AWS Singapura, namun melayani transaksi pelanggan dari Jakarta, Indonesia, dan Tokyo, Jepang.
Titik waktu yang dikembalikan oleh `time.Now()` secara asali (*default*) mengikat dan menempel kuat pada sistem operasi fisik (WIB/SGT/JST).

Saat Anda menyimpan "10 Januari 2025 jam 09:00" ke dalam database Postgres dari server Singapura, server dari Jerman yang menarik data yang sama mungkin salah menginterpretasikannya sebagai "Jam 9 Pagi versi Jerman (CET)", memundurkan waktu acara sebanyak 7 jam dari realita aslinya!

**Solusi Mutlak (*Standard Industry Best Practice*):**
Anda **HARUS SELALU** menormalisasi waktu ke titik jangkar netral absolut (Titik Nol Bumi): **UTC** (*Coordinated Universal Time*).

```go
// BURUK! Mengirimkan Waktu Server ke Klien (Berbahaya jika pindah benua server)
transaksiPembayaran := time.Now()
simpanKeDatabase(transaksiPembayaran)

// BENAR! Paksa mutlak ke titik tengah waktu bumi (Bebas dari zona waktu geografis OS)
transaksiAbsolut := time.Now().UTC()
simpanKeDatabase(transaksiAbsolut)

// Menyiapkan Waktu Lokal Target (Misal: Melayani Pengunjung Website Jakarta)
// Muat database lokasi internal IANA dari OS
lokasiJakarta, errLoad := time.LoadLocation("Asia/Jakarta")
if errLoad != nil {
    panic(errLoad) // Berbahaya, ini artinya database IANA timezone tidak terinstall di Linux server ini (biasanya tzdata package)
}

// Konversi titik absolut UTC database menjadi wajah yang dipahami klien Frontend Jakarta!
waktuTampilLayar := transaksiAbsolut.In(lokasiJakarta)
fmt.Println(waktuTampilLayar.Format("15:04:05 WIB"))
```

### 2. Memutar Balik Waktu dengan *Unix Timestamp*

Pada antarmuka API (*RESTful API* / JSON) atau penulisan kontrak token JWT, waktu jarang dikirimkan dalam bentuk kalimat *String* kompleks panjang seperti `"2025-01-14T09:00:00Z"`. Bahasa pemograman C (zaman 1970) menetapkan standar abadi yang digunakan hingga detik ini: **Unix Time** atau *Epoch Timestamp*.

Itu adalah **angka tunggal integer murni** yang melambangkan: *"Berapa detik persisnya bumi ini telah berlalu terhitung mundur sejak malam Tahun Baru 1 Januari 1970 di Greenwich London?"*

Fungsi di package `time` sangat andal mengkonversi hal ini bolak-balik:
*   **`time.Now().Unix()`**: Mendapatkan *integer* detik panjang (contoh: `1705228800`). Sangat dianjurkan dikirimkan sebagai tipe numerik di JSON, membiarkan aplikasi klien Frontend (Browser JS / Aplikasi Android Kotlin) mengubah angka integer itu menjadi waktu tampilan kalender visual sesuai jam HP milik klien!
*   **`time.Unix(detik_integer, nanodetik)`**: Menghidupkan kembali roh objek jam Go yang telah menjadi *integer* tadi, dan merakitnya utuh menjadi Obyek utuh `time.Time` Go.

```go
sekarang := time.Now()

// Eksekusi Pembuatan Angka Sandi Epoch Biner Tunggal untuk Transmisi JSON
token_expired_at := sekarang.Add(2 * time.Hour).Unix()
fmt.Println("Kirimkan integer ringkas ini ke Database JWT:", token_expired_at)

// Menerjemahkan angka itu kembali dari Database ke Waktu Mesin Go:
// Angka 0 di parameter kedua merepresentasikan tingkat mikrodetik. Jika tidak butuh ketelitian mili/mikro-detik, cukup isikan angka nol.
titikWaktuKembali := time.Unix(token_expired_at, 0)

if time.Now().After(titikWaktuKembali) {
    fmt.Println("Access Denied: Token sudah basi melewati batas umurnya.")
}
```

### 3. Jebakan `time.Parse` dan Format Teks

Ini adalah rahasia yang banyak membuat frustasi: Apabila Anda menyuapi teks ke dalam **`time.Parse(layout, teks)`**, objek `time.Time` yang tercipta akan diatur oleh kompilator seolah-olah waktu itu terjadi di koordinat absolut **UTC+0000**, KECUALI teks tersebut mengandung instruksi zona waktu eksplisit (misal tulisan `-0500` atau huruf `Z` khusus standar ISO-8601 `time.RFC3339`).

Apabila Anda menelan teks `2025-01-01 10:00:00` dan menganggap itu terjadi di Jakarta (karena klien ada di Jakarta), lalu Anda menyimpan ke DB, Anda telah meleset 7 jam lebih lambat!
Penyelesaian masalah ini dilakukan dengan cara injeksi lokasi eksplisit saat proses kompilasi string parse:

```go
// Teks masuk ke API dari sebuah Formulir Web Orang Indonesia.
teksTanggalLahir := "1999-08-17 10:00:00"

// Langkah 1: Kunci titik jangkar wilayah geografis!
zonaJakarta, _ := time.LoadLocation("Asia/Jakarta")

// Langkah 2: Jangan gunakan `time.Parse()`, MELAINKAN GUNAKAN FUNGSI SAKTI KHUSUS INI:
waktuLahirDiJakarta, errPenyusupan := time.ParseInLocation("2006-01-02 15:04:05", teksTanggalLahir, zonaJakarta)

if errPenyusupan == nil {
    // Kini, titik 10 Pagi itu di-cap secara resmi bersemayam di koordinat WIB (UTC+7)!
    // Sistem tak akan kebingungan membandingkannya lagi dengan waktu Server Singapura/Amerika Serikat.
}
```

Trik `ParseInLocation` ini menjamin tidak adanya perbedaan komparasi absolut logika (`.Before()` atau `.After()`) di tengah sistem microservices yang terdistribusi terpecah belah melintasi samudera dan benua. Kemahiran mengorkestrasi waktu ini menyelamatkan integritas pelaporan data di mata klien di ranah produksi global yang rentan akan anomali waktu.

## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
