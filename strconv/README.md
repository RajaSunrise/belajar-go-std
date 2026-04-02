# Modul: `strconv`

## Ringkasan
Package `strconv` (*String Conversion*) memuat sekumpulan perkakas fungsi penterjemah untuk memigrasikan tipe data primitif Golang biner murni secara instan aman ke wujud manifestasi Teks *String* perwujudan aksara yang bisa dibaca representasi visual manusia ASCII, pun menerjemahkannya kembali (balik-putar) secara presisi ke nilai dasar biner perhitungan internal aritmatika.

## Penjelasan Lengkap (Fungsi & Tujuan)
Di belantara arsitektur Web dan *Software*, Anda berhadapan dengan sebuah realita keras yang canggung: Sistem transmisi *Frontend HTML Forms*, URL Parameter (contoh `?page=5`), Variabel Peranti OS Environment (`PORT=8080`), hingga respon HTTP *Body JSON Text API* kesemuanya saling bertukar oper transmisi ke sistem Peladen Golang murni dalam Wujud 100% balutan baju Tipe "Teks / String".

Anda tak mungkin melakukan kalkulasi Matematika (`page + 1`) apabila variabel `page` tersebut bertipe *String* `"5"`. Jika di dalam bahasa seperti JavaScript / PHP bahasa kompilator ajaib memaafkan kebebasan tersebut dan dengan sok pintar diam-diam memanipulasi di belakang punggung ("Oh maksudnya `5`+1 itu `6` kali ya..."), kompilator Go tidak sebodoh itu. Golang memiliki aturan keras kepatuhan kasta (*Strongly and Statically Typed Language*). Angka murni biner `5` berwujud (Integer tipe) berbeda alam kasta dimensi beda spesies memori absolut mutlak ketimbang balutan karakter rentetan `byte("5")` (String Tipe).  `strconv` hadir mendelegasikan penterjemahan legal, mualaf yang presisi dan menginformasikan Error gagal konversi tersembunyi dengan tegas.

**Tujuan dan Fungsi Utama:**
1.  **Parse Intepretasi String ke Biner (String $\rightarrow$ Int/Float/Bool):** Memaksa paksa pembacaan rentetan untaian aksara berwujud *"true"*, atau *"3.14159"*, dan juga *"150"* dan mengubahnya menjelma berevolusi sebagai Integer, Float 64-bit desimal maupun nilai logikal pembenaran `Boolean`. (Terpusat di keluarga Fungsi berawalan `Parse...` dan `Atoi`).
2.  **Format Formulasi Tipe Primitif ke Teks Aksara Visibilitas (Int/Float/Bool $\rightarrow$ String):** Mensintesis kembali balik logika hitungan angka pecahan Desimal Internal Biner ke Teks cetak aksara untuk dilaporkan kembali ke HTML/URL, (Dikeluarkan kelompok Fungsi berawalan `Format...` dan `Itoa`).
3.  **Trik Kutipan & Modifikasi Basis Angka Radiks:** Menyulap angka Biner beralih wujud Basis Okta Heksadesimal maupun merumuskan konversi perakitan String pengutip (`Quote/Append`).

**Mengapa menggunakan `strconv`?**
Jika sistem Backend web Server HTTP Anda memiliki fungsionalitas memuat rute berwujud `/items/98`, rute tersebut di mata Go *Controller Router Mux* diterima sepenuhnya bulat sebagai Teks String mutlak `"98"`. Mustahil API Anda bisa mengirimkan ke mesin *query SQL Driver Database* dengan aman sebelum memigrasikan wujud kebenaran sejatinya sebagai Integer (`98`). Inilah package penterjemah yang nyaris muncul 100% membungkus disetiap *Controller Input API Validation Gateway Server*.

---

## Daftar Fungsi Umum dan Cara Penggunaannya Secara Mendalam

### 1. Rantai Evolusi Sederhana Angka Bulat Klasik (Atoi & Itoa)

