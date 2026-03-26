# Modul: `strings`

## Ringkasan
Package `strings` adalah perpustakaan operasi bawaan (Standard Library) Go yang didesain secara brilian untuk memanipulasi, mencari, mengganti, dan menganalisis string berbasis UTF-8.

## Penjelasan Lengkap (Fungsi & Tujuan)
Di dalam bahasa pemrograman Go, data dengan tipe `string` tidak bisa diubah isinya (*immutable*). Artinya, setelah sebuah string diciptakan ke dalam tumpukan memori, kita sama sekali tidak dapat memodifikasi huruf di dalamnya tanpa membongkar ulang seluruh kerangkanya. Setiap kali Anda menggunakan package `strings` (misalnya untuk mengubah teks kecil menjadi Kapital), Anda **tidak memodifikasi** variabel asli tersebut. Fungsi itu sebenarnya melakukan alokasi memori baru, menyalin data aslinya, lalu mengubah karakter hasil salinan itu, dan akhirnya mengembalikan nilai variabel string *baru* kepada Anda.

Package ini dirancang untuk melakukan tugas ini secepat dan se-efisien mungkin, mengatasi limitasi *immutability* tersebut dengan optimisasi tinggi (*Assembly language* di balik layar untuk arsitektur CPU tertentu).

**Tujuan dan Fungsi Utama:**
1.  **Pencarian Pola dan *Matching*:** Menyediakan arsenal deteksi—apakah kalimat ini mengandung kata rahasia? Apakah awalan kalimat ini sesuai format yang diminta API? Menemukan letak (indeks) spesifik suatu karakter di dalam ratusan teks.
2.  **Konversi Kasus Huruf (Casing):** Menyelaraskan ukuran huruf teks tak rapi dari input pengguna eksternal menjadi `UPPERCASE`, `lowercase`, atau sekadar *Title Case*. Ini krusial untuk pembandingan nilai data.
3.  **Restrukturisasi dan Pembongkaran (Split & Join):** Menghancurkan kalimat utuh menjadi deretan potongan elemen kata dalam format *Slice* (*Array*), dan menempelkannya kembali dengan sembarang elemen separator yang dinginkan programmer. Sangat penting ketika *parsing* file spesifik seperti *Comma Separated Values* (CSV).
4.  **Sterilisasi Data (Trim):** Membersihkan karakter nakal, spasi maya, atau jeda *newline* sisa dari input tak terduga (*Data Sanitization*), yang lazim dilakukan sebelum kita merekam entri form klien menuju repositori *Database SQL*.
5.  **Replikasi dan Penggantian Lanjut (Replace & Repeat):** Menghapus kalimat "kotor" dan menyensornya dengan kalimat aman secara masif melalui mekanisme ganti (*ReplaceAll*).

**Mengapa menggunakan `strings`?**
Jika `net/http` adalah pengangkut datanya, maka `strings` adalah mesin bedah di sisi belakang. Mulai dari menormalkan format email saat registrasi (`strings.ToLower`), mengekstrak token Bearer dari Header HTTP (`strings.TrimPrefix`), hingga merutekan *bot command* di Discord atau Telegram (`strings.HasPrefix`). Tak ada hari bagi programmer Go di mana mereka tak memanggil fungsi dari package vital ini.

---

## Daftar Fungsi Umum dan Cara Penggunaannya Secara Mendalam

### 1. Mencari Fakta di dalam String (Inspeksi & Boolean)

Fungsi berikut membedah isi teks dan mengembalikan nilai kebenaran `bool` atau posisi angka pasti (*int index*).

*   **`strings.Contains(s, substr)`**: Mengembalikan nilai `true` jika variabel `s` mengandung/memiliki teks `substr`. Peka huruf besar-kecil (Case Sensitive!).
*   **`strings.ContainsAny(s, chars)`**: Mirip dengan contains, namun ia tidak mencari kata utuh, melainkan mencari **setidaknya satu** dari karakter apa pun yang ada di string `chars` berada di dalam teks tersebut.
*   **`strings.HasPrefix(s, prefix)`**: Sangat efisien mengecek apakah string diawali persis dengan potongan `prefix`.
*   **`strings.HasSuffix(s, suffix)`**: Mengecek apakah string diakhiri dengan potongan `suffix` tersebut.
*   **`strings.Count(s, substr)`**: Menghitung secara eksak seberapa banyak pola string tersebut berulang kali muncul di dalam variabel string `s`.

