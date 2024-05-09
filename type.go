package bec_uts

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tim struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Pemain   []Pemain           `bson:"pemain,omitempty" json:"pemain,omitempty"`
	Pelatih  string             `bson:"pelatih,omitempty" json:"pelatih,omitempty"`
	Stadion  string             `bson:"stadion,omitempty" json:"stadion,omitempty"`
	Prestasi []string           `bson:"prestasi,omitempty" json:"prestasi,omitempty"`
}

type Pemain struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Posisi       string             `bson:"posisi,omitempty" json:"posisi,omitempty"`
	Tinggi       float64            `bson:"tinggi,omitempty" json:"tinggi,omitempty"`
	Berat        float64            `bson:"berat,omitempty" json:"berat,omitempty"`
	TanggalLahir primitive.DateTime `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir,omitempty"`
	Negara       string             `bson:"negara,omitempty" json:"negara,omitempty"`
	NoPunggung   int                `bson:"no_punggung,omitempty" json:"no_punggung,omitempty"`
	Statistik    Statistik          `bson:"statistik,omitempty" json:"statistik,omitempty"`
}

type Statistik struct {
	Penampilan   int `bson:"penampilan,omitempty" json:"penampilan,omitempty"`
	Gol          int `bson:"gol,omitempty" json:"gol,omitempty"`
	Assist       int `bson:"assist,omitempty" json:"assist,omitempty"`
	KartuKuning  int `bson:"kartu_kuning,omitempty" json:"kartu_kuning,omitempty"`
	KartuMerah   int `bson:"kartu_merah,omitempty" json:"kartu_merah,omitempty"`
	SelisihGol   int `bson:"selisih_gol,omitempty" json:"selisih_gol,omitempty"`
	Performa     int `bson:"performa,omitempty" json:"performa,omitempty"`
	PersentaseUm int `bson:"persentase_um,omitempty" json:"persentase_um,omitempty"`
}
