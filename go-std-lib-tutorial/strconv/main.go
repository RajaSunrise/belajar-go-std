package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// 1. String ke Integer (Atoi = ASCII to Integer)
	strInt := "12345"
	num, err := strconv.Atoi(strInt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("String ke Int: %d (tipe: %T)\n", num, num)

	// 2. Integer ke String (Itoa = Integer to ASCII)
	intNum := 9876
	str := strconv.Itoa(intNum)
	fmt.Printf("Int ke String: %s (tipe: %T)\n", str, str)

	// 3. String ke Float
	strFloat := "3.14159"
	f, err := strconv.ParseFloat(strFloat, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("String ke Float: %f (tipe: %T)\n", f, f)

	// 4. String ke Boolean
	strBool := "true"
	b, err := strconv.ParseBool(strBool)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("String ke Bool: %t (tipe: %T)\n", b, b)
}
