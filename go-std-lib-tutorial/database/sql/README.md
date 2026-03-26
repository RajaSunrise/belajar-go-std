# Modul: `database/sql`

## Ringkasan
Package `database/sql` menyediakan antarmuka generik yang independen-pemasok (*vendor-agnostic*) untuk mengakses dan berinteraksi secara aman dengan basis data relasional tingkat *Enterprise* (seperti PostgreSQL, MySQL, SQLite, Oracle, hingga MS SQL Server).

## Penjelasan Lengkap (Fungsi & Tujuan)
Di dalam ekosistem bahasa pemrograman lain (misalnya Java dengan JDBC atau PHP dengan PDO), pengembang sering kali harus berhadapan dengan kerumitan antarmuka yang sangat padat dan kaku untuk berkomunikasi dengan *Database Management System* (DBMS). Di sisi lain, bahasa pemrograman Go (Golang) menghadirkan inovasi brilian melalui arsitektur `database/sql`.

Hal terpenting yang wajib dipahami oleh seorang pemrogram Go pemula adalah: **Package `database/sql` INI BUKANLAH DRIVER DATABASE!** Package ini sama sekali tidak tahu bagaimana cara mengirim bit jaringan TCP ke server MySQL atau Postgres. Sebaliknya, package ini hanyalah seperangkat spesifikasi cetak biru (*Interface Contracts*) dan sebuah mesin pengelola *Connection Pool* yang sangat tangguh di balik layar. Agar program Go Anda dapat benar-benar membaca data dari server MySQL, Anda harus mengimpor **Driver MySQL pihak ketiga** (misalnya `github.com/go-sql-driver/mysql`), yang secara diam-diam (*blank identifier import*) mendaftarkan dirinya ke dalam pelukan peladen mesin `database/sql`.

**Tujuan dan Fungsi Utama:**
1.  **Pengelolaan Kolam Koneksi (Connection Pooling):** Ini adalah fitur pembunuh (*killer feature*). Anda cukup memanggil `sql.Open()` satu kali di awal jalannya *server*. Di balik layar, Go akan mengurus buka-tutup ribuan koneksi jaringan paralel ke basis data (*multiplexing*), memastikan tidak ada satupun *request* HTTP yang macet karena kehabisan saluran koneksi, sekaligus mencegah server basis data lumpuh (*Connection Refused*) akibat serbuan goroutine.
2.  **Abstraksi Vendor Universal:** Jika hari ini arsitektur Anda memakai *SQLite* ringan, lalu tahun depan bisnis meroket dan tim infrastruktur memaksa migrasi besar-besaran ke peladen awan *PostgreSQL*, Anda HAMPIR TIDAK PERLU MERUBAH fungsi-fungsi logika baris program (seperti `db.Query` atau `db.Exec`) Anda di Go! Anda cukup mengubah baris impor nama driver di atas. (*Agnostic Portability*).
3.  **Pelindung Injeksi Mematikan (Prepared Statements):** Menghentikan celah keamanan fatal legendaris *SQL Injection* secara mutlak. Dengan menggunakan parameter pelindung (`?` di MySQL/SQLite atau `$1` di Postgres) saat fungsi `.Query()` dieksekusi, peretas tidak dapat menelan tabel Anda dengan sisipan trik sintaks mematikan (seperti `'; DROP TABLE users;--`).
4.  **Transaksi Integritas (ACID Transactions):** Menyediakan mekanisme untuk mengeksekusi puluhan kueri rumit (seperti memindahkan saldo Uang dari Rekening A ke B). Jika satu saja kueri gagal di tengah jalan, seluruh rentetan proses tersebut akan otomatis dibatalkan (`Rollback`) secara atomik, melindungi keabsahan laporan keuangan perusahaan (*melalui `db.BeginTx()`*).

**Mengapa menggunakan `database/sql`?**
Walaupun saat ini banyak beredar Pustaka *Object-Relational Mapping (ORM)* mewah di Go seperti *GORM* atau *Ent*, mereka semua secara mutlak mengandalkan dan membungkus `database/sql` ini di dasar perut mesin mereka. Merancang *Microservices Backend* mutakhir yang diklaim berkinerja tinggi hanya akan sah diakui apabila Anda mengerti murni tulisan bahasa SQL asali (*Raw SQL Query*) menggunakan struktur efisien dari paket ini, ketimbang membebani sistem dengan terjemahan ORM berat.

