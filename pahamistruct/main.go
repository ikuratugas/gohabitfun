package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Mahasiswa struct {
	Nama string `json:"nama mahasiswa"`
	Nim  string `json:"nim mahassiwa"`
}

// func sapaMahasiswa(m *Mahasiswa) {
// 	fmt.Printf("halo, %s di alamat %p\n", m.nama, &m.nama)
// }

// func (m *Mahasiswa) ubahNamaMahasiswa(namaBaru string) {
// 	m.nama = namaBaru
// 	fmt.Printf("nama diperbarui jadi %s di alamat %p\n", m.nama, &m.nama)
// }

// func sebutMahasiswa(mhs Mahasiswa) {
// 	fmt.Printf("MAHASISWA TEKNIK INFORMATIKA\n")
// 	fmt.Printf("nama %s\n", mhs.nama)
// 	fmt.Printf("nim %s\n", mhs.nim)
// }

func showToJson(w http.ResponseWriter, r *http.Request) {
	mahasiswa1 := Mahasiswa{
		Nama: "ikrar aprianto",
		Nim:  "2065028",
	}

	jsonData, err := json.Marshal(mahasiswa1)
	if err != nil {
		http.Error(w, "Gagal Mengubah Data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", showToJson)
	fmt.Println("sistem berjalan di server localhsot:8000")
	http.ListenAndServe(":8000", mux)
}
