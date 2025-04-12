package main

import (
	"fmt"
	"golangserver/handlers"
	"log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Inisialisasi Fiber
	app := fiber.New()
	
	// Tambahkan Middleware CORS agar Vue bisa akses API
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Mengizinkan semua domain
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type, Accept",
	}))

	// Menambahkan route untuk mengambil komentar
	app.Get("/comments", handlers.GetCommentsHandler)

	// Rute untuk mendapatkan komentar berdasarkan postId
	app.Get("/comments/filter", handlers.GetCommentsByPostIDHandler)

	// Jalankan server pada port 3000
	fmt.Println("🚀 Server berjalan di http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
