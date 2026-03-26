import os
import textwrap

# Dictionary mapping module path to a tuple of (description in Indonesian, example filename, example code)
modules = {
    "fmt": (
        "Package `fmt` mengimplementasikan I/O berformat dengan fungsi-fungsi yang mirip dengan `printf` dan `scanf` pada bahasa C.",
        "main.go",
        """package main

import "fmt"

func main() {
\tname := "Go Developer"
\t// Menggunakan verb %s untuk string dan %d untuk integer
\tfmt.Printf("Halo, %s! Selamat datang di Go %d.\\n", name, 1)
\t
\t// Menggabungkan string
\tmsg := fmt.Sprintf("Ini adalah string hasil format: %s", name)
\tfmt.Println(msg)
}
"""
    ),
    "os": (
        "Package `os` menyediakan antarmuka independen-platform untuk fungsionalitas sistem operasi, seperti membaca file, menulis file, dan variabel environment.",
        "main.go",
        """package main

import (
\t"fmt"
\t"log"
\t"os"
)

func main() {
\t// Mengambil variabel environment
\tpath := os.Getenv("PATH")
\tfmt.Println("Isi dari variabel PATH:", path)

\t// Membuat file baru
\tfile, err := os.Create("contoh.txt")
\tif err != nil {
\t\tlog.Fatal(err)
\t}
\tdefer file.Close()
\tdefer os.Remove("contoh.txt") // Hapus file setelah selesai

\t// Menulis ke file
\tfile.WriteString("Halo dari package os!\\n")
\tfmt.Println("Berhasil menulis ke file contoh.txt")
}
"""
    ),
    "math": (
        "Package `math` menyediakan konstanta dasar dan fungsi matematika.",
        "main.go",
        """package main

import (
\t"fmt"
\t"math"
)

func main() {
\tfmt.Println("Nilai Pi:", math.Pi)

\t// Mencari nilai akar kuadrat
\tfmt.Println("Akar dari 16 adalah:", math.Sqrt(16))

\t// Membulatkan angka
\tfmt.Println("Ceil 3.14:", math.Ceil(3.14))
\tfmt.Println("Floor 3.14:", math.Floor(3.14))
\tfmt.Println("Round 3.14:", math.Round(3.14))

\t// Nilai absolut
\tfmt.Println("Absolut -10:", math.Abs(-10))
}
"""
    ),
    "strings": (
        "Package `strings` mengimplementasikan fungsi-fungsi sederhana untuk memanipulasi string UTF-8.",
        "main.go",
        """package main

import (
\t"fmt"
\t"strings"
)

func main() {
\ttext := "belajar golang standard library"

\t// Mengubah ke huruf kapital
\tfmt.Println("ToTitle:", strings.ToTitle(text))
\tfmt.Println("ToUpper:", strings.ToUpper(text))

\t// Mengecek apakah string mengandung substring tertentu
\tfmt.Println("Contains 'golang':", strings.Contains(text, "golang"))

\t// Menghitung jumlah kemunculan
\tfmt.Println("Count 'a':", strings.Count(text, "a"))

\t// Memisahkan string menjadi slice
\tsplitted := strings.Split(text, " ")
\tfmt.Printf("Split: %#v\\n", splitted)

\t// Menggabungkan slice menjadi string
\tjoined := strings.Join(splitted, "-")
\tfmt.Println("Join:", joined)
}
"""
    ),
    "time": (
        "Package `time` menyediakan fungsionalitas untuk mengukur dan menampilkan waktu.",
        "main.go",
        """package main

import (
\t"fmt"
\t"time"
)

func main() {
\t// Waktu saat ini
\tnow := time.Now()
\tfmt.Println("Waktu saat ini:", now)

\t// Format waktu (Go menggunakan waktu referensi khusus: Mon Jan 2 15:04:05 MST 2006)
\tfmt.Println("Format (YYYY-MM-DD):", now.Format("2006-01-02"))

\t// Menambahkan waktu
\ttomorrow := now.Add(24 * time.Hour)
\tfmt.Println("Besok:", tomorrow)

\t// Parsing waktu
\tparsedTime, _ := time.Parse("2006-01-02", "2024-08-17")
\tfmt.Println("Hari Kemerdekaan 2024:", parsedTime)

\t// Sleep / Jeda eksekusi
\tfmt.Println("Tunggu 1 detik...")
\ttime.Sleep(1 * time.Second)
\tfmt.Println("Selesai!")
}
"""
    ),
    "net/http": (
        "Package `net/http` menyediakan implementasi klien dan server HTTP.",
        "main.go",
        """package main

import (
\t"fmt"
\t"io"
\t"log"
\t"net/http"
)

func main() {
\t// Menggunakan ServeMux (Go 1.22+ mendukung method based routing)
\tmux := http.NewServeMux()
\t
\tmux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
\t\tfmt.Fprintln(w, "Halo, selamat datang di server Go!")
\t})

\tmux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
\t\tname := r.PathValue("name")
\t\tfmt.Fprintf(w, "Halo, %s!\\n", name)
\t})

\tfmt.Println("Server berjalan di port 8080...")
\t// Uncomment baris di bawah untuk menjalankan server
\t// log.Fatal(http.ListenAndServe(":8080", mux))

\t// Contoh HTTP Client sederhana
\tresp, err := http.Get("https://httpbin.org/get")
\tif err != nil {
\t\tlog.Fatal(err)
\t}
\tdefer resp.Body.Close()

\tbody, _ := io.ReadAll(resp.Body)
\tfmt.Println("Status Code:", resp.StatusCode)
\tfmt.Println("Response Body (dipotong):", string(body[:50]), "...")
}
"""
    ),
    "io": (
        "Package `io` menyediakan antarmuka dasar untuk I/O (Input/Output).",
        "main.go",
        """package main

import (
\t"fmt"
\t"io"
\t"log"
\t"os"
\t"strings"
)

func main() {
\t// Menggunakan io.Reader (membaca dari string)
\treader := strings.NewReader("Ini adalah contoh penggunaan io.Reader\\n")
\t
\t// Menggunakan io.Writer (menulis ke os.Stdout)
\t_, err := io.Copy(os.Stdout, reader)
\tif err != nil {
\t\tlog.Fatal(err)
\t}

\t// io.ReadAll untuk membaca seluruh isi dari Reader
\treader2 := strings.NewReader("Membaca semua data sekaligus.")
\tdata, err := io.ReadAll(reader2)
\tif err != nil {
\t\tlog.Fatal(err)
\t}
\tfmt.Println("Hasil ReadAll:", string(data))
}
"""
    ),
    "context": (
        "Package `context` mendefinisikan tipe Context, yang membawa deadline, sinyal pembatalan, dan nilai terkait permintaan lainnya melintasi batas API dan antar proses.",
        "main.go",
        """package main

import (
\t"context"
\t"fmt"
\t"time"
)

func main() {
\t// Membuat context dasar
\tctx := context.Background()

\t// Menambahkan value ke context
\tctxWithValue := context.WithValue(ctx, "userID", 12345)
\tfmt.Println("UserID dari context:", ctxWithValue.Value("userID"))

\t// Membuat context dengan pembatalan (Cancel)
\tctxCancel, cancel := context.WithCancel(ctx)
\t
\tgo func() {
\t\ttime.Sleep(2 * time.Second)
\t\tcancel() // Membatalkan operasi setelah 2 detik
\t}()

\tselect {
\tcase <-time.After(3 * time.Second):
\t\tfmt.Println("Proses selesai tanpa dibatalkan")
\tcase <-ctxCancel.Done():
\t\tfmt.Println("Proses dibatalkan:", ctxCancel.Err())
\t}

\t// Membuat context dengan timeout
\tctxTimeout, cancelTimeout := context.WithTimeout(ctx, 1*time.Second)
\tdefer cancelTimeout()

\tselect {
\tcase <-time.After(2 * time.Second):
\t\tfmt.Println("Proses lambat selesai")
\tcase <-ctxTimeout.Done():
\t\tfmt.Println("Proses timeout:", ctxTimeout.Err())
\t}
}
"""
    ),
    "encoding/json": (
        "Package `encoding/json` mengimplementasikan encoding dan decoding JSON yang didefinisikan dalam RFC 7159.",
        "main.go",
        """package main

import (
\t"encoding/json"
\t"fmt"
\t"log"
)

// User merepresentasikan data pengguna
type User struct {
\tID       int      `json:"id"`
\tName     string   `json:"name"`
\tEmail    string   `json:"email"`
\tIsActive bool     `json:"is_active"`
\tRoles    []string `json:"roles,omitempty"` // omitempty: diabaikan jika kosong
}

func main() {
\t// 1. Marshal: Mengubah struct/map menjadi JSON string (byte slice)
\tuser := User{
\t\tID:       1,
\t\tName:     "Budi",
\t\tEmail:    "budi@example.com",
\t\tIsActive: true,
\t\tRoles:    []string{"admin", "user"},
\t}

\tjsonData, err := json.MarshalIndent(user, "", "  ")
\tif err != nil {
\t\tlog.Fatal(err)
\t}
\tfmt.Println("Hasil Marshal (Struct ke JSON):\\n", string(jsonData))

\t// 2. Unmarshal: Mengubah JSON string menjadi struct/map
\tjsonString := `{"id":2,"name":"Siti","email":"siti@example.com","is_active":false}`
\tvar newUser User

\terr = json.Unmarshal([]byte(jsonString), &newUser)
\tif err != nil {
\t\tlog.Fatal(err)
\t}
\tfmt.Printf("\\nHasil Unmarshal (JSON ke Struct): %+v\\n", newUser)
}
"""
    ),
    "sync": (
        "Package `sync` menyediakan primitive sinkronisasi dasar seperti mutexes saling eksklusif.",
        "main.go",
        """package main

import (
\t"fmt"
\t"sync"
)

// Counter aman untuk konkurensi (thread-safe)
type Counter struct {
\tmu    sync.Mutex
\tvalue int
}

func (c *Counter) Increment() {
\tc.mu.Lock()   // Kunci sebelum memodifikasi
\tc.value++
\tc.mu.Unlock() // Buka kunci setelah memodifikasi
}

func (c *Counter) Value() int {
\tc.mu.Lock()
\tdefer c.mu.Unlock() // defer memastikan Unlock selalu dipanggil
\treturn c.value
}

func main() {
\tvar wg sync.WaitGroup
\tcounter := Counter{}

\t// Menjalankan 1000 goroutine untuk increment counter
\tfor i := 0; i < 1000; i++ {
\t\twg.Add(1)
\t\tgo func() {
\t\t\tdefer wg.Done()
\t\t\tcounter.Increment()
\t\t}()
\t}

\t// Menunggu semua goroutine selesai
\twg.Wait()

\tfmt.Println("Nilai akhir counter:", counter.Value()) // Seharusnya 1000
}
"""
    ),
    "errors": (
        "Package `errors` mengimplementasikan fungsi-fungsi untuk memanipulasi error.",
        "main.go",
        """package main

import (
\t"errors"
\t"fmt"
)

// Membuat error custom sebagai variabel global
var ErrNotFound = errors.New("data tidak ditemukan")

func findData(id int) (string, error) {
\tif id == 0 {
\t\t// Menggunakan fmt.Errorf dengan %w untuk membungkus (wrap) error
\t\treturn "", fmt.Errorf("pencarian id %d gagal: %w", id, ErrNotFound)
\t}
\treturn "Data Rahasia", nil
}

func main() {
\t// Error biasa
\terr1 := errors.New("ini adalah error sederhana")
\tfmt.Println("Error:", err1)

\t// Membungkus dan mengecek error
\tdata, err := findData(0)
\tif err != nil {
\t\tfmt.Println("Terjadi kesalahan:", err)

\t\t// errors.Is digunakan untuk mengecek apakah suatu error merupakan/membungkus error tertentu
\t\tif errors.Is(err, ErrNotFound) {
\t\t\tfmt.Println("=> Penanganan khusus: Menampilkan halaman 404 Not Found")
\t\t}
\t} else {
\t\tfmt.Println("Data:", data)
\t}

\t// errors.As digunakan untuk mengekstrak tipe error custom tertentu (jika ada)
}
"""
    ),
    "sort": (
        "Package `sort` menyediakan primitif untuk mengurutkan slice dan koleksi yang didefinisikan pengguna.",
        "main.go",
        """package main

import (
\t"fmt"
\t"sort"
)

// Person struct
type Person struct {
\tName string
\tAge  int
}

func main() {
\t// Mengurutkan slice of int
\tnumbers := []int{5, 2, 7, 1, 9, 3}
\tsort.Ints(numbers)
\tfmt.Println("Sorted numbers:", numbers)

\t// Mengurutkan slice of string
\tfruits := []string{"pisang", "apel", "mangga", "jeruk"}
\tsort.Strings(fruits)
\tfmt.Println("Sorted fruits:", fruits)

\t// Mengurutkan slice of struct menggunakan sort.Slice (Go 1.8+)
\tpeople := []Person{
\t\t{"Budi", 25},
\t\t{"Andi", 30},
\t\t{"Citra", 22},
\t}

\t// Urutkan berdasarkan umur secara ascending
\tsort.Slice(people, func(i, j int) bool {
\t\treturn people[i].Age < people[j].Age
\t})
\tfmt.Println("Sorted people by age:", people)

\t// Mengecek apakah sudah berurutan
\tisSorted := sort.SliceIsSorted(people, func(i, j int) bool {
\t\treturn people[i].Age < people[j].Age
\t})
\tfmt.Println("Apakah people sudah diurutkan berdasarkan usia?", isSorted)
}
"""
    ),
    "os/exec": (
        "Package `exec` menjalankan perintah eksternal. Ini membungkus os.StartProcess untuk membuatnya lebih mudah di-remap.",
        "main.go",
        """package main

import (
\t"fmt"
\t"log"
\t"os/exec"
)

func main() {
\t// Menjalankan perintah 'ls -la' (atau 'dir' di Windows)
\t// Catatan: contoh ini berasumsi OS adalah varian Unix/Linux/macOS
\tcmd := exec.Command("ls", "-la")
\t
\t// Mengambil output dari command
\toutput, err := cmd.Output()
\tif err != nil {
\t\t// Jika gagal, coba perintah Windows 'cmd /c dir'
\t\tcmd = exec.Command("cmd", "/c", "dir")
\t\toutput, err = cmd.Output()
\t\tif err != nil {
\t\t\tlog.Fatal(err)
\t\t}
\t}

\tfmt.Println("Output perintah:")
\tfmt.Println(string(output))
}
"""
    ),
    "path/filepath": (
        "Package `filepath` mengimplementasikan fungsi utilitas untuk memanipulasi jalur nama file (filepath) yang kompatibel dengan sistem operasi target.",
        "main.go",
        """package main

import (
\t"fmt"
\t"path/filepath"
)

func main() {
\t// Menggabungkan path dengan aman (menghindari masalah slash/backslash antar OS)
\tpath := filepath.Join("users", "admin", "documents", "file.txt")
\tfmt.Println("Joined Path:", path)

\t// Mendapatkan direktori dan nama file
\tdir := filepath.Dir(path)
\tfile := filepath.Base(path)
\tfmt.Println("Directory:", dir)
\tfmt.Println("Filename:", file)

\t// Mendapatkan ekstensi file
\text := filepath.Ext(path)
\tfmt.Println("Extension:", ext)

\t// Mengecek apakah path absolute
\tfmt.Println("Is absolute?", filepath.IsAbs(path))
\t
\tabsPath, _ := filepath.Abs("main.go")
\tfmt.Println("Absolute path of main.go:", absPath)
}
"""
    ),
    "regexp": (
        "Package `regexp` mengimplementasikan pencarian ekspresi reguler (regular expression).",
        "main.go",
        """package main

import (
\t"fmt"
\t"regexp"
)

func main() {
\t// Compile regex (Gunakan MustCompile untuk inisialisasi global agar panic jika salah)
\tre := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$`)

\temails := []string{
\t\t"user@example.com",
\t\t"invalid-email",
\t\t"admin@domain.co.id",
\t}

\tfor _, email := range emails {
\t\t// MatchString mengembalikan boolean
\t\tmatch := re.MatchString(email)
\t\tfmt.Printf("%s -> Valid? %v\\n", email, match)
\t}

\t// Mencari dan mengganti text
\ttext := "Nomor telepon saya adalah 0812-3456-7890 dan 0899-8888-7777"
\trePhone := regexp.MustCompile(`\\d{4}-\\d{4}-\\d{4}`)
\t
\t// FindAllString mencari semua kecocokan
\tphones := rePhone.FindAllString(text, -1)
\tfmt.Println("\\nNomor telepon yang ditemukan:", phones)

\t// ReplaceAllString mengganti kecocokan
\tanonymized := rePhone.ReplaceAllString(text, "[SENSOR]")
\tfmt.Println("Teks disensor:", anonymized)
}
"""
    ),
    "strconv": (
        "Package `strconv` mengimplementasikan konversi ke dan dari representasi string tipe data dasar.",
        "main.go",
        """package main

import (
\t"fmt"
\t"log"
\t"strconv"
)

func main() {
\t// 1. String ke Integer (Atoi = ASCII to Integer)
\tstrInt := "12345"
\tnum, err := strconv.Atoi(strInt)
\tif err != nil {
\t\tlog.Fatal(err)
\t}
\tfmt.Printf("String ke Int: %d (tipe: %T)\\n", num, num)

\t// 2. Integer ke String (Itoa = Integer to ASCII)
\tintNum := 9876
\tstr := strconv.Itoa(intNum)
\tfmt.Printf("Int ke String: %s (tipe: %T)\\n", str, str)

\t// 3. String ke Float
\tstrFloat := "3.14159"
\tf, err := strconv.ParseFloat(strFloat, 64)
\tif err != nil {
\t\tlog.Fatal(err)
\t}
\tfmt.Printf("String ke Float: %f (tipe: %T)\\n", f, f)

\t// 4. String ke Boolean
\tstrBool := "true"
\tb, err := strconv.ParseBool(strBool)
\tif err != nil {
\t\tlog.Fatal(err)
\t}
\tfmt.Printf("String ke Bool: %t (tipe: %T)\\n", b, b)
}
"""
    ),
    "database/sql": (
        "Package `sql` menyediakan antarmuka generik di sekitar basis data (atau semu) SQL.",
        "main.go",
        """package main

import (
\t"database/sql"
\t"fmt"
\t"log"
\t// "github.com/mattn/go-sqlite3" // Contoh driver yang harus di-import
)

// Catatan: Program ini tidak dapat langsung dijalankan tanpa driver database (misal SQLite, Postgres, MySQL)
// Untuk menjalankannya, Anda butuh `go get github.com/mattn/go-sqlite3` dan import di atas.

func main() {
\tfmt.Println("Package database/sql membutuhkan driver khusus untuk berjalan.")
\tfmt.Println("Contoh di bawah ini menunjukkan cara penggunaannya secara umum.")

\t/*
\t// 1. Membuka koneksi database
\tdb, err := sql.Open("sqlite3", "./example.db")
\tif err != nil {
\t\tlog.Fatal(err)
\t}
\tdefer db.Close()

\t// 2. Membuat tabel
\t_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)`)
\tif err != nil {
\t\tlog.Fatal(err)
\t}

\t// 3. Insert data menggunakan Prepare Statement
\tstmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
\tif err != nil {
\t\tlog.Fatal(err)
\t}
\tdefer stmt.Close()
\t
\t_, err = stmt.Exec("Golang Developer")
\t
\t// 4. Query data
\trows, err := db.Query("SELECT id, name FROM users")
\tdefer rows.Close()
\t
\tfor rows.Next() {
\t\tvar id int
\t\tvar name string
\t\terr = rows.Scan(&id, &name)
\t\tfmt.Printf("ID: %d, Name: %s\\n", id, name)
\t}
\t*/
}
"""
    ),
    "html/template": (
        "Package `template` (html/template) mengimplementasikan template yang digerakkan oleh data untuk menghasilkan keluaran HTML yang aman terhadap injeksi kode.",
        "main.go",
        """package main

import (
\t"html/template"
\t"log"
\t"os"
)

type PageData struct {
\tTitle   string
\tHeading string
\tItems   []string
}

func main() {
\t// Template HTML dengan sintaks aksi {{ }}
\ttmplString := `
<!DOCTYPE html>
<html>
<head>
\t<title>{{.Title}}</title>
</head>
<body>
\t<h1>{{.Heading}}</h1>
\t<ul>
\t{{range .Items}}
\t\t<li>{{.}}</li>
\t{{else}}
\t\t<li>Tidak ada item.</li>
\t{{end}}
\t</ul>
</body>
</html>
`

\t// Parsing template
\ttmpl, err := template.New("webpage").Parse(tmplString)
\tif err != nil {
\t\tlog.Fatal(err)
\t}

\t// Data yang akan dimasukkan ke dalam template
\tdata := PageData{
\t\tTitle:   "Belajar HTML Template di Go",
\t\tHeading: "Daftar Bahasa Pemrograman",
\t\tItems:   []string{"Go", "Python", "JavaScript", "Rust"},
\t}

\t// Mengeksekusi template dan menulis hasilnya ke stdout (bisa juga ke http.ResponseWriter)
\terr = tmpl.Execute(os.Stdout, data)
\tif err != nil {
\t\tlog.Fatal(err)
\t}
}
"""
    )
}

