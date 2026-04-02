package main

import (
	"fmt"
	"sync"
)

// Counter aman untuk konkurensi (thread-safe)
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()   // Kunci sebelum memodifikasi
	c.value++
	c.mu.Unlock() // Buka kunci setelah memodifikasi
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock() // defer memastikan Unlock selalu dipanggil
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// Menjalankan 1000 goroutine untuk increment counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	// Menunggu semua goroutine selesai
	wg.Wait()

	fmt.Println("Nilai akhir counter:", counter.Value()) // Seharusnya 1000
}
