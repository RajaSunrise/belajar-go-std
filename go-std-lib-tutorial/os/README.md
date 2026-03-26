# Modul: `os`

## Ringkasan
Package `os` (Operating System) menyediakan antarmuka independen-platform untuk berinteraksi secara mulus dengan sistem operasi yang mendasari aplikasi Anda. Fungsionalitas utamanya meliputi pembacaan argumen *command-line*, manipulasi variabel lingkungan (*environment variables*), serta operasi mendasar pada *file* dan direktori.

## Penjelasan Lengkap (Fungsi & Tujuan)
Di dunia nyata, kode Anda tidak pernah berjalan di ruang hampa; ia berjalan di atas Windows, Linux, atau macOS. Terkadang, ia berjalan di dalam kontainer Docker di dalam sebuah mesin virtual Linux. Untuk dapat mengambil informasi tentang tempatnya berjalan, atau untuk berinteraksi dengan penyimpanan permanen (*disk*), aplikasi Anda wajib menggunakan package `os`.

Package `os` dirancang sedemikian rupa agar API-nya bersifat seragam, mirip dengan desain Unix. Ini menjamin abstraksi yang baik: jika Anda menulis kode `os.Mkdir()` untuk membuat folder, kode tersebut akan sukses di-*compile* dan berjalan, tidak peduli apakah program akhirnya dieksekusi di *Server Ubuntu* atau di laptop *Windows* Anda. Di belakang layar, package `os` yang akan menerjemahkannya ke dalam *System Call* yang tepat untuk masing-masing sistem operasi target.

**Tujuan dan Fungsi Utama:**
1.  **Akses Argument dan Variabel:** Mengambil daftar argumen saat *executable* dijalankan dari shell CLI (`os.Args`), dan membaca rahasia/konfigurasi sistem dari lingkungan shell (`os.Getenv`).
2.  **Siklus Hidup Proses:** Memeriksa identitas pengguna yang menjalankan program (`os.Getuid`), nama mesin tempat program berjalan (`os.Hostname`), dan mematikan program secara darurat dengan mengirimkan kode kesalahan kembali ke OS (`os.Exit`).
3.  **Abstraksi File Tingkat Rendah:** Membuat tipe `os.File`, yang merupakan objek sakral pembuka gerbang untuk semua pembacaan dan penulisan teks biner ke dalam *hard disk*.
4.  **Manipulasi Direktori:** Membaca isi folder, mengubah hak akses (*chmod/chown*), membuat *symlink*, dan memindahkan file.

**Mengapa menggunakan `os`?**
Jika Anda sedang menulis skrip otomatisasi server (*DevOps*), mengembangkan aplikasi CLI yang menerima *flag*, atau membuat sistem *backend* *cloud-native* di mana semua konfigurasi *database* harus dibaca dari *Environment Variables* (sesuai kaidah *12-Factor App*), package ini adalah alat primer yang harus Anda kuasai.

---

## Daftar Fungsi Umum dan Cara Penggunaannya Secara Mendalam

### 1. Membaca Variabel Lingkungan (Environment Variables)

Cara modern untuk mengonfigurasi aplikasi *server* (terutama di Kubernetes atau Docker) adalah dengan menyuntikkan *Environment Variables*, bukan menyimpannya langsung di dalam *source code* (sangat berbahaya jika kode bocor!).

*   **`os.Getenv(key)`**: Mengambil nilai berdasarkan kunci. Akan mengembalikan string kosong `""` jika kunci tidak ditemukan.
*   **`os.LookupEnv(key)`**: Sama dengan Getenv, tetapi ini mengembalikan `(string, bool)`. Fungsi ini sangat penting jika Anda ingin membedakan apakah variabel itu tidak diatur, atau apakah variabel itu memang sengaja diatur namun nilainya kosong.
*   **`os.Setenv(key, value)`**: Mengubah nilai variabel (hanya berpengaruh pada proses Go yang sedang berjalan, tidak memengaruhi sistem OS aslinya).