```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    artikel := "Golang (Go) adalah bahasa pemrograman open source yang dikembangkan oleh tim engineer di Google."

    // Pencarian utuh
    fmt.Println("Apakah artikel mengandung kata 'bahasa'?", strings.Contains(artikel, "bahasa"))
    fmt.Println("Apakah artikel mengandung kata 'google'? (Case sensitive!)", strings.Contains(artikel, "google"))

    // Pengecekan Awalan / Akhiran
    command := "!kick user123"
    if strings.HasPrefix(command, "!") {
        fmt.Println("Ini adalah bot command valid.")
    }

    filename := "laporan_keuangan_januari.pdf"
    if strings.HasSuffix(filename, ".pdf") {
        fmt.Println("Dokumen siap dibaca sebagai PDF.")
    }

    // Menghitung kemunculan (jumlah huruf 'a' dalam artikel)
    jumlah_a := strings.Count(artikel, "a")
    fmt.Printf("Huruf 'a' muncul sebanyak %d kali di dalam artikel.\n", jumlah_a)
}
```

---

### 2. Mengubah Dimensi Huruf (Manipulasi Casing)

Sangat penting untuk *database normalization* dan pengecekan kata sandi, mengubah ukuran semua karakter di dalam string secara terpadu.

*   **`strings.ToLower(s)`**: Merendahkan semua karakter menjadi huruf kecil. (Sering digunakan sebelum melakukan komparasi `==`).
*   **`strings.ToUpper(s)`**: Menaikkan semua karakter menjadi huruf besar/kapital sempurna.
*   **`strings.ToTitle(s)`**: Mengubah seluruh string menjadi Titel besar, yang untuk banyak tipe bahasa ASCII (seperti Latin Inggris/Indonesia) akan memproduksi hasil yang mirip dengan *ToUpper*.
*   *(Catatan: Jika Anda ingin format Kapital Tiap Kata awal ("Title Case"), Anda perlu mengimpor modul luar `golang.org/x/text/cases` pada standar Go versi modern.)*

```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    emailInputUser := "  Budi.SANTOSO@Perusahaan.CO.ID  "

    // 1. Membersihkan spasi dulu (Trim)
    emailBersih := strings.TrimSpace(emailInputUser)

    // 2. Merubah seluruhnya menjadi lowercase standar untuk database
    emailFinal := strings.ToLower(emailBersih)

    fmt.Println("Data asli      : ", emailInputUser)
    fmt.Println("Data disimpan  : ", emailFinal)

    // Uppercase untuk peringatan UI
    pesan := "perhatian, dilarang melintas!"
    fmt.Println("Sistem Peringatan: ", strings.ToUpper(pesan))
}
```

---

### 3. Membelah dan Mempersatukan String (Split & Join)

Mengubah satu string masif menjadi daftar (`slice []string`) yang elemennya bisa Anda interogasi (looping). Atau, Anda sudah memiliki array data yang ingin dicetak ke layar menjadi format satu rentet kalimat yang dibatasi karakter koma.

*   **`strings.Split(s, sep)`**: Membelah string utama di setiap kemunculan *separator* `sep`, membuang separatornya, dan meletakkan sisa bongkahan tersebut berurutan ke dalam sebuah Array *Slice*. (Jika `sep` diset kosong `""`, ia akan membelah setiap huruf karakter individu).
*   **`strings.Join(a []string, sep)`**: Kebalikan dari Split. Menerima *Slice of String* dan merekatkannya kembali menjadi teks menggunakan semen penengah `sep`.

```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    // 1. Kasus Split: Membaca format direktori
    alamatPath := "/usr/local/bin/golang"
    komponenFolder := strings.Split(alamatPath, "/")

    fmt.Println("Hasil membelah URL path:")
    for i, p := range komponenFolder {
        fmt.Printf("Indeks [%d] -> '%s'\n", i, p)
        // Perhatikan bahwa indeks 0 kosong ('') karena '/' pertama terletak di awal kalimat
    }

    // 2. Kasus Join: Menyatukan database query
    kolomPilih := []string{"id", "nama", "email", "tanggal_lahir"}

    // Merekatkan slice di atas dengan koma spasi
    stringQuery := strings.Join(kolomPilih, ", ")

    hasilSQL := fmt.Sprintf("SELECT %s FROM tabel_user;", stringQuery)
    fmt.Println("\nHasil rakitan SQL Query:")
    fmt.Println(hasilSQL)
}
```

---

### 4. Membersihkan Kotoran Teks (Trim)