---

## Daftar Fungsi Umum dan Cara Penggunaannya Secara Mendalam

### 1. Kunci Utama Kestabilan (Membuka dan Menutup Connection Pool)

Fungsi `sql.Open` sering disalahpahami. Ia **TIDAK** seketika menyambung atau menelepon *Database Server* Anda. Ia hanya merumuskan konfigurasi (Kredensial, Alamat IP, *Password*). Untuk membuktikan apakah pangkalan data benar-benar hidup, bernafas, dan merespons, Anda wajib melanjutkannya dengan ujian fungsi denyut jantung `db.Ping()`.

```go
package main

import (
    "database/sql"
    "fmt"
    // HARUS MENGIMPOR DRIVER! (Menggunakan garis bawah _ agar fungsi init()-nya jalan tanpa diprotes Go)
    // _ "github.com/go-sql-driver/mysql"
)

// Simulasi: Tolong hiraukan _ import di atas jika tak memiliki mesin driver asli di laptop Anda.

func main() {
    // 1. Inisialisasi Kredensial (Format tulisan URI ini sangat spesifik tergantung Driver apa yang dipakai)
    // Format MySQL: user:password@tcp(127.0.0.1:3306)/dbname
    alamatRahasiaDB := "root:rahasia@tcp(127.0.0.1:3306)/toko_online"

    dbPeladen, errBuka := sql.Open("mysql", alamatRahasiaDB)
    if errBuka != nil {
        panic("Format string koneksi Ditolak Sistem!")
    }

    // 2. ATURAN EMAS: Selalu pastikan Kolam Koneksi Utama ini ditutup saat Aplikasi Mati!
    defer dbPeladen.Close()

    // 3. PRAKTIK INDUSTRI: Menyetel Batasan Ketahanan (Tuning Performance)
    // Jangan biarkan koneksi membengkak! Batasi misalnya 100 koneksi max, 10 yang idle (santai).
    dbPeladen.SetMaxOpenConns(100)
    dbPeladen.SetMaxIdleConns(10)

    // 4. Momen Pembuktian (Mencoba menelepon sungguhan)
    errTesKoneksi := dbPeladen.Ping()

    if errTesKoneksi != nil {
        fmt.Println("GAWAT: Server MySQL mungkin mati atau Password Salah!", errTesKoneksi)
    } else {
        fmt.Println("Pangkalan Data terhubung dengan selamat sentosa.")
    }
}
```

---

### 2. Modifikasi Data Mutlak (Exec - INSERT, UPDATE, DELETE)

Bila Anda berniat merubah isi, menggeser tabel, atau melumat data (DDL / DML) dan sama sekali **TIDAK MENGHARAPKAN balasan berupa rentetan data pembacaan** selain konfirmasi sukses, gunakan fungsi eksekusi `db.Exec()`.

Sangat diwajibkan menggunakan sistem parameter *Placeholder* (contoh: tanda tanya `?`) demi alasan keamanan *Anti-SQL-Injection* dan percepatan mesin pangkalan data (*Prepared Statements Cache*).

```go
package main

import (
    "database/sql"
    "fmt"
)

func main() {
    // Asumsikan dbPeladen sudah berhasil di-Open di langkah sebelumnya
    var dbPeladen *sql.DB
    if dbPeladen == nil {
        fmt.Println("Simulasi Modifikasi Eksekusi Kueri...")
    }

    // A: Kueri Insert yang Aman dari Serangan Peretas
    namaMemberBaru := "Joko Santoso"
    umurMember := 27

    // Mesin Go akan secara mandiri membersihkan tanda kutip aneh dari variabel namaMemberBaru
    // hasilLaporanEksekusi, errAksi := dbPeladen.Exec("INSERT INTO member (nama_lengkap, usia) VALUES (?, ?)", namaMemberBaru, umurMember)

    // B: Contoh Penggunaan Update (Perubahan Data)
    // _, errUpdate := dbPeladen.Exec("UPDATE member SET usia = ? WHERE id = ?", 28, idTercipta)
}
```

