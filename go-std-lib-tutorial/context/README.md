# Modul: `context`

## Ringkasan
Package `context` menyediakan mekanisme standar industri untuk menyebarluaskan sinyal pembatalan (*cancellation signals*), mengatur batas waktu (*deadlines* / *timeouts*), serta membawa nilai-nilai khusus (*request-scoped values*) ke seluruh batasan antarmuka API (*API boundaries*) dan fungsi-fungsi *goroutine* yang saling terikat (konkuren) di dalam satu alur eksekusi aplikasi (*call graph*).

## Penjelasan Lengkap (Fungsi & Tujuan)
Dalam arsitektur *microservice* modern, sebuah HTTP request yang masuk dari klien web (browser) seringkali tidak hanya diselesaikan oleh satu fungsi tunggal. Request tersebut memicu rentetan pemanggilan API: fungsi *Controller* memanggil fungsi *Service*, fungsi *Service* memanggil dua database secara paralel (bercabang ke dalam 2 *Goroutine*), dan juga memanggil API pihak ketiga (*payment gateway*).

Bayangkan jika setelah 2 detik berlalu, pengguna (klien web) tiba-tiba bosan dan menutup tab browser mereka. Koneksi HTTP terputus. Jika server Go Anda tidak memiliki mekanisme koordinasi, kedua *Goroutine* yang sedang men-query database dan memanggil *payment gateway* itu akan **terus berjalan tanpa henti** membuang-buang sumber daya CPU, koneksi database, dan memori RAM server, meskipun hasilnya nanti tidak akan dikirimkan kepada siapapun! Inilah masalah yang diselesaikan secara elegan oleh package `context`.

Context bertindak sebagai "tali pengikat" transparan yang digenggam oleh semua goroutine tersebut. Jika "ujung tali" di level teratas (misalnya koneksi HTTP) dipotong atau dibatalkan, sinyal tersebut merambat sangat cepat ke seluruh rantai fungsi di bawahnya, memerintahkan mereka semua untuk segera berhenti bekerja (*fail-fast*).

**Tujuan dan Fungsi Utama:**
1.  **Penerusan Sinyal Batal (Propagated Cancellation):** Menghentikan operasi paralel yang tidak lagi diperlukan secara serentak, mencegah kebocoran *Goroutine* (*Goroutine Leaks*).
2.  **Pemaksaan Batas Waktu (Deadlines & Timeouts):** Menetapkan SLA (*Service Level Agreement*) yang ketat. Misalnya: "Panggilan database ini harus selesai dalam 5 detik, jika tidak, langsung batalkan query-nya dan kembalikan error *Timeout* ke klien".
3.  **Pengangkutan Data Lintas Lapisan (Context Values):** Membawa data otentikasi (misal: ID pengguna yang sedang *login* hasil dari *middleware*) masuk menembus lapisan kode bisnis (*business logic*) hingga ke lapisan basis data tanpa perlu menambahkan banyak parameter ekstra pada definisi *function signature* Anda.

**Aturan Emas Penggunaan Context di Go:**
*   **Jangan menyimpan `Context` di dalam tipe Struct!** Context harus secara eksplisit diteruskan sebagai **parameter pertama** di setiap fungsi yang membutuhkannya (biasanya dinamakan `ctx`).
*   **Jangan meneruskan `nil` Context**, meskipun sebuah fungsi meminta Context tapi Anda belum membutuhkannya, gunakan selalu `context.TODO()`.
*   **Gunakan *Values* (nilai bawaan) HANYA untuk data cakupan-request (*request-scoped data*)** seperti *User ID* atau *Trace ID*. Jangan sekali-kali menggunakannya untuk memasukkan dependensi opsional atau konfigurasi global (seperti *koneksi database*), karena sifatnya yang *untyped* (menyebabkan program kehilangan *type-safety*).

---

## Daftar Fungsi Umum dan Cara Penggunaannya Secara Mendalam

### 1. Titik Awal (Background & TODO)

Setiap *Context* merupakan sebuah rantai yang memiliki akar/induk. Akar dari seluruh *context* selalu diciptakan dari `Background` atau `TODO`. Mereka tidak memiliki batas waktu dan tidak pernah dibatalkan.

