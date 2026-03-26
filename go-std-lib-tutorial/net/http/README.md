# Modul: `net/http`

## Ringkasan
Package `net/http` menyediakan implementasi klien dan server HTTP (*Hypertext Transfer Protocol*) tingkat-produksi (*production-grade*) terpasang bawaan. Ia merupakan fondasi dari nyaris keseluruhan ekosistem kerangka kerja (*Framework*) web di Go dan memungkinkan pengembang membuat aplikasi mikroservis yang mampu melayani jutaan koneksi serentak tanpa bergantung pada perangkat lunak (*software*) peladen web eksternal apa pun (seperti Nginx atau Apache).

## Penjelasan Lengkap (Fungsi & Tujuan)
Sebagian besar bahasa pemrograman mensyaratkan Anda memasang kerangka kerja eksternal besar yang kompleks untuk sekadar membuka port peladen web guna melayani lalu lintas HTTP. Misalnya *Express.js* di ekosistem Node, *Django/Flask* di Python, atau *Tomcat* di ekosistem Java. Secara bertolak belakang, para pembuat Go (yang banyak didominasi insinyur Google pembangun infrastruktur *cloud*) telah merancang *Standard Library* HTTP di Go agar sangat andal, optimal, dan aman. Anda secara harfiah dapat me-raik (*compile*) program Go biner tunggal, lalu membiarkannya terekspos langsung ke jaringan internet ganas, dan peladen `net/http` itu akan menyerap hantaman jutaan koneksi HTTP sembari melindungi diri dari serangan DDoS dangkal berkat penjadwal goroutine berbiaya sangat rendah yang menangani masing-masing *request* dengan *thread* paralelnya sendiri-sendiri secara otomatis.

**Tujuan dan Fungsi Utama:**
1.  **Pelayan Aplikasi (HTTP Server):** Menjalankan simpul server yang "mendengarkan" koneksi TCP pada port jaringan yang ditentukan, mengurai protokol teks HTTP primitif menjadi paket rapih (`http.Request`), lalu melemparkannya kepada Fungsi Pengendali Anda (*Handler/Controller*), di mana tugas Anda cukup merakit balasan (`http.ResponseWriter`).
2.  **Pemetaan Rute Lanjut (ServeMux):** Menyediakan sistem perutean terpusat bawaan (`http.ServeMux` - *multiplexer*). Mulai rilis bersejarah versi **Go 1.22**, Mux secara drastis berevolusi untuk membaca struktur deklarasi URL REST berbasiskan kata Metode dan pola dinamis pendefinisian variabel rute (*Method & Wildcard Path Variables*).
3.  **Klien Agen (HTTP Client):** Melakukan eksekusi pengambilan sinyal keluar. Jika Anda memerlukan integrasi API (misalnya Anda server A, perlu menelepon API Server B pihak ketiga Stripe atau Mailchimp), modul klien `http.Get/Do` mengelola secara penuh penanganan rekoneksi, negosiasi keamanan soket TLS (HTTPS), negosiasi proksi transparan, dan pengelolaan penampungan (*connection pooling*).
4.  **Penanganan File Statis Berkinerja Tinggi:** Menyajikan serangkaian data gambar, direktori HTML/CSS statis langsung ke klien peramban web sekelas Nginx menggunakan fasilitas bawaan `http.FileServer` atau `http.ServeContent` yang paham negosiasi parsial respons HTTP.

**Mengapa menggunakan `net/http`?**
Jika Anda tidak mengerti fondasi objek standar seperti `http.ResponseWriter`, `http.Request`, dan kontrak `http.HandlerFunc`, Anda akan buta arah tatkala harus menggunakan ekosistem *Framework* tingkat tinggi di luar sana seperti *Gin*, *Fiber*, atau *Echo*. Karena nyatanya, hampir seluruh kerangka arsitektural Go yang memfasilitasi komunikasi antarsistem di era Internet murni dibangun berlapis-lapis di atas bahu perkasa `net/http`.

---

## Daftar Fungsi Umum dan Cara Penggunaannya Secara Mendalam

### 1. Merakit HTTP Server Paling Esensial
Cara paling purba namun solid untuk mendirikan *Backend Web Server*. Selalu mulai dengan mendaftarkan alamat rute ke fungsi pengendali (`HandleFunc`), lalu hidupkan Server yang "abadi" mengunci jalannya proses program (`ListenAndServe`).

