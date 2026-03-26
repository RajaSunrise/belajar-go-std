# Modul: `math`

## Ringkasan
Package `math` menyediakan sekumpulan fungsi utilitas tinggi untuk memecahkan masalah matematika tingkat lanjut dan perhitungan algoritmik. Ini mencakup operasi geometri dasar, aritmetika tingkat presisi (*floating point*), hingga konstanta numerik alam semesta.

## Penjelasan Lengkap (Fungsi & Tujuan)
Komputer pada dasarnya hanyalah sebuah alat untuk melakukan kalkulasi angka raksasa secara berulang-ulang. Bahasa pemrograman Go secara alamiah sudah menyediakan operator standar seperti penjumlahan (`+`), pengurangan (`-`), perkalian (`*`), pembagian (`/`), dan modulus sisa bagi (`%`). Namun, kalkulasi ilmu komputer modern yang terlibat di dalam simulasi, kecerdasan buatan, visualisasi game, atau aplikasi analitik finansial menuntut kapabilitas operasi matematika yang lebih berat (seperti menghitung Logaritma atau Trigonometri). Inilah tempat di mana package `math` beraksi.

Desain standar package `math` di Go sangat mirip dengan standar kalkulasi *floating point* IEEE-754. Karena itu, hampir seluruh fungsi di dalam package ini menerima (*parameter input*) dan mengembalikan (*return value*) variabel bertipe **`float64`** (angka desimal presisi ganda). Jika Anda memiliki angka bulat (*integer*), Anda harus secara eksplisit mengonversinya ke `float64` terlebih dahulu sebelum mengumpankannya ke fungsi-fungsi ini.

**Tujuan dan Fungsi Utama:**
1.  **Konstanta Presisi Tingkat Tinggi:** Anda tak perlu mencari nilai Pi di internet lalu menuliskannya secara manual (`3.14159...`), karena package ini menyediakan konstanta `math.Pi` dan konstanta lain dengan presisi penuh arsitektur 64-bit yang paling optimal secara mikroskopis.
2.  **Operasi Geometri & Trigonometri:** Memungkinkan perhitungan sudut dan panjang menggunakan sinus (`math.Sin`), kosinus (`math.Cos`), serta fungsi-fungsi yang mengelola perhitungan sudut ke satuan radian.
3.  **Pangkat, Akar, dan Eksponensial:** Menghitung perhitungan aljabar lanjut seperti Logaritma berbasis eksponensial alam (`math.Log`), atau perhitungan bunga bank majemuk menggunakan pemangkatan (`math.Pow`).
4.  **Kontrol Desimal:** Mengatur secara presisi ke mana arah suatu angka desimal pecahan harus melompat (dibulatkan) menjadi bilangan rasional utuh terdekat (Fungsi `Round`, `Ceil`, dan `Floor`).

**Mengapa menggunakan `math`?**
Jika Anda sedang mengembangkan sebuah game dua dimensi (2D) di Go (misal dengan ebitengine) dan perlu menggerakkan karakter pemain dalam arah diagonal berdasarkan *joystick input*, Anda wajib menghitung trigonometrinya dengan `math`. Jika Anda melakukan analisis data yang mencari selisih jarak standar dari nilai absolut (*Absolute Value*), package `math` akan menjaga Anda dari kebocoran memori (Memory Corruption) akibat kesalahan aritmetika yang rumit.

---

## Daftar Fungsi Umum dan Cara Penggunaannya Secara Mendalam

### 1. Pembulatan Pecahan Desimal (Rounding)

Ini adalah masalah klasik: Anda memiliki nilai perhitungan transaksi `Rp 15000.5`, apakah aplikasi harus menyimpannya sebagai `15000` atau `15001`? Go memberikan kontrol penuh melalui 3 metode pembulatan.

*   **`math.Round(x)`**: Membulatkan `x` ke bilangan bulat rasional terdekat. Di tengah-tengah pertengahan eksak (misal `0.5`), nilainya akan menjauhkan nilai dari titik 0 (yakni membulat ke atas untuk bilangan positif).
*   **`math.Ceil(x)`**: (Ceiling / Langit-langit). Selalu membulatkan angka desimal ke arah nilai yang **lebih besar** atau lebih atas.
*   **`math.Floor(x)`**: (Floor / Lantai). Selalu membulatkan angka desimal ke arah nilai yang **lebih kecil** atau lebih bawah.
*   **`math.Trunc(x)`**: (Truncate). Menghapus secara paksa bagian desimal di belakang titik, mengabaikan segala hal, dan hanya menyisakan nilai integer dasar.

