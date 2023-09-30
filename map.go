package main

import "fmt"

func main() {

	siswa := make(map[int]string)

	siswa[114] = "Kelvin"
	siswa[152] = "Ali usman"
	siswa[101] = "Arifin ilham"

	fmt.Println("Data Siswa:")
	for nim, nama := range siswa {
		fmt.Printf("NIM: %d, Nama: %s\n", nim, nama)
	}

	siswa[101] = "Rifal Gibran"

	fmt.Println("\nData Siswa setelah perubahan:")
	for nim, nama := range siswa {
		fmt.Printf("NIM: %d, Nama: %s\n", nim, nama)
	}

	delete(siswa, 114)

	fmt.Println("\nData Siswa setelah penghapusan:")
	for nim, nama := range siswa {
		fmt.Printf("NIM: %d, Nama: %s\n", nim, nama)
	}

	dosen := map[string]string{
		"Nama Dosen": "Azis ST.MT",
		"pendidikan": "S2",
		"alamat":     "Purwokerto",
		"umur":       "42",
	}

	fmt.Println("\nData Dosen:")
	for key, value := range dosen {
		fmt.Printf("%s: %s\n", key, value)
	}
	fmt.Println("Jumlah data yang ada di Map dosen yaitu :", len(dosen))

	dosenbaikhati := map[string]int{
		"nomor": 07654321,
		"umur":  42,
	}
	fmt.Println("\ndosenbaikhati:")
	fmt.Println(dosenbaikhati)
	fmt.Println("Jumlah data yang ada di dosenbaikhati person yaitu :", len(dosenbaikhati))

	nama, exists := siswa[101]
	if exists {
		fmt.Printf("\nData siswa dengan NIM 101 ditemukan. Nama: %s\n", nama)
	} else {
		fmt.Println("\nData siswa dengan NIM 101 tidak ditemukan.")
	}
}
