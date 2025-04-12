package handlers

import (
	"golangserver/models"
	"golangserver/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetCommentsHandler menangani permintaan untuk mendapatkan daftar komentar
func GetCommentsHandler(c *fiber.Ctx) error {
	// Panggil service untuk mendapatkan komentar
	comments, err := services.FetchComments()
	if err != nil {
		// Jika ada kesalahan, kembalikan respon error ke user
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data komentar",
		})
	}

	// Berikan response sukses ke user
	return c.JSON(comments)
}

// GetCommentsByPostIDHandler mengambil komentar berdasarkan postId
func GetCommentsByPostIDHandler(c *fiber.Ctx) error {
	// Ambil query parameter postId (jika ada)
	postIDParam := c.Query("postId")

	// Panggil service untuk mendapatkan komentar
	comments, err := services.FetchComments()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data komentar",
		})
	}

	// Jika user memberikan postId, filter komentar berdasarkan postId
	if postIDParam != "" {
		postID, err := strconv.Atoi(postIDParam)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "postId harus berupa angka",
			})
		}

		// Filter komentar yang sesuai dengan postId
		var filteredComments []models.Comment
		for _, comment := range comments {
			if comment.PostID == postID {
				filteredComments = append(filteredComments, comment)
			}
		}

		return c.JSON(filteredComments)
	}

	// Jika tidak ada postId, kembalikan semua komentar
	return c.JSON(comments)
}
