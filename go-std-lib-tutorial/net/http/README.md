# Modul: `net/http`

## Ringkasan
Package `net/http` menyediakan implementasi klien dan server HTTP yang tangguh dan ringan.

## Penjelasan Lengkap (Fungsi & Tujuan)
Package `net/http` adalah "Jantung" atau senjata utama bahasa Go di era *cloud-native*. Package ini memungkinkan Go untuk membuat Web Server berkinerja amat tinggi, atau melakukan panggilan Web (HTTP API Call) ke server lain **tanpa memerlukan framework eksternal** sama sekali (seperti Express, Django, atau Spring Boot). Pustaka standar Go ini sudah teruji tangguh menghadapi *traffic* produksi.

**Tujuan dan Fungsi Utama:**
*   **Membuat Web Server Terkelola:** Menggunakan `http.ListenAndServe`, Anda dapat menghidupkan *port* server untuk melayani jutaan HTTP *Request* dari browser atau aplikasi *mobile* sekaligus karena setiap koneksi yang masuk otomatis ditangani oleh sebuah *Goroutine* baru (Sangat Concurrent dan Ringan!).
*   **Routing API (ServeMux):** Mendefinisikan rute atau *endpoint URL* yang akan ditangani oleh sebuah fungsi spesifik (Contoh: request ke `/api/login` akan dialihkan ke fungsi `HandleLogin()`). Sejak Go 1.22, *ServeMux* telah berevolusi dan mendukung **method-based routing** secara bawaan (bisa membedakan `GET /items/{id}` dan `DELETE /items/{id}` secara langsung, mengambil variabel {id} dari path-nya).
*   **HTTP Client Cepat:** Untuk aplikasi yang bekerja sebagai perantara (*Microservice / BFF*), Anda dapat mengirim Data *Request* (GET, POST, PUT, DELETE) ke sistem lain via jaringan internet menggunakan objek `http.Client`.
*   **Membaca/Menulis Request dan Response:** Menyediakan struktur untuk membedah *Header*, *Query Parameter* URL, *Cookie*, isi Formulir Multipart, dan isi utuh (Body) dari sebuah *Request* masuk.

**Mengapa menggunakan `net/http`?**
Jika Anda sedang membuat API backend, Anda **harus** menguasainya. Bahkan jika kelak Anda menggunakan *framework* populer pihak ketiga di Go seperti Gin, Fiber, atau Echo, pemahaman tentang paket ini wajib karena hampir seluruh *framework* tersebut dibangun di atas objek standar `http.ResponseWriter` dan `*http.Request` bawaan `net/http`.

## Daftar Fungsi Umum dan Cara Penggunaannya

### 1. Membuat Server HTTP Sederhana
Membuat web server yang merespons teks polos atau JSON kepada klien. Objek `http.ResponseWriter` berfungsi sebagai wadah di mana Anda menulis teks balasan. Objek `*http.Request` berisi semua informasi yang dikirim oleh klien (pengunjung web).
```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    // Menyetel header agar klien tahu balasan ini adalah JSON
    w.Header().Set("Content-Type", "application/json")

    // Memberikan kode status 200 OK
    w.WriteHeader(http.StatusOK)

    // Mengirim response
    w.Write([]byte(`{"pesan": "Selamat datang di API Go!"}`))
}

func main() {
    // Daftarkan route ke handler
    http.HandleFunc("/", homeHandler)

    fmt.Println("Server berjalan di http://localhost:8080")
    // Jalankan server yang "memblokir/mendengarkan" selamanya pada port 8080
    // log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### 2. Fitur Baru Go 1.22: Method & Path Variables
Sebelum Go 1.22, developer terpaksa menggunakan pustaka *router* pihak ketiga (seperti `gorilla/mux` atau `chi`) hanya untuk membaca ID dari URL (seperti `/users/123`). Sekarang, semua itu didukung bawaan!
```go
mux := http.NewServeMux()

// Hanya melayani method POST pada URL /users
mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("User berhasil dibuat!"))
})

// Melayani method GET dan mengekstrak Path Variable "id"
mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
    // PathValue langsung mengambil string ID dari URL
    userID := r.PathValue("id")
    pesan := fmt.Sprintf("Anda meminta profil user dengan ID: %s", userID)
    w.Write([]byte(pesan))
})

// http.ListenAndServe(":8080", mux)
```

### 3. Menggunakan HTTP Client (Mengambil Data dari Luar)
Go membuat pengambilan data dari API luar menjadi satu baris kode dengan `http.Get`. Harap diingat bahwa kita **wajib menutup body response** (`defer resp.Body.Close()`) agar memori jaringan tidak bocor!
```go
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
```

### 4. Membuat HTTP Client Kustom (Dengan Timeout)
Fungsi `http.Get()` di atas tidak memiliki *timeout*. Jika API pihak luar sedang bermasalah dan tidak merespons selama 10 menit, maka program Anda juga akan macet total 10 menit. Selalu gunakan *Custom Client* di ranah produksi!
```go
// Membuat klien dengan batas waktu respon 5 detik
client := &http.Client{
    Timeout: 5 * time.Second,
}

req, err := http.NewRequest("GET", "https://api.github.com", nil)
// Menambahkan header khusus ke request kita
req.Header.Add("Authorization", "Bearer token-rahasia")

// Melakukan pemanggilan ke server
resp, err := client.Do(req)
// ... dan seterusnya ...
```


## Menjalankan Contoh
Silakan jalankan contoh kode simulasi di dalam folder ini dengan perintah:
```bash
go run main.go
```