*   **`http.ResponseWriter (w)`**: Antarmuka magis untuk menulis ke selang transmisi jaringan klien. Apa pun `Write()` yang disemburkan kepadanya akan dienkode lalu dipancarkan balik melintasi jaringan internet TCP ke layar penjelajah Firefox/Chrome pihak klien. Anda juga menanam kode status (200 OK, 404, 500) dan *Header Config* (tipe konten JSON) lewat *interface* ini.
*   **`*http.Request (r)`**: Penunjuk data raksasa mengenai segala atribut si Klien pengirim permintaan masuk. Menampung data *Method* (Apakah ini operasi `GET` membaca data atau `POST` mengubah data?), Parameter *Query String* URL, alamat IP Klien asal (`r.RemoteAddr`), dan utamanya: muatan data paket badan (*Payload Body*) yang bisa disedot dan diproses oleh Anda.

```go
package main

import (
    "fmt"
    "net/http"
)

// 1. Buat Fungsi Pengendali Utama (Controller/Handler)
func haloPengunjung(w http.ResponseWriter, r *http.Request) {
    // Menyetel header agar klien tahu balasan ini adalah JSON
    w.Header().Set("Content-Type", "application/json")

    // Memberikan kode status 200 OK
    w.WriteHeader(http.StatusOK)

    // Mengirim response
    w.Write([]byte(`{"pesan": "Selamat datang di API Go!"}`))
}

func main() {
    // 2. Petakan Rute Mux Dasar
    // Tiap kali ada request ke URL path "/", lemparkan pekerjaannya ke fungsi haloPengunjung
    http.HandleFunc("/", haloPengunjung)

    // 3. Hidupkan Server di port mesin (misal port 8080)
    // Fungsi ini akan terus-menerus terblokir dan melayani request di belakang layar (loop abadi).
    fmt.Println("Server mengudara memantau alamat http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err) // Tangani kepanikan jika port jaringan tsb terbukti sedang diduduki aplikasi lain
    }
}
```

---

### 2. Keajaiban REST API: *Routing* Berbasis Objek ServeMux (Fitur Rilis Go 1.22+)

Dahulu kala, ServeMux standar sangat bodoh karena tak memahami deklarasi variabel URL. Pengembang terpaksa memakai pustaka `gorilla/mux`. Kini Go merespons keluhan itu!

