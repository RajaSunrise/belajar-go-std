# Modul: `html/template`

## Ringkasan
Package `html/template` memproduksi kode HTML secara dinamis dari data yang aman terhadap *Code Injection*.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `html/template` digerakkan oleh data, artinya ini adalah mesin yang menggabungkan desain statis/template (`HTML`) dengan data program dari Go (`struct` / `map`), kemudian memuntahkan hasil akhir halaman Web yang lengkap. (Package saudaranya `text/template` digunakan selain untuk HTML).

**Tujuan dan Fungsi Utama:**
*   **Merender UI Web (Server Side Rendering):** Membangun UI secara klasik tanpa butuh React atau Vue.
*   **Pengikatan Data (Data Binding):** Data struct dari aplikasi diletakkan ke dalam penanda aksi (seperti tanda `{{.Title}}` atau `{{.Nama}}`).
*   **Kontrol Alur (Control Flow):** Anda dapat menulis percabangan dasar IF-ELSE, atau melakukan *Looping* Array (seperti fungsi `range`) langsung di dalam template.
*   **Keamanan Ekstra (Auto Escaping):** Ini fitur paling mutakhir; Package ini paham konteks HTML (HTML Contextual Autoescaping). Jika user mengisikan script jahat (`<script>alert(1)</script>`) ke *form* nama mereka, `html/template` akan mensterilkan atau *escape* *tags* tersebut secara otomatis (menjadi `&lt;script&gt;`). Ini melindungi situs dari serangan *Cross Site Scripting* (XSS) berbahaya tanpa upaya dari pemrogram.

**Mengapa menggunakan `html/template`?**
Untuk mengembangkan panel admin kecil, situs web yang bergantung pada performa SEO, rendering formulir statis, atau memproduksi dokumen/surat elektronik (Email) yang bergaya HTML rapi. Semua ini dapat dilakukan murni dari dalam program Go.

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run main.go
```
*(Catatan: Anda harus berada di dalam direktori modul ini atau menyertakan path penuh file tersebut).*

## Referensi Dokumentasi Resmi
Untuk informasi lebih detail mengenai fungsi dan tipe apa saja yang tersedia, silakan merujuk pada dokumentasi Go (Golang) resmi di:
[https://pkg.go.dev/html/template](https://pkg.go.dev/html/template)
