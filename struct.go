package main

import "fmt"

// Struct Mahasiswa
type Mahasiswa struct {
	Nama    string
	NIM     string
	Jurusan string
}

// Method
func (ini Mahasiswa) TampilkanInfoMahasiswa() {
	fmt.Printf("Mahasiswa:\nNama: %s\nNIM: %s\nJurusan: %s\n", ini.Nama, ini.NIM, ini.Jurusan)
}

// Struct Dosen
type Dosen struct {
	Nama       string
	NIP        string
	Pendidikan string
}

// Method
func (aku Dosen) TampilkanInfoDosen() {
	fmt.Printf("Dosen:\nNama: %s\nNIP: %s\nPendidikan: %s\n", aku.Nama, aku.NIP, aku.Pendidikan)
}

func main() {
	// penggunaan struct Mahasiswa
	mahasiswa := Mahasiswa{
		Nama:    "Rifal Gibran",
		NIM:     "2202214",
		Jurusan: "Teknik Elektro",
	}
	fmt.Println("Informasi Mahasiswa:")
	mahasiswa.TampilkanInfoMahasiswa()

	// penggunaan struct Dosen
	dosen := Dosen{
		Nama:       "Dr. Pangestu alam ST.MT",
		NIP:        "0231322",
		Pendidikan: "S3",
	}
	fmt.Println("\nInformasi Dosen:")
	dosen.TampilkanInfoDosen()
}
