package main

import (
	"fmt"
	"os"
	"strings"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var peserta = []Biodata{
	{"Thomas", "Jakarta", "Programmer", "Ingin mempelajari Golang untuk pengembangan aplikasi backend"},
	{"Nicko", "Bandung", "Data Analyst", "Tertarik memperdalam ilmu hubungan antara Data Analyst dan Golang"},
	{"Rex", "Bali", "Full Stack Developer", "Mau memperdalam bahasa pemrograman baru"},
	{"Chloe", "Yogyakarta", "DevOps Engineer", "Golang cocok untuk sistem yang skalabel dan efisien"},
	{"Kath", "Surabaya", "Mobile Developer", "Penasaran dengan performa tinggi Golang untuk aplikasi mobile"},
}

func findPerson(nama string) *Biodata {
	for _, biodata := range peserta {
		if strings.ToLower(biodata.Nama) == strings.ToLower(nama) {
			return &biodata
		}
	}
	return nil
}

/*
 Cara lain pakai EqualFold (hanya catatan saja) - auto recommendation dari extensions Go

func findPerson(nama string) *Biodata {
	for _, biodata := range peserta {
		if strings.EqualFold(biodata.Nama, nama) {
			return &biodata
		}
	}
	return nil
}
*/

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Silahkan insert nama peserta di akhir argumen. Contoh: go run biodata.go nama_peserta")
		return
	}

	nama := os.Args[1]

	biodata := findPerson(nama)

	if biodata != nil {
		fmt.Printf("Nama: %s\n", biodata.Nama)
		fmt.Printf("Alamat: %s\n", biodata.Alamat)
		fmt.Printf("Pekerjaan: %s\n", biodata.Pekerjaan)
		fmt.Printf("Alasan Memilih Kelas Golang: %s\n", biodata.Alasan)
	} else {
		fmt.Printf("Biodata untuk nama %s tidak ditemukan.\n", nama)
	}
}