def generate_markdown(module_name, description, example_filename):
    return f"""# Modul: `{module_name}`

## Deskripsi Singkat
{description}

## Cara Penggunaan
Silakan jalankan contoh kode di bawah ini dengan perintah:
```bash
go run {example_filename}
```

## Referensi Dokumentasi Resmi
Untuk informasi lebih lanjut, silakan baca dokumentasi resmi di:
[https://pkg.go.dev/{module_name}](https://pkg.go.dev/{module_name})
"""

def main():
    base_dir = "go-std-lib-tutorial"
    if not os.path.exists(base_dir):
        os.makedirs(base_dir)

    # Create a root README
    root_readme = f"""# Panduan Lengkap Belajar Go Standard Library

Repositori ini berisi kumpulan tutorial dan contoh kode penggunaan Standard Library (Pustaka Standar) bawaan dari bahasa pemrograman [Go (Golang)](https://go.dev/).

Semua penjelasan ditulis dalam **Bahasa Indonesia** dengan menggunakan sintaks dan fitur terbaru dari bahasa Go.

## Daftar Modul yang Dicakup:

"""
    for mod in sorted(modules.keys()):
        root_readme += f"- [{mod}](./{mod}/README.md)\n"

    with open(os.path.join(base_dir, "README.md"), "w") as f:
        f.write(root_readme)

    for mod_path, (desc, file_name, code) in modules.items():
        # Create directory
        full_dir = os.path.join(base_dir, mod_path)
        os.makedirs(full_dir, exist_ok=True)

        # Write README.md
        readme_path = os.path.join(full_dir, "README.md")
        with open(readme_path, "w") as f:
            f.write(generate_markdown(mod_path, desc, file_name))

        # Write Go file
        go_file_path = os.path.join(full_dir, file_name)
        with open(go_file_path, "w") as f:
            f.write(code)

    # Initialize go mod in the base_dir
    os.system(f"cd {base_dir} && go mod init github.com/user/go-std-lib-tutorial")
    print(f"Berhasil meng-generate struktur di direktori '{base_dir}'")

if __name__ == "__main__":
    main()