Operasi konversi string ke angka bulat mendominasi sekitar 80% penggunaan murni harian pengembang pemrogram. Karena itu sang pencipta bahasa di Google memberikan 2 rute jalan tol alias *shortcut* yang super praktis nan ringkas:

*   **`strconv.Atoi(s)`**: Singkatan dari historis bahasa *C (ASCII to Integer)*. Ini mengubah String `"88"` menjadi hitungan *integer* murni `88`. Jangan lupa bahwa ia mereturn Nilai Ganda (Angka, dan Error). Ia memuntahkan *Error* jika Teks itu terbukti "ngawur" tak berbau Angka (contoh: `"88apel"`).
*   **`strconv.Itoa(i)`**: Singkatan kebalikannya *(Integer to ASCII)*. Ia hanya mereturn Murni *String* tunggal! Angka 88 diubah menjadi teks `"88"`. (Tak ada *return Error* pada proses ini karena secara fisika Komputer, merubah angka Integer biner ke Teks *String* sudah dijamin takkan pernah Gagal!)

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // ---- KASUS 1: Atoi (Teks String Input ke Angka Perhitungan) ----
    inputTeksUmurFormulirWeb := "24"

    // Jangan sekali-kali mengabaikan Penanganan _ Error nya
    umurAsliHitungan, errMigrasi := strconv.Atoi(inputTeksUmurFormulirWeb)

    if errMigrasi != nil {
        fmt.Println("Peringatan! Hacker klien mengirim karakter huruf ke kotak isian Umur angka!")
    } else {
        tahunLahir := 2024 - umurAsliHitungan // Kita kini bisa leluasa pakai operator Hitungan '-' !
        fmt.Printf("Pelanggan sah meminum Alkohol! Terdeteksi tahun kelahirannya: %d\n", tahunLahir)
    }

    // ---- KASUS 2: Itoa (Menyusun Angka kembali menjadi String ke Cetakan Pesan) ----
    jumlahStokGudangBase := 450

    // Anggap aja Kita ingin merakit SQL Log Teks Dinamis, tidak pakai fmt.Sprintf, kita memigrasi eksplisit!
    stringPesanDigabung := "Terdapat total sisa " + strconv.Itoa(jumlahStokGudangBase) + " Dus di dalam kontainer."

    fmt.Println(stringPesanDigabung)
}
```

---

### 2. Penterjemahan Parse Kaliber Berat Tipe Presisi Presisi Kompleks (ParseFloat, ParseBool)

Jika dunia perhitungan matematika presisi tinggi melanda (Misal perhitungan koordinat peta Lintang Bumi `Latitude`, atau perhitungan Bunga KPR Uang Desimal pecahan). Jalan Tol `Atoi` primitif di atas lumpuh tak bersenjata. Kita harus memanggil keluarga inti `Parse`.

*   **`strconv.ParseFloat(s, bitSize)`**: Menelan paksa Teks menjadi Bilangan Pecahan Biner. *bitSize* didefinisikan 64 (demi `float64`) agar menahan presisi angka maksimal.
*   **`strconv.ParseBool(s)`**: Ini fitur jenius yang sangat lunak hatinya. Go menerima aneka rupa bentuk tulisan Teks `1`, `t`, `T`, `TRUE`, `true`, `True` sebagai wujud keabsahan logika Benar `True`. Di lain sisi rentetan `0`, `f`, `F`, `FALSE`, `false`, `False` diinterpretasikan Salah/`False`.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // A: Kasus Memparsing Pecahan Titik Mata Uang (Float)
    hargaInputStringKatalog := "98550.75" // Menggunakan format titik Inggris standard

    hargaFloatHitungan, errFloat := strconv.ParseFloat(hargaInputStringKatalog, 64)
    if errFloat != nil {
        panic(errFloat)
    }

    pajakPPn := hargaFloatHitungan * 0.11
    fmt.Printf("Operasi Hitung PPN Beres! Tambahan Biaya: Rp %.2f\n\n", pajakPPn)


    // B: Kasus Logika Toggle Switch URL parameter (Parse Bool)
    // Coba saja Anda kirim "T", "true", "1" dari sisi Web Frontend, pasti akan diterima valid True.
    paramMintaDiskon := "1"

    apakahBerhakDapatPromo, errPromo := strconv.ParseBool(paramMintaDiskon)

    if errPromo != nil {
        fmt.Println("Error format salah! Anda harus kirim 'true' atau 'false' murni.")
    } else {
        if apakahBerhakDapatPromo {
            fmt.Println("Mengeksekusi Algoritma Pemotongan Diskon Khusus!")
        } else {
            fmt.Println("Mengeksekusi Transaksi Normal reguler.")
        }
    }
}
```

