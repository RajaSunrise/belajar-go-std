# Modul: `sort`

## Ringkasan
Package `sort` menyediakan primitif fungsional algoritma *pengurutan* (sorting) dan *pencarian biner* (binary search) yang sangat dioptimasi (menggunakan implementasi algoritma *Pattern-Defeating Quicksort* atau PDQSort yang mutakhir) untuk menyortir himpunan data dinamis (*Slice*) maupun koleksi struktur data kompleks buatan pengguna secara efisien.

## Penjelasan Lengkap (Fungsi & Tujuan)
Mengurutkan sebuah daftar (misalnya daftar harga barang dari termurah ke termahal, atau daftar nama kontak sesuai abjad) adalah operasi fundamental algoritma ilmu komputer yang seringkali menjadi momok pemicu pelambatan performa jika ditulis secara sembarangan (*misalnya menggunakan Bubble Sort manual yang memakan waktu eksponensial $O(n^2)$*).

Package `sort` menghilangkan kebutuhan pengembang untuk menciptakan kembali roda yang bundar. Dengan memanggil fungsinya, Anda langsung memperoleh tenaga algoritma kelas kakap $O(n \log n)$ yang teruji stabilitasnya oleh Google. Uniknya, di awal penciptaannya, Go tidak menyediakan pengurutan universal satu-baris-jadi karena batasan ketiadaan *Generics* (Tipe data generik tipe-T). Seiring evolusi perbaikan yang agresif (khususnya rilis terobosan `sort.Slice` di Go 1.8), fitur pada package ini semakin merampingkan alur penulisan kode tanpa mengorbankan keamanan tipe (*Type-Safety*).

**Tujuan dan Fungsi Utama:**
1.  **Pengurutan Tipe Primitif Bawaan:** Fungsi kilat (seperti `sort.Ints` dan `sort.Float64s`) siap mengeksekusi pengurutan daftar Array angka dan teks secara langsung secara menaik (*Ascending* / dari terkecil hingga terbesar) secara langsung.
2.  **Kustomisasi Urutan Objek Kompleks (*sort.Slice*):** Kapabilitas untuk mengambil daftar struktur kompleks (seperti *Slice of Struct User* yang berisi Umur, Nama, Jabatan) dan mendidik kompilator bagaimana cara mengukur porsi perbandingan mana dari properti obyek tersebut yang lebih "berat", lebih "panjang" atau lebih "tua" dengan memberikan injeksi fungsi pendefinisi pembanding kustom.
3.  **Membalik Urutan (*Reverse*):** Mengubah hasil urutan naik menjadi terbalik (*Descending* / Menurun) dengan menggunakan pelindung pembungkus Antarmuka terbalik (`sort.Reverse`).
4.  **Pencarian Biner Berkecepatan Cahaya (*Binary Search*):** Apabila Anda memiliki satu juta jejak *Log Database ID* dalam satu *Slice* di Memori RAM—asalkan daftarnya *SUDAH dalam keadaan diurutkan!*—Anda bisa menemukan di index mana satu jejak ID tertentu bersembunyi. Dibanding mengeceknya satu demi satu dari awal hingga akhir (memakan waktu maksimal sejuta iterasi), algoritma *Binary Search* Package `sort` hanya memakan waktu sekedipan debu (maksimal ~20 iterasi tebakan memotong tengah).

**Mengapa menggunakan `sort`?**
Jika aplikasi backend Anda tidak menggunakan SQL `ORDER BY` untuk mengambil urutan (karena datanya hasil agregasi perbandingan banyak API secara In-Memory), menyajikan daftar Peringkat Nilai Siswa (Leaderboard) atau mengurutkan katalog Keranjang Belanja Pelanggan di layar Android membutuhkan sentuhan mutlak dan performa solid dari fungsi di package ini.

---

## Daftar Fungsi Umum dan Cara Penggunaannya Secara Mendalam

### 1. Eksekusi Pengurutan Primitif Instan (Ints, Float64s, Strings)

