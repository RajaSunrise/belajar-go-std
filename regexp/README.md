# Modul: `regexp`

## Ringkasan
Package `regexp` (*Regular Expression*) mengimplementasikan pencarian, validasi identitas teks, ekstraksi potongan sub-data, serta manipulasi *string* tingkat ekstrem (*Advanced Pattern Matching*) yang digerakkan oleh mesin penafsir sintaks bahasa "Reguler Ekspresi" varian RE2.

## Penjelasan Lengkap (Fungsi & Tujuan)
Di dalam pengolahan data teks konvensional, package dasar `strings` milik Go sudah sangat piawai untuk mencari apakah "Teks A mengandung Kata B". Namun, dunia nyata jarang sesederhana itu. Seringkali Anda tidak dihadapkan pada pencarian *kata mati*, melainkan Anda disodorkan sebuah pola abstrak. Contoh masalah: "Temukan semua nomor telepon lokal yang bersembunyi di dalam paragraf panjang ini, dengan kriteria bahwa nomor tersebut harus selalu diawali dengan 08, diikuti oleh antara 8 hingga 11 deret angka, dan bisa jadi diselipi karakter garis pisah (-)". Ini adalah masalah yang mustahil diselesaikan dengan efisien menggunakan pencarian `strings` konvensional.

Bahasa pemrograman Go mengusung pustaka standar *Regexp* yang diotaki oleh algoritma *execution engine* bernama RE2. Secara historis, banyak bahasa pemrograman lain (seperti Perl, Ruby, atau Java awal) memiliki implementasi *Regexp* yang cacat, di mana jika seorang peretas merancang "teks jebakan" tertentu (dikenal sebagai serangan ReDoS - *Regexp Denial of Service*), CPU server akan terjebak dalam putaran pencarian yang tak berujung miliaran kali lipat lamanya dan menyebabkan *server crash*. Hebatnya, RE2 di Go dirancang spesifik untuk **selalu berlari dalam kecepatan waktu linear ($O(n)$)**, mencegah serangan eksekusi abadi tersebut secara teknis mutlak di tingkat arsitektur. Anda boleh memproses jutaan pola tanpa khawatir Server Anda meledak!

**Tujuan dan Fungsi Utama:**
1.  **Validasi Cerdas Berbasis Pola:** Mendeteksi kesahihan format isian identitas (Format Alamat Email, *Username* tanpa spasi/karakter aneh, KTP, Nomor Kartu Kredit, dsb) dalam hitungan milidetik secara asertif melalui komparasi `.MatchString()`.
2.  **Pemecah (*Splitter*) Lanjut:** Memecah sebuah string kalimat bukan menggunakan 1 karakter koma konstan biasa, melainkan memecahnya berlandaskan pola abstrak (misal: "Pisahkan kalimat ini di setiap tempat di mana Anda menemukan 3 huruf vokal berjejeran").
3.  **Ekstraksi Data Tersembunyi (Scraping):** Menyedot dan mengumpulkan segala serpihan kata dari sebuah dokumen raksasa yang *cocok* (match) dengan kriteria pola ke dalam struktur Array *Slice* `.FindAllString()`.
4.  **Sensor dan Pergantian Ganas (Anonymization):** Melewati satu blok naskah dokumen panjang, menemukan pola alamat email / IP *Server*, lalu menggantinya serentak seketika menjadi karakter bintang (`***`) demi meloloskan Log *Privacy Policy* GDPR di ranah produksi `.ReplaceAllString()`.

**Mengapa menggunakan `regexp`?**
Jika sistem API Peladen Web yang Anda gagas harus mensterilkan format data sensitif kiriman klien (*Sanitization*), mengurai data acak (Web *Scraping* Data Mentah), hingga merancang *Bot Parser Chat* Discord yang cerdas di Go, package ini ibarat tongkat sihir serba bisa yang memotong ribuan baris perintah validasi *if-else* IF panjang Anda menjadi sekadar sejentik rumusan baris *Pattern Regexp* elegan.

---

## Daftar Fungsi Umum dan Cara Penggunaannya Secara Mendalam

### 1. Inisialisasi Pola Kompilasi: MustCompile vs Compile

Satu aturan paling diagungkan di Go: *Kompilasi String Pola Regex adalah pekerjaan berat memeras CPU!* Karena itu, Anda **HARUS dan WAJIB** melakukan deklarasi fungsi "Kompilasi Pola Regex" tersebut hanya sekali saja (contohnya menempatkannya di wilayah variabel *Global*, bukan meletakkannya di dalam *loop* iterasi 1000 kali perulangan fungsi).