---

### 3. Ekstremitas Merajut (Format) Konfigurasi Radiks Basis Angka Biner / Heksa

Di sudut kasus langka pengerjaan alat Kriptografi Hashing, Logika Pewarnaan *Web Pixel CSS* (#AABBCC), Anda harus membungkus nilai Integer biasa dikempeskan dimanipulasi perwujudan kulitnya menjadi hitungan Basis Radiks bukan 10 (*Hexadecimal / Octal*). Keluarga `.FormatInt` lah senjatanya.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Angka Integer Desimal Murni (Basis Kehidupan Manusia radiks 10)
    angkaSpesialWarnaInternal := int64(255) // Nilai Mentok Maksimum Putih

    // Mari merubahnya menjadi perwujudan teks representasi Radiks Kode CSS Heksadesimal (Basis 16)!
    // Parameter kedua menandakan (base / radiksnya)
    wujudTeksHexadecimal := strconv.FormatInt(angkaSpesialWarnaInternal, 16)

    // Merubah wujud representasi mesin primitif Mesin (Basis 2 / Binery / 0101)
    wujudTeksBineryPita := strconv.FormatInt(angkaSpesialWarnaInternal, 2)

    fmt.Println("Angka Asal Fisik Manusia :", angkaSpesialWarnaInternal)
    fmt.Printf("Wujud Cetakan Web CSS Hex : #%s (huruf ff)\n", wujudTeksHexadecimal)
    fmt.Printf("Wujud Transmisi Binery 01 : %s\n", wujudTeksBineryPita)
}
```

---

## Bagian Lanjutan: Kecepatan Konversi Ekstrem (*Zero Allocation*), Optimasi Penggabungan String, dan Kesenjangan Komputasi Basis (*Parsing Bounds Limit*)

Di belakang keluguan fungsi manipulasi string dasar (`Itoa` dan `Atoi`) dari paket `strconv` terdapat pusaran pemborosan memori komputasi asali bila Anda sedang menyusun respons peladen bertrafik maut (contoh: 50.000 Konversi Integer Identitas Transaksi per Detik pada *Microservice HTTP Server*). Setiap satu kali Anda menyuruh golang mengubah angka 500 menjadi kalimat `"500"`, sistem pangkalan `Garbage Collector` Go menghela napas panjang sebab ia ditekan untuk membeli (Alokasi Heap) sepotong memori array baru yang sebentar lagi bakal dibuang hancur!

Menangani performa ekstrem pada format pencetakan (Serialization JSON / Protobuf Manual) tak bisa dilakukan tanpa bantuan keluarga punggawa *Append* (*Sisipkan Langsung ke Memori Biner Laci Berjalan*).

### 1. Penghancur Batas Memori: Rumpun Senjata "Append" (Zero-Allocation Conversions)

Insinyur Go papan atas menolak keras pemanggilan kombinasi naif `strconv.Itoa()` atau dewa pemboros memori `fmt.Sprintf("%d")` kala mereka menggabungkan angka ke dalam rakitan *String Raksasa* beruntun.

Fungsi sakti **`strconv.AppendInt(sliceTujuan, angkaMasukan, radix_basis)`** mendeteksi alamat dasar Array Penampung Byte kita (`[]byte`), mengunyah angka internal Integer murni, menterjemahkannya ke huruf teks angka, **lalu MENEMPELKANNYA (MENYAMBUNGNYA)** ke bagian buritan belakang *Slice* penampung tersebut secara mutlak IN-PLACE (Merubah array di letak posisi yang sama tanpa mendelegasikan alokasi pembuatan wadah Memori yang baru dari Pabrik RAM OS).

```go
// PRAKTIK INDUSTRI BERPERFORMA DEWA (Level Framework HTTP seperti Gin/Fiber)