```go
package main
import (
    "fmt"
    "os"
)

func main() {
    // Cara standar: os.Getenv
    dbPass := os.Getenv("DATABASE_PASSWORD")
    if dbPass == "" {
        fmt.Println("Peringatan: Password DB kosong!")
    }

    // Cara aman: os.LookupEnv
    // Berguna saat kita butuh tahu pasti apakah env memang di-set
    port, exists := os.LookupEnv("PORT")
    if !exists {
        fmt.Println("Env PORT tidak ditemukan. Menggunakan default 8080.")
        port = "8080"
    }
    fmt.Println("Server akan berjalan di port:", port)

    // Mengekspor ulang semua env yang sedang aktif ke dalam bentuk slice
    semuaEnv := os.Environ()
    fmt.Printf("Ada %d env variables yang aktif di sistem ini.\n", len(semuaEnv))
}
```

---

### 2. Argumen Command-Line (`os.Args`)

Ketika Anda menjalankan aplikasi dari terminal, misalnya `go run main.go user create --force`, semua kata setelah pemanggilan program tersebut akan ditangkap oleh variabel *slice of string* raksasa bernama `os.Args`.

Perlu diingat: **Indeks ke-0 (`os.Args[0]`) selalu berisi path lengkap menuju *file executable* itu sendiri**. Argumen asli pengguna baru dimulai pada indeks ke-1.

```go
package main
import (
    "fmt"
    "os"
)

func main() {
    // Mengecek apakah user memberikan argumen yang cukup
    if len(os.Args) < 2 {
        fmt.Println("Penggunaan: aplikasiku <perintah>")
        fmt.Println("Contoh: aplikasiku status")
        os.Exit(1)
    }

    perintah := os.Args[1]

    switch perintah {
    case "status":
        fmt.Println("Sistem berjalan dengan normal.")
    case "hapus":
        if len(os.Args) > 2 && os.Args[2] == "--force" {
            fmt.Println("Menghapus data secara paksa!")
        } else {
            fmt.Println("Gagal: Butuh flag --force untuk menghapus.")
        }
    default:
        fmt.Println("Perintah tidak dikenali:", perintah)
    }
}
```

---

### 3. Memanipulasi File (Create, Open, Close)

Membaca dan menulis *file* adalah operasi I/O krusial. Struktur dasar fungsi di sini selalu mengembalikan dua nilai: pointer bertipe `*os.File`, dan `error`. Ingat aturan emas Go: **Setiap kali file dibuka secara manual, ia harus ditutup dengan `defer file.Close()`** agar *file descriptor* tidak terkunci/bocor di memori OS.

#### Menulis ke File Baru
Fungsi `os.Create` akan membuat file baru. Jika nama file tersebut sudah ada di *hard disk*, file lama tersebut akan langsung ditimpa (*truncate*) menjadi kosong. Berhati-hatilah!

```go
package main
import (
    "fmt"
    "os"
)

func main() {
    // 1. Membuka / Membuat file
    f, err := os.Create("laporan.txt")
    if err != nil {
        panic(err)
    }
    // 2. Pastikan file ditutup saat fungsi selesai!
    defer f.Close()

    // 3. Menulis teks (kembali berupa slice of byte)
    teks := []byte("Baris 1: Transaksi Selesai.\n")
    _, err = f.Write(teks)
    if err != nil {
        panic(err)
    }

    // 4. Atau menggunakan cara praktis khusus string
    f.WriteString("Baris 2: Data disimpan.\n")

    // 5. Menyuruh sistem operasi mem-flush antrean memori fisik ke hard disk (Opsional tapi aman)
    f.Sync()

    fmt.Println("Berhasil menulis file.")
}
```

#### Membaca dari File
Ada dua cara mendasar: membaca secara utuh (jika Anda yakin ukuran file kecil), atau membacanya sepotong demi sepotong (chunking). Sejak Go 1.16, package `os` memperkenalkan `os.ReadFile` untuk cara instan.

```go
package main
import (
    "fmt"
    "os"
)

func main() {
    // CARA INSTAN: Membaca seluruh file dalam satu napas
    // Hanya gunakan jika ukuran file < beberapa MB!
    dataBytes, err := os.ReadFile("laporan.txt")
    if err != nil {
        if os.IsNotExist(err) {
            fmt.Println("File laporan.txt belum ada.")
            return
        }
        panic(err)
    }
    fmt.Println("Isi File:\n", string(dataBytes))
}
```

---

### 4. Membuat Direktori dan Mengatur Hak Akses (Permissions)

