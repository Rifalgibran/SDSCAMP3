package main

import "fmt"

type Mahasiswa struct {
	Nama    string
	NIM     string
	Jurusan string
}

func (m *Mahasiswa) TampilkanInfo() {
	fmt.Printf("Mahasiswa:\nNama: %s\nNIM: %s\nJurusan: %s\n", m.Nama, m.NIM, m.Jurusan)
}

func (m *Mahasiswa) UbahNama(namaBaru string) {
	m.Nama = namaBaru
}

func main() {

	mahasiswa := &Mahasiswa{
		Nama:    "Sueb junaedi ",
		NIM:     "022014",
		Jurusan: "Teknik Elektro",
	}

	fmt.Println("Informasi Mahasiswa sebelum perubahan:")
	mahasiswa.TampilkanInfo()

	// Mengubah nama mahasiswa menggunakan method
	mahasiswa.UbahNama("Rifal Gibran")

	fmt.Println("\nInformasi Mahasiswa setelah perubahan:")
	mahasiswa.TampilkanInfo()
}