Sekarang Anda diwajibkan menginisiasi struktur `mux := http.NewServeMux()` dan memanfaatkannya sebagai sentral perutean yang dapat membedah langsung identitas parameter liar di tengah-tengah rentetan URL Web (menggunakan simbol `{}`). Dan mencomot *String* parameter tersebut cukup lewat panggilan `r.PathValue()`.

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Bangun Router Terisolasi
    router := http.NewServeMux()

    // Endpoint 1: Hanya menanggapi Verb Metode GET ke daftar benda (Dilarang metode POST/PUT/DELETE, otomatis ditendang)
    router.HandleFunc("GET /api/buku", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Membuka halaman indeks seluruh katalog buku Perpustakaan...")
    })

    // Endpoint 2: Fitur Menawan Path Variables! (menangkap identitas buku abstrak dari URL `{isbn}`)
    router.HandleFunc("GET /api/buku/{isbn}", func(w http.ResponseWriter, r *http.Request) {
        // Tarik variabel ekstrak URL secara instan lewat fungsi mutakhir:
        nomorISBN := r.PathValue("isbn")

        kalimatLaporan := fmt.Sprintf("Mencari rincian detil buku dengan Kode Registrasi ISBN: %s", nomorISBN)
        w.Write([]byte(kalimatLaporan))
    })

    // Endpoint 3: Menangkap metode HTTP Spesifik
    router.HandleFunc("POST /api/buku", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusCreated) // Kembalikan kode 201 Created (Buku baru dirilis dan dimasukkan)
        w.Write([]byte("Buku baru sukes ditambahkan ke gudang data!"))
    })

    fmt.Println("Peladen Berbasis Mux Modern Go 1.22 beroperasi pada Port: 8080")
    // Kirim objek router Mux yang telah dibekali kecerdasan khusus kita ke parameter peladen.
    // Uncomment baris di bawah untuk menjalankan server ini:
    // http.ListenAndServe(":8080", router)
}
```

---

### 3. Menggunakan HTTP Client (Mengambil Data dari Luar)
Go membuat pengambilan data dari API luar menjadi satu baris kode dengan `http.Get`. Harap diingat bahwa kita **wajib menutup body response** (`defer resp.Body.Close()`) agar memori jaringan tidak bocor!
```go
package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
    if err != nil {
        panic("Gagal request: " + err.Error())
    }
    // PENTING: Jangan lupa ditutup setelah dibaca
    defer resp.Body.Close()

    // Membaca seluruh isi balasan dari server
    body, _ := io.ReadAll(resp.Body)
    fmt.Println("Status Respon:", resp.StatusCode) // misal: 200
    fmt.Println("Isi:", string(body))
}
```

---

## Bagian Lanjutan: Keamanan *Router Multiplexer*, Performa Tinggi, dan Kustomisasi Timeout Kritis

Secara harafiah, bahasa `Go` dapat menjalankan server HTTP yang fungsional penuh hanya dengan tiga baris kode murni (`http.HandleFunc`, lalu panggil `http.ListenAndServe`). Akan tetapi, di dunia produksi Enterprise (Level Perusahaan) arsitektur itu sangat rapuh jika diekspos ke antarmuka internet luar tanpa dipersenjatai lapisan pengaman batas waktu jaringan!

### 1. Pembunuh Senyap: Masalah "The Empty Server Timeout"

Fungsi legendaris `http.ListenAndServe(":8080", nil)` yang terdapat di ribuan artikel tutorial dasar di internet sebenarnya **menyimpan satu masalah kritis bagi sistem Production!**
Fungsi peladen bawaan asali itu sama sekali **TIDAK MEMILIKI TIMEOUT (Batas Waktu Tunggu) KONEKSI!**

Jika peretas atau bot jaringan iseng sengaja membuka koneksi TCP ke port `8080` Anda, mengirim sebuah HTTP Request, tetapi dengan sengaja membaca balasan `Response` Server secara teramat lambat (1 byte per 1 menit) -- Server Go Anda akan dengan senang hati menunggu dengan sabar meladeni koneksi tersebut tanpa menendangnya. Praktik serangan ini dikenal sebagai *Slowloris Attack*.

Hasilnya? Setelah diserang oleh 10.000 bot paralel lamban, server RAM Anda habis karena *File Descriptor* dan goroutine menumpuk, menyebabkan kelumpuhan *Gateway Server*.
Solusinya adalah Anda HARUS merakit Obyek `http.Server` kustom sebelum menghidupkannya!

```go
package main

import (
    "net/http"
    "time"
    "fmt"
)

func main() {
    routerPusat := http.NewServeMux()
    routerPusat.HandleFunc("GET /halo", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Salam kenal dari Server Tangguh!"))
    })

    // INI ADALAH CARA YANG BENAR STANDAR INDUSTRI ENTERPRISE
    // Kita merakit sendiri mesin peladen Web-nya dengan batasan toleransi sadis!
    serverKokoh := &http.Server{
        Addr:         ":8080",
        Handler:      routerPusat,

        // 1. Memaksa Server memutuskan koneksi jika Klien tak becus menyelesaikan bacaannya (Request Body) dalam 5 detik
        ReadTimeout:  5 * time.Second,

        // 2. Membatasi Server agar tak menahan proses Goroutine menulis Response balasan melebihi 10 detik
        WriteTimeout: 10 * time.Second,

        // 3. Waktu maksimal koneksi menganggur/sepi (Keep-Alive) sebelum ditendang oleh peladen
        IdleTimeout:  120 * time.Second,
    }

    fmt.Println("Peladen Enterprise Level mengangkasa memantau port :8080 ...")

    // Alih-alih fungsi bawaan, kita memanggil eksekutor mesin Server spesifik yang sudah dirakit:
    // err := serverKokoh.ListenAndServe() ...
}
```

### 2. Membangun Penengah Penyeleksi Jalan (Middleware Interceptor)

Paket `net/http` tidak memiliki sistem *Middleware* magis terintegrasi bawaan ala Express.js Node, namun filosofi *Duck Typing* antarmuka (Interface) Go secara alami mendukung perakitan *Middleware Pattern* tingkat dewa.

*Middleware* di Go pada dasarnya hanyalah sebuah Fungsi yang menerima parameter bertipe penengah `http.Handler`, lalu membungkus *request* itu dengan pekerjaan inspeksi verifikasi internal (seperti Cek Validasi Otorisasi / Token JWT), baru kemudian merestui fungsinya dioper diteruskan kepada *Handler* selanjutnya di ujung tujuan.

```go
// Tipe definisi alias fungsi bungkus Middleware
func VerifikasiScurity(nextHandler http.Handler) http.Handler {

    // Kita mengembalikan sebuah Fungsi Penyelundup Identitas Baru (Wrapper)
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        // 1. PRA-PEMROSESAN SEBELUM MENUJU KE TUJUAN UTAMA
        kunciRahasia := r.Header.Get("X-API-KEY")

        if kunciRahasia != "SandiSuperSakti" {
            // Mencegat! Menendang Klien Pencuri!
            http.Error(w, "Harap mendaftar API KEY sah sebelum memasuki Istana", http.StatusUnauthorized)

            // RETURN mutlak memberhentikan aliran request di titik poin ini.
            return
        }

        // 2. MENERUSKAN RESTU KE FUNGSI SELANJUTNYA
        // Memanggil nextHandler agar ia dieksekusi normal (Karena lulus pemeriksaan satpam)
        nextHandler.ServeHTTP(w, r)
    })
}

