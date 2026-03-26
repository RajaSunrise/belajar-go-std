# Modul: `fmt`

## Ringkasan
Package `fmt` (Format) mengimplementasikan operasi Input/Output (I/O) dasar yang berformat. Package ini sangat krusial dan merupakan salah satu package yang paling pertama dipelajari oleh programmer Go, karena menyediakan cara standar untuk mencetak teks ke layar dan memformat string.

## Penjelasan Lengkap (Fungsi & Tujuan)
Di dalam bahasa pemrograman C, Anda mungkin mengenal fungsi `printf` dan `scanf`. Package `fmt` di Go adalah evolusi dari konsep tersebut, dirancang agar lebih aman terhadap tipe data (type-safe), lebih mudah digunakan, dan terintegrasi penuh dengan sistem *interface* Go (khususnya `io.Writer` dan `io.Reader`).

**Tujuan dan Fungsi Utama:**
1.  **Mencetak Output (Printing):** Menampilkan data ke standar output (`os.Stdout`), yang biasanya adalah layar konsol terminal Anda. Ini adalah cara utama untuk memberikan informasi kepada pengguna CLI (Command Line Interface) atau sekadar menampilkan log sederhana saat proses *debugging*.
2.  **Pemformatan String (String Formatting):** Menyusun string kompleks dari berbagai tipe data (angka, boolean, struktur data) menggunakan "Verbs" atau penanda format tanpa harus mengubah semuanya menjadi string secara manual menggunakan konversi tipe data.
3.  **Membaca Input (Scanning):** Mengambil input teks yang diketikkan pengguna melalui keyboard (atau dari *pipeline* teks lainnya) dan memasukkannya langsung ke dalam variabel Go yang sesuai dengan tipe datanya.
4.  **Menulis ke Stream (F-Printing):** Keluarga fungsi yang diawali dengan huruf `F` (seperti `fmt.Fprintf`) memungkinkan Anda untuk menulis teks yang sudah diformat langsung ke dalam objek apa pun yang mengimplementasikan `io.Writer`, seperti file di *hard disk*, koneksi jaringan HTTP, atau *buffer* memori.

**Mengapa menggunakan `fmt`?**
Kapan pun Anda perlu membuat teks yang dinamis (mengandung nilai dari variabel yang bisa berubah-ubah), atau kapan pun Anda ingin melihat nilai sebenarnya dari sebuah *struct* bersarang tanpa harus membuat fungsi `ToString()` kustom, package `fmt` menyediakan solusinya dalam satu baris kode.

---

## Daftar Fungsi Umum dan Cara Penggunaannya Secara Mendalam

### 1. Keluarga Fungsi Print (`Print`, `Println`, `Printf`)

Ini adalah fungsi yang paling sering dipanggil. Semuanya menulis ke `os.Stdout`.

#### a. `fmt.Println()`
Mencetak teks dan otomatis menambahkan karakter *newline* (baris baru) di akhir. Spasi otomatis ditambahkan di antara argumen.
```go
package main
import "fmt"

func main() {
    nama := "Alice"
    umur := 25
    // Otomatis memberi spasi antara nama dan umur
    fmt.Println("Halo,", nama, "umur Anda", umur)
    // Output: Halo, Alice umur Anda 25
}
```

#### b. `fmt.Print()`
Mirip dengan `Println`, namun tidak menambahkan baris baru di akhir.
```go
package main
import "fmt"

func main() {
    fmt.Print("Loading")
    for i := 0; i < 3; i++ {
        fmt.Print(".") // Output akan bersambung di baris yang sama
    }
    // Output: Loading...
}
```

#### c. `fmt.Printf()` (Print Formatted)
Fungsi paling kuat untuk mencetak. Membutuhkan string format khusus yang berisi *Verbs* (seperti `%s`, `%d`), yang kemudian akan digantikan oleh argumen yang diberikan.
```go
package main
import "fmt"

func main() {
    suhu := 36.5
    // %f digunakan untuk float. %.2f artinya 2 angka di belakang koma
    fmt.Printf("Suhu tubuh rata-rata manusia adalah %.2f derajat Celcius.\n", suhu)
    // Output: Suhu tubuh rata-rata manusia adalah 36.50 derajat Celcius.
}
```

