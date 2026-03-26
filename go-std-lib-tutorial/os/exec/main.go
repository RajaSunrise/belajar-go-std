package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Menjalankan perintah 'ls -la' (atau 'dir' di Windows)
	// Catatan: contoh ini berasumsi OS adalah varian Unix/Linux/macOS
	cmd := exec.Command("ls", "-la")

	// Mengambil output dari command
	output, err := cmd.Output()
	if err != nil {
		// Jika gagal, coba perintah Windows 'cmd /c dir'
		cmd = exec.Command("cmd", "/c", "dir")
		output, err = cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Output perintah:")
	fmt.Println(string(output))
}