// ... Penggunaan (Asumsikan di fungsi main):
// mux.Handle("GET /harta_rahasia", VerifikasiScurity(http.HandlerFunc(PencetakDataEmasRahasia)))
```

### 3. File Server Statik (Asset Tembakan Browser Web)

Kejeniusan utama `net/http` adalah dukungan file server super berkecepatan kencang yang bisa membalap fungsi Server NGINX statik! Saat membuat antarmuka Aplikasi React SPA atau Vue Frontend statis, Anda tidak butuh sistem Apache/Nginx eksternal yang ruwet *config*-nya.

Golang menyajikan dukungan pemancaran direktori `images/` HTML/CSS dengan balutan satu baris perintah sakti fungsi `http.FileServer`. Paket bawaan ini sudah memiliki pemahaman mendalam tentang header Range Partial HTTP (Bisa me*resume* unduhan jika putus), penanganan header Modifikasi (If-Modified-Since Cache HTTP *Status 304*), lengkap perlindungan kebal serapan ancaman peretasan mundur (*Directory Traversal Hack `../../../etc/password`*) ke file sistem sensitif Linux.

```go
// Meminta Go melayani Folder bernama "file_gambar_static" yang ada di sebelahnya
folderPublikSistem := http.FileServer(http.Dir("./file_gambar_static"))

// Mendaftarkan Router dengan penghapusan rute Prefix bawaan (StripPrefix)
// Maksudnya: Jika ada Request ke alamat IP /aset/logo.png,
// Suruh Go mencari file "logo.png" secara murni langsung di dalam folder "./file_gambar_static" tanpa awalan "/aset".
// routerMuxKita.Handle("/aset/", http.StripPrefix("/aset/", folderPublikSistem))
```

Peladen HTTP Go adalah dewa di industri layanan awan (Cloud Native Services) modern karena struktur perpaduan komponen fungsi bawaan tanpa dependensi eksternal (100% Go native) yang stabil bak bongkahan logam abadi, melayani rute permintaan dari penjuru angkasa jagat raya secara efisien tinggi. Menguasainya berarti menguasai jembatan pertukaran arus komunikasi HTTP secara komprehensif.

## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```

---

## Studi Kasus Dunia Nyata: Pola Desain Pemutusan Panggilan Berantai (Timeout Context Propagation) di Sistem Terdistribusi

Bayangkan masalah terparah (*Cascading Failure*) yang menimpa arsitektur mikroservis raksasa (Netflix/Gojek).
Klien dari *Mobile App* (HP Android) memanggil API Utama Anda (`Microservice A`).
API Utama `A` memanggil Layanan Produk `B`. Layanan Produk `B` lantas memanggil Database Gudang `C`.

Bagaimana jika Database Gudang `C` tiba-tiba mati lampu atau internet kabelnya terputus dan mengambang diam tanpa merespons balik?
Jika tak ada sistem pelindung, `B` akan menunggu selamanya. Karena `B` menunggu selamanya, `A` pun menunggu selamanya. Dan Klien HP pun layar loadingnya berputar tanpa henti selamanya! Jutaan *Request* masuk, tertumpuk diam di memori RAM, lalu seluruh Server dunia hancur lebur mati kehabisan kapasitas.

Mekanisme penawar surga dari Go adalah **Pengangkutan Context Batal (Context Timeout Propagation) melewati batas-batas Klien `net/http` !**