*   **`context.Background()`**: Ini adalah *context* dasar utama. Biasanya digunakan di level tertinggi program (dalam fungsi `main()`, fungsi init, atau pada tes unit). Context ini adalah "orang tua" dari semua operasi turunan di bawahnya.
*   **`context.TODO()`**: Secara fungsional sama persis dengan `Background`. Ini digunakan hanya sebagai "penanda semantik" bagi programmer bahwa "Saya belum tahu *context* apa yang harus di-passing di sini dari fungsi pemanggil, jadi sementara saya gunakan TODO dulu".

```go
package main

import (
    "context"
    "fmt"
)

func prosesData(ctx context.Context, data string) {
    // Simulasi fungsi yang menerima context sebagai parameter PERTAMA (aturan emas)
    fmt.Println("Memproses data:", data)
}

func main() {
    // Memulai rantai eksekusi root dari sistem
    ctxRoot := context.Background()

    // Meneruskan context tersebut ke bawah
    prosesData(ctxRoot, "Muatan Transaksi")
}
```

---

### 2. Membatalkan Eksekusi secara Manual (WithCancel)

Digunakan ketika kita memiliki sekumpulan pekerja (*workers*), dan kita ingin memiliki "tombol panik" untuk menghentikan mereka semua secara manual apabila suatu kondisi terpenuhi (misalnya salah satu pekerja gagal, maka batalkan semua pekerja lainnya agar tidak membuang resource).

