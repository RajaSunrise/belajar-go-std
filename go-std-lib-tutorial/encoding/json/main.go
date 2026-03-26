package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User merepresentasikan data pengguna
type User struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	IsActive bool     `json:"is_active"`
	Roles    []string `json:"roles,omitempty"` // omitempty: diabaikan jika kosong
}

func main() {
	// 1. Marshal: Mengubah struct/map menjadi JSON string (byte slice)
	user := User{
		ID:       1,
		Name:     "Budi",
		Email:    "budi@example.com",
		IsActive: true,
		Roles:    []string{"admin", "user"},
	}

	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hasil Marshal (Struct ke JSON):\n", string(jsonData))

	// 2. Unmarshal: Mengubah JSON string menjadi struct/map
	jsonString := `{"id":2,"name":"Siti","email":"siti@example.com","is_active":false}`
	var newUser User

	err = json.Unmarshal([]byte(jsonString), &newUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nHasil Unmarshal (JSON ke Struct): %+v\n", newUser)
}