### Menjalin Rantai Ikatan Pembatalan Menggunakan `http.NewRequestWithContext`

Alih-alih merakit pesanan HTTP keluar biasa `http.NewRequest()`, Anda HAKUL YAKIN wajib menambatkan *Context Request Klien Induk* pada pesan surat *Client* yang akan Anda terbangkan pergi!
Begitu HP klien ditutup aplikasinya, Klien putus koneksi $\rightarrow$ Request Induk Batal $\rightarrow$ Surat Request Anak ke Servis B Batal $\rightarrow$ Penggalian Database C langsung otomatis mati diputus dari memori! Sangat Brilian!

```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

// 1. FUNGSI HANDLER UTAMA API KITA (Melayani HP Klien Android)
func GerbangPusatAPIHandler(w http.ResponseWriter, r *http.Request) {

    fmt.Println("[SERVER A] Menerima ketukan pintu masuk dari HP Klien.")

    // Trik Mahadewa Golang: Kita cabut "Roh Nyawa" (Context) bawaan dari sang Klien HP ini!
    // Apabila Klien tiba-tiba menutup App Androidnya, Roh Nyawa ini langsung Teriak "BATAL!" ke segala penjuru.
    rohNyawaKlien := r.Context()

    // Namun kita juga tak mau diakali: Kalau Klien sabar menunggu, KITA BATASI SERVER HANYA MAKSIMAL 3 DETIK KERJA!
    // Kita bungkus roh tersebut dengan rompi Bom Waktu 3 Detik.
    ctxMaksimal3Detik, cabutBomWaktu := context.WithTimeout(rohNyawaKlien, 3*time.Second)
    defer cabutBomWaktu() // Bereskan sisa timer jika kerjaan selesai lebih awal

    // 2. MEMANGGIL SERVIS TETANGGA (SERVIS B)
    // Trik Keselamatan Kiamat Server: JANGAN PAKAI http.Get() BIASA !!
    // Pakailah NewRequestWithContext, dan selundupkan/kalungkan Roh Nyawa Klien Ber-Bom Waktu tadi kepadanya!
    suratPanggilanKeluar, _ := http.NewRequestWithContext(ctxMaksimal3Detik, "GET", "https://situs-tujuan-eksternal.com/api/katalog", nil)

    klienJaringan := &http.Client{}

    fmt.Println("[SERVER A] Berangkat Terbang mengambil data ke Servis Luar B...")

    // MENGEKSEKUSI PENERBANGAN
    // Jika situs-eksternal mati total macet diam 1 menit, Klien ini secara AJAIB di detik ke-3
    // akan memotong kabel LAN secara sepihak dan mereturn ERROR!
    responLuar, errTelepon := klienJaringan.Do(suratPanggilanKeluar)

    if errTelepon != nil {
        // Melacak Jejak Dosa Kematian: Kenapa dia gagal?
        // Apakah karena Situs Eksternal B nya jelek, atau karena Timeout kita yg meledak duluan?
        if ctxMaksimal3Detik.Err() == context.DeadlineExceeded {
            fmt.Println("[PERTAHANAN SISTEM] Panggilan dimatikan secara kejam oleh Pelindung Timeout Server A! RAM kita terselamatkan dari kemacetan maut!")
            http.Error(w, "Mohon bersabar sistem sedang sibuk luar biasa (Gateway Timeout)", http.StatusGatewayTimeout)
            return
        }

        fmt.Println("Situs eksternal menolak telepon koneksi awal (Connection Refused).")
        http.Error(w, "Gagal menghubungi layanan kawan", http.StatusBadGateway)
        return
    }

    defer responLuar.Body.Close() // Pastikan ditutup!

    w.Write([]byte("Sukses mengambil Harta Karun dari Server Luar dengan Selamat Tepat Waktu!"))
}

func main() {
    // Andaikan kita menyalakan Peladen A.
    // http.HandleFunc("/toko", GerbangPusatAPIHandler)
    // http.ListenAndServe(":8080", nil)
}
```

Pengaplikasian terpadu tali korelasi nyawa (*Context Bound*) pada modul pengangkut pesawat jet `net/http` klien akan memastikan arsitektur ekosistem *Backend Server* Go (Golang) milik organisasi Anda mustahil bisa dihancurkan tumbang di bawah tekanan kemacetan lajur latensi jaringan inter-servis *Microservices* lintas benua terburuk sekalipun!
