package main

import "fmt"

type Mahasisawa struct {
	nama string
	nim  string
}

func main() {
	fmt.Println("mencoba lagi")
	baru := Mahasisawa{nama: "ikrar aprianto", nim: "20 650 028"}
	fmt.Printf("MAHASISWA TEKNIK INFORMATIKA\n")
	fmt.Printf("nama %s\n", baru.nama)
	fmt.Printf("nim %s\n", baru.nim)
}