Di Linux/Unix, setiap file dan direktori memiliki hak akses 3-digit oktal, contohnya `0755` (pemilik bisa membaca/menulis/mengeksekusi, pengguna lain hanya bisa membaca/mengeksekusi). Go merepresentasikan angka oktal ini dengan tipe `os.FileMode`. Penting: Awali angka oktal dengan angka nol `0`.

```go
package main
import (
    "fmt"
    "os"
)

func main() {
    // os.Mkdir membuat satu folder spesifik
    err := os.Mkdir("folder_baru", 0755)
    if err != nil && !os.IsExist(err) { // IsExist mengecek apakah error karena folder sudah ada
        panic(err)
    }

    // os.MkdirAll akan membuat seluruh folder secara berantai (seperti `mkdir -p` di terminal)
    // Jika folder "rahasia" belum ada, ia akan membuatnya juga.
    err = os.MkdirAll("folder_rahasia/data_tahun_ini", 0755)
    if err != nil {
        panic(err)
    }
    fmt.Println("Berhasil membuat hirarki direktori.")
}
```

---

### 5. Memeriksa Informasi File (Stat)

Terkadang sebelum kita melakukan sesuatu dengan file, kita hanya ingin sekadar tahu informasinya: apakah dia sebuah file biasa atau direktori? Berapa ukuran MB-nya? Kapan ia terakhir dimodifikasi? Fungsi `os.Stat` akan mengembalikan objek antarmuka `os.FileInfo` yang berisi seluruh metadata tersebut.

```go
package main
import (
    "fmt"
    "os"
)

func main() {
    info, err := os.Stat("main.go")
    if err != nil {
        if os.IsNotExist(err) {
            fmt.Println("File main.go tidak ditemukan!")
            return
        }
        panic(err)
    }

    fmt.Println("Nama Asli:", info.Name())
    fmt.Println("Ukuran (Bytes):", info.Size())
    fmt.Println("Mode Akses:", info.Mode())
    fmt.Println("Waktu Terakhir Dimodifikasi:", info.ModTime())
    fmt.Println("Apakah dia Direktori (Folder)?", info.IsDir())
}
```

---

## Bagian Lanjutan: Fitur Tingkat Lanjut, Keamanan File, dan Sinyal OS

Manipulasi file dan direktori adalah tulang punggung dari banyak aplikasi tingkat sistem, kontainerisasi, dan DevOps. Di bagian ini, kita mengeksplorasi penggunaan tingkat ahli dari package `os`.

### 1. Memahami Hak Akses (Permissions) Unix / FileMode

Ketika Anda menggunakan `os.Mkdir("folder", 0755)` atau `os.WriteFile("rahasia.txt", data, 0600)`, angka oktal tersebut (diawali dengan 0) bukanlah angka sembarangan. Di Unix/Linux, itu mewakili 3 kelompok hak: Pemilik (User), Grup (Group), dan Orang Lain (Others).
Tiap kelompok dihitung dengan penjumlahan:
*   **4** = Read (Membaca)
*   **2** = Write (Menulis)
*   **1** = eXecute (Menjalankan/Memasuki Folder)

**Contoh Umum di Go:**
*   `0644`: (Read/Write untuk Pemilik, Read-Only untuk semua orang). Ini adalah standar aman untuk hampir semua file teks/log yang Anda buat dengan `os.Create`.
*   `0755`: (Read/Write/Execute untuk Pemilik, Read/Execute untuk semua orang). Standar wajib untuk membuat *Folder* (`os.Mkdir`). Tanpa hak *Execute* pada sebuah direktori, Anda bahkan tidak bisa melihat isinya (`ls`)!
*   `0600`: (Read/Write HANYA untuk Pemilik, orang lain dilarang). Wajib digunakan saat program Go Anda menulis file sensitif seperti Kunci RSA, Token JWT, atau file `.env` ke disk.

```go
// Cara aman menulis kredensial
// token := []byte("TOKEN_JWT_SUPER_RAHASIA")
// err := os.WriteFile("/etc/app/config.key", token, 0600)
// File ini sekarang tidak bisa dibaca oleh hacker yang login sebagai user biasa (bukan root)
```

### 2. Komunikasi Lintas Proses: `os.Pipe`

