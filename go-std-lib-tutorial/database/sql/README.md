# Modul: `database/sql`

## Ringkasan
Package `database/sql` menyediakan antarmuka standar untuk berinteraksi dengan database relasional (SQL).

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `database/sql` adalah abstraksi brilian di dalam Go. Package ini **bukanlah** implementasi koneksi ke spesifik database, melainkan *interface umum* yang mendefinisikan cara program Go seharusnya berkomunikasi ke *driver* database manapun (seperti MySQL, PostgreSQL, SQLite, SQL Server, Oracle).

**Tujuan dan Fungsi Utama:**
*   **Connection Pooling (Otomatis):** Fungsi `sql.Open` akan merawat banyak koneksi database (Pool) di balik layar. Package ini mengurus koneksi mana yang *idle* dan mana yang sedang terpakai agar performa *query* maksimal.
*   **Pengeksekusian Perintah:** Membaca (Query) dengan `db.Query` yang mengembalikan banyak row, atau sekadar menulis / mengubah data (DDL/DML) dengan `db.Exec`.
*   **Scanning Data:** Memetakan (Scan) setiap kolom *row* di database yang terpilih ke dalam struktur variabel atau tipe di Golang (`rows.Scan`).
*   **Prepared Statements:** Sangat memudahkan penggunaan *Parameterized Query* (`?` atau `$1`) untuk keamanan *anti SQL Injection*.
*   **Transaksi:** Mendukung memulai, melakukan *Commit*, atau *Rollback* sebuah Transaksi Relasional (ACID) secara terpadu melalui fungsi `db.BeginTx`.

**Mengapa menggunakan `database/sql`?**
Jika Anda mengembangkan *microservice* tanpa menggunakan ORM (Object Relational Mapping) raksasa dan lebih suka mengontrol *raw SQL queries*, package ini wajib dipahami. Karena menggunakan standar antarmuka yang sama, jika besok Anda bermigrasi dari SQLite ke Postgres, kode Go yang ditulis nyaris tidak perlu banyak perubahan.

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/database/sql](https://pkg.go.dev/database/sql)