Apabila *Slice* daftar himpunan kita murni primitif standar satu tipe utuh, jangan bersusah pikir panjang. Eksekusi utilitas praktis siap seduh ini! (*Perlu diperhatikan: Fungsi pengurutan pada Go akan MEMODIFIKASI SECARA LANGSUNG objek aslinya/In-Place mutation, jadi list aslinya akan tergantikan secara ajaib!*)

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    // KASUS 1: Mengurutkan angka murni (Ranking Ujian)
    skorUjian := []int{85, 45, 92, 70, 60}

    // Tembakan perintah.
    sort.Ints(skorUjian)

    // Perhatikan, ia termodifikasi secara otomatis di tempat aslinya dari kecil ke besar!
    fmt.Println("Peringkat Skor Terendah-Tertinggi:", skorUjian)

    // KASUS 2: Mengurutkan daftar kata berdasarkan alfabetis kamus A-Z
    daftarKlien := []string{"Xaverius", "Bambang", "Alfonso", "Zulkarnain"}

    sort.Strings(daftarKlien)

    fmt.Println("\nSusunan Absen Huruf A-Z:", daftarKlien)
}
```

---

### 2. Pengurutan Kustom Struktur Lanjut Berdasarkan Parameter Abstrak (`sort.Slice`)

Ini adalah penggunaan yang paling mendominasi panggung harian pemrograman *Backend*. Anda memiliki kumpulan data *Struct Database* Obyek `Produk`, dan pengunjung Web UI mengklik menu sortir: "Urutkan berdasarkan Harga Termurah" lalu mengklik "Urutkan berdasarkan Bintang Ulasan Terbanyak". Fungsi `sort.Slice` adalah pahlawan pembawa solusi ajaib ini.

Anda diharuskan menyertakan fungsi pendefinisi tak bernama (*Anonymous Callback Function*) yang mengajari sang Mesin *Sorter* cara membedakan mana obyek yang lebih pantas berada di depan berdasarkan patokan dua buah variabel indeks fiktif (kiri `i`, lawan kanan `j`). Semudah memberikan jawaban *True/False*.

```go
package main

import (
    "fmt"
    "sort"
)

type BarangDagang struct {
    KodeItem string
    HargaRp  float64
    Terjual  int
}

func main() {
    katalogToko := []BarangDagang{
        {KodeItem: "Sepatu Kets", HargaRp: 450000, Terjual: 120},
        {KodeItem: "Topi Caping", HargaRp:  35000, Terjual: 900},
        {KodeItem: "Tas Kulit  ", HargaRp: 890000, Terjual: 45},
        {KodeItem: "Kaos Sablon", HargaRp:  75000, Terjual: 250},
    }

    // SCENARIO A: Pelanggan Meminta Urutan Harga dari Termurah hingga Termahal (Ascending)
    // Aturannya: Kembalikan TRUE Jika Barang Kiri (i) LEBIH MURAH (<) dari Barang Kanan (j)
    sort.Slice(katalogToko, func(i, j int) bool {
        return katalogToko[i].HargaRp < katalogToko[j].HargaRp
    })

    fmt.Println("=== KATALOG FILTER: HARGA TERMURAH ===")
    for _, barang := range katalogToko {
        fmt.Printf("- %s | Rp %.0f\n", barang.KodeItem, barang.HargaRp)
    }

    // SCENARIO B: Pelanggan Meminta Barang Paling Laris (Terjual Paling Banyak / Descending Menurun)
    // Aturannya DIBALIK: Kembalikan TRUE Jika Barang Kiri (i) LEBIH BESAR PENJUALANNYA (>) dari Kanan (j)
    sort.Slice(katalogToko, func(i, j int) bool {
        return katalogToko[i].Terjual > katalogToko[j].Terjual
    })

    fmt.Println("\n=== KATALOG FILTER: PALING LARIS TERJUAL (BEST SELLER) ===")
    for _, barang := range katalogToko {
        fmt.Printf("- %s | Laku: %d buah\n", barang.KodeItem, barang.Terjual)
    }
}
```

---

### 3. Pencarian Biner Terarah (`sort.Search`)

Bayangkan permainan anak sekolahan: "Saya memikirkan satu angka dari 1 hingga 100. Tebak angka berapa. Saya hanya akan menjawab: Kebesaran atau Kekecilan."
Jika Anda *bodoh/naif*, Anda akan menebak berurutan: "1? 2? 3? 4?". Jika angkanya 99, Anda perlu 99 pertanyaan. Membosankan dan lambat (*Linear Search*).

Jika Anda orang jenius, tebakan pertama Anda pastilah pertengahan: "50!".
Jawaban: "Kekecilan."
Di detik pertama itu juga Anda menyapu habis mencoret kemungkinan angka 1-49! Tebakan kedua: "75!".
Maka cuma perlu maksimal 7 tebakan, angka 100 pun terpecahkan. Kecepatan magis ini disebut algoritma *Binary Search Logaritmik*. Package `sort` mendedikasikan fitur pembunuh naga ini.

*SYARAT MUTLAK: Tumpukan Slice Data Array WAJIB sudah disortir sebelumnya agar fungsi pencarian biner ini tidak tersesat salah arah!*

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    // 1. Array SUDAH diurutkan menaik! Syarat mutlak!
    daftarIDKartuKependudukan := []int{
        1100, 1420, 2055, 3091, 5800, 6211, 8880, 9102, 9991, 10500,
    }

    targetIDPencarian := 8880

    // 2. Eksekusi Tembakan Algoritma Pencarian Biner
    // Ingat, kita bertugas memberikan petunjuk: Kapan tebakannya sudah mencukupi target yang dicari? (Tebakan >= Target)
    indexHasilDitemukan := sort.Search(len(daftarIDKartuKependudukan), func(tebakanIndex int) bool {
        // Apakah nilai pada tebakan indeks tengah ini, SAMA ATAU LEBIH BESAR dibanding Angka yg kita cari?
        return daftarIDKartuKependudukan[tebakanIndex] >= targetIDPencarian
    })

    // 3. Pengecekan Keabsahan Temuan (Karena Search akan memuntahkan angka total panjang len(array) bila ia tak berhasil menemukannya)
    if indexHasilDitemukan < len(daftarIDKartuKependudukan) && daftarIDKartuKependudukan[indexHasilDitemukan] == targetIDPencarian {
        fmt.Printf("BINGO! Target Pelaku ID %d berhasil ditemukan bersembunyi di Laci Indeks Nomor [%d].\n", targetIDPencarian, indexHasilDitemukan)
    } else {
        fmt.Println("Buronan DPO ID tersebut TIDAK TERDAFTAR di dalam tumpukan Kartu Induk.")
    }
}
```

