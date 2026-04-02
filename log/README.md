# Panduan Komprehensif: Standard Library `log` dan `log/slog` di Golang

Selamat datang di panduan paling komprehensif, terperinci, dan mendalam mengenai standard library Golang untuk kebutuhan *logging*. Di dalam ekosistem pengembangan perangkat lunak modern, logging bukan sekadar fitur tambahan untuk menampilkan pesan teks di konsol, melainkan tulang punggung dari *observability*, *monitoring*, *debugging*, dan *auditing* pada sistem produksi skala besar.

Bahasa pemrograman Go (Golang) pada awalnya menyediakan package `log` yang sangat sederhana, cepat, dan *thread-safe*. Seiring dengan perkembangan paradigma *microservices* dan infrastruktur *cloud-native*, kebutuhan akan *structured logging* (logging terstruktur, biasanya dalam format JSON) menjadi sangat krusial. Merespons kebutuhan ini, mulai versi Go 1.21, tim inti Go memperkenalkan package `log/slog` yang membawa fitur *structured logging* berkinerja tinggi langsung ke dalam *standard library*, tanpa memerlukan dependensi pihak ketiga (seperti `logrus` atau `zap`).

Dokumentasi ini disusun sedemikian rupa agar semua orang, mulai dari pemula hingga *Senior Software Engineer*, dapat memahami secara cepat dan menyeluruh bagaimana memanfaatkan kedua package ini di Golang. Kita akan mengupas arsitektur dasar, implementasi, praktik terbaik (*best practices*), serta studi kasus di dunia nyata.

---