Validasi *input* klien adalah mimpi buruk pemrograman jika Anda tidak melakukan *sanitasi*. Terutama mengatasi klien yang gemar menambahkan "Spasi kosong / karakter Tab" di depan nama pengguna mereka saat mengisi kolom pendaftaran.

*   **`strings.TrimSpace(s)`**: Menghapus seluruh karakter spasi yang *tak terlihat* (` `), spasi tabulasi (`\t`), serta jeda pemotong baris (*newline* `\n` atau *carriage return* `\r`) dari SISI DEPAN dan SISI BELAKANG string. Karakter spasi di tengah antar kata tidak disentuh.
*   **`strings.TrimPrefix(s, prefix)`**: Melucuti spesifik suatu string awal secara paksa. Hanya jika kalimat tersebut benar-benar diawali dengan string awal (prefix) tersebut. Sangat berguna untuk mensterilkan URL API.
*   **`strings.TrimSuffix(s, suffix)`**: Kebalikan dari TrimPrefix, melucuti akhiran secara absolut.

```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    komentarRaw := "\n\t\t   Ini adalah komentar yang sangat buruk formatnya.  \n\r "
    komentarBersih := strings.TrimSpace(komentarRaw)
    fmt.Printf("Komentar Raw: '%s'\n", komentarRaw)
    fmt.Printf("Komentar Bersih: '%s'\n", komentarBersih)

    // Contoh Ekstraksi HTTP Header menggunakan TrimPrefix
    authHeader := "Bearer eyJhbGciOiJIUzI1NiIsInR5c" // Token Autentikasi JWT

    // Menghapus tulisan "Bearer " sehingga kita hanya mendapatkan string token murni
    tokenAman := strings.TrimPrefix(authHeader, "Bearer ")
    fmt.Println("Token JWT ter-ekstrak:", tokenAman)
}
```

---

### 5. Substitusi dan Perbaikan (Replace)

Fungsi utama untuk mengganti sub-kalimat secara spesifik di dalam paragraf raksasa.

*   **`strings.Replace(s, old, new, n)`**: Mengganti teks asli `old` dengan karakter baru `new` sebanyak `n` kali kemunculan. Hal terpenting: **Jika angka `n` dideklarasikan sebagai -1, artinya Go akan merobak dan menimpa SELURUH letak kata tersebut di dalam string sampai tak bersisa.**
*   **`strings.ReplaceAll(s, old, new)`**: Singkatan modern (diperkenalkan sejak Go 1.12) yang melakukan hal yang persis sama dengan memanggil `.Replace` dengan nilai *count* -1. Penulisannya lebih elegan, meminimalisir kemungkinan Anda kelupaan menulis -1 di belakang koma.

```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    cerita := "Kucing saya lucu. Anjing peliharaan tetangga suka menyalak pada Kucing saya. Kucing adalah hewan imut."

    // Mengganti hanya 1 kata Kucing pertama yang ditemukan
    gantiSatu := strings.Replace(cerita, "Kucing", "Singa", 1)
    fmt.Println("Ganti Pertama Saja:", gantiSatu)

    // Merubah seluruh kata "Kucing" menjadi "Harimau" di satu naskah
    // Ini sama fungsinya dengan: strings.Replace(cerita, "Kucing", "Harimau", -1)
    gantiSemua := strings.ReplaceAll(cerita, "Kucing", "Harimau")
    fmt.Println("Ganti Semua Secara Massal:", gantiSemua)
}
```

---

## Bagian Lanjutan: Optimasi Ekstrem dan Keamanan String (Immutability)

String di bahasa Go adalah "raja" tipe data. Mereka digunakan di mana-mana: nama variabel sistem, *payload body HTTP*, query database. Namun, banyak *engineer* yang tidak memahami bagaimana mesin *Runtime* Go mengelola String di belakang layar, berakibat pada pembengkakan tagihan server *Cloud* (AWS/GCP) yang tidak perlu karena program memakan *RAM* bergiga-giga.

### 1. Rahasia *Immutability* (Sifat Kekal Teks)

Di Go, `string` adalah sepotong memori pembacaan murni (*Read-Only byte slice*). Setelah string tercipta, isi hurufnya **TIDAK BISA** dirubah!
Jika Anda mencoba merubah huruf pertama: `nama[0] = 'B'`, *Compiler Go* akan menjerit menolak kompilasi.