Di sistem Linux, Anda sering menggunakan perintah `|` (Piping) di terminal, misalnya `ls -l | grep "go"`. Di dalam program Go, Anda bisa menciptakan simulasi pipa yang sama di dalam memori menggunakan `os.Pipe()`.

Fungsi ini mengembalikan sepasang file: satu untuk dibaca (ujung keluar pipa) dan satu untuk ditulis (ujung masuk pipa). Ini sangat vital ketika Anda perlu menghubungkan output dari satu modul ke input modul lainnya secara *real-time* tanpa membuat file fisik sementara di *hard disk*!

```go
// r, w, err := os.Pipe()
// if err != nil {
//    panic(err)
// }

// Menjalankan goroutine untuk memompa data ke dalam Pipa
// go func() {
//    w.Write([]byte("Data mengalir dari ujung satu..."))
//    w.Close() // Wajib ditutup setelah selesai memompa!
// }()

// Goroutine utama menunggu di ujung yang lain dan menyerapnya
// dataKeluar, _ := io.ReadAll(r)
// fmt.Println(string(dataKeluar))
```

### 3. Graceful Shutdown (Menangkap Sinyal Kematian OS)

Ketika Anda menekan `Ctrl+C` di terminal untuk mematikan Server Go Anda, atau ketika Kubernetes memutuskan untuk membunuh Pod *Docker* Anda karena *scaling down*, Sistem Operasi sebenarnya mengirimkan sinyal bernama `SIGINT` atau `SIGTERM`.

Jika aplikasi Anda langsung mati detik itu juga, transaksi database yang sedang berlangsung akan *corrupt*, dan uang pelanggan bisa hilang!
Sebagai *Developer* Senior, Anda diwajibkan menjerat sinyal kematian dari OS ini dan memberikan waktu bagi aplikasi Go untuk "membereskan barang-barangnya" (Graceful Shutdown) sebelum benar-benar keluar via `os.Exit`.

```go
import (
    "os"
    "os/signal"
    "syscall"
    "fmt"
)

func main() {
    // 1. Buat channel penangkap sinyal kematian
    sinyalKematian := make(chan os.Signal, 1)

    // 2. Beritahu package signal untuk membajak sinyal SIGINT (Ctrl+C) dan SIGTERM (Kill Kubernetes)
    // dan membuangnya ke channel kita alih-alih mematikan program secara instan!
    signal.Notify(sinyalKematian, os.Interrupt, syscall.SIGTERM)

    // ... Jalankan Server Web Anda di Background (Goroutine) ...

    // 3. Program Utama Anda akan "Membeku" (Block) di sini, menunggui datangnya malaikat maut (Sinyal OS)
    sig := <-sinyalKematian

    // Jika sampai baris ini, berarti ada orang/sistem yang menyuruh program mati!
    fmt.Printf("\n[PERINGATAN] Menerima sinyal %v! Memulai prosedur Shutdown Darurat!\n", sig)

    // Lakukan pembersihan: Tutup Koneksi Database, Tunggu transaksi selesai, Flush Memori Cache ke Disk
    // ... proses clean-up (misal 5 detik) ...
    fmt.Println("Semua transaksi berhasil diselamatkan. Menutup aplikasi dengan aman. Sampai jumpa!")

    // 4. Barulah Anda matikan diri Anda secara terhormat.
    os.Exit(0)
}
```

### 4. Menghindari Kunci File Bertabrakan (File Locking)

Jika dua goroutine (atau dua program Go terpisah di mesin yang sama) mencoba menulis ke file log (`os.OpenFile`) yang sama persis di milidetik yang sama, data Anda akan saling menimpa, menghasilkan file log yang cacat (Gibberish/Corrupted).

Go tidak memberikan fitur *File Lock* ajaib secara *native* pada package OS (karena desain Windows dan Unix sangat berbeda soal ini).
Solusi Industri: Selalu buka file log Anda di SATU goroutine pusat, dan gunakan *Channel* agar goroutine lain mengirim pesan string log mereka ke goroutine pusat tersebut. Atau, delegasikan pembukaan file pada library logging tangguh seperti `go.uber.org/zap` atau `logrus`. Jangan pernah memanggil `os.Create` ke target yang sama di ratusan tempat paralel! Pahami ini demi keandalan server monolitik Anda yang tak tertembus.

## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
