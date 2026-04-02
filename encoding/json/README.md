# Modul: `encoding/json`

## Ringkasan
Package `encoding/json` menyediakan konversi timbal-balik berkecepatan tinggi yang menjembatani struktur data biner statis di dalam kode Go (berupa *Struct*, *Map*, atau *Slice*) menjadi format tekstual populer JSON (*JavaScript Object Notation* RFC 7159) yang dapat dikirim melintasi batas jaringan internet (sebagai respons Web API RESTful), dan sebaliknya.

## Penjelasan Lengkap (Fungsi & Tujuan)
Di era pengembangan piranti lunak komputasi awan modern, JSON telah menggusur XML sebagai *lingua franca* (bahasa perantara tunggal) untuk pertukaran data antarsistem komputer, baik itu antarmuka Aplikasi Web (React/Vue), Aplikasi Mobile, maupun komunikasi *Backend-to-Backend* (Microservices).

Keunggulan terbesar bahasa Go adalah bagaimana ia mengimplementasikan dukungan format JSON ini: tidak dibutuhkan kelas *Mapper* raksasa yang rumit seperti di bahasa Java. Cukup dengan menempelkan atribut *Struct Tags* kecil bertanda *backtick* (```) di belakang deklarasi *Field Struct* Anda, mesin `encoding/json` bawaan secara ajaib akan menggunakan ilmu *Reflection* di balik layar untuk membaca peta struktur data statis Go Anda, lalu menenun atau membongkarnya menjadi susunan JSON yang luwes. Kesederhanaan dan ketangguhan arsitektural ini merupakan nilai jual mutlak mengapa banyak perusahaan rintisan global menggunakan Go sebagai pilar dasar pembentukan gerbang API (*API Gateway*).

**Tujuan dan Fungsi Utama:**
1.  **Marshalling (Go $\rightarrow$ JSON):** Mengonversi, mengepak, dan mengubah struktur obyek data dalam memori program Go (contoh: *Struct User*) ke wujud aliran byte teks (String JSON) untuk selanjutnya dikirim meretas ruang internet sebagai HTTP Response Body, lewat punggung fungsi kilat `json.Marshal()`.
2.  **Unmarshalling (JSON $\rightarrow$ Go):** Operasi dekonstruksi; menguraikan aliran teks JSON asing yang datang dari klien *Frontend* tak tepercaya, lalu menyerap dan menanamkan nilainya ke dalam cetakan rangka statis *Struct* memori memori Go secara terstruktur menggunakan `json.Unmarshal()`.
3.  **Kendali Pengecualian Format (Omit Empty & Dash):** Menyediakan mekanisme protektif *Struct Tags* untuk menyembunyikan properti berskala rahasia (seperti Kolom Kata Sandi Akun) agar tidak ikut terkonversi terkirim keluar (*data leaking*), atau mengabaikan pencetakan variabel kosong/nihil untuk menghemat bandwith jaringan.
4.  **Dekoder Aliran Jaringan (Streaming Encoder/Decoder):** Jika sistem dihadapkan oleh fail JSON berukuran monster (Gigabytes) atau aliran respon HTTP tebal yang bila ditaruh ke RAM secara konvensional akan mematikan *Server*, fungsi kelas atas `json.Decoder` membaca obyek JSON langsung sekeping demi keping dari lubang pipa `io.Reader`.

**Mengapa menggunakan `encoding/json`?**
Anda mustahil luput dari package sentral ini jika Anda merintis karir sebagai perekayasa Peladen Web *Backend* menggunakan Go. Mengonsumsi hasil panggilan API pembayaran Midtrans, mengirimkan data keranjang katalog produk e-commerce menuju App ReactJS konsumen, atau bahkan menanam rekaman data tak-terstruktur (*Unstructured NoSQL Data*) ke dalam gudang kolom JSON Postgres—keseluruhannya bersandar sepenuhnya pada pemanggilan pustaka sakti nan ringan ini.

---

## Daftar Fungsi Umum dan Cara Penggunaannya Secara Mendalam

### 1. Proses Marshalling: Dari Obyek *Go Struct* ke Teks JSON Murni

Aturan Emas Konversi JSON: Mesin *Marshaller* Go beroperasi di luar bungkus (*package*) asal Anda. Itu berarti, **hanya nama *Field* (kolom) Struct yang huruf awalnya Kapital (Diekspor / *Public*) yang bisa diterjemahkan ke JSON!** Jika Anda menamai kolom dengan huruf kecil (`password string`), paket *json* tak bisa melihat keberadaannya secara mutlak dan ia takkan tercetak.

```go
package main

import (
    "encoding/json"
    "fmt"
)

// Menyiapkan cetakan (Struct)
// Seluruh elemen Field ini HURUF KAPITAL di awal, karenanya ia bersifat "Public" dapat diconvert JSON.
type PegawaiPerusahaan struct {
    NomorInduk    int      `json:"id_karyawan"`   // Tag: Nama kunci JSON saat dicetak akan berubah!
    NamaLengkap   string   `json:"nama"`
    DivisiKerja   string   `json:"divisi"`
    SedangCuti    bool     `json:"status_cuti"`
    DaftarKeahlian []string `json:"skills"`       // Slice Array biasa langsung didukung
}

func main() {
    // 1. Membuat Obyek asali di memory internal Go
    karyawanBaru := PegawaiPerusahaan{
        NomorInduk:    9901,
        NamaLengkap:   "Siti Aisyah",
        DivisiKerja:   "Sistem Informasi",
        SedangCuti:    false,
        DaftarKeahlian: []string{"Golang", "Kubernetes", "PostgreSQL"},
    }

    // 2. Marshalling (Mengepak ke wujud biner Byte JSON)
    // Terdapat versi json.MarshalIndent() jika Anda menginginkan spasi cetakan cantik bagi mata Manusia.
    // Namun json.Marshal() normal lebih gesit di server produksi.

    hasilJSON, errPengepakan := json.MarshalIndent(karyawanBaru, "", "  ") // Indentasi 2 spasi

    if errPengepakan != nil {
        panic("Format memori Obyek tak masuk akal / Gagal Dikonversi!")
    }

    // 3. Tampilkan teks output.
    // Variabel hasilJSON bentuk aslinya adalah []byte, kita ubah ke wujud teks demi Konsol Monitor.
    fmt.Println("Hasil Pengepakan JSON Cantik:\n")
    fmt.Println(string(hasilJSON))
}
```

---

### 2. Atribut Ajaib *Struct Tags* (Rahasia, Pengabaian, dan String Konversi)

Di dunia riil, struktur internal database Anda seringkali memuat sandi akun (*Password*) dan Kunci Autentikasi Rahasia (*Token Key*). Sangat berbahaya bila benda tersebut dikirim terselip tanpa sengaja ke peramban Web via respons API. Gunakan manipulasi tag JSON.

*   **`json:"-"`**: Menginstruksikan modul JSON untuk bersikap buta: *baikan field baris ini sepenuhnya dari alam raya (Sembunyikan mutlak, tak boleh dikirim dan ditangkap!)*.
*   **`json:"nama_key,omitempty"`**: *(Omit Empty / Hilangkan jika Kosong)*. Jika variabel internal tsb dalam keadaan Nilai Dasar Bawaan (*string "" kosong*, *angka 0*, atau *array kosong/nil*), ia tidak akan disertakan pada hasil akhir cetakan teks JSON. Mengirit kuota *bandwidth*.
*   **`json:",string"`**: Terkadang, format API pihak ketiga aneh, ia mengirim nilai Kuantitas Stok barang numerik tapi dengan wujud petik tali di dalam JSON (`{"stok": "15"}`). Tag ini akan memaksakan konversi tipe paksa pada saat pengepakan/pembongkaran antara Teks murni dan Integer/Float asali Struct.

```go
package main

import (
    "encoding/json"
    "fmt"
)

type TransaksiRahasia struct {
    KodeOrder    string  `json:"order_id"`
    NominalUang  float64 `json:"nominal"`

    // TAG RAHASIA: Selamanya tidak akan pernah diikutsertakan ke dalam JSON, super aman!
    SandiKartu   string  `json:"-"`

    // TAG OMITEMPTY: Apabila String Diskon Kosong, atau bernilai nihil bawaan, tak akan dikirim agar irit.
    KuponDiskon  string  `json:"kupon,omitempty"`

    // TAG STRING CONVERT: Secara memori Go, ID database adalah angka Murni 12891,
    // Tapi akan dicetak JSON jadi petik ganda {"id_database": "12891"} demi kompatibilitas klien aneh
    IDDatabase   int     `json:"id_database,string"`
}

func main() {
    // Simulasi data dari Database
    dataAsli := TransaksiRahasia{
        KodeOrder:    "ORD-9990A",
        NominalUang:  500000.50,
        SandiKartu:   "PIN_RAHASIA_KARTU_KREDIT_1234", // Tak terkirim!
        KuponDiskon:  "",                             // Kosong! Akan terbuang dr struktur
        IDDatabase:   9090,
    }

    hasilCetak, _ := json.Marshal(dataAsli)

    fmt.Println("Lihatlah kengerian magic tag: Password dan Kupon Kosong menghilang ditelan bumi!")
    fmt.Println(string(hasilCetak))
}
```

---

### 3. Proses Unmarshalling: Mencerap String JSON Menjadi *Struct Go*

Menerjemahkan dari klien ke bahasa sistem. Teks JSON hanyalah deretan karakter string kosong tanpa makna. Proses dekontruksi `json.Unmarshal()` membutuhkan wujud sebuah *Pointer Variabel Kosong Tujuan* untuk ia injeksikan (isi) dengan nilai data setelah ia membedahnya.

```go
package main

import (
    "encoding/json"
    "fmt"
)

// Struct tujuan penerima tembakan data (Receiver Structure)
type LokasiKordinat struct {
    Kota        string  `json:"kota"`
    LintangUtara float64 `json:"latitude"`
    BujurTimur  float64 `json:"longitude"`
    KodePos     int     `json:"zip_code"`
}

func main() {
    // 1. Data mentah (Skenario nyata berasal dari Body Request HTTP / Data Base)
    pesanMentahKlienJSON := []byte(`{
        "kota": "Jakarta Pusat",
        "latitude": -6.1751,
        "longitude": 106.8272,
        "zip_code": 10110
    }`)

    // 2. Persiapkan wadah tujuan yang MASIH KOSONG / Terisi nilai Default.
    var lokasiHasil LokasiKordinat

    // 3. Proses Tembakan Dekripsi Unmarshall.
    // Aturan Krusial: HARUS memberikan wujud Alamat Penunjuk Memory Pointers (tanda &) pada wadah lokasiHasil!
    // Jika tidak diberikan pointer ampersand, paket json tidak bisa mengakses dan merekayasa wadah tsb.
    errDekripsi := json.Unmarshal(pesanMentahKlienJSON, &lokasiHasil)

    if errDekripsi != nil {
        panic(fmt.Sprintf("Wujud struktur kiriman JSON ditolak mentah-mentah: %v", errDekripsi))
    }

    // 4. Sukses! Manipulasi variabel obyek Go secara natural
    fmt.Printf("Data masuk tervalidasi menyeberang Jaringan: Kota %s (Kode Wilayah %d)\n", lokasiHasil.Kota, lokasiHasil.KodePos)
}
```

---

### 4. Mengelabui Data JSON Tak Terstruktur dengan Peta `map[string]interface{}` (Any)

Anda mengonsumsi API cuaca publik *asing* yang struktur desain format JSON responsenya tak jelas, tak konsisten, dan teramat panjang (ratusan properti), dan Anda hanya butuh 1 biji informasi saja tanpa mau bersusah payah menuliskan deklarasi raksasa *Struct Go* yang ribet? Go menyediakan fleksibilitas penyelesaian secara instan tanpa mendefinisikan *Struct* di awal menggunakan koleksi dinamis `Map`.

*(Catatan Modern: Di versi rilis Go terkini (1.18+), Anda bisa meringkas syntax tua `interface{}` menjadi kata kunci `any`)*.

```go
package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    // Teks tak jelas dari respon API acak, ada isi String Teks, ada Array, ada Boolean. Campur Aduk!
    dataAsingAneh := []byte(`{
        "status": "sukses",
        "angka_ajaib": 42.5,
        "is_active": true,
        "items": ["apel", "jeruk", "mangga"]
    }`)

    // Buat Peta penampung Buta (Peta Dinamis dengan Kunci String, berisikan Nilai berwujud Bebas Apa Saja 'any')
    var emberPenampung map[string]any

    // Tembak Data tersebut masuk membobol Ember Peta Buta!
    json.Unmarshal(dataAsingAneh, &emberPenampung)

    // Pengambilan data. Hati-hati, karena nilainya bertipe buta `any`,
    // Anda harus melakukan pemaksaan asersi tipe (Type Assertion) secara sadar diri agar bisa mengolahnya.

    // Mengambil Teks String Biasa
    statusPesan := emberPenampung["status"].(string)

    // Peraturan Keras Go: Semua angka hasil Unmarshall JSON Peta Buta SELALU ditelan sebagai format Float64!
    // Meski awalnya di JSON telihat seperti integer murni tanpa koma.
    angkaNilai := emberPenampung["angka_ajaib"].(float64)

    // Mengambil Elemen Array bersarang (Lebih pusing asersinya!)
    arrayBuahMentah := emberPenampung["items"].([]any)
    buahPertama := arrayBuahMentah[0].(string)

    fmt.Printf("Status Ekstraksi: %s\n", statusPesan)
    fmt.Printf("Angka Diekstrak: %.2f\n", angkaNilai)
    fmt.Printf("Isi keranjang index awal: %s\n", buahPertama)
}
```

---

## Bagian Lanjutan: Decoding Aliran Masif (Streaming), Penanganan JSON Fleksibel, dan Pola Arsitektur API Ganas

Sementara fungsi `json.Marshal` dan `json.Unmarshal` biasa menangani >90% beban kerja konversi JSON pada arsitektur perangkat lunak berbasis REST API, insinyur *backend* tingkat mahir menyadari batasan fisik yang mematikan dari kedua fungsi tersebut. Fungsi-fungsi dasar ini bersifat manipulasi *In-Memory* penuh; mereka membutuhkan alokasi memori secara bongkahan utuh (*Bulk Memory Allocation*) agar bisa bekerja dengan sempurna. Di ranah pemrosesan tingkat dewa—seperti menelan balasan fail JSON raksasa berbobot GigaByte (*Streaming Big Data*), struktur API kacau yang isinya bisa berubah-ubah, dan kecepatan latensi hitungan mikrodetik—Anda harus mengerahkan arsenal fungsi kelas berat milik paket `encoding/json`.

### 1. Merobohkan Limit RAM dengan Dekoder Berbasis Aliran (*Streaming Encoder & Decoder*)

Skenario Kiamat Memori: Server Go Anda adalah sistem perantara (API Gateway). Ia harus mengambil arsip transaksi log JSON pelanggan dari sebuah server penagihan (*Billing Server*) yang ukurannya membesar secara dramatis mencapai 5 Gigabytes murni JSON, lalu mengekstraknya ke database Anda.
Jika Anda mendownload respon dari server penagihan, lalu memanggil `io.ReadAll` (membakar 5GB RAM), kemudian memanggil `json.Unmarshal` (Membakar dan menyalin lagi 5GB RAM ke Struct Go), dalam sepersekian detik server tumpuan utama perusahaan Anda meledak (*OOM Kill Panic*) karena kehabisan RAM.

**Senjata Solusi Mutlak: `json.Decoder`**
Berbeda dengan Unmarshal, Dekoder diciptakan untuk *"Menempel langsung pada Pipa Mulut"* aliran TCP HTTP (`io.Reader`). Ia menyedot bit byte jaringan, merakit objek JSON pertama yang ia temukan, menelannya, dan membuang memorinya tanpa harus menunggu miliaran obyek sisa JSON di bawahnya terunduh! Ini memungkinkan Anda memakan file JSON 5 Terabytes sekalipun dengan memori RAM yang diam stabil di ukuran beberapa Megabyte ringan. Sangat ajaib dan hemat daya.

```go
// Simulasi:
// aliranGilaPipaInternet := strings.NewReader(`[{"id_trx": 1}, {"id_trx": 2}]`)
// mesinSedotDekoder := json.NewDecoder(aliranGilaPipaInternet)