---

### 2. Memahami *Verbs* (Penanda Format)

Kunci dari `fmt` adalah menguasai *Verbs*. Berikut adalah daftar verb yang paling penting dan wajib diketahui:

#### General (Umum)
*   **`%v`**: (Value) Menampilkan nilai dalam format default. Sangat berguna jika Anda malas memikirkan tipe datanya (apakah int, string, atau bool).
*   **`%+v`**: Jika digunakan pada *Struct*, ini akan mencetak nilai sekaligus nama *field*-nya (sangat berguna untuk debugging).
*   **`%#v`**: Mencetak representasi sintaks Go (seolah-olah Anda menulis kode untuk membuat objek tersebut).
*   **`%T`**: (Type) Hanya mencetak *tipe data* dari variabel, bukan nilainya.

```go
type Point struct {
    X, Y int
}
p := Point{10, 20}

fmt.Printf("%v\n", p)   // Output: {10 20}
fmt.Printf("%+v\n", p)  // Output: {X:10 Y:20}
fmt.Printf("%#v\n", p)  // Output: main.Point{X:10, Y:20}
fmt.Printf("%T\n", p)   // Output: main.Point
```

#### Integer (Bilangan Bulat)
*   **`%d`**: (Decimal) Basis 10.
*   **`%b`**: (Binary) Basis 2 (contoh: 1010).
*   **`%o`**: (Octal) Basis 8.
*   **`%x` / `%X`**: (Hexadecimal) Basis 16, menggunakan a-f atau A-F.

#### Floating Point (Desimal)
*   **`%f`**: Menampilkan bilangan desimal standar tanpa eksponen (contoh: 123.456).
*   **`%e` / `%E`**: Menampilkan bilangan dalam notasi ilmiah (contoh: 1.234560e+02).

#### String & Byte Slice
*   **`%s`**: Menampilkan teks murni.
*   **`%q`**: (Quoted) Menampilkan string yang diapit tanda kutip ganda, dengan karakter khusus di-*escape* (contoh: `"Hello \n World"`).

---

### 3. Keluarga Fungsi Sprint (`Sprint`, `Sprintln`, `Sprintf`)

Fungsi dengan awalan `S` ini cara kerjanya 100% sama dengan keluarga `Print`, namun alih-alih mencetaknya ke layar konsol, mereka **merakit dan mengembalikan hasilnya dalam bentuk sebuah variabel tipe `string`**.

Ini sangat penting untuk membangun *query* database, memformat respons API, atau menyusun template *email* sederhana.

```go
package main
import "fmt"

func main() {
    nama := "Bob"
    saldo := 150000.75

    // Menyusun string laporan ke dalam variabel baru
    laporan := fmt.Sprintf("Pelanggan %s memiliki saldo sebesar Rp %.2f.", nama, saldo)

    // String ini kemudian bisa disimpan ke database, dikirim via HTTP, dll.
    fmt.Println("Teks yang disimpan:", laporan)
}
```

---

### 4. Keluarga Fungsi Fprint (`Fprint`, `Fprintln`, `Fprintf`)

Huruf `F` di sini singkatan dari *File*, namun secara teknis ini mengarah pada *Interface* `io.Writer`. Ini berarti fungsi ini tidak membatasi dirinya pada *file* di hardisk saja, tetapi objek apa pun yang bisa "ditulisi" di Go. Argumen pertamanya selalu objek target `io.Writer`.

