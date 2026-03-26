package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Menggunakan ServeMux (Go 1.22+ mendukung method based routing)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Halo, selamat datang di server Go!")
	})

	mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		fmt.Fprintf(w, "Halo, %s!\n", name)
	})

	fmt.Println("Server berjalan di port 8080...")
	// Uncomment baris di bawah untuk menjalankan server
	// log.Fatal(http.ListenAndServe(":8080", mux))

	// Contoh HTTP Client sederhana
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Response Body (dipotong):", string(body[:50]), "...")
}
