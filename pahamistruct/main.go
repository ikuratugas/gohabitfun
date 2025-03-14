package main

import "fmt"

type Mahasisawa struct {
	nama string
	nim  string
}

func sebutMahasiswa(mhs Mahasisawa) {
	fmt.Printf("MAHASISWA TEKNIK INFORMATIKA\n")
	fmt.Printf("nama %s\n", mhs.nama)
	fmt.Printf("nim %s\n", mhs.nim)
}

func main() {
	fmt.Println("mencoba lagi")
	baru := Mahasisawa{nama: "ikrar aprianto", nim: "20 650 028"}

	sebutMahasiswa(baru)
}
