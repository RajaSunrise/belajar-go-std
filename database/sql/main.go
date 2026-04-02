package main

import (
	"fmt"
	// "database/sql"
	// "log"
	// "github.com/mattn/go-sqlite3" // Contoh driver yang harus di-import
)

// Catatan: Program ini tidak dapat langsung dijalankan tanpa driver database (misal SQLite, Postgres, MySQL)
// Untuk menjalankannya, Anda butuh `go get github.com/mattn/go-sqlite3` dan import di atas.

func main() {
	fmt.Println("Package database/sql membutuhkan driver khusus untuk berjalan.")
	fmt.Println("Contoh di bawah ini menunjukkan cara penggunaannya secara umum.")

	/*
	// 1. Membuka koneksi database
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. Membuat tabel
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)`)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Insert data menggunakan Prepare Statement
	stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("Golang Developer")

	// 4. Query data
	rows, err := db.Query("SELECT id, name FROM users")
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
	*/
}