---

## Bagian Lanjutan: Stabilitas Algoritma, Sortir Paralel Konkurensi, dan Pola Pencarian Lanjut Tingkat Produksi

Package `sort` pada pandangan pertama terlihat sederhana, namun algoritma pengurutan (sorting) adalah salah satu operasi paling mahal (*computationally expensive*) di ilmu komputer. Di sistem terdistribusi skala besar, di mana Anda mengurutkan memori RAM yang berisi ratusan ribu baris data log atau ribuan objek produk katalog E-Commerce, penggunaan API dasar package `sort` tanpa pemahaman mendalam tentang stabilitas dan komputasi paralel dapat mengundang *bottleneck* (kemacetan CPU) parah.

### 1. Rahasia Algoritma: *Pattern-Defeating Quicksort* (PDQSort) dan Pengurutan Stabil (`sort.Stable`)

Pada rilis versi-versi terbarunya, Go membuang algoritma *Quicksort* tradisional dari pustaka standarnya dan beralih ke algoritma dewa yang revolusioner bernama **PDQSort**. Kehebatan PDQSort adalah kemampuannya mendeteksi pola data yang *sudah hampir terurut* (atau urutan terbalik) secara instan, dan menghindari jebakan kasus terburuk (*Worst Case* $O(n^2)$) yang sering menimpa Quicksort klasik. Ia secara mulus beralih ke algoritma *Insertion Sort* untuk irisan kecil atau *Heap Sort* bila partisi membandel.

Namun, algoritma PDQSort biasa **TIDAK STABIL (Unstable)**.
Apa artinya algoritma "Tidak Stabil"?

Bayangkan Anda memiliki daftar produk yang sudah diurutkan berdasarkan WAKTU RILIS (dari yang paling baru).
Kemudian, Anda meminta package `sort.Slice` untuk mengurutkannya berdasarkan HARGA (dari termurah).
Jika ada 3 produk yang memiliki harga persis *SAMA* (misalnya Rp 100.000), algoritma *Unstable* (seperti `sort.Slice`) bisa saja mengacak-acak urutan ketiga produk tersebut secara tak menentu! Urutan WAKTU RILIS mereka yang sudah rapi sebelumnya akan rusak berantakan.

**Solusi Industri: Pengurutan Stabil Multilapis (Multi-tier Sorting)**
Jika Anda ingin mempertahankan *urutan relatif asli* elemen-elemen yang bernilai sama, Anda **WAJIB** menggunakan `sort.SliceStable`.
Algoritma ini menggunakan *Merge Sort* yang sedikit lebih rakus memori, tetapi ia menjamin bahwa urutan aslinya tidak akan pernah diusik/diputar balik.

```go
// Simulasi antrean
// type PenumpangPesawat struct {
//    Nama        string
//    StatusMember string // "Platinum", "Gold", "Silver"
//    WaktuCheckIn int    // Jam check in (Misal 1, 2, 3)
// }

// KITA PAKAI sort.SliceStable() untuk menjaga keadilan urutan Check-In!
// sort.SliceStable(antrean, func(i, j int) bool {
//    // Logika prioritas: Gold menang atas Silver
//    if antrean[i].StatusMember == "Gold" && antrean[j].StatusMember == "Silver" {
//        return true
//    }
//    return false // Biarkan posisi stabil jika sama-sama gold atau silver
// })
```