**Dampak bagi Package `strings`:**
Karena sifat kekal (*immutability*) ini, setiap kali Anda memanggil fungsi manipulasi di package `strings`, sebut saja `strings.ToUpper("hello")`, sistem operasi harus **menyewa blok memori RAM baru yang kosong**, menyalin kata "HELLO" ke sana, dan memberikan alamat memori baru itu kepada Anda. Memori lama "hello" akan dibiarkan mati menjadi sampah untuk dipungut *Garbage Collector* (GC).

Jika Anda menggunakan fungsi `strings.ReplaceAll()` pada sebuah dokumen ensiklopedia bervolume 10 MB sebanyak seratus kali, Anda baru saja membuang 1 Gigabyte RAM tanpa tujuan, menyiksa server secara instan!

### 2. Penyelamat Kinerja: `strings.Builder`

Pada zaman dahulu, orang menggabungkan kata dengan memanggil fungsi di atas secara berulang, atau bahkan lebih parah, memakai operator tambah:

```go
// SANGAT LAMBAT DAN BOROS MEMORI (Menghasilkan sampah string baru di setiap loop!)
// pesanBesar := ""
// for _, kata := range daftarRibuanKata {
//    pesanBesar += kata + " "
// }
```

Mulai dari Go versi modern, para dewa di Google memperkenalkan struktur dewa penolong: `strings.Builder`. Ini adalah mesin traktor memori mutakhir.
`strings.Builder` menyewa SATU blok tanah memori besar sejak awal. Setiap kali Anda menambahkan kata, ia hanya meletakkannya bersebelahan di tanah lapang tersebut (In-Place Mutation). Tidak ada "Copy" berlebih, tidak ada siksaan terhadap sistem *Garbage Collector* (GC).

```go
import "strings"

// var rakitanBaja strings.Builder

// PRAKTIK TERBAIK: Kita tebak (prediksi) kira-kira berapa ukuran akhirnya (misal 10.000 byte)
// Dengan fungsi Grow, Builder langsung memborong kavling RAM 10KB di detik pertama!
// rakitanBaja.Grow(10000)

// for _, kata := range daftarRibuanKata {
//    rakitanBaja.WriteString(kata)
//    rakitanBaja.WriteString(" ")
// }

// Proses Finalisasi mutlak: Merubah kavling tersebut menjadi wujud String Abadi.
// Ini dieksekusi HANYA 1 KALI di paling akhir, tanpa memakan waktu nanodetik sama sekali!
// hasilAkhir := rakitanBaja.String()
```
Selalu gunakan `strings.Builder` untuk menyatukan puluhan baris kueri *SQL*, menyusun hasil render *Template Web HTML*, atau merakit pesan file CSV!

### 3. Kecepatan Kilat Validasi tanpa Alokasi

Go menyediakan fungsi perbandingan yang luar biasa efisien yang tidak mengonsumsi memori tambahan nol-byte (0 Allocations).

*   **`strings.EqualFold(a, b)`**: Fungsi ini adalah "Permata Tersembunyi". Saat Anda mencocokkan kredensial surel (*email login*), Anda butuh membandingkannya tanpa mempedulikan besar/kecil huruf (*Case Insensitive*).
    Pemula sering melakukan ini:
    `if strings.ToLower(emailInput) == strings.ToLower(emailDB)`
    Itu **SANGAT BURUK** karena `.ToLower()` menghasilkan 2 string memori baru secara gratis!

    Cara yang benar, elegan, dan *Zero-Memory-Allocation*:
    `if strings.EqualFold(emailInput, emailDB)`
    Sistem Go membandingkan mereka secara langsung dalam taraf huruf biner per karakter secara silang tanpa membikin objek baru!

*   **`strings.Clone(s)`**: Fitur langka (Go 1.18+). Berfungsi untuk melepas keterikatan memori. Jika Anda memiliki dokumen Novel raksasa 5 GB (*String Induk*), dan Anda menggunakan trik Slice `cuplikan := novelBesar[0:10]` untuk mengambil 10 karakter pertamanya.
    Selama variabel `cuplikan` masih eksis, maka **Novel Induk raksasa 5 GB itu TIDAK AKAN BISA DIHAPUS oleh GC** (karena `cuplikan` kecil itu diam-diam menunjuk alamat ujung ke memori raksasa tersebut).
    Solusi: `cuplikan = strings.Clone(novelBesar[0:10])`. Ini memaksa `cuplikan` digandakan mandiri, dan membiarkan Novel raksasa 5GB mati dilepas (dihapus).

Penguasaan teknik di atas membedakan kode *Scripting* biasa dengan *Microservices Backend System* kelas atas, menjadikan aplikasi backend memakan jumlah RAM sekecil-kecilnya.

## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
