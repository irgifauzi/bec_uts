package bec_uts

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

// Fungsi untuk terhubung ke MongoDB dan mengembalikan koneksi ke database
func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

// Fungsi untuk menyisipkan satu dokumen ke dalam koleksi tertentu di database MongoDB
func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

// Fungsi untuk menyisipkan data pemain ke dalam koleksi "pemain" di database
func InsertPemain(nama string, posisi string, tinggi float64, berat float64, tanggalLahir time.Time, negara string, noPunggung int, statistik Statistik) (insertedID interface{}) {
	var pemain Pemain
	pemain.Nama = nama
	pemain.Posisi = posisi
	pemain.Tinggi = tinggi
	pemain.Berat = berat
	pemain.TanggalLahir = primitive.NewDateTimeFromTime(tanggalLahir.UTC())
	pemain.Negara = negara
	pemain.NoPunggung = noPunggung
	pemain.Statistik = statistik
	return InsertOneDoc("football_db", "pemain", pemain)
}

// Fungsi untuk mendapatkan data pemain berdasarkan nama
func GetPemainByNama(nama string) (pemain Pemain) {
	collection := MongoConnect("football_db").Collection("pemain")
	filter := bson.M{"nama": nama}
	err := collection.FindOne(context.TODO(), filter).Decode(&pemain)
	if err != nil {
		fmt.Printf("GetPemainByNama: %v\n", err)
	}
	return pemain
}

// Fungsi untuk mendapatkan semua data pemain
func GetAllPemain() (data []Pemain) {
	collection := MongoConnect("football_db").Collection("pemain")
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllPemain: ", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