Fungsi `context.WithCancel(parent)` menciptakan anak dari *context parent* yang diberikan, dan mengembalikan anak tersebut **beserta** sebuah fungsi kecil `cancel()`. Fungsi `cancel()` ini jika dipanggil, akan menutup saluran (channel) `Done()` milik si anak dan *seluruh* cucu-cucunya.

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func pekerja(ctx context.Context, id int) {
    for {
        select {
        case <-ctx.Done():
            // Sinyal Batal tertangkap!
            fmt.Printf("Pekerja %d menerima perintah berhenti. Error: %v\n", id, ctx.Err())
            return // Keluar dari goroutine, mencegah kebocoran memori

        default:
            // Jika belum dibatalkan, terus bekerja
            fmt.Printf("Pekerja %d sedang bekerja...\n", id)
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctxInduk := context.Background()

    // Kita buat anak yang bisa dibatalkan secara manual
    ctxBisaBatal, tombolBatal := context.WithCancel(ctxInduk)

    // PASTIKAN tombolBatal DITEKAN ketika fungsi main selesai, ini sangat krusial
    // agar anak-anak goroutine tidak menjadi zombie yatim piatu di memory!
    defer tombolBatal()

    // Memutar 3 Goroutine Pekerja
    for i := 1; i <= 3; i++ {
        go pekerja(ctxBisaBatal, i)
    }

    // Biarkan mereka bekerja 2 detik
    time.Sleep(2 * time.Second)

    fmt.Println("-> Waktunya menghentikan semua secara serentak!")
    // MENEKAN TOMBOL PANIK!
    tombolBatal()

    // Tunggu sedikit agar pesan log mereka sempat tercetak sebelum program keluar
    time.Sleep(1 * time.Second)
}
```

---

### 3. Eksekusi dengan Batasan Waktu Otomatis (WithTimeout / WithDeadline)

Ini adalah penggunaan paling populer dari package `context` dalam operasi mikroservis (*API calls*, *Database queries*). Anda ingin operasi tersebut otomatis gagal/batal jika tidak selesai dalam $X$ detik. Hal ini mencegah aplikasi "membeku/hang" tanpa henti.

*   **`context.WithTimeout(parent, duration)`**: Anak context otomatis batal jika waktu *duration* (misal 5 detik) telah berlalu.
*   **`context.WithDeadline(parent, time)`**: Mirip timeout, tapi menggunakan waktu jam absolut (misal: otomatis batal persis pada pukul 12:00 siang ini).

```go
package main

import (
    "context"
    "fmt"
    "time"
)

// Simulasi fungsi pemanggilan API ke pihak eksternal (Stripe/Payment) yang lambat
func panggilAPIPaymentLuar(ctx context.Context) error {
    saluranSukses := make(chan string)

    // Goroutine yang melakukan tugas sebenarnya (memakan waktu 4 detik)
    go func() {
        time.Sleep(4 * time.Second)
        saluranSukses <- "API Payment Berhasil Merespons!"
    }()

    // Pemilihan sinyal balapan (Race condition handling)
    select {
    case hasil := <-saluranSukses:
        fmt.Println("Laporan:", hasil)
        return nil

    case <-ctx.Done():
        // Terjadi jika waktu timeout habis SEBELUM saluranSukses mengirimkan balasan
        fmt.Println("OPERASI DIGAGALKAN! API eksternal terlalu lamban.")
        // Mengembalikan alasan pembatalan (misal: "context deadline exceeded")
        return ctx.Err()
    }
}

func main() {
    // Kita buat aturan ketat: Operasi APAPUN di bawah ini MAKSIMAL HANYA BOLEH 2 DETIK!
    ctxTimeout, fungsiBatalkan := context.WithTimeout(context.Background(), 2 * time.Second)

    // Jangan lupa defer! Membatalkan lebih awal jika eksekusi selesai kurang dari 2 detik,
    // akan membebaskan memory timer di sistem Go lebih cepat.
    defer fungsiBatalkan()

    fmt.Println("Mulai mencoba memanggil API...")

    // Kita oper context ber-timeout ini ke fungsi lamban di atas.
    // Karena API memakan waktu 4 detik, sedangkan batas kita 2 detik, pasti akan GAGAL (Timeout).
    err := panggilAPIPaymentLuar(ctxTimeout)
    if err != nil {
        fmt.Println("Hasil Akhir Transaksi: ERROR ->", err)
    } else {
        fmt.Println("Hasil Akhir Transaksi: SUKSES!")
    }
}
```

---

### 4. Menyelipkan Muatan Data Lintas Batas (WithValue)

Sangat berguna saat Anda membuat Web Server, di mana Anda memiliki fungsi penengah (*Middleware*) yang memverifikasi kata sandi/JWT token pengguna. Setelah Middleware tahu siapa pengguna yang berhasil masuk (misalnya `UserID: 888`), Anda ingin meneruskan data `UserID` ini menembus jauh ke dalam fungsi Controller atau Database Repository Anda tanpa merusak *Interface Function* aslinya.

**Peringatan Tipe (Type Safety):** Karena struktur internal *context* memakai tipe kosong `interface{}`, penggunaan tipe data bawaan murni (seperti `string` "userId") sebagai kunci (key) dilarang keras karena bisa menyebabkan bentrokan (*collision*) jika ada *package* luar lain yang juga secara kebetulan menggunakan *string key* "userId". Selalu buat tipe data khusus kustom untuk mengamankan Kunci.

```go
package main

import (
    "context"
    "fmt"
)

// 1. CARA AMAN & BENAR: Definisikan tipe kustom tak diekspor untuk Kunci Context.
// Hal ini mencegah tabrakan/collision Kunci dari package-package asing lain.
type kunciKonteksTipe string
const (
    kunciIdentitasUser kunciKonteksTipe = "userID"
    kunciStatusAkses   kunciKonteksTipe = "peran"
)

// Fungsi lapisan paling dalam (Simulasi Database Repository)
func cariDataSensitifDiDB(ctx context.Context) {
    // Membongkar dan mengeluarkan nilai yang diselipkan.
    // .Value() mengembalikan nilai bertipe `any` (interface kosong).
    // Kita harus memaksakan konversinya (Type Assertion) ke tipe aslinya.
    idPengguna := ctx.Value(kunciIdentitasUser)

    // Pastikan nilainya ada!
    if idPengguna == nil {
        fmt.Println("Gagal ke database: Penelusup tak dikenal!")
        return
    }

    peran := ctx.Value(kunciStatusAkses).(string) // Type assertion ke String

    fmt.Printf("Memasuki Database -> Menjalankan Query khusus untuk UserID [%d] dengan Hak Akses [%s].\n", idPengguna, peran)
}

func main() {
    // Lapis Pertama: Terima Kunjungan Web Kosong (Siklus awal Request)
    ctxRoot := context.Background()

    // Lapis Kedua: Melewati Middleware Autentikasi
    // JWT divalidasi, dan diketahui identitas penelpon.
    // Kita bungkus context lama menjadi context baru yang "mengandung ID"
    ctxDenganID := context.WithValue(ctxRoot, kunciIdentitasUser, 54321)

    // Kita bungkus sekali lagi menjadi context berlapis yang "juga mengandung Peran"
    ctxLengkapPenuh := context.WithValue(ctxDenganID, kunciStatusAkses, "Administrator Utama")

    // Lapis Terakhir: Kita terjun dalam panggil Controller ke Database.
    // Data "menembus" masuk secara elegan tanpa perlu parameter baru `cariData(id, peran)`
    cariDataSensitifDiDB(ctxLengkapPenuh)
}
```

---

## Bagian Lanjutan: Anti-Pola (Anti-Patterns), Desain API, dan Jebakan Memori Goroutine

Dalam pengembangan sistem terdistribusi, *package* `context` adalah tulang punggung dari stabilitas *Service Level Agreement* (SLA). Meskipun sederhana secara konseptual, penggunaannya yang salah sering kali menjadi penyebab utama masalah produksi yang sulit dilacak (seperti *deadlock* sporadis, koneksi database yang mengambang, atau variabel *tracing* yang hilang di tengah jalan).

### 1. Aturan Emas Penyampaian Context (Pewarisan Parameter)

Perdebatan paling umum bagi pengembang Go pemula adalah: *"Di mana saya harus menyimpan objek context ini? Bisakah saya menaruhnya di dalam `struct` layanan saya?"*

**SANGAT DILARANG MENYIMPAN CONTEXT DI DALAM STRUCT!**
```go
// ANTI-PATTERN SANGAT BERBAHAYA (JANGAN LAKUKAN INI)
type UserService struct {
    // db  *sql.DB
    // ctx context.Context // SALAH BESAR!
}

// Saat HTTP Request masuk, Anda mengubah state global atau state objek
// func (s *UserService) HandleUser(r *http.Request) {
    // s.ctx = r.Context() // Kondisi Balapan (Race Condition) jika 100 user mengakses bersamaan!
// }
```

Context bersifat **Request-Scoped** (Siklus hidupnya eksklusif terikat pada SATU permintaan spesifik pengguna). Jika Anda menyimpannya di struct `UserService` (yang biasanya bertindak sebagai *Singleton* untuk melayani semua request), pengguna A akan membatalkan request pengguna B secara tidak sengaja, dan program Anda akan panik di bawah beban bersamaan.

**Pola Praktik Terbaik (Best Practice):**
Context HARUS SELALU menjadi **parameter pertama** secara eksplisit di setiap definisi fungsi (*Function Signature*) yang melakukan operasi berpotensi lambat (seperti I/O jaringan, kueri database, atau RPC eksternal).

```go
// POLA BENAR DAN AMAN TINGKAT INDUSTRI
// type UserService struct {
//    db *sql.DB // Koneksi DB bersifat global/singleton, aman di struct
// }

// Fungsi selalu bersih, mandiri, dan MENGOPER context dari pemanggilnya
// func (s *UserService) GetUserProfile(ctx context.Context, userID int) (*Profile, error) {
//    row := s.db.QueryRowContext(ctx, "SELECT nama FROM users WHERE id=?", userID)
//    ...
// }
```

### 2. Bahaya Membocorkan Memori dengan `WithTimeout` (Lupa Memicu Defer Cancel)

Salah satu bug memori paling klasik di bahasa Go berkaitan dengan pembuatan pengatur waktu (Timer).
Ketika Anda memanggil `ctx, cancel := context.WithTimeout(parent, 5*time.Second)`, sistem Go di balik layar akan mengalokasikan memori untuk sebuah objek pengatur waktu (Timer) internal yang akan tertidur menunggu selama 5 detik.

Bayangkan fungsi Anda ternyata sukses memanggil database dengan sangat cepat dalam waktu 0.1 detik.
Fungsi Anda kemudian `return` dan menganggap pekerjaan selesai. Lantas apa yang terjadi pada *Timer 5 detik* yang masih tertidur di memori Go tadi?
Timer tersebut **AKAN TERUS HIDUP MENJADI SAMPAH ZOMBIE** selama 4.9 detik berikutnya, menyandera RAM server Anda!

Jika Anda melayani 10.000 *Request per Second* (RPS), Anda baru saja menciptakan 50.000 timer zombie berjalanan berpotensi menyebabkan *Out Of Memory (OOM) Kill* dari sistem operasi Linux.

**Hukum Mutlak Pencegahan Kebocoran (Leak Prevention):**
ANDA WAJIB MAHA WAJIB MEMANGGIL FUNGSI `cancel()` BEGITU PEKERJAAN ANDA SELESAI LEBIH AWAL!
Cara teraman adalah menggunakan perintah sakti `defer` persis di baris berikutnya.

```go
// func (s *OrderService) CekGudangLuarNegeri() error {
//    // Kita membuat context dengan batas waktu 5 detik
//    ctxBatasWaktu, batalkanCepat := context.WithTimeout(context.Background(), 5*time.Second)
//
//    // ATURAN BESI: Jamin bahwa begitu fungsi ini berakhir sukses/gagal secara instan,
//    // batalkanCepat() dipanggil sehingga Go seketika itu juga menghancurkan objek Timer 5 detiknya!
//    defer batalkanCepat()
//
//    // Operasi ini misal cuma butuh waktu 0.2 detik selesai!
//    err := s.apiPihakKetiga.CekStok(ctxBatasWaktu)
//    return err
// }
```

### 3. Mengakali Pembatalan Jaringan (Detached Contexts)

Terkadang, ada sebuah operasi di mana Anda *ingin* meneruskan nilai data otentikasi (JWT User ID) dari context HTTP Request klien, **TETAPI** Anda **TIDAK INGIN** operasi asinkron latar belakang Anda ikut mati dibatalkan manakala klien tersebut tiba-tiba iseng menekan tombol silang "Batal / Close Tab" di peramban Browser-nya.

Misal: Klien web mengeklik tombol "Bayar Pesanan". Klien tidak sabaran dan menutup aplikasi Androidnya.
Jika sistem Pembayaran Anda terikat pada *Context Klien*, maka panggilan "Potong Saldo Kartu Kredit" di *goroutine background* Anda akan ikut Batal di tengah jalan! Klien barangnya tidak terbayar, tapi API Bank mungkin sudah separuh jalan memproses. Kekacauan Status (*Inconsistent State*) terjadi!

**Solusi: Konsep Pelepasan Keterikatan Pembatalan (WithoutCancel)**
Di Go versi modern (1.21+), disediakan fungsi elegan `context.WithoutCancel(parentCtx)`.
Fungsi ini menyalin sebuah Context Baru yang **tetap membawa** semua data titipan otentikasi (*Values*) dari induk lamanya, namun ia **memotong** tali pembatalan dan batas waktunya!

```go
// func HandlerBayarTagihan(w http.ResponseWriter, r *http.Request) {
//    ctxKlienTerikat := r.Context()
//
//    // Kita ingin melanjutkan proses potong tagihan ke Pihak Bank di Goroutine Background.
//    // KITA HARUS MEMOTONG TALI PEMBATALAN KLIEN, TAPI MEMBAWA SERTA DATA "User_ID" nya!
//
//    // Go 1.21+:
//    ctxAmanLatarBelakang := context.WithoutCancel(ctxKlienTerikat)
//
//    go func(ctxAman context.Context) {
//        // Biar klien menutup HP-nya berkali-kali,
//        // Proses eksekusi pemotongan uang ini HANTAM TERUS PANTANG MUNDUR!
//        // Dan UserID masih bisa dibaca dengan sukses!
//        userID := ctxAman.Value("user_id")
//        // ProsesKeBankGatewayLuar(ctxAman, userID, tagihanHarga)
//    }(ctxAmanLatarBelakang)
//
//    // Mengembalikan sukses instan ke layar Klien!
//    w.Write([]byte("Pembayaran sedang di proses di antrian latar belakang Server!"))
// }
```
*Catatan Historis: Sebelum Go 1.21, programmer tingkat lanjut harus membuat tipe Struct kustom mereka sendiri secara manual yang membungkus murni isi metode `Value()`, namun sengaja menghilangkan implementasi metode `Done()` dan `Deadline()`.*

Penguasaan teknik pembelahan tali `context` seperti contoh di atas menandakan bahwa pemahaman Anda terhadap arsitektur eksekusi bersilangan asinkron di bahasa Go telah mencapai keunggulan paripurna, siap menghadapi anomali kondisi kemacetan jaringan skala Enterprise yang brutal di lapangan produksi.

## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