```go
package main
import (
    "fmt"
    "math"
)

func main() {
    nilai := 12.34

    // Demonstrasi Dasar
    fmt.Printf("Angka asli: %f\n", nilai)
    fmt.Printf("Round (Terdekat): %.0f\n", math.Round(nilai)) // Output: 12
    fmt.Printf("Ceil  (Selalu Naik): %.0f\n", math.Ceil(nilai))  // Output: 13
    fmt.Printf("Floor (Selalu Turun): %.0f\n", math.Floor(nilai)) // Output: 12

    // Perhatikan ketika di angka kritis x.5
    nilaiKritis := 12.50
    fmt.Printf("\nAngka kritis: %f\n", nilaiKritis)
    fmt.Printf("Round: %.0f\n", math.Round(nilaiKritis)) // Otomatis naik jadi 13

    // Truncate vs Floor di nilai Negatif
    negatif := -12.75
    fmt.Printf("\nAngka negatif: %f\n", negatif)
    fmt.Printf("Floor (-12.75 turun berarti mundur ke -13): %.0f\n", math.Floor(negatif))
    fmt.Printf("Trunc (hanya membuang pecahan .75): %.0f\n", math.Trunc(negatif)) // Output: -12
}
```

---

### 2. Akar Kuadrat, Pemangkatan, dan Absolut

Operasi-operasi ini adalah inti dari aljabar.

*   **`math.Sqrt(x)`**: (Square Root). Mendapatkan nilai akar kuadrat dari bilangan. Mengembalikan `NaN` (Not a Number) jika variabel `x` adalah negatif.
*   **`math.Cbrt(x)`**: (Cube Root). Mendapatkan nilai akar pangkat tiga.
*   **`math.Pow(x, y)`**: (Power). Memangkatkan bilangan `x` sebanyak pangkat `y` (jadi $x^y$). Ini juga bisa digunakan untuk menghitung akar kuadrat lanjutan dengan trik matematika (menggunakan `math.Pow(x, 1.0/3.0)` setara dengan `math.Cbrt`).
*   **`math.Abs(x)`**: (Absolute). Menghilangkan tanda hubung negatif (`-`). Mengubah semua nilai menjadi non-negatif. Berguna untuk membandingkan jarak tanpa peduli apakah arahnya ke kiri atau ke kanan dari grafik nol.

```go
package main
import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("=== Akar dan Pangkat ===")

    luasTanah := 144.0
    sisiTanah := math.Sqrt(luasTanah) // Jika persegi 144, berarti sisinya 12
    fmt.Printf("Luas Tanah: %.0f, Panjang Sisinya: %.0f\n", luasTanah, sisiTanah)

    angkaDasar := 2.0
    pangkat := 8.0
    hasilPangkat := math.Pow(angkaDasar, pangkat) // 2 pangkat 8
    fmt.Printf("%.0f pangkat %.0f adalah: %.0f\n", angkaDasar, pangkat, hasilPangkat)

    fmt.Println("\n=== Nilai Absolut ===")
    suhuPagi := 25.0
    suhuMalam := 18.0

    // Jika kita mengurangi 18 - 25 = -7. Tapi kita hanya peduli pada "Selisih Jarak Suhu" saja, bukan nilainya.
    selisihSuhu := math.Abs(suhuMalam - suhuPagi)
    fmt.Printf("Perbedaan suhu ekstrem adalah: %.0f derajat.\n", selisihSuhu)
}
```

---

### 3. Evaluasi Minimum dan Maksimum

Kadang kita hanya perlu mengambil nilai batas, seperti membatasi bahwa diskon maksimal aplikasi adalah 100 ribu, tidak boleh lebih besar dari angka tersebut walau rumus aslinya mencapai 200 ribu.

*   **`math.Max(x, y)`**: Membandingkan 2 bilangan *float* dan mengembalikan yang ukurannya lebih besar.
*   **`math.Min(x, y)`**: Membandingkan 2 bilangan *float* dan mengembalikan yang ukurannya lebih kecil.

*(Catatan: Sejak Go versi terbaru (Go 1.21), fungsi `max()` dan `min()` telah ditambahkan ke dalam bahasa bawaan Go (*built-in*), sehingga bisa digunakan pada sembarang tipe, tidak hanya *float64*. Namun secara arsitektural dan historis, fungsi di package math ini tetap ada untuk mendukung tipe floating point khusus).*

