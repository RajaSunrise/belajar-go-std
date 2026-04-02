# Modul: `errors`

## Ringkasan
Package `errors` mengimplementasikan serangkaian fungsi fundamental untuk merakit, mengevaluasi konteks, membongkar rantai pembungkusan (*Error Wrapping*), dan memvalidasi silsilah dari entitas *error* di dalam bahasa pemrograman Go.

## Penjelasan Lengkap (Fungsi & Tujuan)
Di dalam bahasa pemrograman berorientasi objek tradisional (seperti Java, C#, atau Python), penanganan kesalahan dilakukan melalui mekanisme *Exception Handling* (blok `try-catch`). Saat sebuah kesalahan fatal terjadi, eksekusi program langsung "terlempar" keluar dari fungsi tersebut, mencari blok `catch` terdekat untuk menanganinya.

Desainer Go mengambil filosofi yang sepenuhnya bertolak belakang. Di Go, **sebuah error hanyalah nilai (*value*) biasa**. Error diperlakukan sama persis dengan variabel *integer* atau *string*. Ia dikembalikan oleh sebuah fungsi sebagai *return value* kedua (atau terakhir), dan menjadi tanggung jawab pemanggil fungsi tersebut untuk langsung memeriksa nilainya (menggunakan blok sakral `if err != nil`). Pendekatan eksplisit ini memaksa *programmer* untuk memikirkan penanganan kegagalan di setiap langkah eksekusi, menjadikan aplikasi Go jauh lebih kokoh, stabil, dan perilakunya sangat mudah diprediksi.

Objek *error* di Go sebenarnya hanyalah sebuah Antarmuka (*Interface*) sangat sederhana yang hanya memiliki satu metode: `Error() string`. Package `errors` menyediakan implementasi standar dan utilitas tingkat lanjut (diperkenalkan pada Go 1.13) untuk mengelola antarmuka ini secara profesional.

**Tujuan dan Fungsi Utama:**
1.  **Pembuatan Error Dasar:** Membuat instance error baru yang mengangkut pesan statis deskriptif melalui fungsi seringan bulu `errors.New()`.
2.  **Rantai Pembungkusan Konteks (Error Wrapping):** Ketika error bergelembung naik dari lapisan bawah (misal: gagal koneksi database) menuju lapisan atas (misal: gagal mendaftarkan user), sangat penting untuk menambahkan "konteks" cerita pada error tersebut tanpa menghilangkan jejak error aslinya.
3.  **Inspeksi Kekerabatan (errors.Is):** Menggantikan operator perbandingan primitif `==`. Fungsi ini sanggup menembus lapisan-lapisan *error* yang telah terbungkus (*wrapped*) untuk mendeteksi apakah di suatu tempat di dalam silsilahnya, terdapat jenis error spesifik yang sedang kita cari (misalnya mencari error `ErrNotFound` atau `ErrTimeout`).
4.  **Ekstraksi Tipe Kustom (errors.As):** Jika kita mendefinisikan tipe *error* kustom buatan sendiri (berupa `struct` yang mengimplementasikan antarmuka error dan memiliki *field* data tambahan seperti `HTTPStatusCode`), fungsi ini memungkinkan kita untuk "memancing" dan mengekstrak *struct* spesifik tersebut dari dalam tumpukan error yang kusut.

**Mengapa menggunakan `errors`?**
Jika Anda tidak menggunakan utilitas `errors.Is` atau `errors.As` di era Go modern, dan masih mengandalkan pemeriksaan string kuno seperti `strings.Contains(err.Error(), "timeout")`, Anda sedang membangun bom waktu. Pesan teks error bisa berubah kapan saja di versi *library* masa depan, namun identitas struktur (Tipe Error) akan tetap abadi. Penggunaan package ini adalah ciri khas *Senior Go Engineer*.

---

## Daftar Fungsi Umum dan Cara Penggunaannya Secara Mendalam

### 1. Menciptakan Error Sederhana & Sentinel Error (errors.New)

Cara paling instan memberitahukan bahwa fungsi Anda menolak input yang diberikan. Pendekatan praktik terbaik (*Best Practice*) di Go adalah mendeklarasikan error-error yang sering terjadi sebagai variabel *Global* (sering disebut *Sentinel Errors*) dengan awalan nama `Err`.

```go
package main

import (
    "errors"
    "fmt"
)

// 1. PRAKTIK TERBAIK: Deklarasikan Sentinel Error di tingkat package.
// Ini memungkinkan fungsi-fungsi lain di proyek Anda mengimpor dan mencocokkan nilainya nanti.
var ErrStokKosong = errors.New("gagal: kuantitas stok barang tidak mencukupi")
var ErrKartuDitolak = errors.New("gagal: bank menolak transaksi kartu kredit")

func prosesPembelian(kuantitas int) error {
    if kuantitas > 10 {
        // Mengembalikan Sentinel Error yang telah disepakati
        return ErrStokKosong
    }
    if kuantitas < 1 {
        // Atau membuat error dadakan yang spesifik
        return errors.New("kuantitas tidak boleh kurang dari 1")
    }
    return nil // Nilai ajaib yang menandakan fungsi sukses sempurna
}

func main() {
    // Uji coba membeli 15 barang
    errTrans := prosesPembelian(15)

    // Pola sakral Go yang akan Anda tulis jutaan kali
    if errTrans != nil {
        fmt.Println("Transaksi Dibatalkan:", errTrans.Error())
        // Mengembalikan HTTP 400 Bad Request, dsb...
    } else {
        fmt.Println("Transaksi Sukses!")
    }
}
```

---

### 2. Membungkus Error dengan Konteks (Wrapping & errors.Unwrap)

Seringkali, error `ErrStokKosong` saja tidak cukup bagi tim *DevOps* untuk melacak masalah di *log server*. "Stok kosong untuk produk yang mana? Kapan?". Kita perlu membungkusnya. Kita memanfaatkan teman karib package ini, yaitu `fmt.Errorf` dengan kata kunci rahasia `%w` (Wrap), bukan `%s`!

```go
package main

import (
    "errors"
    "fmt"
)

var ErrKoneksiDB = errors.New("terjadi kegagalan koneksi ke database")

// Lapisan bawah (Misal fungsi Database)
func simpanUserKeDB() error {
    // Simulasi DB mati
    return ErrKoneksiDB
}

// Lapisan tengah (Misal fungsi Bisnis)
func registrasiUser(nama string) error {
    errDB := simpanUserKeDB()
    if errDB != nil {
        // KITA MEMBUNGKUSNYA!
        // %w akan menjaga struktur asli ErrKoneksiDB tetap hidup di dalam perut error baru ini.
        return fmt.Errorf("gagal mendaftarkan user [%s] karena: %w", nama, errDB)
    }
    return nil
}

func main() {
    err := registrasiUser("Budi")

    if err != nil {
        // Output akan menyambung: "gagal mendaftarkan user [Budi] karena: terjadi kegagalan..."
        fmt.Println("LOG SYSTEM:", err)

        // Membongkar lapisannya 1 tingkat (Jarang digunakan secara langsung, tapi bagus untuk dipahami)
        errorAsli := errors.Unwrap(err)
        fmt.Println("Penyebab Akar (Root Cause):", errorAsli)
    }
}
```

---

### 3. Inspeksi Identitas Kekerabatan (errors.Is)

Melanjutkan contoh di atas, bayangkan Anda memanggil `registrasiUser()`. Anda mendapat *error* kembalian, dan Anda ingin memunculkan halaman "Mohon Coba Lagi" HANYA JIKA error tersebut disebabkan oleh masalah `ErrKoneksiDB`.

Bagaimana cara mengeceknya? Menggunakan `==` akan gagal karena error itu sudah berubah wujud karena dibungkus oleh *fmt.Errorf*! Solusinya adalah `errors.Is`.

Fungsi `errors.Is(errYgDidapat, targetErr)` akan membongkar seluruh lapisan *wrapping* secara rekursif (*looping* ke dalam perut error), satu per satu, dan mengecek apakah ada bagian dari rantai tersebut yang identik dengan `targetErr`.

```go
package main

import (
    "errors"
    "fmt"
)

// (Menggunakan Sentinel Error dari contoh sebelumnya)
var ErrKoneksiDB = errors.New("db terputus")

func eksekusi() error {
    // Membungkus 2 lapis!
    errLapis1 := fmt.Errorf("level repository gagal: %w", ErrKoneksiDB)
    errLapis2 := fmt.Errorf("level controller gagal: %w", errLapis1)
    return errLapis2
}

func main() {
    errMisterius := eksekusi()

    // PENGECEKAN KUNOM YANG SALAH:
    // if errMisterius == ErrKoneksiDB { ... } -> PASTI FALSE! Karena wujudnya sudah berbeda.

    // PENGECEKAN MODERN YANG BENAR:
    if errors.Is(errMisterius, ErrKoneksiDB) {
        fmt.Println(">> TINDAKAN: Menampilkan notifikasi 'Maaf server sedang Maintenance' kepada pengguna.")
    } else {
        fmt.Println(">> TINDAKAN: Menampilkan pesan error umum.")
    }
}
```

---

### 4. Ekstraksi Data Tipe Kustom Berwujud Error (errors.As)

Terkadang sebuah teks *string* tidaklah cukup. Anda butuh error yang membawa muatan data kaya fitur (*Rich Data*), misalnya membawa kode status HTTP spesifik (404, 500) agar lapisan REST API Anda tahu kode apa yang harus dikembalikan ke *Browser*.

Anda membuat tipe *Struct* sendiri yang mengimplementasikan `Error() string`. Lalu di lapisan atas, Anda ingin mengekstrak (mencopot paksa) *struct* kaya fitur tersebut keluar dari bungkusannya untuk Anda ambil data kodenya menggunakan `errors.As`.

```go
package main

import (
    "errors"
    "fmt"
)

// 1. Membuat Tipe Error Kustom (Rich Error Struct)
type HttpError struct {
    StatusCode int
    Message    string
}

// 2. Wajib membuat metode Error() agar struct ini sah dipandang sebagai "error" oleh bahasa Go.
func (e *HttpError) Error() string {
    return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

// Fungsi simulasi pencarian
func cariProfil() error {
    // Terjadi error, kita mengembalikan struct pointer kustom kita
    errAsli := &HttpError{StatusCode: 404, Message: "Profil pengguna tidak ditemukan di server"}

    // Kemudian dibungkus oleh fungsi pemanggil di atasnya
    return fmt.Errorf("gagal memuat dashboard: %w", errAsli)
}

func main() {
    errTerima := cariProfil()
    if errTerima == nil {
        return
    }

    // Kita sediakan wadah KOSONG dengan tipe pointer kustom target kita
    var targetKustom *HttpError

    // Kita perintahkan errors.As untuk menelisik ke dalam perut errTerima.
    // Jika ia menemukan struktur yang tipenya SAMA dengan wadah kita,
    // ia akan MENG-COPY nilainya ke dalam wadah tersebut (via referensi '&'), lalu mereturn True!

    cocok := errors.As(errTerima, &targetKustom)

    if cocok {
        // HORE! Kita berhasil mengambil datanya, sekarang kita bisa mengakses field kustomnya!
        fmt.Printf("Tertangkap Error Spesifik! Mengirim kode HTTP %d ke Browser.\n", targetKustom.StatusCode)
        fmt.Printf("Pesan Rahasia: %s\n", targetKustom.Message)
    } else {
        // Jika tidak cocok, berarti ini adalah error string biasa yang membosankan
        fmt.Println("Error Biasa:", errTerima)
    }
}
```

---

## Bagian Lanjutan: Rantai Kesalahan Kompleks Berkondisi (Error Joining), Stack Trace, dan Praktik Industri Kritis

Pemahaman konsep `errors.Is` dan `errors.As` di Go ibarat menguasai pisau bedah tingkat menengah. Akan tetapi, realita sistem mikroservis korporat (*Enterprise Microservices*) tidak hanya memberikan diagnosis tunggal. Sebuah permintaan pengunggahan fail (*file upload*) ke peladen dari pengguna dapat berhadapan pada tiga kesalahan serentak: "Format Fail Terlarang", ditambah "Ukuran Batas Kouta Sisa Terlampaui", dan diperburuk oleh kegagalan peladen penyimpanan Cloud S3 internal yang melemparkan "I/O Connection Timeout". Menghadapkan ketiga malapetaka tersebut secara bersamaan utuh kepada mesin pengawas sistem log membutuhkan peralatan kelas dewa di ekosistem Go.

### 1. Evolusi Agregasi Kesalahan Massal di Era Modern: `errors.Join` (Fitur Go 1.20+)

Sebelum Go versi rilis 1.20 diterbitkan, programmer terpaksa mengandalkan paket eksternal canggih pihak ketiga bermerk seperti `hashicorp/go-multierror` atau `uber-go/multierr` demi merangkai dan "menggabungkan menjahit" aneka rupa *error* terpisah ke dalam satu tubuh struktur obyek tunggal raksasa.

Kini, Go telah secara rasmi mematenkan konsep brilian ini ke badan Pustaka Standar intinya melalui fungsi `errors.Join`.
Fungsi sakti ini mengambil serakan puluhan *error* lepasan, merekatkannya erat-erat, dan mengembalikan wujud satu obyek *error* majemuk padat. Yang menakjubkan dari implementasi ini adalah: Fungsi pembongkar identitas ajaib `errors.Is` maupun `errors.As` terbukti **TETAP SANGGUP SECARA REKURSIF TEMBUS MENYELAM** membedah rongga dalam *error* gabungan multiseluler ini dan mengenali satu per satu secara independen jejak asal muasal anak cucu error di dalamnya!

```go
// Menyiapkan sekumpulan Sentinel Errors Statis
// var ErrValidasiUsiaMuda   = errors.New("Validasi Form Gagal: Pendaftar di bawah 18 tahun.")
// var ErrValidasiEmailAneh  = errors.New("Validasi Form Gagal: Format e-mail pendaftar invalid.")

// func prosesVerifikasiFormulirPendaftaranLengkap(usiaInput int, emailInput string) error {
//    var KumpulanLaporanDosa error // Wadah kosong awal pengumpul dosa

//    if usiaInput < 18 {
//        KumpulanLaporanDosa = errors.Join(KumpulanLaporanDosa, ErrValidasiUsiaMuda)
//    }

//    if len(emailInput) < 5 {
//        KumpulanLaporanDosa = errors.Join(KumpulanLaporanDosa, ErrValidasiEmailAneh)
//    }

//    return KumpulanLaporanDosa
// }
```

### 2. Bahaya Miskonsepsi Ekstraksi *Error Tipe Dasar* yang Tersembunyi (The Implicit Trap)

Programmer senior yang merantau migrasi hijrah ke bahasa Go sering tersandung celaka pada operasi instrumen asertif (Type Assertion) error tradisional yang menggunakan sintaks kaku lawas (misal: `errKonkret, apakahBenar := errAcak.(*MyCustomError)`).

Memang cara tradisional itu bekerja bila *error* tersebut masih perawan polos (*naked*). Tapi, begitu *error* kustom tersebut dilempar berjenjang di dalam rantai lapisan arsitektur (dari Database $\rightarrow$ Repository $\rightarrow$ Service Business Logic $\rightarrow$ Controller API) dan diselimuti jaket bungkus pembungkus (via `fmt.Errorf("Gagal ambil profil: %w", errAcak)`), metode tradisional *Type Assertion* di atas **SANGAT DIPASTIKAN AKAN GAGAL MENYEDIHKAN** dan mereturn `False` secara membabi buta. Mengapa? Karena kulit luarnya sudah bukan *MyCustomError* lagi, melainkan sebuah entitas pelindung kusam pembungkus bawaan format (`*fmt.wrapError`).

Karena itulah kehadiran `errors.As` diperjuangkan gigih menjadi standar hukum syariat (*Best Practices*) di bahasa Go. Fungsi ia akan merobek-robek menguliti jaket pembungkus tersebut sampai mencapai inti berliannya.

### 3. Ekosistem Pelacakan Tumpukan Eksekusi Fatal (Stack Traces)

Hingga naskah ini diketik pada versi terbarunya, Go secara pragmatis *Sengaja* menolak keras implementasi pencetakan Tumpukan Jejak Kesalahan Tumpukan Log memori otomatis penuh baris kode (*Stack Traces Generation*) layaknya Java/Node pada paket `errors` bawaannya ini demi menjaga aplikasi biner kompiler Go tetap ringan super ramping irit megabyte RAM eksekusi.

"Jika saya ingin jejak asal baris (*Stack Trace*), apa yang harus saya lakukan?"
Praktik pengembang elit (*Elite Developer Practice*): Jangan mengandalkan paket standar biasa bilamana Anda menyusun Kerangka Proyek (*Boilerplate Framework*) Perusahaan dari nol! Gunakanlah injeksi tambahan pengait pustaka raksasa komunitas mutakhir ternama dunia seperti **`github.com/pkg/errors`** (atau cabangnya yang terkini) di lapisan terendah fungsi Basis Data peladen Anda.

Setiap Anda memanggil sintaks luar semacam `errors.Wrap(err_asli, "Terputus saat query X")`, perpustakaan eksternal magis itu akan secara siluman menangkap potret sidik jari Baris Barisan Eksekusi Kode (File `auth_repository.go` Baris Ke 142) menempel kuat membeku bersamaan di dalam obyek kesalahan *(Error Value)* tersebut, sehingga saat *error* itu akhirnya mendaki ditangkap oleh penjaga gerbang *Controller* Lapis Teratas dan dicetak membentang panjang di layar monitor, ia akan memuntahkan denah navigasi baris tumpukan utuh komprehensif nan mengerikan (`%+v`), menyelamatkan harga diri pengembang perusahaan Anda dari ancaman misteri berburu bug yang membingungkan.

## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