// Kita pesan Kavling tanah di awal (Kapasitas Buffer awal pre-allocated 100 byte).
// Tak ada lagi permohonan memori tambahan ke OS saat komputasi berlangsung! Hemat waktu mutlak!
// wadahAjaibBebasAlokasi := make([]byte, 0, 100)

// wadahAjaibBebasAlokasi = append(wadahAjaibBebasAlokasi, []byte("Laporan Transaksi ID: ")...)

// INILAH KUNCI RAHASIA: APPEND-INT!
// Fungsi ini tidak memuntahkan String "99008821" Baru!
// Ia malah memahat huruf demi huruf angka itu LANGSUNG nempel bersambung ke ekor wadahAjaibBebasAlokasi.
// wadahAjaibBebasAlokasi = strconv.AppendInt(wadahAjaibBebasAlokasi, 99008821, 10)

// Hasil jadi (Satu kali tembak final cetak layar)
// fmt.Println(string(wadahAjaibBebasAlokasi))
```

Rumpun keluarga fungsi `Append` ini mencakup saudaranya yang komplit: `AppendBool`, `AppendFloat`, `AppendQuote`. Apabila Anda menyelidiki jerohan di balik mesin *JSON Encoder* bawaan bahasa Go yang super cepat, Anda mendapati mereka memanfaatkan kumpulan perintah *Append* ini jutaan kali untuk merakit Teks Bodi respons JSON HTTP balasan.

### 2. Bahaya Maut *Integer Overflow* (Sistem Pembatas Parsial Ambang Komputasi Biner)

Kelemahan paling kronis programmer junior adalah menganggap operasi pembacaan `strconv.Atoi()` (atau keluarga *ParseInt*) tidak akan gagal apabila kalimat Teks yang dimasukkan klien Browser memang semuanya murni Angka Numerik.

Bagaimana jika Klien iseng mengisi kotak formulir URL Web umur dengan Angka Murni: `"999999999999999999999999999999999999999"`?
Apakah Penterjemah Angka Go mereturn sukses?
**TIDAK!** Mesin Go akan panik memberikan Error (ErrRange). Mengapa? Karena batasan dimensi Memori Kertas 64-Bit *Integer* di mesin komputer manapun tidak sanggup menahan ledakan digit sebesar angka 9 raksasa tadi (*Integer Overflow*).

Mencegah meledaknya program perbankan membutuhkan kompromi perlindungan `ParseInt` yang ketat.

```go
// Sintaks Fundamental: strconv.ParseInt(InputTeks, RadixBasis, BitSize_BatasMaksimalTipeData)
// Nilai BitSize = 32 artinya kita membatasi ia setara dengan tipe 'int32' maksimal angkanya sekitar 2 Milyar lebih.
// Apabila ia melebihi Batas Max Int32 tersebut, ParseInt akan LANGSUNG MENOLAK DAN ERROR, melindungi Data Server!

// teksTransferHacker := "99999999999999" // 99 Trilyun!

// angkaAmanHitungan, errKebesaran := strconv.ParseInt(teksTransferHacker, 10, 32)

