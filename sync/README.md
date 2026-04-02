# Modul: `sync`

## Ringkasan
Package `sync` menyediakan piranti pelindung sinkronisasi mekanik instrumen (*low-level synchronization primitives*) yang esensial tingkat dewa untuk mengekang keliaran kebebasan Goroutine (*Go's Concurrent Threads*), menjinakkan tabrakan data balapan memori buta (*Race Conditions*), dan mengoordinasi ritme harmonis pergerakan eksekusi jutaan rentetan blok paralelisasi di dalam tubuh server aplikasi yang sibuk.

## Penjelasan Lengkap (Fungsi & Tujuan)
Semboyan kampanye paling legendaris dari bahasa pemrograman Go di muka bumi ini berbunyi: *"Jangan saling berkomunikasi dengan cara membagi-bagikan memori; bagikan memori dengan cara saling berkomunikasi"* (*Don't communicate by sharing memory; share memory by communicating via Channels*). Secara murni, idiom Go modern akan menuntut Anda menggunakan fasilitas obrolan pipa asinkron (`Channels`) untuk memandu perpindahan kendali data yang rumit antar-goroutine.

Akan tetapi, di arena pertempuran sejati, terkadang membungkus sebuah entitas obyek yang sederhana (misalnya: Peta Data Penyimpanan Sesion sementara / `map` In-Memory, ataupun satu hitungan Counter Integer Global pengunjung situs Web) ke dalam arsitektur labirin `Channels` merupakan hal *sangat* berlebihan (*Overkill*), memberatkan alokasi RAM, dan melambatkan pemrosesan kecepatan murni CPU. Di sinilah primitif klasik perlindungan instrumen `sync` bersinar mengamankan fondasi memori secara primitif dan absolut cepat!

**Tujuan dan Fungsi Utama:**
1.  **Penguncian Pintu Brankas Memori (*Mutex*):** Tipe alat `sync.Mutex` adalah gembok pengunci eksklusif mutlak (*Mutual Exclusion*). Manakala dua atau sepuluh ribu goroutine pekerja (*workers*) datang berkerumun berupaya menyuntikkan dan mengubah variabel nilai Map/Slice di lokasi memori RAM tunggal yang sama secara bebarengan hitungan detik, gembok Mutex ini akan meredam kekacauan itu dengan memaksa mereka "Berbaris antre menunggui Pintu Gembok yang terkunci, masuk satu per satu". Mencegah kecelakaan mematikan server.
2.  **Rombongan Penunggu (*WaitGroup*):** Menahan fungsi pusat pemanggil (misal: blok kode `main()`) agar mau diam terpaku bersabar menunggui proses eksekusi keseratus pekerja latar belakang *goroutine* yang masih memproses tugas asinkronnya masing-masing (*Fork & Join Model*), hingga mereka semua menyatakan laporan: "Kami seluruh rombongan seratus pekerja sudah tuntas!".
3.  **Pelindung Inisiasi Sekali Seumur Hidup (*Once*):** Alat pemaksa saklar `sync.Once`. Fungsi pemicu ini menjamin secara absolut 100% bahwa tugas inisialisasi yang berat (seperti membentuk jembatan koneksi pangkalan Data SQL raksasa Server) hanya akan digarap dan dikerjakan secara eksak **Satu Kali Saja**, tak perduli berapa puluh ribu rentetan *request* klien goroutine memanggil gedoran fungsi penawaran tersebut bebarengan membanjiri di awal mulanya.

**Mengapa menggunakan `sync`?**
Ketidaktahuan terhadap utilitas krusial bawaan package pelindung `sync` saat Anda membangun sebuah layanan server asinkron Go adalah hal ihwal bunuh diri massal aplikasi (yang berakhir ditandai kelumpuhan Fatal sistem di *Terminal Linux* berbunyi: `panic: fatal error: concurrent map writes` atau `concurrent map read and map write`). Anda dituntut memahami mekanismenya sebelum merangkul desain arsitektur Paralelisasi apapun.

---

## Daftar Fungsi Umum dan Cara Penggunaannya Secara Mendalam

### 1. Pelindung Anti Kecelakaan Memori Bersama (Mutual Exclusion - `sync.Mutex`)

Gembok keamanan klasik. Bayangkan `Mutex` seperti "Kunci Toilet Umum 1 Kamar". Walau ada 5 orang *(goroutine)* yang ingin membuang hajat / merombak modifikasi nilai data, orang pertama yang sampai akan merenggut memegang kuncinya (`mu.Lock()`), masuk dan merubah nilai di dalamnya. Keempat orang lainnya yang menyusul di luar terpaksa *Terkunci Membeku* bersabar sampai orang tadi selesai modifikasi dan melepas keluar membuka gemboknya lagi (`mu.Unlock()`).

*Aturan Emas: SELALU kawinkan perintah `.Unlock()` pelucutan gembok bersamaan dekripsi pernyataan aman `defer`, tepat setelah sangasat Anda memanggil `Lock()`, agar pintu Toilet tidak sampai terkunci tergembok selamanya diakibatkan di tengah fungsi tersebut ada kecelakaan Panik sistem (*Panic Crash*).*

```go
package main

import (
    "fmt"
    "sync"
)

// Obyek Brankas Terpadu kita. Berisikan gembok dan isinya
type RekeningBankAman struct {
    gembokAman sync.Mutex
    totalSaldo float64
}

// Menambahkan uang lintas dimensi Goroutine
func (r *RekeningBankAman) SetorUangRombongan(tambahan float64) {
    r.gembokAman.Lock()         // 1. REBUT KUNCI TOILET / GEMBOK EKSKLUSIF! (Goroutine lain yg melintas akan beku antre)
    defer r.gembokAman.Unlock() // 2. JAMINAN ABSOLUT: Pastikan pintu dibuka saat pulang selesai tugas.

    // 3. Modifikasi isi memori suci
    r.totalSaldo = r.totalSaldo + tambahan
}

// Mengecek Uang, karena mengambil bacaan saja juga sangat dilarang kala orang lain sedang menulisnya!
func (r *RekeningBankAman) CekSaldoAman() float64 {
    r.gembokAman.Lock()
    defer r.gembokAman.Unlock()
    return r.totalSaldo
}

func main() {
    akunKita := RekeningBankAman{}
    var pasukanTunggu sync.WaitGroup

    // Kita kirim 500 Mesin Penambang Uang Goroutine serentak menyerbu Akun yang sama
    for i := 0; i < 500; i++ {
        pasukanTunggu.Add(1)
        go func() {
            defer pasukanTunggu.Done()
            akunKita.SetorUangRombongan(100.0) // Setor 100 perak x 500 = 50Ribu
        }()
    }

    // Tunggu mereka selesai semua merampok bank...
    pasukanTunggu.Wait()

    // Mari kita cek sisa saldo: DIJAMIN 100% adalah pas presisi 50,000.
    // Andaikan Anda merancang kode ini TANPA `gembokAman.Lock()`, uangnya pasti bocor nilainya tidak presisi!
    fmt.Printf("Buku Tabungan Valid Total Berisi: Rp %.2f\n", akunKita.CekSaldoAman())
}
```

*(Catatan Modern Lanjut: Ada jenis gembok spesialis `sync.RWMutex` (Read-Write Mutex). Sangat optimal jika Obyek memori tersebut 90% waktu hanya untuk DILIHAT nilainya (Read), bukan Dirubah (Write). Kunci pembaca `RLock` memperbolehkan puluhan orang goroutine membacanya masuk bersamman, gembok aslinya khusus menahan akses Modifikasi saja).*

---

### 2. Penghimpun Rombongan Pekerja (Wait Group - `sync.WaitGroup`)

Konsep asali di belakang bahasa Go menuntut Goroutine yang dilahirkan dari kata sakti *keyword `go`* adalah entitas latar belakang yang bebas berlarian bagaikan debu di angkasa. Jika fungsi panggung teater sentral `main()` Anda selesai ter-eksekusi, *semua* pekerja goroutine anak-anak yang belum selesai beres-beres juga akan ikut dibantai dan dimusnahkan secara zalim oleh eksekutor OS saat itu juga. Rombongan *WaitGroup* dirancang menahan nafas fungsi utama sampai "Seluruh rombongan melaporkan pekerjaan mereka lunas".

Ada tiga tarian sekuensial fungsi yang tak terpisahkan: `Add`, `Done`, dan `Wait`.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Fungsi Karyawan biasa yang akan lari di Background
func laksanakanTugasLama(IDPekerja int, rombongan *sync.WaitGroup) {
    // SYARAT PENTING: Karyawan melapor tuntas mencoret namanya dari daftar absen
    defer rombongan.Done()

    time.Sleep(1 * time.Second) // Beban tugas asinkron berat
    fmt.Printf("Agen %d menyatakan ekstraksi tuntas dan mundur!\n", IDPekerja)
}

func main() {
    var satgasRombongan sync.WaitGroup

    fmt.Println("MABES Pusat mendelegasikan perintah 3 unit Agen Paralel.")

    // Menjalankan Loop untuk 3 Unit
    for agen := 1; agen <= 3; agen++ {
        // PERATURAN BESI: Panggil fungsi .Add(1) di baris utama SEBELUM Goroutine itu melesat melompat!
        satgasRombongan.Add(1)

        // Kita wajib mengoper memori Referensi Pointer &satgasRombongan kepada Fungsi
        go laksanakanTugasLama(agen, &satgasRombongan)
    }

    fmt.Println("MABES Pusat menanti kepulangan utuh seluruh anggota Satgas...")

    // Fungsi ini membekukan sistem, menahan proses `main` tidak boleh tamat,
    // sambil menghitung mundur sinyal `Done()` dari 3 ke angka KOSONG (0).
    satgasRombongan.Wait()

    fmt.Println("\nMABES Pusat: Alhamdulillah seluruh Agen berkumpul usai tugas dengan selamat.")
}
```

---

### 3. Pengaman Saklar Tunggal Anti Panik (*Run Once - `sync.Once`*)

Anda merancang sebuah arsitektur Sistem API yang berat di mana untuk merespons kunjungan ke halaman utama, diperlukan pembuatan "Keran Kolektor Sambungan Terowongan Jaringan Redis In-Memory" yang cukup menyita beban *Load CPU Server*. Celakanya, Anda tak tahu goroutine punggawa mana yang pertama kali akan terpicu oleh pelanggan web, tapi yang Anda tahu pasti, siapapun yg datang duluan, koneksi raksasa Redis itu **HANYA BOLEH DICIPTAKAN TEPAT SATU KALI DI AWAL KEHIDUPAN SERVER**, selebihnya pemanggilan pengunjung yang lain hanya perlu me-reuse koneksinya. Panggil *Superman* kita: `sync.Once`.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

var pengamanSaklarTunggal sync.Once

// Fungsi Inisiasi berat (misal membuat Engine Mesin Database, atau baca config File raksasa)
func nyalakanMesinAplikasi() {
    fmt.Println("\n[PEMBERITAHUAN ALAMAT SERVER] >>> BOOMM!! Mesin V8 DB Dihidupkan memakan waktu 1 Detik!!!")
    time.Sleep(1 * time.Second)
    fmt.Println("[PEMBERITAHUAN ALAMAT SERVER] >>> Mesin menyala panas dan Siap berlayar.\n")
}

// Simulasi Permintaan Pengunjung Web yang Datang Bersama-Sama Berserabut
func layaniKunjunganTamu(IDTamu int, wg *sync.WaitGroup) {
    defer wg.Done()

    // KEAJAIBAN:
    // Sekalipun ada 100 orang beramai-ramai mengetuk memanggil pintu saklar Do() ini,
    // di saat milidetik yang persis saling salip,
    // paket Go akan menjamin absolut HANYA SATU TAMU yang sukses mengeksekusi fungsi "nyalakanMesinAplikasi"
    // Sedang 99 Tamu lainnya seketika tak terdeteksi memanggil fungsi itu / meleset melewatinya secara instan!
    pengamanSaklarTunggal.Do(nyalakanMesinAplikasi)

    // Sisa eksekusi normal biasa tamu (Semua 100 tamu menikmati mesin yang telah menyala)
    fmt.Printf("- Tamu [%d] telah terlayani dan meminum kopi dingin dengan nyaman.\n", IDTamu)
}

func main() {
    var kawananTamu sync.WaitGroup

    // Tembakan massal serbu dari 10 pengunjung beruntun!
    for orang := 1; orang <= 10; orang++ {
        kawananTamu.Add(1)
        go layaniKunjunganTamu(orang, &kawananTamu)
    }

    kawananTamu.Wait()
    fmt.Println("Siklus Tamu Harian habis ter-resolusi sempurna.")
}
```

---

## Bagian Lanjutan: Pola Arsitektur Konkurensi Canggih, Mutex Pembaca/Penulis, dan Kolam Memori (sync.Pool)

Menyandarkan seluruh harapan keamanan sistem hanya pada `sync.Mutex` konvensional ibarat menggunakan palu godam raksasa untuk memukul paku payung; ia berhasil secara teknis (mencegah benturan/balapan data panik memori *Race Condition*), namun ia berpotensi besar merusak efisiensi kinerja CPU paralel (mengubah sistem multi-core Anda yang hebat menjadi sempit macet seperti jalan tol satu lajur).

Pada tataran arsitektur sistem skala besar (seperti merancang *Cache In-Memory* lokal atau sistem Penyangga Sesi *Session Store*), pemahaman mengenai utilitas khusus di package ini sangat krusial untuk mempertahankan kecepatan merespons peladen web (*Latency Response*).

### 1. Kunci Spesialis Pembaca (*sync.RWMutex*) - Mencegah Kemacetan Pembacaan Data

Di dunia nyata, sangat lazim ditemui bahwa sebuah struktur data besar (misal Peta `map` daftar harga diskon Katalog Produk) sangat jarang diubah harganya (mungkin hanya di-Update setiap 1 Jam sekali), namun data Peta tersebut ditanyakan dan **DIBACA secara serentak bersamaan oleh 50.000 Goroutine pembeli** setiap menitnya.

Jika Anda melindungi Peta Harga tersebut menggunakan `sync.Mutex` standar (`.Lock()`), maka ketika 50.000 pembeli mencoba mengecek harga, mereka semua **DIHARUSKAN ANTRE** satu per satu masuk ke gerbang fungsi karena Gembok Standar melarang lebih dari 1 orang masuk ke memori meskipun mereka *hanya numpang melihat/membaca nilai tanpa merubah apapun*! Performa server akan anjlok drastis (*Bottleneck*).

Kehadiran `sync.RWMutex` (*Read-Write Mutex*) merupakan solusi surgawi. Gembok ini memiliki 2 jenis Kunci Ganda yang terpisah: Kunci Pembaca Murni (`RLock`) dan Kunci Penulis Perusak (`Lock`).

**Aturan RWMutex:**
*   **`RLock()` (Gembok Melihat):** Ribuan, jutaan goroutine boleh memegang kunci RLock secara BERSAMAAN tanpa perlu saling antre! Semuanya bisa masuk mengekstrak data Peta beramai-ramai di detik yang sama, memaksimalkan kekuatan multi-core CPU.
*   **`Lock()` (Gembok Menulis Mutlak):** Ketika petugas Admin tiba-tiba datang ingin Merubah Harga Diskon, ia akan memanggil fungsi eksekusi ini. Gembok Penulis Mutlak ini akan menyalakan Alarm Berbahaya: "Tutup semua pintu! Larang pembaca baru masuk! Tunggu pembaca yang masih di dalam keluar! Lalu masukkan Petugas Admin sendirian secara eksklusif absolut!".
    Seketika ia mengunci, tidak ada goroutine lain (baik pembaca maupun penulis) yang boleh melintas masuk.

```go
// Simulasi Gudang Diskon Pusat
// type GudangDiskonPusat struct {
//    kunciPintu  sync.RWMutex // Gembok Canggih Dua Fungsi
//    petaHargaDB map[string]int
// }

// func (g *GudangDiskonPusat) DapatkanHarga(namaKatalog string) int {
//    // RIBUAN GOROUTINE BISA LEWAT BARIS INI BERSAMAAN SEKALIGUS!
//    g.kunciPintu.RLock()
//    defer g.kunciPintu.RUnlock()
//    return g.petaHargaDB[namaKatalog]
// }

// func (g *GudangDiskonPusat) RevisiHargaBaru(namaKatalog string, hargaBaru int) {
//    // MENENDANG SEMUA GOROUTINE PEMBACA DEMI KEAMANAN INTEGRITAS MODIFIKASI!
//    g.kunciPintu.Lock()
//    defer g.kunciPintu.Unlock()
//    g.petaHargaDB[namaKatalog] = hargaBaru
// }
```

### 2. Kolam Penampungan Memori Bebas Sampah (*sync.Pool*)

Sistem Pengutip Sampah (*Garbage Collector / GC*) Go adalah salah satu mesin otomatisasi *memory management* tercepat dan terdepan di kelas industri. Namun kecepatan GC Go bukan berarti Anda bisa memproduksi sampah memori sesuka hati tanpa konsekuensi hukuman pelambatan (*CPU GC Pause/Overhead*).

Jika Anda membuat sebuah API Server pencetak PDF Resi Transaksi, yang di dalamnya setiap kali sebuah Request masuk, Anda mengalokasikan (menciptakan) penyangga memori *Array Slice Bytes* sebesar 5 MB untuk merakit kerangka gambar laporan PDF tersebut.
Bayangkan 1000 Transaksi / Detik (RPS). Anda menciptakan dan membuang sampah 5 GB RAM setiap detiknya! GC Go akan terbakar sibuk membersihkan sisa jasad 1000 PDF yang telah terkirim tadi, menyiksa beban Server yang sangat kritis.

**Jawaban dari Surga: `sync.Pool` (Daur Ulang Obyek Memori Temporer)**

Alih-alih membuang kerangka penampung Array Bytes PDF tadi setelah dikirim ke Web Browser klien, kita *"Mencucinya hingga kosong"* lalu *"Memasukkannya kembali ke dalam Laci Gudang / Pool"* agar Request Pengunjung HTTP berikutnya bisa *Meminjamnya* untuk dipakai ulang, sehingga Server Anda **TIDAK PERLU MENGALOKASIKAN MEMORI BARU SAMA SEKALI DARI OS** (Mewujudkan utopia magis *Zero-Allocation Memory Profile*).

```go
// MENDESAIN GUDANG DAUR ULANG MEMORI GLOBAL
// var penampungWadahBuffer = sync.Pool{
//    New: func() interface{} {
//        return new(bytes.Buffer)
//    },
// }

// Saat butuh:
// wadahPinjaman := penampungWadahBuffer.Get().(*bytes.Buffer)
// defer penampungWadahBuffer.Put(wadahPinjaman) // Wajib dikembalikan
// wadahPinjaman.Reset() // Dicuci bersih dulu
```

Penguasaan `sync.RWMutex` dan kejeniusan optimisasi memori alokasi-nol (`sync.Pool`) menjadi standar verifikasi utama dalam sesi wawancara perekayasa senior (*Senior Backend Engineer*) di ranah pengembangan ekosistem Go (Golang) bertaraf arsitektur perusahaan skala global.

## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