### 2. Tantangan *Interface* Kuno `sort.Interface` (Sebelum Era Go 1.8)

Jika Anda membedah pustaka (*library*) Go pihak ketiga ciptaan lama, Anda akan menemui pendekatan aneh dalam mengurutkan data kustom. Sebelum fungsi `sort.Slice` (yang menggunakan *Reflection* lambat di belakang layar) diciptakan, para sesepuh Go mengurutkan data secara manual dengan membuat *Custom Type* yang wajib menuruti kontrak 3 metode absolut dari `sort.Interface`: `Len()`, `Less(i, j)`, dan `Swap(i, j)`.

Meskipun terlihat sangat bertele-tele dan menyebalkan untuk diketik, **menggunakan `sort.Interface` manual ini TERBUKTI JAUH LEBIH CEPAT secara *Benchmark Performance* CPU** dibandingkan menggunakan fungsi praktis `sort.Slice()`!
Bagi pengembang aplikasi perbankan berjuta transaksi per detik (*High Frequency Trading*), setiap nanodetik bernilai emas, dan menulis struktur ini secara manual adalah harga pantas untuk dipertaruhkan demi kinerja maksimal server Linux.

```go
type KoleksiSaham []int // Misal struct aslinya

// WAJIB IMPLEMENTASI 3 FUNGSI INI AGAR DIAKU SEBAGAI SORT.INTERFACE!
func (k KoleksiSaham) Len() int           { return len(k) }
func (k KoleksiSaham) Swap(i, j int)      { k[i], k[j] = k[j], k[i] } // Manual menukar posisi memory
func (k KoleksiSaham) Less(i, j int) bool { return k[i] > k[j] }

// ... Saat di Fungsi Main, Pemanggilannya Sangat Brutal Cepat:
// sort.Sort(dataTrading) // Eksekusi dengan Kecepatan Dewa!
```

### 3. Pencarian Presisi (Binary Search Lanjutan: .SearchInts & .SearchStrings)

Terkadang kita tidak perlu mencari data di koleksi struktur yang rumit. Jika array kita primitif, pemanggilan `sort.Search()` biasa dengan *callback function* `func(i int) bool` terasa terlalu panjang dan menguras baris kode.
Go menyediakan jalan tol *shortcut* bagi pencarian primitif tersebut:

*   **`sort.SearchInts(slice, angka)`**: Langsung mengembalikan indeks angka tersebut.
*   **`sort.SearchStrings(slice, kata)`**: Sangat berguna mengecek keberadaan username di *Whitelist Array* berukuran 100.000 user yang sudah disortir alfabetis di memori.

**Perangkap Binary Search Klasik (Wajib Tahu!):**
Fungsi `.Search` ini bekerja dengan mentalitas *Insertion Point* (Titik Penyisipan). Apa artinya?
Jika data yang Anda cari **TIDAK ADA** di dalam barisan, fungsi `.Search` di Golang **TIDAK AKAN MENGEMBALIKAN -1 ATAU ERROR!**
Ia justru akan mengembalikan *Indeks bayangan* letak di mana seharusnya elemen asing itu pantas disisipkan kalau ia memang mau dimasukkan ke array agar urutannya tetap rapi!

Inilah sebabnya mengapa Anda wajib memvalidasi kembaliannya.

```go
// KASUS JEBAKAN:
// daftarMemberVVIP := []string{"Andi", "Budi", "Charlie", "Zoro"} // Sudah Urut A-Z
// namaDicari := "Diana"

// Pencarian Kilat Logaritmik Biner
// indexDiana := sort.SearchStrings(daftarMemberVVIP, namaDicari)

// JIKA ANDA TIDAK WASPADA, indexDiana akan bernilai 3! (Karena letak yang pas untuk sisip huruf D adalah sesudah C dan sebelum Z).
// Jika Anda langsung memanggil daftarMemberVVIP[indexDiana], Anda akan salah menyapa Tuan Zoro!

// KODE AMAN STANDAR INDUSTRI VALIDASI MUTLAK:
// if indexDiana < len(daftarMemberVVIP) && daftarMemberVVIP[indexDiana] == namaDicari {
//    fmt.Println("Otentikasi Sukses! Silahkan masuk Nona Diana.")
// }
```
Kelalaian menangani keluaran *insertion point* ini bertanggung jawab atas ribuan *bug* produksi di hari-hari perdana karir programmer pemula Go.

## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
