package main

import (
	"fmt"
	"os"
)

// Person adalah contoh struct untuk demonstrasi formatting
type Person struct {
	Name string
	Age  int
}

// Implementasi interface Stringer
func (p Person) String() string {
	return fmt.Sprintf("Person(Name: %s, Age: %d)", p.Name, p.Age)
}

func main() {
	fmt.Println("--- 1. Basic Print & Format ---")
	name := "Go Developer"
	// Menggunakan verb %s untuk string dan %d untuk integer
	fmt.Printf("Halo, %s! Selamat datang di Go %d.\n", name, 1)

	// Menggabungkan string
	msg := fmt.Sprintf("Ini adalah string hasil format: %s", name)
	fmt.Println(msg)

	fmt.Println("\n--- 2. Struct Formatting ---")
	p := Person{Name: "Budi", Age: 25}
	fmt.Printf("Value (%%v): %v\n", p)
	fmt.Printf("Value with field names (%%+v): %+v\n", p)
	fmt.Printf("Go-syntax representation (%%#v): %#v\n", p)
	fmt.Printf("Type (%%T): %T\n", p)

	fmt.Println("\n--- 3. Width & Precision ---")
	f := 3.14159265
	fmt.Printf("Default float (%%f): %f\n", f)
	fmt.Printf("Precision 2 (%%.2f): %.2f\n", f)
	fmt.Printf("Width 10, Precision 2 (%%10.2f): |%10.2f|\n", f)
	fmt.Printf("Width 10, Left Justified (%%-10.2f): |%-10.2f|\n", f)

	str := "Golang"
	fmt.Printf("String padding (%%10s): |%10s|\n", str)
	fmt.Printf("String padding left (%%-10s): |%-10s|\n", str)

	fmt.Println("\n--- 4. Error Formatting ---")
	// fmt.Errorf digunakan untuk membuat atau membungkus error (Error wrapping dengan %w pada Go 1.13+)
	originalErr := fmt.Errorf("file tidak ditemukan")
	wrappedErr := fmt.Errorf("gagal membuka konfigurasi: %w", originalErr)
	fmt.Println(wrappedErr)

	fmt.Println("\n--- 5. Fprint ---")
	// Menulis ke io.Writer tertentu, misal os.Stdout atau file
	fmt.Fprintln(os.Stdout, "Pesan ini dicetak menggunakan fmt.Fprintln ke os.Stdout")

	fmt.Println("\n--- 6. Scanning (Input) ---")
	// Contoh penggunaan Scan (biasanya menunggu input dari user di terminal).
	// Dalam contoh ini, di-comment untuk menghindari pause saat program dijalankan secara otomatis.
	/*
		var inputName string
		var inputAge int
		fmt.Print("Masukkan nama dan umur (dipisahkan spasi): ")
		fmt.Scan(&inputName, &inputAge)
		fmt.Printf("Anda memasukkan: %s, umur %d\n", inputName, inputAge)
	*/

	// Scan dari string menggunakan Sscan
	inputStr := "Andi 30"
	var parsedName string
	var parsedAge int
	fmt.Sscan(inputStr, &parsedName, &parsedAge)
	fmt.Printf("Hasil Sscan -> Nama: %s, Umur: %d\n", parsedName, parsedAge)
}