// if errKebesaran != nil {
//    if numErr, benar := errKebesaran.(*strconv.NumError); benar && numErr.Err == strconv.ErrRange {
//        fmt.Println("[SECURITY BLOCK] Transaksi Ditolak: Nominal Melanggar Batas 32-Bit!")
//    }
// }
```

### 3. Eksekusi Pengutipan Teks Ganda Bebas Cacat Kutipan (Quote)

Fungsi sakti perakitan Log Administrator di server adalah **`strconv.Quote(s)`**. Saat pengunjung Web liar tak dikenal menyusup mencoba-coba alamat rute tersembunyi seperti `http://server/app/%22rahasia_bocor%0A`, Anda berniat mencetak peringatan log penolakan ke terminal layar Command Line.

Masalah kiamatnya: Jika string input tamu itu mengandung karakter ajaib loncat baris Enter (`\n`) maupun Tanda Kutip jebakan (`"`), mencetaknya secara langsung `fmt.Printf("Tamu masuk ke: %s", ruteLiar)` akan merusak susunan format teks barisan Log Terminal Anda yang sudah rapi, bahkan menipu mata (*Terminal Injection / ANSI Escape Sequence Attack*).

Gunakan tameng penjarap pengutip `Quote`:
```go
jejakRuteKlienLiar := "admin/dashboard"
 MENGHAPUS_DB!"

// PERTAHANAN GO: MENGURUNG TEKS KE DALAM PENJARA KUTIPAN AMAN!
teksPenjaraAman := strconv.Quote(jejakRuteKlienLiar)

fmt.Println("Laporan Penjaga: Seseorang mencoba masuk menuju rute:", teksPenjaraAman)
// Output Hasil Terjamin Lurus 1 Baris Cetakan Mutlak:
// Laporan Penjaga: Seseorang mencoba masuk menuju rute: "admin/dashboard"
 MENGHAPUS_DB!"
```
Pemaparan fungsi paket konversi tingkat atas dari `strconv` ini merupakan rahasia dibalik kokohnya server backend mikro layanan berbasis bahasa Go. Modul ini secara konstan menyuplai tenaga penengah antara logika biner mesin yang rigid (kaku) di belakang layar bersanding harmonis dengan lautan pertukaran teks visual di garis perbatasan luar internet HTTP.

## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```

---

## Studi Kasus Dunia Nyata: Dekonstruksi Permintaan Klien di Gerbang Pintu API (API Router Validation Guard)

Sembilan puluh persen dari waktu kerja fungsional di dalam rutinitas peladen API (*Backend Server*) dihabiskan untuk satu proses primitif: "Menerima lemparan teks mentah JSON/URL dari perangkat seluler Android (*Frontend Klien*), mem-validasinya menjadi bentuk Angka Tipe Kuat (*Strongly-Typed Numbers*), kemudian baru diizinkan masuk ke ruang rapat operasi pangkalan data SQL".

Bila gerbang Pintu API kita lalai mengkonversi dan langsung menyuapkan teks mentah tersebut, malapetaka injeksi data sampah (Type Mismatch) niscaya akan menghancurkan relasi integritas antar-Tabel Basis Data Anda. Pengamanan terpadu menggunakan paket komprehensif `strconv` merupakan pertahanan perimeter absolut API kita.

### Penyaringan URL Parameter Ekstrem Berlapis Validasi (URL Query Extractor Guard)

Mari kita merakit contoh arsitektur *Controller* penangkap Filter Katalog Penjualan yang menerima 3 paramater liar URL dari internet bebas:
Alamat Klien mengetik di peramban: `https://api.toko.com/produk?limit=50&harga_minimal=150000.50&hanya_aktif=true`

Fungsi Go kita harus merombak ke-tiga tipe string mentah tersebut dengan tiga jajaran alat pengubah wujud sakti paket ini.