// Kita baca dulu kurung siku Pembuka Array '['
// tokenAwal, errT := mesinSedotDekoder.Token()

// SELAMA mesin masih mendeteksi ada obyek {..} di depan mata, kita sedot satu demi satu!
// for mesinSedotDekoder.More() {
//    var wadahTunggal TransaksiKecil
//    errSerap := mesinSedotDekoder.Decode(&wadahTunggal)
//    fmt.Println(wadahTunggal)
// }
```

### 2. Kustomisasi Transformasi (JSON Marshaler / Unmarshaler Interface)

Ada kalanya format representasi memori internal di kode Go dan wujud luarnya (format JSON) harus sangat berbeda (terkena manipulasi format spesifik).
Contoh klasik: Di dalam Go, umur pendaftar direkam menggunakan format `time.Time` murni yang kaku, tetapi aplikasi seluler (*Frontend*) memaksa API Anda menelurkan balasan JSON di mana nilai kelahirannya menjadi format UNIX detik integer kuno.

Anda tidak boleh memaksa mengotori Struct inti Go yang elegan dengan merubahnya menjadi tipe `int`! Sebagai kompromi cerdas, Go menyediakan pintu belakang (*Interface*) agar tipe Anda bisa mencegat (Intercept) dan memalsukan wujud format datanya detik-detik sebelum modul `encoding/json` mengepaknya.

Anda cukup menempelkan metode pertukaran siluman **`MarshalJSON()`** dan **`UnmarshalJSON()`** di tipe tersebut!

```go
// type SandiWaktuKustom time.Time