*   **`regexp.Compile()`**: Mengembalikan `(*Regexp, error)`. Digunakan jika pola Regex Anda berasal dari *input* User dinamik yang belum tentu valid ketikannya.
*   **`regexp.MustCompile()`**: Mengembalikan `*Regexp`. (Ini yang paling disarankan dipakai). Ia akan langsung mematikan aplikasi seketika (*Panic*) saat pertama di-*run* jika ternyata rumusan pola yang Anda tulis (secara *hardcoded*) mengalami salah ketik *typo* sintaks. (Menghindari *bug* tersembunyi).

```go
package main

import (
    "fmt"
    "regexp"
)

// PRAKTIK TERBAIK (BEST PRACTICE):
// Inisialisasikan variabel Regexp yang ter-compile HANYA SATU KALI di bagian Global Package!
// Pola: ^ (diawali), [a-zA-Z0-9_] (Hanya Boleh Huruf/Angka/Underscore), {3,16} (Panjang 3-16), $ (diakhiri)
var validUsernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,16}$`)

func main() {
    daftarTesPendaftar := []string{
        "pemain_hebat",    // VALID: Masuk kriteria
        "12",              // INVALID: Terlalu pendek (kurang dari 3 huruf)
        "budi@santoso",    // INVALID: Ada karakter '@' yang dilarang
        "aku_raja_dunia",  // VALID: Masuk kriteria
    }

    fmt.Println("--- VALIDASI NAMA PENGGUNA BARU ---")
    for _, username := range daftarTesPendaftar {
        // .MatchString langsung menyemburkan nilai True/False keabsahan kecocokannya.
        if validUsernameRegex.MatchString(username) {
            fmt.Printf("[DITERIMA] '%s' adalah username format yang sah.\n", username)
        } else {
            fmt.Printf("[DITOLAK]  '%s' melanggar standar format.\n", username)
        }
    }
}
```

---

### 2. Memburu Teks Tersembunyi Ekstraksi Data (FindAllString)

Asumsikan ada sebuah dokumen teks tebal berantakan yang Anda sedot dari halaman Wikipedia atau Twitter. Bos Anda memerintahkan Anda untuk merangkum seluruh Nomor Induk yang ada di teks acak tersebut ke dalam satu barisan *Array Slice*.

Fungsi keluarga `.Find` siap mengemban operasi pemburuan data masif tersebut:

```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    // Teks target yang berantakan ruwet
    paragrafMentah := "Halo, silahkan hubungi saya di nomor pribadi 0812-333-4444 atau telp rumah 021-999-8888. Jika urgen kontak asisten saya di talian 0855-111-2222 ya. Angka 1234 bukan telepon."

    // Kita menyusun rumus Pola Deteksi "Kacamata Telepon"
    // Pola: Cari urutan 3-4 Angka, dihubungkan tanda Strip, 3 Angka, Strip, 4 Angka.
    polaKacamataTelp := regexp.MustCompile(`\d{3,4}-\d{3}-\d{4}`)

    // Operasi Perburuan Massal
    // .FindAllString membutuhkan 2 parameter:
    // Parameter kedua adalah "Batas Maksimum Tangkapan". Jika diisi Angka -1,
    // ia akan menelusuri secara rakus liar sampai habis seisi paragraf dokumen ke dasar tanpa ampun!
    hasilTangkapanSlice := polaKacamataTelp.FindAllString(paragrafMentah, -1)

    fmt.Printf("Misi berburu selesai! Berhasil menyiduk total %d data nomor kontak valid.\n", len(hasilTangkapanSlice))
    fmt.Println("Daftar Aset Diekstrak: ")
    for idx, nomor := range hasilTangkapanSlice {
        fmt.Printf(" [%d] -> %s\n", idx+1, nomor)
    }
}
```

---

### 3. Substitusi Radikal dan Keamanan Penyensoran (ReplaceAllString)

Setiap *System Engineer* Backend sangat mengerti kewajiban hukum yang mengikat perusahaan untuk menjaga rahasia log kredensial alamat surel pengguna (Email) di dalam tumpukan sistem pemantau Log Kibana/Elastic.

`.ReplaceAllString()` menelusuri seluruh pola *string target* dan melumatnya habis menimpanya dengan kata sandi masker ganti string buatan Anda.

```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    // Data Log Sistem Mentah yang tercemar data identitas Klien! Bahaya privasi!
    logSistem := `[ERROR 500] Terjadi kegagalan pembayaran dari Akun john.doe19@gmail.com dengan alamat tagihan kantorceo@yahoo.co.id. Hubungi sistem pusat.`

    // Merumuskan Detektor Alamat Email Universal sederhana
    polaEmailAncaman := regexp.MustCompile(`[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

    // Tembakan Lumat Timpa Massal (Gantikan Semua Email yang kau temui dengan Tanda Bintang Sensor Masking!)
    logAmanTerpublikasi := polaEmailAncaman.ReplaceAllString(logSistem, "[SUREL-DISEMBUNYIKAN-DEMI-KEAMANAN]")

    fmt.Println("--- Log Mentah Asli (Privasi Bocor) ---")
    fmt.Println(logSistem)

    fmt.Println("
--- Log Masking Modifikasi Rilis Bersih ---")
    fmt.Println(logAmanTerpublikasi)
}
```

---

## Bagian Lanjutan: Pembedahan Pola Abad Pertengahan, Tangkapan Penamaan Kelompok Spesifik (*Named Capturing Groups*), dan Batasan Kuota Penggantian Komputasi

Pemahaman awal mengenai fungsi `regexp` (Reguler Ekspresi) seringkali terbentur di tembok tebal ketika Anda dihadapkan oleh struktur pola rumit raksasa, misalnya sebuah teks log peladen web `Nginx` atau hasil ekstrak balasan protokol mesin surel *SMTP* di mana informasi yang Anda buru bukan lagi sekadar "Ya atau Tidak", melainkan sepotong parameter dinamis (Nomor ID Pelanggan, atau Kode Resi) yang terselip berhimpitan di antara kata sandi dan tanggal acak. Di arena tingkat dewa inilah fungsi ekstraksi Grup Sub-Kecocokan (*Capturing Groups*) mengambil alih panggung penguasaan data bahasa *Go*.

### 1. Rahasia Terbesar Perburuan Data: Grup Sub-Tangkapan Bersarang (*Submatches*)

Fungsi dasar `.FindString()` maupun `.FindAllString()` hanya mencaplok satu kalimat gelondongan secara kasar secara mutlak dari awal hingga akhir kata. Namun jika kita disodorkan rentetan teks rumit seperti: *"Halo, nomor kartu perdana telpon milik istri Bapak Ahmad adalah: 0812-9999-5555. Silahkan topup 50000 segera."*

Lalu instruksi manajer Anda: "Sistem Go saya tak peduli urusan kalimat Halo atau Silahkan Topup! Saya HANYA INGIN NAMA (Bapak Ahmad) dan NOMOR TELEPON-nya (0812-9999-5555) masuk langsung terpisah ke dua variabel Go yang berbeda!"

Teknik penyelesaian: **`(Grup Kurung)`** dan perintah andalan **`.FindStringSubmatch()`**

Setiap kali Anda meletakkan tanda kurung biasa `(...)` di dalam rumus *Regex*, mesin Go *RE2* menginterpretasikannya sebagai arahan "Tangkap dan Simpan bagian spesifik yang cocok dengan aturan di dalam kurung ini ke Array Nomor 1, lalu yang kurung berikutnya ke Array Nomor 2".

```go
// 1. Data Mentah Acak berantakan
// kalimatCSO := "Laporan keluhan terbaru dari pelanggan VIP atas nama Tuan Ahmad. Nomor kontaknya 0812-9999-5555."

// 2. Rumusan Pola Ekstraksi Bersarang (Submatch Groups)
// Pola bacaan kita: "nama Tuan <NAMA>, Nomor kontaknya <ANGKA-STRIP-ANGKA>"
// Tanda Kurung (...) artinya: INI ADALAH BAGIAN YANG SAYA TANGKAP KE ARRAY!
// polaRadarTangkapan := regexp.MustCompile(`nama Tuan ([a-zA-Z]+)\.\s*Nomor kontaknya (\d{4}-\d{4}-\d{4})`)

// 3. Eksekusi Pemburuan (Submatch bukan String biasa)
// arrayKepinganData := polaRadarTangkapan.FindStringSubmatch(kalimatCSO)

// 4. MEMBEDAH RAHASIA SLICE TANGKAPAN (Wajib Mengerti!)
// Indeks ke-[0] : SELALU BERISI hasil tangkapan kalimah utuh penuh (Gelondongan kasar)!
// Indeks ke-[1] : HASIL TANGKAPAN KURUNG PERTAMA KITA (NAMA)
// Indeks ke-[2] : HASIL TANGKAPAN KURUNG KEDUA KITA (NOMOR TELEPON)
```

### 2. Kustomisasi Transformasi Fungsi Ekstrem (*ReplaceAllStringFunc*)

Keluarga *Replace* `.ReplaceAllString("Kata Asli", "Kata Baru")` sangat berguna untuk menyensor atau mengganti kata secara statis mati konstan berulang ulang.
Tapi apa yang harus Anda perbuat jika Anda disuruh "Mencari semua kata yang salah format huruf kecil di sebuah paragraf HTML, namun setiap kali Anda menemukannya, Anda harus mengubah HANYA KATA ITU menjadi HURUF BESAR SEMUA secara dinamis dan menaruhnya kembali ke posisi tempatnya semula"?

Senjata rahasia fungsi penyusupan modifikasi tingkat lanjut (*Advanced String Substitution Function*): **`.ReplaceAllStringFunc(stringAwal, FungsiManipulatorCallback)`**.

Setiap kali mesin pencari Regex mendeteksi satu kecocokan pola di kalimat itu, Go akan menghentikan waktu (Pause), menyodorkan potongan kata asli tersebut kepada fungsi injeksi kustom buatan Anda sendiri (`func`), mempersilakan Anda berkreasi memodifikasi huruf potongan teks itu dengan logika *If-Else* di Go sesuka hati, lalu Anda wajib mereturn `String` modifikasi baru yang oleh mesin Go otomatis dijahit lem dimasukkan kembali mengganti teks asal ke posisi kalimat semula! Ajaib bukan kepalang!

```go
// Paragraf Dokumen HTML Kacau (Beberapa judul kota tidak kapital)
// dokumenWisata := "Rencana perjalanan menuju kota jakarta, lalu ke kota bandung."

// Menggunakan kelas pencari huruf kecil [a-z] agar kita cuma nangkep kota yang penulisan hurufnya salah kecil semua.
// radarKataKotaSalah := regexp.MustCompile(`kota\s+([a-z]+)`)

// KITA MULAI PESTA BEDAH KUSTOM!
// dokumenDiperbaikiSempurna := radarKataKotaSalah.ReplaceAllStringFunc(dokumenWisata, func(kataTarget string) string {
//    // Tugas kita adalah merubah "kota jakarta" -> menjadi "KOTA JAKARTA"
//    hasilPerbaikan := strings.ToUpper(kataTarget)
//    return hasilPerbaikan
// })
// Cetakan Akhir: "Rencana perjalanan menuju KOTA JAKARTA, lalu ke KOTA BANDUNG..."
```

Membongkar kekuatan fungsional dinamis `regexp` serta teknik cengkraman Sub-tangkapan Kurung (*Submatch Grouping*) tidak sebatas menyajikan kemudahan, melainkan menyokong pondasi sistem kecerdasan orkestrasi otomatis (*Orchestration Parsing*) bilamana Go dituntut untuk menerjemahkan masukan format asing berwujud teks log mentah pangkalan data, menjadikan Peladen *Backend* Anda perkasa menangani badai arus data dari sistem warisan (*Legacy Mainframes*).

## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```

---

## Studi Kasus Dunia Nyata: Penganalisis Berkas Jejak Sistem Log Raksasa (Log Parser Pipeline) Berbasis Ekspresi Reguler

Penggunaan paling krusial dari package `regexp` dalam arsitektur perusahaan perangkat lunak adalah saat tim keandalan situs (*Site Reliability Engineering / SRE*) dipaksa untuk menyelidiki dan mengekstrak informasi rahasia dari ratusan Gigabyte berkas *Syslog NGINX* web server terenkripsi demi melacak peretas gelap alamat IP liar dari negeri seberang.

Anda tak mungkin mengekstrak jutaan baris teks log dengan paket `strings` manual. Rumusan Regular Expression menangani hal ini secara elegan. Di studi kasus ini, kita akan menangkap 3 komponen sakti dari 1 baris teks log *Apache/NGINX*: **Alamat IP Klien**, **Waktu Stempel**, dan **URL Target Sasaran**.

### Pemetaan Pola Rumit dengan Kelompok Tangkapan (*Capturing Group*) Bernama Spesifik

Untuk mempermudah pengambilan Array, Go mendukung fitur ajaib `(?P<nama_grup>pola)` yang menamai laci array tangkapan regex Anda, sehingga Anda tak perlu repot lagi menghafal Index Nomor [1], [2], dst.

```go
package main

import (
    "bufio"
    "fmt"
    "regexp"
    "strings"
)

func main() {
    // 1. Simulasikan 3 baris Log Nginx murni (Biasanya ini dibaca dari file via os.Open)
    dataMentahLog := `192.168.0.1 - - [25/Jan/2025:10:30:15 +0000] "GET /api/v1/auth/login HTTP/1.1" 200 512
10.0.0.99 - - [25/Jan/2025:10:31:01 +0000] "POST /api/v1/checkout HTTP/1.1" 500 128
45.33.22.11 - - [25/Jan/2025:10:32:44 +0000] "HACK /phpmyadmin/config HTTP/1.1" 403 0`

    // 2. RUMUSAN POLA MATA DEWA:
    // Mencari IP:      (?P<IPAddress>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
    // Mengabaikan "- - [" dan mencari waktu: \[ (?P<WaktuTamu>[^\]]+) \]
    // Mengabaikan Kutip dan mengambil Metode & URL Target: "(?P<Metode>\w+) (?P<URLTarget>[^\s]+)

    rancanganMataDewa := `^(?P<IPAddress>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}) - - \[(?P<WaktuTamu>[^\]]+)\] "(?P<Metode>\w+) (?P<URLTarget>[^\s]+)`
    mesinParserSakti := regexp.MustCompile(rancanganMataDewa)

    // Ambil daftar nama kelompok yang kita tulis di atas (Mendapatkan slice: ["", "IPAddress", "WaktuTamu", "Metode", "URLTarget"])
    kunciNamaKelompok := mesinParserSakti.SubexpNames()

    fmt.Println("=== MEMULAI EKSTRAKSI INTELEJEN LOG SERVER ===")

    // Kita gunakan Scanner agar hemat memori (Membaca baris per baris, bukan sekaligus!)
    pemindaiPipa := bufio.NewScanner(strings.NewReader(dataMentahLog))

    for pemindaiPipa.Scan() {
        barisLogTunggal := pemindaiPipa.Text()

        // Sedot dan tangkap submatch ke dalam Array
        tangkapanArray := mesinParserSakti.FindStringSubmatch(barisLogTunggal)

        if tangkapanArray == nil {
            fmt.Println("[DIABAIKAN] Baris sampah tak beraturan:", barisLogTunggal)
            continue
        }

        // KEAJAIBAN GO: Kita susun Peta (Map) agar lebih elok diakses berdasarkan nama!
        hasilTangkapanBersih := make(map[string]string)

        for i, namaKunci := range kunciNamaKelompok {
            // Abaikan Index ke 0 (karena index 0 isinya teks gelondongan full) dan namaKunci yang kosong
            if i != 0 && namaKunci != "" {
                hasilTangkapanBersih[namaKunci] = tangkapanArray[i]
            }
        }

        // Tampilkan Intelijen yang Didapat
        fmt.Println("\n>> KECURIGAAN AKTIVITAS DITEMUKAN:")
        fmt.Println("  - Pelaku IP Asal  :", hasilTangkapanBersih["IPAddress"])
        fmt.Println("  - Jam Kejahatan   :", hasilTangkapanBersih["WaktuTamu"])
        fmt.Println("  - Operasi Target  :", hasilTangkapanBersih["Metode"], hasilTangkapanBersih["URLTarget"])

        // Pengecekan Bahaya Otomatis
        if hasilTangkapanBersih["Metode"] == "HACK" {
            fmt.Println("    [ALARM MERAH] Blokir segera alamat IP ini dari Firewall Router Induk AWS!")
        }
    }
}
```

Penyelarasan paduan fungsi *Regexp Submatch* bertenaga kecepatan penafsiran waktu linier *RE2* ini membukakan tirai wawasan mutlak perihal identifikasi celah anomali, mengkonversikan bahasa peladen Go (Golang) Anda bukan sebatas sebagai pembuat Aplikasi Peladen Web belaka, melainkan mampu dioperasikan sebagai Mesin Anti-Virus Penganalisis Intrusi Jaringan tingkat dewa berkapasitas besar (*Massive Intrusion Detection Analyst System*).
