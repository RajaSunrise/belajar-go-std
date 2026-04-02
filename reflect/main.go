package main

import (
	"fmt"
	"reflect"
)

// User merepresentasikan entitas pengguna dalam sistem.
// Struct ini menggunakan struct tags untuk mendemonstrasikan bagaimana metadata
// dapat disematkan dan dibaca saat runtime menggunakan package "reflect".
type User struct {
	ID        int    `json:"id" db:"user_id" validate:"required"`
	Username  string `json:"username" db:"username" validate:"required,min=3"`
	IsActive  bool   `json:"is_active" db:"active"`
	hidden    string // field unexported (huruf kecil) tidak dapat diakses atau diubah secara eksternal oleh reflect
}

// DemoTypeAndValue mendemonstrasikan fungsi dasar refleksi:
// mengekstrak informasi tipe (Type) dan nilai (Value) dari sebuah interface{}.
func DemoTypeAndValue() {
	var number float64 = 3.14159
	var text string = "Belajar Golang Reflection"

	// reflect.TypeOf() mengembalikan objek reflect.Type yang berisi informasi tentang tipe data.
	typeOfNumber := reflect.TypeOf(number)
	// reflect.ValueOf() mengembalikan objek reflect.Value yang berisi nilai aktual di memori.
	valueOfText := reflect.ValueOf(text)

	fmt.Printf("Variabel 'number': Tipe = %v, Kind = %v\n", typeOfNumber.Name(), typeOfNumber.Kind())
	fmt.Printf("Variabel 'text': Nilai = %v, Kind = %v\n", valueOfText.String(), valueOfText.Kind())
}

// DemoStructReflection mendemonstrasikan cara melakukan inspeksi pada struktur data (Struct).
// Ini adalah fondasi dari library seperti encoding/json atau ORM database.
func DemoStructReflection() {
	u := User{
		ID:       101,
		Username: "gopher_master",
		IsActive: true,
		hidden:   "rahasia_internal",
	}

	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	fmt.Println("--- Inspeksi Struct 'User' ---")
	fmt.Printf("Nama Struct: %s, Total Field: %d\n", t.Name(), t.NumField())

	// Iterasi setiap field dalam struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)      // Mendapatkan reflect.StructField (metadata)
		value := v.Field(i)      // Mendapatkan reflect.Value (data aktual)

		// Periksa apakah field bisa diekspor (huruf kapital)
		if field.IsExported() {
			fmt.Printf("Field: %-10s | Tipe: %-10s | Nilai: %v\n", field.Name, field.Type, value.Interface())
		} else {
			fmt.Printf("Field: %-10s | (Unexported field, tidak bisa diekstrak nilainya dengan aman)\n", field.Name)
		}
	}
}

// DemoStructTags mendemonstrasikan cara membaca dan mem-parsing Struct Tags.
// Struct tags sangat krusial untuk validasi data, serialisasi (JSON/XML), dan pemetaan kolom database.
func DemoStructTags() {
	u := User{}
	t := reflect.TypeOf(u)

	fmt.Println("--- Membaca Struct Tags ---")

	// Kita akan mencari field 'Username'
	field, found := t.FieldByName("Username")
	if found {
		fmt.Printf("Field '%s' ditemukan:\n", field.Name)
		// field.Tag.Get() akan mengembalikan string sesuai dengan key yang diminta
		fmt.Printf("  Tag 'json'     : %s\n", field.Tag.Get("json"))
		fmt.Printf("  Tag 'db'       : %s\n", field.Tag.Get("db"))
		fmt.Printf("  Tag 'validate' : %s\n", field.Tag.Get("validate"))
	}
}

// DemoModifyValue mendemonstrasikan salah satu fitur paling kuat namun berbahaya dari refleksi:
// mengubah nilai memori secara dinamis saat runtime menggunakan pointer.
func DemoModifyValue() {
	fmt.Println("--- Memodifikasi Nilai Melalui Reflection ---")

	score := 50
	fmt.Printf("Nilai awal score: %d\n", score)

	// PENTING: Untuk mengubah nilai, Anda HARUS mengoper pointer ke reflect.ValueOf().
	// Jika Anda mengirim 'score' langsung (pass by value), reflect hanya akan menerima
	// salinannya dan tidak bisa mengubah variabel asli.
	v := reflect.ValueOf(&score)

	// v saat ini adalah pointer. Kita gunakan Elem() untuk men-dereference pointer
	// dan mendapatkan nilai aslinya yang bisa dimodifikasi (settable).
	elem := v.Elem()

	// Periksa apakah elemen tersebut bisa diubah nilainya (CanSet)
	if elem.CanSet() {
		// Ubah nilai integer menggunakan SetInt()
		elem.SetInt(99)
		fmt.Println("Berhasil mengubah nilai!")
	}

	fmt.Printf("Nilai akhir score setelah dimodifikasi: %d\n", score)
}

func main() {
	println("=== Demonstrasi Basic Type & Value ===")
	DemoTypeAndValue()
	println("\n======================================\n")

	println("=== Demonstrasi Inspeksi Struct ===")
	DemoStructReflection()
	println("\n======================================\n")

	println("=== Demonstrasi Struct Tags ===")
	DemoStructTags()
	println("\n======================================\n")

	println("=== Demonstrasi Modifikasi Nilai (Set) ===")
	DemoModifyValue()
	println("\n======================================\n")
}