#### Contoh Menulis ke File Disk:
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("log.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Menulis teks yang diformat langsung ke dalam file "log.txt"
    fmt.Fprintf(file, "Sistem gagal melakukan booting pada port %d.\n", 8080)
}
```

#### Contoh Menulis ke HTTP Response (Web Server):
Ini adalah teknik standar jika Anda membuat Web Server di Go tanpa bantuan *framework*.
```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // w adalah sebuah io.Writer! Jadi kita bisa langsung menggunakan Fprintf kepadanya.
    nama := r.URL.Query().Get("nama")
    if nama == "" {
        nama = "Tamu"
    }
    fmt.Fprintf(w, "<h1>Selamat datang di Website kami, %s!</h1>", nama)
}
```

---

### 5. Membaca Input Pengguna (Scanning)

Meski jarang digunakan di aplikasi modern berbasis Web/API, fungsi *scanning* sangat berguna saat Anda menulis *tool* CLI kecil yang membutuhkan jawaban langsung dari pengguna di terminal.

Fungsi `fmt.Scan` (dan variannya `Scanln`, `Scanf`) membaca teks dari `os.Stdin` dan mencoba mengubahnya menjadi tipe data variabel yang *pointer*-nya kita berikan.

#### Contoh Penggunaan `Scanln`:
`Scanln` berhenti membaca setelah menemukan karakter baris baru (saat pengguna menekan tombol Enter).
```go
package main
import "fmt"

func main() {
    var nama string
    var umur int

    fmt.Print("Masukkan nama dan umur Anda (pisahkan dengan spasi): ")
    // Penting: Kita harus memberikan pointer (&) agar Scanln bisa mengubah nilai aslinya
    // _, err := fmt.Scanln(&nama, &umur)
}
```

---

### 6. Implementasi Interface `fmt.Stringer` (Custom Formatter)

Ini adalah trik lanjutan. Bagaimana jika Anda memiliki *struct* kompleks, dan Anda ingin Go mencetaknya dengan rapi setiap kali Anda memberikan struct tersebut ke `fmt.Println`, tanpa harus memanggil `fmt.Printf` setiap saat?

Anda bisa melakukannya dengan membuat tipe data Anda mengimplementasikan antarmuka `fmt.Stringer`. Caranya hanyalah dengan menambahkan sebuah method bernama `String()` yang mengembalikan `string`.

```go
package main
import "fmt"

type IPAddr [4]byte

// Kita membuat method khusus String() untuk tipe IPAddr.
// Dengan begini, IPAddr sekarang mengimplementasikan fmt.Stringer.
func (ip IPAddr) String() string {
    return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
    // Membuat sebuah map dengan nilai berupa IPAddr kustom kita
    hosts := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }

    // Saat dicetak, fmt.Println akan mendeteksi method String()
    // dan menggunakannya secara otomatis!
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}
```

---

## Bagian Lanjutan: Fitur Tingkat Lanjut dan Kinerja Ekstrem

Meskipun terlihat sederhana, package `fmt` menyembunyikan kekuatan dan pertimbangan kinerja yang luar biasa. Bagian ini mengeksplorasi penggunaan `fmt` yang lebih dalam dan sering ditemui di lingkungan produksi.

### Perbedaan Krusial `%v`, `%+v`, dan `%#v` dalam Analisis Data

Ketika mencetak struct untuk keperluan log, Anda akan sering berhadapan dengan tiga varian verb nilai (value verb) ini:

*   **`%v`**: Hanya mencetak nilai mentah dari field-field struct. Output: `{John Doe 30 true}`
*   **`%+v`**: Sangat direkomendasikan untuk logging error! Ini akan mencetak nama field bersama nilainya. Output: `{Nama:John Doe Umur:30 IsAdmin:true}`
*   **`%#v`**: Format ini adalah representasi bahasa Go (Go-syntax representation). Ini mencetak sintaks yang valid seolah-olah Anda bisa men-copy-paste output tersebut kembali ke dalam file `.go`. Menampilkan tipe struct lengkap, dan string diapit tanda kutip. Output: `main.User{Nama:"John Doe", Umur:30, IsAdmin:true}`. Sangat bermanfaat untuk *Unit Testing* saat membandingkan *expected* vs *actual* structs.

### Menyembunyikan Data Sensitif dengan `fmt.Formatter`

Bagaimana jika Anda memiliki struct `User` yang memiliki field `Password`, dan Anda takut tidak sengaja mencetaknya dengan `fmt.Printf("%+v", user)` sehingga passwordnya bocor ke sistem log (Kibana/Datadog)?
Anda bisa memutus rantai tersebut dengan mengimplementasikan interface `fmt.Formatter` atau cukup `fmt.Stringer`.

