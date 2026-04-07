package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("--- 1. Konstanta Matematika Dasar ---")
	fmt.Println("Nilai Pi (math.Pi):", math.Pi)
	fmt.Println("Nilai E (math.E):", math.E)
	fmt.Println("Nilai Phi (math.Phi):", math.Phi)
	fmt.Println("Max Int64:", math.MaxInt64)
	fmt.Println("Max Float64:", math.MaxFloat64)

	fmt.Println("\n--- 2. Operasi Dasar & Pangkat ---")
	// Mencari nilai akar kuadrat
	fmt.Println("Akar dari 16 (Sqrt):", math.Sqrt(16))
	// Pangkat
	fmt.Println("2 pangkat 3 (Pow):", math.Pow(2, 3))
	// Pangkat 10
	fmt.Println("10 pangkat 2 (Pow10):", math.Pow10(2))
	// Nilai absolut
	fmt.Println("Absolut -10.5 (Abs):", math.Abs(-10.5))

	fmt.Println("\n--- 3. Fungsi Pembulatan ---")
	fmt.Println("Ceil 3.14 (Membulatkan ke atas):", math.Ceil(3.14))
	fmt.Println("Floor 3.14 (Membulatkan ke bawah):", math.Floor(3.14))
	fmt.Println("Round 3.5 (Membulatkan ke nilai terdekat):", math.Round(3.5))
	fmt.Println("Trunc 3.14 (Memotong desimal):", math.Trunc(3.14))

	fmt.Println("\n--- 4. Logaritma & Eksponensial ---")
	fmt.Println("Natural Logaritma dari E (Log):", math.Log(math.E))
	fmt.Println("Logaritma basis 10 dari 100 (Log10):", math.Log10(100))
	fmt.Println("Logaritma basis 2 dari 8 (Log2):", math.Log2(8))
	fmt.Println("Eksponensial dari 1 (Exp):", math.Exp(1)) // e^1

	fmt.Println("\n--- 5. Trigonometri ---")
	// Parameter menggunakan radian
	rad := 90.0 * (math.Pi / 180.0) // Konversi 90 derajat ke radian
	fmt.Printf("Sin(90 derajat): %.4f\n", math.Sin(rad))
	fmt.Printf("Cos(90 derajat): %.4f\n", math.Cos(rad))
	fmt.Printf("Tan(45 derajat): %.4f\n", math.Tan(45.0*(math.Pi/180.0)))

	fmt.Println("\n--- 6. Penanganan NaN dan Infinity ---")
	// NaN (Not a Number) terjadi misalnya karena operasi akar negatif atau 0/0
	nanValue := math.NaN()
	fmt.Println("Apakah nanValue adalah NaN?", math.IsNaN(nanValue))
	fmt.Println("Apakah 0.0 adalah NaN?", math.IsNaN(0.0))

	// Infinity terjadi misalnya karena pembagian dengan angka mendekati 0 atau melewati batas float
	infValue := math.Inf(1) // Positif infinity
	fmt.Println("Nilai positif infinity:", infValue)
	fmt.Println("Apakah infValue infinity?", math.IsInf(infValue, 0)) // sign 0 berarti positif atau negatif inf
	fmt.Println("Apakah -infValue infinity negatif?", math.IsInf(math.Inf(-1), -1))

	fmt.Println("\n--- 7. Max dan Min ---")
	fmt.Println("Max dari 10.5 dan 20.3:", math.Max(10.5, 20.3))
	fmt.Println("Min dari 10.5 dan 20.3:", math.Min(10.5, 20.3))
}
