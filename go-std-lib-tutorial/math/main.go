package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Nilai Pi:", math.Pi)

	// Mencari nilai akar kuadrat
	fmt.Println("Akar dari 16 adalah:", math.Sqrt(16))

	// Membulatkan angka
	fmt.Println("Ceil 3.14:", math.Ceil(3.14))
	fmt.Println("Floor 3.14:", math.Floor(3.14))
	fmt.Println("Round 3.14:", math.Round(3.14))

	// Nilai absolut
	fmt.Println("Absolut -10:", math.Abs(-10))
}
