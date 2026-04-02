# Modul: `html/template`

## Ringkasan
Package `html/template` menyediakan mesin perakit templat berbasis data (*Data-Driven Template Engine*) teramat tangguh untuk menyintesis konstruksi balasan teks antarmuka HTML dinamis yang sepenuhnya tersaring aman dan kebal terhadap peretasan injeksi kode berbahaya (*Cross-Site Scripting / XSS Attack*).

## Penjelasan Lengkap (Fungsi & Tujuan)
Di era awal perintisan halaman web dinamis (seperti bahasa PHP lawas), seringkali dijumpai kelemahan fatal: pengembang menyambungkan (*concat*) nama input pengguna secara membabi buta ke dalam teks koding HTML `echo "<h1>Halo " + username + "</h1>"`. Masalah kiamat terjadi ketika seorang peretas (hacker) mengisi namanya dengan teks `<script>curi_password()</script>`. Alih-alih mencetak nama, penjelajah browser secara buta mengeksekusi kode rahasia virus tersebut—sebuah kelemahan mematikan XSS.

Package `html/template` milik Go hadir dengan landasan konseptual revolusioner: "Penyaringan Kesadaran Kontekstual Lanjutan" (*Contextual Auto-Escaping*). Ini bermakna, saat kerangka mesin templat mendaratkan (*inject*) variabel teks nama sang peretas ke dalam templat, program Go cukup cerdas untuk memahami bahwa, *"Oh, variabel ini mendarat di dalam atribut Tag HTML/JavaScript, bukan sekadar teks biasa!"* Kompilator Go akan secara otomatis membungkus mensterilkan teks jahat peretas tersebut dengan entitas amannya menjadi tulisan jinak `&lt;script&gt;`, melumpuhkan serangannya menjadi tak berdaya seketika.

*(Peringatan: Go juga mempunyai package saudara bernama `text/template` dengan sintaks kembar, HANYA gunakan varian `text/` untuk memproduksi surel e-mail teks polos biasa atau merakit file manifest Konfigurasi `yaml`, namun selalu MUTLAK gunakan `html/template` manakala memproduksi output yang dieksekusi oleh peramban Browser Web).*

