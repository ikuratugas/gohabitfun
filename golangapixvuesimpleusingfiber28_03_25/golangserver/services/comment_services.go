package services

import (
	"encoding/json"
	"errors"
	"golangserver/models"
	"io"
	"net/http"
)

// FetchComments mengambil data komentar dari API JSONPlaceholder
func FetchComments() ([]models.Comment, error) {
	// URL API JSONPlaceholder
	url := "https://jsonplaceholder.typicode.com/comments"

	// Lakukan HTTP GET request ke API
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.New("gagal melakukan request ke API")
	}
	defer response.Body.Close()

	// Baca isi response body dengan io.ReadAll
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("gagal membaca response dari API")
	}

	// Buat slice untuk menampung hasil JSON
	var comments []models.Comment

	// Parsing JSON ke struct Comment
	err = json.Unmarshal(body, &comments)
	if err != nil {
		return nil, errors.New("gagal melakukan parsing JSON")
	}

	return comments, nil
}
