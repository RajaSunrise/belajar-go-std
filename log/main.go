package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"os"
	"time"
)

// DemoStandardLog mendemonstrasikan penggunaan package "log" standar di Go.
// Package "log" klasik menyediakan fitur logging yang simpel, thread-safe,
// dan cukup untuk aplikasi kecil atau CLI sederhana.
func DemoStandardLog() {
	// 1. Logging standar langsung ke stdout/stderr.
	log.Println("Ini adalah pesan log standar menggunakan log.Println")

	// 2. Custom Logger menggunakan os.Stdout dengan prefix dan flag tertentu.
	// Flag log.Ldate | log.Ltime | log.Lshortfile menambahkan tanggal, waktu, dan lokasi file/baris.
	customLogger := log.New(os.Stdout, "[CUSTOM-INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	customLogger.Println("Pesan ini dicetak menggunakan custom logger")

	// 3. Kita juga dapat mengatur flag pada logger default secara global
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetPrefix("[GLOBAL] ")
	log.Println("Pesan ini dicetak dengan logger default yang telah dimodifikasi flag dan prefixnya")

	// Reset default logger agar tidak mempengaruhi demo selanjutnya
	log.SetPrefix("")
	log.SetFlags(log.LstdFlags)
}

// DemoStructuredLog mendemonstrasikan penggunaan package "log/slog" yang diperkenalkan di Go 1.21.
// "log/slog" menyediakan structured logging yang sangat berguna untuk aplikasi produksi
// modern, microservices, dan agregasi log tersentralisasi (seperti ELK, Datadog, dll).
func DemoStructuredLog() {
	// 1. TextHandler: Mencetak log terstruktur dalam format teks yang mudah dibaca manusia (key=value).
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug, // Atur level minimum ke Debug agar semua log muncul
	})
	textLogger := slog.New(textHandler)

	// Set textLogger sebagai logger default secara global.
	slog.SetDefault(textLogger)

	slog.Info("Aplikasi dimulai menggunakan TextHandler",
		slog.String("version", "1.0.0"),
		slog.Int("port", 8080),
	)

	slog.Debug("Pesan debug, berguna untuk troubleshooting mendalam",
		slog.Any("config", map[string]string{"env": "development"}),
	)

	// 2. JSONHandler: Mencetak log dalam format JSON. Sangat ideal untuk mesin (log aggregators).
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true, // Menambahkan info file dan baris asal log
	})
	jsonLogger := slog.New(jsonHandler)

	jsonLogger.Info("Memproses transaksi baru",
		slog.String("transaction_id", "TRX-987654321"),
		slog.Float64("amount", 150.50),
		slog.Time("timestamp", time.Now()),
	)

	// 3. Menggunakan slog.Group untuk mengelompokkan atribut log
	jsonLogger.Info("Detail pengguna",
		slog.Group("user",
			slog.String("id", "USR-123"),
			slog.String("name", "Budi Santoso"),
			slog.String("role", "admin"),
		),
	)

	// 4. Log dengan Error
	err := errors.New("koneksi database terputus")
	jsonLogger.Error("Gagal melakukan query",
		slog.String("query", "SELECT * FROM users"),
		slog.String("error", err.Error()),
		// Anda juga bisa menggunakan slog.Any("error", err)
	)

	// 5. Contextual Logging: Membawa logger di dalam context
	// Ini hanyalah contoh konsep, biasanya dilakukan di middleware HTTP/gRPC.
	ctx := context.WithValue(context.Background(), "request_id", "REQ-ABCDEF")

	// Dalam praktiknya, kita membuat helper function untuk mengekstrak atau menyuntikkan logger ke context.
	// Di sini kita tunjukkan bagaimana kita bisa menambahkan context log.
	jsonLogger.LogAttrs(ctx, slog.LevelInfo, "Memproses request HTTP",
		slog.String("request_id", ctx.Value("request_id").(string)),
		slog.String("path", "/api/v1/users"),
	)
}

func main() {
	println("=== Demonstrasi Standard Log (log) ===")
	DemoStandardLog()
	println("\n======================================\n")

	println("=== Demonstrasi Structured Log (log/slog) ===")
	DemoStructuredLog()
	println("\n======================================\n")
}