**Tujuan dan Fungsi Utama:**
1.  **Rendering UI Klasik (*Server Side Rendering/SSR*):** Di tengah membludaknya *Framework* raksasa *Frontend Single Page Application* (seperti *React/Vue*), merakit halaman antarmuka Admin Panel / Sistem Dashboard (*Dashboard Internal*) atau Laman Pendaratan statis (*Landing Page SEO*) konvensional dengan *SSR* mesin murni Golang lebih ringan, super cepat dirilis, dan tidak memerlukan kerumitan ekosistem *Node.js* (*Webpack*).
2.  **Pemetaan Variabel Bawaan Data Cerdas (*Data Binding*):** Mengikutsertakan Objek `Struct` atau Kamus `Map` sebagai suntikan asupan darah sumber data, lalu membongkarnya di bagian depan UI Template menggunakan notasi aksi deklaratif sederhana seperti `{{.JudulBerita}}` atau `{{.NamaBapak}}`.
3.  **Kendali Percabangan & Rombongan Looping Internal:** Templat ini bukanlah teks bodoh statis belaka. Di dalam badan cetakan templat Anda berkesempatan menyusun kendali cabang percabangan logika `{{if .ApakahAktif}}`, maupun memutar mengulang rentetan tabel rekaman daftar pelanggan berkalang deret (*Array Slices*) menggunakan tata bahasa fungsional sakti `{{range .DaftarPeserta}}`.
4.  **Komposisi Fragmen (Templat Bersarang Berserikat):** Memfasilitasi praktik baik rekondisi blok (*DRY - Don't Repeat Yourself*). Anda bisa mendefinisikan blok rangka umum (*Layout Base*) yang menampung komponen Header Navigasi, kemudian komponen Anak halaman utama menyuntikkan isi badannya ke ruang kosong Induknya itu (*Template Inheritance & Includes*).

**Mengapa menggunakan `html/template`?**
Jika sistem peladen Anda dituntut menerbitkan lembar cetak Struk/Faktur Belanja (*Invoice* Digital) yang anggun untuk dilampirkan ke surel PDF klien setiap ada transaksi e-commerce, atau merilis panel *CMS* administrasi supercepat independen 0 kilobyte dependensi untuk operator perusahaan, pengetahuan menyusun rakitan arsitektur *template* terintegrasi Go ini adalah harta kekayaan langka mempesona yang menguntungkan efektivitas karier Anda.

---

## Daftar Fungsi Umum dan Cara Penggunaannya Secara Mendalam

### 1. Deklarasi Fundamental dan Evaluasi Dasar (*Parse & Execute*)

Batu pijakan penyusunan template menuntut dua tahapan tak tepisahkan. Tahapan *Pertama*: Menyiapkan dan "Mencerna/Memparsing" teks templat telanjang menjadi Objek Cetak Struktur Mesin Memori (Seringkali dilaksanakan 1x saja kala Server Go dihidupkan (*Booting*)). Tahapan *Kedua*: Mengisi, "Mengeksekusi" Objek cetakan yang telah hidup tersebut dengan data pasien spesifik klien, sembari memuntahkannya langsung merasuk ke keran aliran lubang TCP antarmuka `io.Writer` (Layar Konsol, File *Disk*, atau langsung ke HTTP *ResponseWriter* Penjelajah Klien Web).

```go
package main

import (
    "html/template"
    "os"
)

// Menyiapkan Kerangka Wadah Struktur Suplai Data Pasokan
type ProfilTokoSitus struct {
    NamaToko string
    Slogan   string
}

func main() {
    // 1. Deklarasi teks Teks Mentah (Menggunakan notasi aksi Tanda Kurung Kurawal Ganda {{ .FieldAjaib }})
    // Tanda titik ( . ) melambangkan Titik Akar Sumber (Yakni Obyek ProfilTokoSitus yang kelak diumpankan)
    teksKerangkaMentah := `
        <h1>Selamat berkunjung menelusuri Katalog kami di {{.NamaToko}}!</h1>
        <p>Motto Kami Hari Ini: <b>{{.Slogan}}</b></p>
    `

    // 2. TAHAP PARSING (Pembuatan Cetakan Mesin Memori Kompilasi)
    // Diwajibkan memberi nama pengenal internal untuk cetakan templat ini (misal "BerandaToko").
    cetakanMesin, errParse := template.New("BerandaToko").Parse(teksKerangkaMentah)
    if errParse != nil {
        panic("Sintaksis Kerangka Kurawal Aksi tak masuk akal / Typo fatal!" + errParse.Error())
    }

    // 3. Menyiapkan Asupan Suntikan Data (Bebas menggunakan Struct murni maupun Peta Map string dinamis)
    dataSuntik := ProfilTokoSitus{
        NamaToko: "Serba Murah Elektronik Sentosa",
        Slogan:   "Kualitas Premium, Harga Subsidi Kerakyatan!",
    }

    // 4. TAHAP EKSEKUSI (Pembuahan Peleburan Templat + Data)
    // Sasaran Akhir Output Kita: Layar Konsole Standar Monitor Terminal (os.Stdout yang merupakan io.Writer Sah).
    errRilisEksekusi := cetakanMesin.Execute(os.Stdout, dataSuntik)

    if errRilisEksekusi != nil {
        panic("Kegagalan memuntahkan rakitan evaluasi penyatuan template")
    }
}
```

---

## Bagian Lanjutan: Injeksi Keamanan Fungsi Khusus (*Custom FuncMaps*), Sanitasi Aman (*Safe HTML/JS*), dan Kecepatan Pemrosesan Kompilasi Memori

Keunggulan telak `html/template` dibanding pustaka pesaing di bahasa dinamis bukan terletak pada struktur *Layouting* bersarang `{{define}}` belaka, namun pada kendali ekstensi pemrosesan cetakan logikanya (*Custom Action Functions*). Di bagian ini, kita akan membongkar bagaimana memperluas perbendaharaan fungsi templat hingga batas akhir kemampuan sistem, merakit Filter Keamanan Teks kustom, serta cara mengelabui pelindung *Auto-Escaping* secara hati-hati bila sistem kita bersikeras menuntut injeksi *String DOM murni*.

### 1. Menyuntikkan Kemampuan Bahasa Pemrograman ke dalam Template (`template.FuncMap`)

Secara *Default*, saat Anda sedang mendesain struktur kerangka kurawal ganda `{{.HargaBeli}}`, templat hanya dapat membaca dan mencetak properti angka bulat mentah (*misal: 15000000*). Bagaimana bila Manajer Pemasaran Anda memohon: "Tolong tampilkan angka jelek itu di layar HTML dengan wujud pemisah ribuan Rupiah yang indah, seperti **Rp 15.000.000**!" ?

Sintaks HTML templat tidak memiliki fungsi bawaan bawaan untuk memanipulasi format teks mata uang kompleks. Solusinya: Anda HAKUL YAKIN harus menyuntikkan (mendaftarkan) fungsi buatan Go ciptaan Anda sendiri ke dalam alam rahim memori sang Mesin Templat SEBELUM templat itu mencerna string HTML (sebelum `.Parse`). Inilah fitur penyeberang dimensi **`FuncMap`**.

```go
package main

import (
    "fmt"
    "html/template"
    "os"
)

// 1. BUAT FUNGSI MANIPULASI GOLANG KUSTOM DI ALAM NYATA
// Fungsi ini menelan angka, memuntahkan Teks Format Cantik
func konversiFormatRupiah(uangMasuk int) string {
    // (Simulasi Sederhana Format Ribuan) Di dunia nyata Anda pakai package "golang.org/x/text/language"
    return fmt.Sprintf("Rp %d.00", uangMasuk)
}

func main() {
    // 2. PENDAFTARAN KAMUS FUNGSI (FuncMap)
    // Berisi "NamaPanggilanDiTemplate" berpasangan dengan "AlamatFungsiGoAsli"
    kamusFungsiSakti := template.FuncMap{
        "CetakDuitIDR": konversiFormatRupiah, // Mendaftarkan fungsi custom kita!
    }

    // 3. Merancang Kerangka Teks Web (HTML String)
    // KITA SEKARANG BISA MENGGUNAKAN SIMBOL PIPA ( | ) UNTUK MENGIRIM VARIABEL HARGA MASUK KE DALAM FUNGSI CetakDuitIDR !
    teksDesain := `
    === RINCIAN TAGIHAN KERANJANG ===
    Barang Anda: {{.NamaItem}}

    Harga Awal Mentah (Jelek): {{.Nominal}}
    Harga Cetak Cantik (Fungsi): {{.Nominal | CetakDuitIDR}}
    =================================
    `

    // 4. ATURAN WAJIB MENGKOMPILASI TEMPLATE DENGAN FUNGSI (Funcs MUST BE CALLED BEFORE Parse!)
    mesinRakitDenganFungsi := template.New("PencetakFaktur")

    // TEMPELKAN KAMUS FUNGSI KITA DULUAN!
    mesinRakitDenganFungsi = mesinRakitDenganFungsi.Funcs(kamusFungsiSakti)

    // Baru Boleh Mem-Parsing Teks HTML nya.
    template.Must(mesinRakitDenganFungsi.Parse(teksDesain))

    // Eksekusi Panggilan Ujicoba Cetak Resi
    dataPembeli := struct{ NamaItem string; Nominal int }{
        NamaItem: "Laptop Gaming Generasi 12",
        Nominal:  24500000,
    }

    mesinRakitDenganFungsi.Execute(os.Stdout, dataPembeli)
}
```
Hasil kemegahan dari trik Pipa (`|`) layaknya Terminal Linux ini memungkinkan perantaian tak berujung (contoh: `{{.Tulisan | HurufKecilkan | HilangkanSpasi | Potong10Karakter}}`), menyulap templat Anda menjadi cerdas dan berdikari dari logika lapis bawah kode Go Server yang ruwet.

### 2. Mematikan Pelindung Sanitasi Secara Sadar Diri (Tipe `template.HTML`)

Kadangkala (sangat langka dan berbahaya), arsitektur menuntut Peladen Go membaca Artikel Blog Postingan *Berita* (yang sudah dikodekan berisi cetakan *Bold* `<b>`, *Miring* `<i>` dan *Gambar* `<img>` aman dari Database SQL), lalu menaruhnya murni telanjang tampil merender grafis di laman Klien.

Sistem *Auto-Escaping* Go otomatis akan mengubah tanda kurung HTML `<b>` berita kita menjadi teks culun cacat rupa: `&lt;b&gt;Berita Terkini&lt;/b&gt;`. Ia mencurigainya sebagai serangan *hacker*, padahal itu HTML resmi milik tim Admin kita!

Anda harus memberikan "*Surat Izin Bebas Tilang / Bebas Cek*" kepada mesin Templat, memaksanya menelan teks itu mentah-mentah.
Kuncinya adalah mengubah Tipe Data variabel `string` Anda menjadi tipe sakti **`template.HTML`** (atau `template.JS` khusus script, `template.URL` khusus *href* tautan).

```go
package main

import (
    "html/template"
    "os"
)

type BeritaPortal struct {
    JudulBesar     string
    IsiArtikelTeks string        // Bakal Dibungkus/Dikebiri oleh Mesin Go (Aman namun merusak tag B/I/U)
    IsiArtikelHTML template.HTML // DILARANG MENYENTUHNYA! Go akan meloloskannya langsung ke Browser Web Klien!
}

func main() {
    desainCMS := `
    JUDUL   : {{.JudulBesar}}
    VERSI 1 : {{.IsiArtikelTeks}}
    VERSI 2 : {{.IsiArtikelHTML}}
    `

    mesinBerita := template.Must(template.New("BeritaKu").Parse(desainCMS))

    // Input Identik (Keduanya berisi Tag Cetak Tebal <b> dan Cetak Miring <i>)
    kontenTeksTag := "Ini adalah cuplikan <b>BENCANA BESAR</b> di daerah <i>Pegunungan</i>."

    suntikanData := BeritaPortal{
        JudulBesar:     "Laporan Siang",
        IsiArtikelTeks: kontenTeksTag,                           // String murni biasa
        IsiArtikelHTML: template.HTML(kontenTeksTag), // SULAP STRING JADI TIPE SPESIAL!
    }

    // Mari Lihat Perbedaan Reaksi Perlakuannya!
    mesinBerita.Execute(os.Stdout, suntikanData)
    // Di Versi 1: Ia dicetak jelek (Termasuk tulisan tagnya "&lt;b&gt;" terlihat ke pembaca)
    // Di Versi 2: Ia dibiarkan menjadi HTML DOM sungguhan (<b>BENCANA BESAR</b>) yang akan membuat tulisan di Layar Browser tercetak Tebal sempurna!
}
```

*PERINGATAN KERAS KEAMANAN:* Menjejalkan *input string formulir dari luar yang diisi oleh klien tamu sembarangan* ke dalam pembungkus tipe sakti `template.HTML(...)` adalah pelanggaran SOP Keamanan Tingkat Kritis! Fitur *bypass* ini eksklusif dan HANYA direstui apabila Teks HTML yang disuntikkan tersebut murni 100% diproduksi aman dari dalam peladen lokal perusahaan Anda (*Trusted Source Origin*).

Penggabungan pengetahuan modul utilitas internal Peta Fungsi Kustom (`FuncMap`) ini melegitimasi *Go* sebagai pilar kokoh arsitektur Peladen Penuh (Full-Stack Monolith Server Web). Ia menjinakkan kekacauan presentasi data visual HTML peramban muka dan melindunginya dengan benteng sanitasi XSS tebal asali, semuanya ditampung ke dalam sebongkah Berkas *Executable* Biner (Binari Go) murni yang mandiri tanpa kebergantungan perangkat lunak pihak ketiga penyedia layanan templat sama sekali.

## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```

---

## Studi Kasus Dunia Nyata: Arsitektur Modular Templat Dinamis Penyelamatan Redundansi Halaman (Server-Side Component Layouting)

Sewaktu Anda membangun aplikasi Web Portal Administrasi (seperti Aplikasi Kasir POS Toko), Anda takkan mungkin mendesain ulang komponen Kode HTML `Sidebar`, `Header Navigation`, serta pemanggilan skrip `JavaScript/CSS` yang berulang persis sama panjang di 50 keping file `.html` laman Anda yang berbeda.

Insinyur Go tulen menangani kerepotan ini dengan memelopori desain "Keping Rangka Payung Induk" (*Master Base Template Layouting*).
Konsepnya: Kita mendirikan HANYA 1 BUAH cetakan rangka rumah kosong `master.html` induk, yang mana di dinding ruang tamu kosong HTML itu kita tinggalkan lubang pemanggil berongga (*Action Hooks*): `{{template "Isi_Konten_Utama" .}}`. Lantas, bila seseorang membuka rute halaman *Dashboard* web, kita gabungkan paksakan leburkan (ParseFiles) File Induk itu dengan file `dashboard.html`, secara ajaib Go akan mengisi rongga kosong Sang Induk dengan isian komponen perut Anak tersebut!

### Rancang Bangun Peladen Templat Canggih Multi-Halaman (*Multi-Page Composition*)

```go
package main

import (
    "bytes"
    "fmt"
    "html/template"
    "os"
)

func main() {
    // 1. FILE MASTER BASE INDUK (Bayangkan ini disimpan di fisik disk file "views/base_master.html")
    // Ia memegang kendali wujud tulang rangka utuh Situs Web dari ujung Kepala <html> hingga Ekor </html>
    kerangkaPusatInduk := `
        <!DOCTYPE html>
        <html lang="id">
        <head>
            <meta charset="UTF-8">
            <title>{{.SitusNamaGlobal}} - Portal Pegawai</title>
            <style> body { font-family: Arial; background: #fafafa; } </style>
        </head>
        <body>
            <nav style="background: black; color: white; padding: 10px;">
                <b>Panel Canggih | Selamat datang, Tuan Pengawas!</b>
            </nav>

            <main style="margin: 20px;">
                <!-- DI SINILAH KEAJAIBAN BERSARANG! -->
                <!-- Ini adalah lubang kunci, menunggu Komponen Anak menyuntikkan badannya ke titik ini! -->
                <!-- Perhatikan titik di ekor (.), artinya kita melempar Meneruskan variabel Data Golang turun ke Anaknya -->
                {{template "RUANG_TAMU_UTAMA" .}}
            </main>

            <footer style="text-align:center; margin-top:50px;">
                <small>&copy; 2025 Divisi Rekayasa Perangkat Lunak Internal Corp</small>
            </footer>
        </body>
        </html>
    `

    // 2. FILE LAMAN ANAK HALAMAN DEPAN (Bayangkan file fisik "views/beranda.html")
    // File Anak ini CUMA BERISI KUMPULAN DEKLARASI {{define}}. Ia tak peduli tag html dasar!
    anakHalamanBeranda := `
        {{define "RUANG_TAMU_UTAMA"}}
            <h2>Metriks Statistik Penjualan Global</h2>
            <p>Sistem ini sedang berjalan di Mode Produksi.</p>
            <ul>
                <li>Total Profit: Rp {{.TotalUang}}</li>
                <li>Status: Sehat.</li>
            </ul>
        {{end}}
    `

    // 3. FILE LAMAN ANAK HALAMAN SETTING (Bayangkan file "views/pengaturan.html")
    anakHalamanPengaturan := `
        {{define "RUANG_TAMU_UTAMA"}}
            <h2>Ruang Ganti Sandi Pegawai</h2>
            <input type="text" placeholder="Ketik Sandi Usang..."/>
            <button>Perbaharui Sandi</button>
        {{end}}
    `

    // TAHAPAN PEMBANGUNAN MESIN MEMORI (Di Dunia Nyata, ParseFiles dipanggil 1 kali saat Server Baru Menyala!)

    // Mesin 1: Menggabungkan Induk dengan Laman Beranda
    mesinRuteBeranda := template.Must(template.New("BASE_BERANDA").Parse(kerangkaPusatInduk))
    template.Must(mesinRuteBeranda.Parse(anakHalamanBeranda))

    // Mesin 2: Menggabungkan Induk dengan Laman Pengaturan (Setup Baru)
    mesinRuteSeting := template.Must(template.New("BASE_PENGATURAN").Parse(kerangkaPusatInduk))
    template.Must(mesinRuteSeting.Parse(anakHalamanPengaturan))

    // DATA ASUPAN PELAYANAN
    dataTamuBeranda := struct{ SitusNamaGlobal string; TotalUang int }{
        SitusNamaGlobal: "Panel Bisnis Super",
        TotalUang:       99000000,
    }

    // 4. MENSIMULASIKAN PENGIRIMAN TEMBAKAN KE BROWSER WEB
    fmt.Println("=== MENGHASILKAN HALAMAN RENDER BERANDA ===")
    // Pada implementasi server HTTP, os.Stdout ini akan diganti oleh parameter w (http.ResponseWriter)
    mesinRuteBeranda.Execute(os.Stdout, dataTamuBeranda)

    fmt.Println("\n\n=== MENGHASILKAN HALAMAN RENDER PENGATURAN ===")
    mesinRuteSeting.Execute(os.Stdout, dataTamuBeranda) // Tetap oper Data, jika dia butuh variabel SitusNama
}
```

Pengoperasian terstruktur modul Templat Multi-Fasad (Multi-Facets Template Rendering) semacam ini tidak semata memuluskan langkah peladen menyajikan arsitektur HTML Dinamis (*Classic Web Architecture*), namun secara hakikat, kepatuhan atas pemisahan Modul Presentasional dan Operasional Logik Bisnis murni di punggung Server Biner Go Anda akan mempesonakan kancah kualitas ketahanan perangkat lunak perusahan raksasa terpercaya tingkat yurisdiksi kelas tinggi.