// func (s *SandiWaktuKustom) MarshalJSON() ([]byte, error) {
//     waktuAsli := time.Time(*s)
//     angkaEpochInteger := waktuAsli.Unix()
//     return []byte(fmt.Sprintf("%d", angkaEpochInteger)), nil
// }
```

### 3. Trik Tipe Data Lentur `json.RawMessage`

Bencana lain saat mengkonsumsi Webhook API Eksternal jahat (Misal API Bank/Media Sosial): Di suatu saat respons payload obyek `data` mereka adalah JSON Obyek Profil yang rumit `{"id":1, "name":"A"}`, namun bila bank tersebut gagal memproses tagihan, mendadak rute payload obyek `data` itu secara menjijikkan berganti wujud menjadi pesan tulisan teks biasa bertipe *String* Murni: `"Sedang Gangguan"`.

Jika Anda membongkarnya (*Unmarshal*) ke struct Obyek, Go seketika akan menelurkan Error panik keras (*Type Mismatch*), meruntuhkan keandalan servis Peladen aplikasi yang Anda bangun bersusah payah!

Solusi magis: Ikat field siluman tersebut dengan tameng pengaman tipe **`json.RawMessage`**. Tipe ini akan menolak membongkar isi field tersebut, membiarkannya diam membeku wujud aslinya sebagai rentetan *byte JSON Mentah*. Barulah di tahap berikutnya Anda bisa menimbang-nimbang sendiri apakah akan mengubahnya sebagai obyek, atau mengekstraknya sebagai Teks murni berdasarkan kode respons. Trik luar biasa elegan mencegah *Crash* API mematikan!

## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
