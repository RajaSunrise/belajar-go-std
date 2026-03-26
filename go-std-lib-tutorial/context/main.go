package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Membuat context dasar
	ctx := context.Background()

	// Menambahkan value ke context
	ctxWithValue := context.WithValue(ctx, "userID", 12345)
	fmt.Println("UserID dari context:", ctxWithValue.Value("userID"))

	// Membuat context dengan pembatalan (Cancel)
	ctxCancel, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(2 * time.Second)
		cancel() // Membatalkan operasi setelah 2 detik
	}()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Proses selesai tanpa dibatalkan")
	case <-ctxCancel.Done():
		fmt.Println("Proses dibatalkan:", ctxCancel.Err())
	}

	// Membuat context dengan timeout
	ctxTimeout, cancelTimeout := context.WithTimeout(ctx, 1*time.Second)
	defer cancelTimeout()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Proses lambat selesai")
	case <-ctxTimeout.Done():
		fmt.Println("Proses timeout:", ctxTimeout.Err())
	}
}