---

### 3. Ekstraksi Penguraian Membaca Banyak Baris Data (Query & Scan)

Bagian ini sedikit rumit namun inilah pekerjaan harian Anda kelak. Manakala Anda mengeksekusi kueri penarikan `SELECT`, peladen tidak memuntahkan semuanya ke RAM Anda secara brutal. Ia mengembalikan objek Pintu Aliran (`sql.Rows`). Anda berkewajiban membuka pintu itu berulang kali melangkah (*Iterate / Loop*) menggunakan `.Next()`, lalu menyerap (*Scan*) cairan data di kolom tersebut menuju gelas-gelas variabel memori di Struct aplikasi Go Anda menggunakan *Pointers*.

**Aturan Sangat Fatal:** Pintu Lorong Data `Rows` WAJIB MAHA WAJIB Ditutup dengan `defer rows.Close()` sesegera mungkin! Jika Anda lupa, koneksi terowongan itu akan macet/menggantung selamanya (*Goroutine Leak / Connection Exhaustion*)!

```go
package main

import (
    "database/sql"
    "fmt"
)

// Merancang cetakan penampung memori
type ProfilMember struct {
    ID     int
    Nama   string
    Umur   int
    Status sql.NullString // Rahasia Menangani Kolom NULL Database (Kosong Tanpa Nilai)
}

func main() {
    var dbPeladen *sql.DB
    if dbPeladen == nil {
        fmt.Println("Simulasi Pembacaan Masif Ekstraksi Pangkalan Data...")
    }

    // lorongHasilBaris, errQuery := dbPeladen.Query("SELECT id_member, nama_lengkap, usia FROM member WHERE usia > ?", 20)
    // defer lorongHasilBaris.Close()

    // var koleksiDaftarOrang []ProfilMember

    // for lorongHasilBaris.Next() {
    //    var obj ProfilMember
    //    errSedot := lorongHasilBaris.Scan(&obj.ID, &obj.Nama, &obj.Umur)
    //    koleksiDaftarOrang = append(koleksiDaftarOrang, obj)
    // }
}
```

---

## Bagian Lanjutan: Misteri Null (Nil/Kosong), Pola Koneksi Abadi (Connection Pooling), dan Jebakan Memori Barisan Kueri

Sekalipun pustaka standar Go `database/sql` adalah peranti penggerak andalan yang menggerakkan basis data berskala raksasa, desainnya memiliki filosofi keras yang menuntut ketelitian sempurna. Kesalahan paling lazim yang menyebabkan ratusan isu kelumpuhan server di hari pertama *Production* Go adalah bocornya memori koneksi dan ledakan *panic* saat berhadapan dengan data "Null" di kolom tabel SQL.

### 1. Ancaman Kiamat NULL (Nil Value) di Kolom Basis Data

Di bahasa pemograman dinamis (PHP, Node.js/JavaScript), jika sebuah sel tabel di MySQL bernilai kosong (`NULL`), ia akan secara halus masuk dan menjelma menjadi tipe bawaan `null` di RAM aplikasi tanpa menyebabkan kepanikan berarti.

Di Go (Golang), tipe primitif seperti `int` atau `string` **TIDAK BISA** dan **TIDAK PERNAH BISA** menampung nilai *nil*. Variabel `string` di Go selalu memiliki wujud (minimal string kosong `""`), sedangkan `NULL` di database adalah "Ketiadaan Ruang Absolut".

Jika Anda memaksa menyedot (`.Scan(&namaUser)`) sebuah kolom *Nama* bernilai `NULL` dari tabel Postgres ke variabel *string* biasa, program Anda akan meledak (*Panic: sql: Scan error on column index 0: unsupported Scan, storing driver.Value type <nil> into type *string*).

**Senjata Penyelamat: `sql.NullString`, `sql.NullInt64`, dsb.**
Anda WAJIB membungkus tipe data yang berpotensi *NULL* dengan tipe pelindung bersyarat bawaan Go.

