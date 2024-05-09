package bec_uts

import (
	"fmt"
	"testing"
	"time"
)

func TestInsertPemain(t *testing.T) {
	nama := "Halaand"
	posisi := "Penyerang"
	tinggi := 170.0
	berat := 72.5
	tanggalLahir := time.Date(1987, time.June, 24, 0, 0, 0, 0, time.UTC)
	negara := "denmark"
	noPunggung := 10
	statistik := Statistik{
		Penampilan: 500,
		Gol:        400,
		Assist:     200,
		KartuKuning: 10,
		KartuMerah:  2,
		SelisihGol:  700,
		Performa:    9,
		PersentaseUm: 80,
	}
	hasil := InsertPemain(nama, posisi, tinggi, berat, tanggalLahir, negara, noPunggung, statistik)
	fmt.Println(hasil)
}

func TestGetPemainByNama(t *testing.T) {
	nama := "Haaland"
	pemain := GetPemainByNama(nama)
	fmt.Println(pemain)
}

func TestGetAllPemain(t *testing.T) {
	data := GetAllPemain()
	fmt.Println(data)
}