```go
package main
import (
    "fmt"
    "math"
)

func main() {
    saldo := 50000.0
    hargaBarang := 80000.0

    // Contoh sederhana penggunaan min/max:
    pengeluaranTerbesar := math.Max(saldo, hargaBarang)
    fmt.Println("Angka terbesar di antara saldo dan harga adalah:", pengeluaranTerbesar)
}
```

---

### 4. Konstanta Matematika

Nilai-nilai konstanta di-hardcode ke dalam compiler sehingga eksekusinya kilat, tidak menuntut CPU menghitung dari nol. Semuanya bertipe `float64`.

*   **`math.Pi`**: Konstanta Pi (3.141592...) untuk perhitungan keliling/luas lingkaran.
*   **`math.E`**: Konstanta bilangan Euler (2.718281...) untuk penghitungan logaritma natural (Ln).
*   **`math.MaxFloat64` / `math.MaxFloat32`**: Nilai numerik terbesar yang **mungkin** ditampung oleh memori fisik 64-bit pada mesin Anda. Berguna untuk menginisialisasi perulangan pencarian nilai paling kecil. (Dengan memulai perbandingan dari `variabelTerkecil = math.MaxFloat64`, lalu kita meloop list tersebut).

```go
package main
import (
    "fmt"
    "math"
)

func main() {
    jariJari := 10.0

    // Rumus Keliling Lingkaran: 2 * pi * r
    keliling := 2 * math.Pi * jariJari

    // Rumus Luas Lingkaran: pi * r * r
    luas := math.Pi * math.Pow(jariJari, 2)

    fmt.Printf("Lingkaran berdiameter %.0f cm memiliki:\n", jariJari*2)
    fmt.Printf("Keliling: %f cm\n", keliling)
    fmt.Printf("Luas Area: %f cm persegi\n", luas)

    fmt.Println("\nBatas mentok memory MaxFloat64:", math.MaxFloat64)
}
```

---

## Bagian Lanjutan: Komputasi Keilmuan, Performa, dan Angka Ajaib (NaN/Inf)

Package `math` Go bukanlah sekadar kumpulan fungsi aritmetika sekolah dasar; ia didesain dengan tingkat keakuratan ilmiah tinggi yang meniru standar operasi *Hardware Floating-Point* yang disepakati secara internasional (IEEE-754). Di bagian ini, kita mengupas tuntas perilaku anomali matematika yang sering mengecoh insinyur data (*Data Engineers*).

### 1. Menghadapi Kiamat Perhitungan: NaN (Not-a-Number) dan Infinity

Berbeda dengan angka integer (`int`) di Go yang akan menyebabkan program PANIK (Crash) seketika jika Anda membaginya dengan nol (`5 / 0`), tipe `float64` di dalam `math` didesain untuk "Tahan Banting". Bukannya membuat sistem meledak, pembagian atau perhitungan tidak valid akan mengembalikan nilai konstanta khusus:

*   **`Infinity` (Tak Terhingga)**: Terjadi jika Anda membagi angka float positif dengan `0.0`.
*   **`-Infinity` (Negatif Tak Terhingga)**: Terjadi jika Anda membagi angka float negatif dengan `0.0`.
*   **`NaN` (Not a Number / Bukan Angka)**: Terjadi pada operasi matematika yang mustahil, misalnya akar kuadrat dari angka negatif (`math.Sqrt(-1)`), atau nol dibagi nol (`0.0 / 0.0`).

Bahaya fatalnya adalah: **`NaN` menular seperti virus!** Jika ada satu saja operasi `NaN` di tengah-tengah rentetan 50 rumus kompleks Anda, hasil akhirnya pasti akan menjadi `NaN` juga, dan masuk merusak database PostgreSQL Anda.

Anda HANYA BISA mendeteksinya menggunakan fungsi utilitas khusus (jangan gunakan operator `==`):

```go
// hasilAneh := math.Sqrt(-25.0)

// CARA DETEKSI YANG BENAR:
// if math.IsNaN(hasilAneh) {
//    fmt.Println("Peringatan: Rumus menghasilkan nilai mustahil (Bukan Angka).")
// }

// Mendeteksi Tak Terhingga (Inf):
// if math.IsInf(hasilAneh, 0) { // Angka 0 berarti cek baik Inf positif maupun negatif
//    fmt.Println("Peringatan: Terjadi pembagian dengan nol! (Tak Terhingga)")
// }
```

### 2. Akar Masalah Presisi Pecahan (Floating Point Precision Loss)