```go
// type ProfilStaf struct {
//    IDKaryawan int
//    NamaTengah sql.NullString // BANYAK ORANG INDONESIA TAK PUNYA NAMA TENGAH (Maka di DB ia NULL)
//    UmurAktif  sql.NullInt64  // MUNGKIN SAJA UMUR BELUM DIISI (Kosong)
// }

// rowPenembakSatu := db.QueryRow("SELECT id, nama_tengah, umur FROM karyawan WHERE id = ?", 55)
// var stafTarget ProfilStaf
// errBedah := rowPenembakSatu.Scan(&stafTarget.IDKaryawan, &stafTarget.NamaTengah, &stafTarget.UmurAktif)

// Pengecekan aman
// if stafTarget.NamaTengah.Valid {
//     fmt.Println(stafTarget.NamaTengah.String)
// }
```

### 2. Bahaya Mematikan Kebocoran Kueri Tunggal (`QueryRow`)

Seringkali Anda tidak perlu mengunduh 10.000 baris. Anda hanya butuh 1 baris (contoh: mengecek sandi milik 1 pengguna saat *Login*). Fungsi `.QueryRow()` diciptakan spesifik untuk kemudahan ini (menghindari kerumitan menulis *looping* `.Next()`).

Namun `QueryRow` memiliki **Jebakan Fatal**: Ia menahan koneksi basis data tersebut, MENGUNCI nya dari *Connection Pool*, sampai Anda benar-benar mengeksekusi panggilan `.Scan()` di akhir kalimat!

Jika Anda memanggil `db.QueryRow(...)` lalu tidak pernah memanggil `.Scan()` (karena mungkin Anda cuma iseng memanggilnya), koneksi itu bocor membusuk di memori, dan jika dilakukan 100 kali, seluruh batas koneksi peladen SQL habis (*Pool Exhaustion*). Server Web macet dan berhenti menanggapi semua orang.

```go
// ATURAN BESI QUERY ROW!

// BARIS INI AKAN MENGGANTUNG KONEKSI RAM DATABASE (JIKA DIBIARKAN BERDIRI SENDIRI):
// db.QueryRow("SELECT status_aktif FROM user WHERE id = ?", 12)

// CARA BENAR: Eksekusi berantai (Chaining) sampai tuntas disedot Scan!
// var statusUser bool

// Dengan memanggil Scan, fungsi QueryRow akan otomatis menutup pintu koneksinya (Close)
// di belakang layar seusai menyedot data! Aman untuk Sistem!
// errSatuBaris := db.QueryRow("SELECT status_aktif FROM user WHERE id = ?", 12).Scan(&statusUser)
```

### 3. Ekstremitas Kecepatan Penyisipan Massal (*Bulk Inserts* & *Prepared Statements*)

Jika Bos Perusahaan meminta Anda mengunggah 5000 baris data tagihan excel klien ke *Database* setiap paginya. Pendekatan junior adalah: Membaca excel, lalu melakukan perulangan `for` 5000 kali, memanggil `db.Exec("INSERT INTO tagihan ...")` 5000 kali pula.

Ini sangat melambatkan peladen basis data. Tiap pemanggilan `Exec` biasa, database SQL (Postgres/MySQL) akan membaca *String SQL* yang Anda kirim, menganalisa sintaksnya (Parsing), membuat Rencana Eksekusi (Query Plan Execution), baru mengisikan datanya. 5000 kali!

**Senjata Perang Performa: `db.Prepare()` (Kueri Bersenjata Disiapkan)**

Anda mengajari Mesin SQL Rencana Kueri **SEKALI SAJA**. Ia menyimpan rencana itu di otaknya. Barulah Anda menembakkan 5000 *Peluru Parameter* variabel data ke mesin tersebut secara beruntun tanpa kompilasi ulang kueri dari awal. Ini menaikkan kecepatan pemrosesan *Bulk Insert* secara gila-gilaan menembus angka ribuan persen.

```go
// TAHAP 1: KITA SIAPKAN CETAKAN SENJATA (PREPARE) - HANYA 1 KALI!
// Senapan mesin Statement ini memegang jalur memori khusus ke Database Server
// senapanMesinPenyisipan, errSiapkan := db.Prepare("INSERT INTO laporan_harian (kode_toko, pendapatan) VALUES (?, ?)")

// JANGAN PERNAH LUPA TUTUP SENAPANNYA JIKA SUDAH SELESAI MENEMBAK
// defer senapanMesinPenyisipan.Close()

// TAHAP 2: EKSEKUSI MASSAL
// for _, setoran := range koleksiSetoranTokoPagiIni {
//    _, errTembak := senapanMesinPenyisipan.Exec(setoran.Kode, setoran.Duit)
// }
```