```go
type Pengguna struct {
    Username string
    Password string // RAHASIA!
}

// Dengan mengimplementasikan method String(), kita mengendalikan persis
// bagaimana objek ini "terlihat" saat dilempar ke fungsi fmt mana pun.
func (p Pengguna) String() string {
    return fmt.Sprintf("Pengguna{Username: %s, Password: [DISEMBUNYIKAN]}", p.Username)
}
```
Setiap kali rekan satu tim Anda mencoba melakukan `fmt.Println(penggunaBaru)`, sistem Go akan secara otomatis memanggil metode `.String()` ini, memastikan `Password` tidak pernah tercetak secara tidak sengaja.

### Implikasi Kinerja `fmt.Sprintf` vs `strings.Builder`

Satu hal yang sering tidak disadari oleh pemula: fungsi-fungsi di dalam package `fmt` menggunakan fitur *Reflection* (paket `reflect`) di belakang layar secara ekstensif untuk menganalisis tipe data apa pun (`any` atau `interface{}`) yang Anda umpankan kepadanya saat runtime.
Proses *reflection* ini memiliki biaya (overhead) performa.

Jika Anda perlu menggabungkan 100.000 string di dalam sebuah *for-loop*, **JANGAN PERNAH** menggunakan:
`hasil = fmt.Sprintf("%s,%s", hasil, dataBaru)`

Sebagai gantinya, gunakan `strings.Builder` yang dirancang khusus untuk alokasi memori efisien tanpa *reflection*:
```go
// var builder strings.Builder
// for i := 0; i < 100000; i++ {
//    builder.WriteString(dataBaru)
//    builder.WriteString(",")
// }
// hasil := builder.String()
```
Gunakan `fmt.Sprintf` HANYA untuk merakit string log, pesan error, atau template teks pendek di mana kemudahan pembacaan kode jauh lebih penting daripada kerugian beberapa nanodetik waktu CPU.

### Keamanan: Pencetakan Format Secara Dinamis

Jangan pernah memasukkan input pengguna langsung sebagai parameter pertama `fmt.Printf`!
```go
userInput := "%s%s%s%s"
// SANGAT BURUK DAN BERBAHAYA!
// fmt.Printf(userInput)

// CARA AMAN:
// Selalu gunakan verb %s eksplisit sebagai parameter pertama
// fmt.Printf("%s", userInput)
```
Jika Anda membiarkan string format dikontrol oleh pengguna luar, mereka berpotensi memaksa program untuk membaca memori internal atau membuat aplikasi panik (panic) karena mengharapkan argumen yang tidak ada.

### Memindai (Scanning) dengan Pola Khusus: `fmt.Sscanf`

Walaupun fungsi `Scan` jarang digunakan untuk aplikasi server modern (REST API/gRPC), `fmt.Sscanf` (String Scan Formatted) adalah utilitas tersembunyi yang sangat kuat untuk membedah string mentah tanpa perlu repot menggunakan Regexp!

```go
dataSuhuSensor := "Suhu Ruang Server: 24.5 Celcius Kelembapan: 60%"

var suhu float64
var kelembapan int

// Kita mendefinisikan "Cetakan Pola" untuk mengekstrak hanya bagian angkanya saja!
_, err := fmt.Sscanf(dataSuhuSensor, "Suhu Ruang Server: %f Celcius Kelembapan: %d%%", &suhu, &kelembapan)

if err == nil {
    // Berhasil diekstrak! suhu = 24.5, kelembapan = 60
}
```
Ini jauh lebih cepat dan lebih mudah dibaca daripada menulis pola *Regular Expression* yang rumit jika format stringnya sudah diketahui dengan pasti dan sangat kaku!

Dengan memahami teknik-teknik canggih ini, penguasaan Anda terhadap I/O format di Go telah mencapai standar industri profesional tingkat mahir. Operasi *Formatting* ini adalah landasan paling absolut ketika merancang instrumen *Logger* performa gila-gilaan pada arsitektur perangkat lunak monolitik berskala ratusan juta Request.

## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