## Daftar Isi
1. [Pengantar Sistem Logging](#1-pengantar-sistem-logging)
2. [Package `log` Klasik: Fondasi Logging di Go](#2-package-log-klasik-fondasi-logging-di-go)
    - [Fungsi-Fungsi Dasar](#fungsi-fungsi-dasar)
    - [Membuat Custom Logger](#membuat-custom-logger)
    - [Konfigurasi Flags dan Prefix](#konfigurasi-flags-dan-prefix)
3. [Evolusi Logging: Mengapa Kita Butuh Structured Logging?](#3-evolusi-logging-mengapa-kita-butuh-structured-logging)
4. [Package `log/slog`: Standard Baru Logging di Go](#4-package-logslog-standard-baru-logging-di-go)
    - [Konsep Dasar `slog`](#konsep-dasar-slog)
    - [TextHandler vs JSONHandler](#texthandler-vs-jsonhandler)
    - [Menggunakan Atribut (Attributes) dan Grup (Groups)](#menggunakan-atribut-attributes-dan-grup-groups)
    - [Menangani Context (Contextual Logging)](#menangani-context-contextual-logging)
5. [Praktik Terbaik (Best Practices) di Production](#5-praktik-terbaik-best-practices-di-production)
6. [Studi Kasus Dunia Nyata](#6-studi-kasus-dunia-nyata)
7. [Kesimpulan](#7-kesimpulan)

---

## 1. Pengantar Sistem Logging

Dalam pengembangan perangkat lunak, log adalah catatan (rekaman) peristiwa yang terjadi selama aplikasi berjalan. Log sangat penting karena ketika aplikasi Anda berjalan di lingkungan produksi (seperti di server AWS, GCP, atau cluster Kubernetes), Anda tidak bisa melakukan *step-by-step debugging* secara langsung menggunakan *debugger*. Satu-satunya jendela Anda untuk melihat apa yang sedang dilakukan oleh aplikasi Anda adalah melalui *log*.

Log yang baik haruslah:
- **Informatif:** Menjelaskan apa yang terjadi, kapan terjadi, dan mengapa terjadi.
- **Dapat Dicari (Searchable):** Memudahkan *engineer* untuk memfilter peristiwa berdasarkan kriteria tertentu (misal: mencari semua log error untuk pengguna dengan ID tertentu).
- **Cepat (Performant):** Proses mencetak log tidak boleh menjadi *bottleneck* (hambatan) yang memperlambat kinerja aplikasi itu sendiri.

Di Go, filosofi "kesederhanaan" sangat dijunjung tinggi. Itulah sebabnya package `log` bawaan dirancang agar sangat mudah digunakan namun tetap memiliki *performance* yang sangat baik.

---

## 2. Package `log` Klasik: Fondasi Logging di Go

Package `log` adalah standard library paling dasar yang selalu ada sejak versi awal Golang. Package ini secara otomatis menyediakan sebuah objek *logger* default (dikenal sebagai *standard logger*) yang mencetak outputnya ke `os.Stderr`.

### Fungsi-Fungsi Dasar

Package `log` menyediakan tiga keluarga fungsi utama:

1. **Print, Printf, Println:** Mencetak pesan log biasa. Alur eksekusi aplikasi akan terus berlanjut.
2. **Fatal, Fatalf, Fatalln:** Mencetak pesan log, kemudian secara otomatis memanggil fungsi `os.Exit(1)`. Fungsi ini akan langsung menghentikan program Anda dengan status error. Sangat berguna ketika aplikasi tidak dapat memulai fungsinya, misalnya karena gagal membaca file konfigurasi atau gagal terhubung ke database utama.
3. **Panic, Panicf, Panicln:** Mencetak pesan log, lalu memanggil fungsi `panic()`. Berbeda dengan `Fatal` yang mematikan aplikasi secara mendadak, `panic` memungkinkan *deferred functions* dieksekusi terlebih dahulu, dan kepanikan ini masih bisa "ditangkap" (recovered) oleh fungsi `recover()` di level pemanggil.

**Contoh Sederhana:**
```go
package main

import "log"

func main() {
    log.Println("Aplikasi sedang diinisialisasi...")

    // log.Fatalf("Gagal terhubung ke database. Aplikasi berhenti.")
    // log.Panicln("Terjadi kesalahan sistem yang tidak terduga!")
}
```

### Membuat Custom Logger

Dalam aplikasi skala menengah hingga besar, sangat disarankan untuk tidak bergantung semata pada *standard logger* bawaan (global), karena pengaturan global dapat berdampak pada seluruh package di aplikasi Anda (termasuk library eksternal yang menggunakan package `log`). Praktik yang lebih baik adalah menginstansiasi (membuat objek) *logger* khusus Anda sendiri.

Kita bisa membuat custom logger menggunakan fungsi `log.New()`. Fungsi ini membutuhkan tiga argumen:
1. `out io.Writer`: Destinasi di mana log akan ditulis. Bisa berupa `os.Stdout`, `os.Stderr`, atau file (misalnya menggunakan `os.OpenFile`).
2. `prefix string`: Teks awalan yang akan selalu dicetak sebelum pesan utama log. Ini sangat berguna untuk membedakan log dari berbagai komponen aplikasi (contoh: `[DATABASE] `, `[HTTP-SERVER] `).
3. `flag int`: Kumpulan bendera (*flags*) untuk menentukan metadata tambahan apa yang perlu disertakan, seperti tanggal, waktu, atau lokasi file.

**Contoh Custom Logger:**
```go
package main

import (
    "log"
    "os"
)

func main() {
    // Membuka atau membuat file untuk menampung log
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal("Gagal membuka file log:", err)
    }
    defer file.Close()

    // Membuat logger khusus yang menulis ke file tersebut
    fileLogger := log.New(file, "APP-INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

    fileLogger.Println("Pesan ini akan ditulis ke dalam app.log, lengkap dengan tanggal, waktu, dan asal file.")
}
```

### Konfigurasi Flags dan Prefix

Package `log` menyediakan beberapa konstanta *flag* bawaan yang bisa digabungkan menggunakan operator *bitwise OR* (`|`):
- `log.Ldate`: Menampilkan tanggal (contoh: `2023/10/25`)
- `log.Ltime`: Menampilkan waktu lokal (contoh: `14:30:05`)
- `log.Lmicroseconds`: Menampilkan waktu hingga tingkat mikrodetik (`14:30:05.123123`)
- `log.LUTC`: Memaksa zona waktu yang ditampilkan menggunakan format UTC, bukan lokal.
- `log.Lshortfile`: Menampilkan nama file dan nomor baris tempat log dipanggil (`main.go:15`)
- `log.Llongfile`: Sama seperti Lshortfile, tetapi dengan path direktori absolut yang lengkap.
- `log.LstdFlags`: Gabungan standar antara `Ldate` dan `Ltime`. Ini adalah pengaturan default pada standard logger.

Anda dapat memodifikasi logger yang sudah berjalan kapan saja menggunakan *methods* seperti `SetFlags()` dan `SetPrefix()`.

---

## 3. Evolusi Logging: Mengapa Kita Butuh Structured Logging?

Meskipun package `log` klasik cepat dan mudah digunakan, ia memiliki satu kelemahan fatal ketika digunakan di infrastruktur modern: outputnya berupa teks mentah (plain text) tak terstruktur.

Mari kita bayangkan pesan log seperti ini:
`2023/10/25 14:30:05 [INFO] User ID 5123 berhasil membeli barang ID 8891 seharga 500000`

Bagi manusia, pesan ini sangat mudah dibaca. Tetapi, bagaimana jika Anda memiliki sistem yang melayani 10.000 permintaan per detik? Anda mungkin akan memompa miliaran baris log ke *log aggregator* seperti Elasticsearch, Splunk, atau Datadog.

Jika manajemen meminta laporan: *"Berapa total nilai pembelian (harga) yang berhasil pada hari ini khusus untuk pengguna di atas ID 5000?"*

Dengan log plain text, Anda terpaksa menulis ekspresi reguler (Regex) kompleks untuk mem-parsing, mengekstrak "ID User", mengekstrak "ID barang", dan mengekstrak "Harga". Regex sangat mahal secara komputasi (memakan CPU) dan sangat rentan patah (misalnya, jika developer iseng mengubah susunan kata di kode sumber).

**Solusinya adalah Structured Logging.**
Daripada mencetak sebuah kalimat cerita, kita mencetak data berupa *key-value pairs* yang didefinisikan secara eksplisit, paling sering direpresentasikan dalam format JSON.

Pesan log sebelumnya jika diubah menjadi JSON:
```json
{
  "time": "2023-10-25T14:30:05Z",
  "level": "INFO",
  "msg": "Transaksi berhasil",
  "user_id": 5123,
  "item_id": 8891,
  "amount": 500000
}
```
Format JSON ini sangat mudah diurai (parsed) secara otomatis oleh mesin-mesin *log aggregator*. Mereka bisa langsung mengindeks field `amount` sebagai tipe data numerik dan mengizinkan Anda melakukan query matematis seperti agregasi, rata-rata, dengan instan dan tanpa regex!

---

## 4. Package `log/slog`: Standard Baru Logging di Go

Melihat betapa mendesaknya kebutuhan industri akan structured logging, komunitas Go selama bertahun-tahun merilis banyak library eksternal sukses (seperti `sirupsen/logrus`, `uber-go/zap`, dan `rs/zerolog`). Namun, banyaknya variasi library membuat ekosistem menjadi terpecah-pecah. Library A mungkin menggunakan antarmuka logger Zap, sementara Library B mengharapkan antarmuka Logrus. Hal ini menyulitkan integrasi komponen di dalam aplikasi yang lebih besar.

Oleh karena itu, package `log/slog` dilahirkan. Dirancang oleh komite Go dan diinkubasi dalam repositori ekperimental sebelum akhirnya dirilis resmi di Go 1.21. `slog` menyediakan API standard, berkinerja tinggi, ringan, dan sangat kuat untuk kebutuhan logging terstruktur.

### Konsep Dasar `slog`

`log/slog` mengusung desain yang memisahkan antara "Frontend" (API yang dipanggil oleh programmer) dan "Backend" (logika yang memformat dan menulis output akhir).

- **Logger:** Ini adalah objek frontend tempat developer memanggil metode seperti `Info()`, `Debug()`, `Warn()`, dan `Error()`.
- **Record:** Representasi internal dari sebuah log event (berisi waktu, level, pesan, dan daftar atribut).
- **Handler:** Antarmuka backend (`slog.Handler`). Handler bertugas mengambil objek `Record` dan memutuskan bagaimana merendernya (misalnya, menjadikannya JSON, lalu mengirimnya ke console atau jaringan).

Hal ini membuat `slog` sangat *extensible*. Anda dapat membuat *custom handler* sendiri, misalnya Handler yang akan mengirim error berlevel fatal secara otomatis sebagai peringatan (alert) ke aplikasi Slack atau Telegram tim on-call Anda.

### TextHandler vs JSONHandler

Standard library `slog` menyediakan dua handler bawaan yang paling sering digunakan:

1. **`slog.TextHandler`**: Cocok digunakan pada lingkungan pengembangan (development environment). Log dirender sebagai daftar key=value yang ramah di mata manusia.

   ```go
   handler := slog.NewTextHandler(os.Stdout, nil)
   logger := slog.New(handler)
   logger.Info("Aplikasi siap menerima koneksi")
   // Output: time=2023-10-25T14:30:05.000Z level=INFO msg="Aplikasi siap menerima koneksi"
   ```

2. **`slog.JSONHandler`**: Wajib digunakan di production. Log dirender sebagai JSON valid dalam satu baris (NDJSON - Newline Delimited JSON).

   ```go
   handler := slog.NewJSONHandler(os.Stdout, nil)
   logger := slog.New(handler)
   logger.Info("Aplikasi siap menerima koneksi")
   // Output: {"time":"2023-10-25T14:30:05.000Z","level":"INFO","msg":"Aplikasi siap menerima koneksi"}
   ```

### Menggunakan Atribut (Attributes) dan Grup (Groups)

Kekuatan utama `slog` terletak pada cara Anda menyisipkan data. Cara paling efisien dan *strongly-typed* adalah menggunakan fungsi konstruktor atribut yang telah disediakan:

- `slog.String("key", "value")`
- `slog.Int("key", 100)`
- `slog.Float64("key", 99.9)`
- `slog.Bool("key", true)`
- `slog.Time("key", time.Now())`
- `slog.Duration("key", time.Second)`
- `slog.Any("key", complexStruct)`

Penggunaan *strongly-typed attributes* ini memastikan alokasi memori (memory allocation) ditekan seminimal mungkin (menghindari proses *reflection* yang lambat di Go).

Selain itu, jika Anda memiliki objek data yang besar dan logis, Anda dapat mengelompokkannya di bawah satu *namespace* menggunakan `slog.Group`.

**Contoh Kasus:**
```go
logger.Info("Memulai proses sinkronisasi data",
    slog.String("task_id", "T-892"),
    slog.Group("database",
        slog.String("host", "db.internal.svc"),
        slog.Int("port", 5432),
        slog.String("driver", "postgres"),
    ),
    slog.Group("metrics",
        slog.Int("records_found", 15000),
        slog.Duration("query_time", 250*time.Millisecond),
    ),
)
```
Output JSON yang dihasilkan akan jauh lebih rapi dan bersarang (*nested object*), yang membuat *indexing* pada sistem monitoring seperti Kibana menjadi sangat terstruktur.

### Menangani Context (Contextual Logging)

Di dalam aplikasi Go modern (terutama web server atau microservices gRPC), kita sering kali mengoper objek `context.Context` ke setiap lapisan aplikasi. `Context` sangat berguna untuk membawa informasi berharga seperti `request_id`, informasi autentikasi user (`user_id`), atau parameter tracing.

Anda dapat mendesain *custom handler* pada `slog` yang secara otomatis mengekstrak nilai-nilai dari *context* dan menyematkannya ke dalam log. Selain itu, Anda juga dapat menggunakan method `LogAttrs`:

```go
func handleRequest(ctx context.Context, logger *slog.Logger) {
    reqID, _ := ctx.Value("request_id").(string)

    // Log level Info dengan menyuntikkan informasi context
    logger.LogAttrs(ctx, slog.LevelInfo, "Memproses data pengguna",
        slog.String("request_id", reqID),
    )
}
```

Metode ini memastikan bahwa setiap rentetan peristiwa (log trail) yang berasal dari satu *HTTP Request* yang sama dapat ditelusuri dengan sangat mudah karena memiliki identitas log (`request_id`) yang selalu sama di setiap baris lognya. Ini adalah teknik mutlak dalam *distributed tracing*.

---

## 5. Praktik Terbaik (Best Practices) di Production

Setelah memahami komponen dari package `log` dan `log/slog`, saatnya menerapkan pengetahuan ini layaknya seorang engineer profesional.

### Jangan Gunakan fmt.Println Untuk Logging!
`fmt.Println` atau `fmt.Printf` tidak memiliki fitur penanda waktu (timestamp), tidak menangani konkurensi (goroutine safety) dengan aman pada output tertentu, dan sama sekali tidak bisa difilter berdasar tingkat *severity* (seperti Info, Warning, Error). Selalu gunakan `log` atau `slog`.

### Hindari Penggunaan Panic/Fatal Secara Sembarangan
Penggunaan `log.Fatal` harus sangat dibatasi, idealnya hanya di dalam fungsi `main()` aplikasi pada saat proses *bootstrap* / inisialisasi awal. Contoh sah menggunakan `Fatal`:
- Konfigurasi environment (env vars) yang diwajibkan ternyata tidak ada.
- Kegagalan saat mem-parsing konfigurasi utama.
- Gagal bind ke port server.

Namun, JANGAN pernah memanggil `Fatal` atau `Panic` di dalam handler rute HTTP. Jika sebuah request gagal mengakses database, itu seharusnya adalah log berlevel `Error` dan mengembalikan respons HTTP `500 Internal Server Error`, dan BUKAN mematikan (crash) keseluruhan server Go Anda!

### Tentukan Aturan Level Log Anda
Standarkan pedoman ini dalam tim kerja Anda:
- **DEBUG (`slog.LevelDebug`):** Sangat rinci (verbose). Informasi variabel, langkah-langkah loop iterasi, payload mentah request/response (hati-hati jangan membocorkan password). Umumnya dimatikan di lingkungan Production untuk menghemat biaya *storage* log.
- **INFO (`slog.LevelInfo`):** Peristiwa bisnis rutin dan siklus hidup sistem yang penting. Misalnya: "Server mulai", "Job harian selesai", "Pembayaran diterima".
- **WARN (`slog.LevelWarn`):** Peristiwa yang tidak biasa atau tidak diinginkan, tetapi sistem masih bisa melanjutkan tugasnya tanpa terganggu secara fungsional. Misalnya: "Mencoba ulang (retry) koneksi ke API pihak ketiga ke-2 kali karena timeout".
- **ERROR (`slog.LevelError`):** Kegagalan fungsional yang memerlukan perhatian engineer. Operasi (misal, sebuah request spesifik) gagal secara total karena kondisi yang tak tertangani.
- **FATAL / PANIC:** Aplikasi tidak bisa dipertahankan hidup lebih lama lagi. Proses mati sepenuhnya.

### Sentralisasi Inisialisasi Logger
Jangan membuat instansiasi `slog.NewJSONHandler` berulang kali di setiap file `main.go` yang Anda miliki. Buatlah satu package kecil di internal Anda, katakanlah `pkg/logger`, yang menyediakan fungsi `InitLogger(environment string)` yang mana fungsinya akan mengembalikan TextHandler untuk argumen `environment="dev"` dan JSONHandler untuk `environment="prod"`.

### Set Default Logger agar Library Eksternal Patuh
Terkadang Anda menggunakan library buatan pihak ketiga yang keras kepala mencetak log mereka menggunakan package `log` klasik (standard log global). Agar format output aplikasi Anda tidak kacau balau (bercampur antara JSON dan text biasa), `slog` menyediakan utilitas luar biasa:
```go
// Buat JSON logger kita
jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

// Ini akan memaksa semua panggilan `log.Println` atau `log.Printf` klasik yang ada di seluruh project (termasuk library pihak ketiga)
// untuk diarahkan melalui engine JSONHandler milik slog. Level defaultnya otomatis diset sebagai INFO.
slog.SetDefault(jsonLogger)
```
Trik kecil ini menjamin bahwa seluruh *stdout* aplikasi Anda adalah 100% JSON yang seragam.

---

## 6. Studi Kasus Dunia Nyata

### Kasus: Menangani PII (Personally Identifiable Information)
Saat menggunakan log berbasis atribut dengan bebas, engineer sering kali secara tak sengaja melog data sensitif seperti password, NIK, atau nomor Kartu Kredit. Ini berbahaya secara keamanan dan melanggar aturan privasi data (GDPR/PDPA).

Menggunakan `log/slog`, kita bisa mencegat dan memanipulasi *Record* log sebelum dicetak dengan mendefinisikan *custom `ReplaceAttr`* saat mengonfigurasi `slog.HandlerOptions`.

```go
opts := &slog.HandlerOptions{
    Level: slog.LevelInfo,
    ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
        // Redaksi (sensor) otomatis field-field terlarang
        if a.Key == "password" || a.Key == "credit_card" {
            a.Value = slog.StringValue("****** [REDACTED] ******")
        }
        return a
    },
}

logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
logger.Info("Update Profil Gagal",
    slog.String("username", "budi123"),
    slog.String("password", "p4ssw0rdRahasia"), // Nilai ini akan disensor!
)
```
Solusi ini elegan, efisien, dan diproses tepat di level arsitektur logging bawah, tanpa Anda harus merombak logika aplikasi bisnis sama sekali.

---

## 7. Kesimpulan

Menguasai ekosistem logging yang disediakan secara bawaan oleh bahasa pemrograman Go (Golang) adalah syarat wajib bagi siapapun yang ingin merancang aplikasi *backend* yang andal, dapat diskalakan (scalable), dan mudah dipantau (observable).

Package `log` menawarkan kecepatan, kesederhanaan ekstrim, dan kemudahan dalam penanganan utilitas CLI atau aplikasi kecil skala lab. Namun, di dunia industri yang didominasi arsitektur cloud, kontainerisasi (Docker/Kubernetes), serta instrumen diagnostik yang kompleks, integrasi dengan struktur data log yang solid mutlak diperlukan. Itulah mengapa transisi dan adaptasi ke package `log/slog` adalah langkah terbaik yang bisa diambil oleh engineer modern.

Dengan memanfaatkan kapabilitas tinggi dari `slog` seperti *JSON formatting*, *key-value attribute nesting* (grup), mitigasi masalah alokasi memori melalui *typed attributes*, serta integrasi *Context* yang tanpa cela, Go membuktikan bahwa kita tidak selalu harus berlari ke ekosistem library open-source eksternal (third-party dependencies) untuk mendapatkan infrastruktur software kelas *enterprise*. Standar library bawaan Go seringkali lebih dari cukup untuk menuntaskan pekerjaan secara gemilang!
