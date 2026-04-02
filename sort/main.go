package main

import (
	"fmt"
	"sort"
)

// Person struct
type Person struct {
	Name string
	Age  int
}

func main() {
	// Mengurutkan slice of int
	numbers := []int{5, 2, 7, 1, 9, 3}
	sort.Ints(numbers)
	fmt.Println("Sorted numbers:", numbers)

	// Mengurutkan slice of string
	fruits := []string{"pisang", "apel", "mangga", "jeruk"}
	sort.Strings(fruits)
	fmt.Println("Sorted fruits:", fruits)

	// Mengurutkan slice of struct menggunakan sort.Slice (Go 1.8+)
	people := []Person{
		{"Budi", 25},
		{"Andi", 30},
		{"Citra", 22},
	}

	// Urutkan berdasarkan umur secara ascending
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sorted people by age:", people)

	// Mengecek apakah sudah berurutan
	isSorted := sort.SliceIsSorted(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Apakah people sudah diurutkan berdasarkan usia?", isSorted)
}