Komputer memproses segala hal dalam basis biner (0 dan 1). Manusia berpikir dalam basis desimal (10). Fraksi sederhana seperti `0.1` di dunia manusia sayangnya **tidak bisa direpresentasikan secara sempurna** di dunia biner. Ia menjadi desimal berulang tak terhingga yang terpotong kapasitas memori.

Ini berakibat fatal pada logika perbandingan finansial:

```go
// uangSistem := 0.1 + 0.2
// uangHarapan := 0.3

// Di bahasa Go (dan semua bahasa pemrograman lain), INI AKAN FALSE!
// Karena 0.1 + 0.2 di memori komputer bernilai 0.30000000000000004
// if uangSistem == uangHarapan { ... }
```

**Solusi Toleransi Epsilon:**
Dalam package `math`, Anda tidak boleh menggunakan operator `==` untuk tipe Float. Anda harus mengecek apakah jarak perbedaan absolut antara kedua angka itu lebih kecil dari batas ambang toleransi mikroskopis (sering disebut *Epsilon*).

```go
const EpsilonToleransi = 1e-9 // Toleransi hingga 9 angka di belakang koma

func periksaSamaDengan(a, b float64) bool {
    // Apakah jarak selisih keduanya lebih tipis dari debu Epsilon?
    return math.Abs(a - b) <= EpsilonToleransi
}
```

### 3. Ekstremitas Kecepatan Hardware (Assembly Built-ins)

Beberapa fungsi di dalam package `math` Go (seperti `math.Sqrt`, `math.Exp`, `math.Log`) sejatinya **tidak murni ditulis dalam kode Go**.
Ketika di-compile ke arsitektur AMD64 atau ARM64 (mesin cloud modern), *compiler* Go dengan cerdas akan membuang kodenya dan langsung menggantinya dengan satu baris instruksi bahasa *Assembly* (Mesin) tingkat perangkat keras CPU yang bersangkutan.

Artinya, memanggil `math.Sqrt()` di Go nyaris secepat memanggilnya di bahasa C tingkat bawah, memakan waktu hanya sekitar 1 hingga 2 nanodetik per pemanggilan! Jangan pernah mencoba menulis fungsi pengakar-kuadrat atau logaritma sendiri menggunakan trik iterasi, karena kecepatan implementasi fungsi `math` bawaan ini sudah menyentuh batas mutlak kecepatan cahaya transistor fisika CPU komputer Anda. Mengapa menginventarisasi ulang roda yang sudah sempurna?

## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```



### 4. Trik Kalkulasi Float Menghindari Pembagian 0 (Divide-By-Zero Panic)

Sebuah jebakan mematikan bagi sistem Microservice Go Anda adalah jika sebuah request web dari *user* memiliki parameter yang bila diolah secara matematis menghasilkan angka Pembagi Nol (`0`). Pada tipe `int`, ini langsung akan memicu PANIC mematikan server seketika.

Namun, tipe `float64` bawaan package `math` tidak akan pernah panik. Karena ia mengadopsi standar IEEE-754, sebuah pembagian `X / 0.0` akan secara ajaib memuntahkan nilai konstanta khusus `+Inf` (Positif Tak Terhingga) atau `-Inf` (Negatif Tak Terhingga).

Masalahnya: Jika nilai `Inf` ini disusupkan masuk ke dalam perintah `json.Marshal` atau dikirim langsung ke kueri `database/sql`, library-library tersebut BISA MENOLAK dan PANIC (karena format JSON RFC resmi tidak mengenali wujud nilai Tak Terhingga Infinity).

```go
// simulasi
// uangKlien := 150000.0
// angkaPembagi := 0.0

// Ini tidak error di go, dia jadi "Inf"
// hasilTagihan := uangKlien / angkaPembagi

// JANGAN LAKUKAN INI:
// bodiJson, err := json.Marshal(map[string]any{"hasil": hasilTagihan})
// JIKA DIJALANKAN: json: unsupported value: +Inf

// SELALU PATUHI PRAKTIK INI SEBELUM MENGIRIM!
// if math.IsInf(hasilTagihan, 0) {
//      fmt.Println("Error: Transaksi memicu Pembagian Tak Terhingga.")
// }
```

Pastikan selalu menyaring input Anda dengan logika verifikasi matematis murni `math.IsInf()` yang dibahas di atas sebelum melempar hasil hitungan tersebut menyeberangi modul lain ke perbatasan sistem yang rapuh (Frontend Web/DB SQL).