Dengan mengkombinasikan pertahanan struktur statis tipe `sql.NullString`, dan kepiawaian menjaga keran pelindung batas koneksi `QueryRow` dengan pendelegasian tutup *Close* mutlak pada `Rows`, Anda menjamin peladen web API mikroservis Anda kokoh berdiri dari hantaman anomali pangkalan data liar (DB *Network Fluctuation*) yang mendera siang malam di alam jaringan bebas korporat (Production Servers).

## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```

---

## Studi Kasus Dunia Nyata: Arsitektur Repositori Bersih (Clean Architecture)

Menggunakan package `database/sql` secara langsung di dalam file `main.go` atau fungsi penanganan rute (`Handler`) HTTP Anda mungkin tampak mudah di awal, tetapi itu akan berubah menjadi neraka kode (*Spaghetti Code*) saat proyek Anda berkembang melebihi 10 tabel.
Di industri komputasi modern, kita membungkus interaksi database ini menggunakan lapisan "Repository Pattern" atau "Data Access Object (DAO)".

Bagaimana memisahkan logika SQL kotor dari logika Bisnis Suci Anda?

### Pola Penyuntikan Dependensi (Dependency Injection)

Alih-alih membuat variabel `db *sql.DB` sebagai variabel Global (yang sangat dilarang karena merusak pengujian *Unit Testing*), Anda meletakkan koneksi DB sebagai komponen dalam *Struct Repository*.

```go
package repository

import (
    "context"
    "database/sql"
    // Import model domain Anda, misal: "proyek/model"
)

// 1. Antarmuka (Interface) Suci
// Lapisan bisnis Controller HANYA boleh melihat interface ini.
// Controller TIDAK BOLEH tahu apakah di belakang layar kita pakai MySQL atau Postgres.
type PenggunaRepository interface {
    CariBerdasarkanID(ctx context.Context, id int) (*model.Pengguna, error)
    SimpanPenggunaBaru(ctx context.Context, data *model.Pengguna) error
}

// 2. Implementasi Konkret Database SQL
type postgresPenggunaRepo struct {
    koneksiPool *sql.DB
}

// 3. Fungsi Konstruktor (Factory)
// Injeksi: Kita "menyuntikkan" objek koneksi SQL ke dalam repository kita.
func NewPostgresPenggunaRepo(db *sql.DB) PenggunaRepository {
    return &postgresPenggunaRepo{
        koneksiPool: db,
    }
}

// 4. Pelaksanaan Operasi Murni
func (r *postgresPenggunaRepo) CariBerdasarkanID(ctx context.Context, id int) (*model.Pengguna, error) {
    var user model.Pengguna

    // SELALU gunakan metode berakhiran "Context" di tingkat industri!
    kueri := "SELECT id_user, nama_lengkap, email, tanggal_daftar FROM t_users WHERE id_user = $1 LIMIT 1"
    barisTunggal := r.koneksiPool.QueryRowContext(ctx, kueri, id)

    errSedot := barisTunggal.Scan(&user.ID, &user.NamaLengkap, &user.Email, &user.TanggalDaftar)
    if errSedot != nil {
        if errSedot == sql.ErrNoRows {
            return nil, nil // Tidak error, memang orangnya tidak ada.
        }
        // Laporkan error database sebenarnya untuk dicatat di log server.
        return nil, errSedot
    }

    return &user, nil
}
```

Dalam skenario *Unit Testing*, karena *Controller* Anda hanya menuntut antarmuka `PenggunaRepository`, Anda dapat dengan mudah membuat Struct Tiruan (Mock Repository) yang tidak pernah memanggil `database/sql` sama sekali, dan hanya mengembalikan Struct `model.Pengguna` palsu seketika. Pemahaman atas pengkotakan tanggung jawab (*Separation of Concerns*) ini akan menjauhkan arsitektur Peladen Go Anda dari kebusukan perangkat lunak (*Software Rot*).