```go
package main

import (
    "fmt"
    "net/url"
    "strconv"
)

// Struktur Suci Lapis Dalam yang WAJIB bersih dari tipe String!
type KriteriaSaringPencarian struct {
    BatasLimit      int
    HargaMinimumIDR float64
    HarusBarangAktif bool
}

func main() {
    // 1. Simulasikan Tangkapan Data Mentah dari Internet (r.URL.Query())
    teksKirimanPeretasInternet := url.Values{
        "limit":         {"50"},
        "harga_minimal": {"150000.50"},
        "hanya_aktif":   {"true"},

        // Simulasikan parameter jahat yang dicoba dimasukkan!
        "id_produk":     {"123_DROP_TABLE_PRODUK"},
    }

    // 2. Persiapkan Wadah Tampungan Bersih
    var kriteriaBersihSistem KriteriaSaringPencarian

    // 3. OPERASI PENYARINGAN KETAT STRCONV

    // A. Penyaringan Angka Bulat Limit
    if nilaiLimitMentah := teksKirimanPeretasInternet.Get("limit"); nilaiLimitMentah != "" {

        limitBersih, errLimit := strconv.Atoi(nilaiLimitMentah)
        if errLimit != nil || limitBersih < 1 || limitBersih > 100 {
            // Mencegat peretas yang iseng memasukkan Limit= Sejuta Milyar
            fmt.Println("GAGAL: Parameter Limit Kuantitas Barang tak rasional atau Format Huruf!")
            return
        }
        kriteriaBersihSistem.BatasLimit = limitBersih
    }

    // B. Penyaringan Nominal Pecahan Sensitif 64-Bit (Uang Rupiah)
    if nominalMentah := teksKirimanPeretasInternet.Get("harga_minimal"); nominalMentah != "" {

        nominalBersih, errHarga := strconv.ParseFloat(nominalMentah, 64)
        if errHarga != nil || nominalBersih < 0 {
            fmt.Println("GAGAL: Harga batas bawah menolak injeksi huruf abjad!")
            return
        }
        kriteriaBersihSistem.HargaMinimumIDR = nominalBersih
    }

    // C. Penyaringan Sinyal Boolean Logika
    if statusMentah := teksKirimanPeretasInternet.Get("hanya_aktif"); statusMentah != "" {

        // Trik sakti Go: Merombak kata-kata "t", "true", "1", "T" sekaligus instan jadi kebenaran mutlak boolean
        statusBersih, errBool := strconv.ParseBool(statusMentah)
        if errBool != nil {
            fmt.Println("GAGAL: Kategori centang status barang cacat format logika!")
            return
        }
        kriteriaBersihSistem.HarusBarangAktif = statusBersih
    }

    // D. Pengujian Serangan Peretas (Injeksi SQL Tipe Angka Bulat)
    idUjicobaJahat := teksKirimanPeretasInternet.Get("id_produk")
    _, errDeteksiHacker := strconv.Atoi(idUjicobaJahat)

    if errDeteksiHacker != nil {
        fmt.Println("\n>> ALARM KEAMANAN SISTEM MENYALA <<")
        fmt.Printf("Mendeteksi Upaya Serangan Cacat Format di Parameter ID Produk: %s\n", idUjicobaJahat)
        fmt.Println("Permintaan Otomatis Dibuang ke Tempat Sampah (HTTP 400 Bad Request)!!\n")
    }

    // 4. Lulus Ujian Perbatasan! Lanjutkan Sistem Eksekusi!
    fmt.Println("---- DATA BEBAS KUMAN BERSIH ----")
    fmt.Printf("Obyek Memori Go Murni: %+v\n", kriteriaBersihSistem)
    fmt.Println("Sistem aman membedah operasi Kueri Lanjutan SQL di Lapis Bisnis Repositori.")
}
```

Pengawalan pintu perbatasan menggunakan perpaduan pemanggilan `Atoi`, `ParseFloat` dan peredam anomali asertif `ParseBool` dari paket mutlak `strconv` ini menyegel kelemahan mutasi (*Type Juggling Flaws*) yang sering mengandaskan kejayaan Peladen *Backend* dari platform *Scripting* dinamis tetangga pudar.
