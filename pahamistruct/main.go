package main

import "fmt"

type Mahasiswa struct {
	nama string
	nim  string
}

func sapaMahasiswa(m *Mahasiswa) {
	fmt.Printf("halo, %s di alamat %p\n", m.nama, &m.nama)
}

func (m *Mahasiswa) ubahNamaMahasiswa(namaBaru string) {
	m.nama = namaBaru
	fmt.Printf("nama diperbarui jadi %s di alamat %p\n", m.nama, &m.nama)
}

func sebutMahasiswa(mhs Mahasiswa) {
	fmt.Printf("MAHASISWA TEKNIK INFORMATIKA\n")
	fmt.Printf("nama %s\n", mhs.nama)
	fmt.Printf("nim %s\n", mhs.nim)
}

func main() {
	fmt.Println("mencoba lagi")
	baru := Mahasiswa{nama: "ikrar aprianto", nim: "20 650 028"}
	fmt.Printf("nilai alamat awal %p\n", &baru.nama)
	sapaMahasiswa(&baru)

	fmt.Println("--")
	baru.ubahNamaMahasiswa("rikar")

	sebutMahasiswa(baru)
}
