package belajar_golang_json // Nama package tempat file ini berada

import (
	"encoding/json" // Package untuk encode & decode data JSON
	"fmt"           // Package untuk output ke console
	"testing"       // Package untuk membuat unit test di Golang
)

func TestMapDecode(t *testing.T) { // Fungsi test untuk decode JSON ke map
	jsonString := `{"id":"1","name":"Macbook Pro","price":20000000}` // Data JSON dalam bentuk string
	jsonBytes := []byte(jsonString)                                 // Mengubah string JSON menjadi byte slice

	var result map[string]interface{} // Deklarasi map dengan value bertipe interface{} agar fleksibel
	err := json.Unmarshal(jsonBytes, &result) // Decode JSON byte ke dalam map
	if err != nil {                            // Mengecek apakah terjadi error saat decode
		panic(err) // Menghentikan program jika terjadi error
	}

	fmt.Println(result)         // Menampilkan seluruh isi map hasil decode
	fmt.Println(result["id"])   // Mengambil value dengan key "id"
	fmt.Println(result["name"]) // Mengambil value dengan key "name"
	fmt.Println(result["price"])// Mengambil value dengan key "price"
}

func TestMapEncode(t *testing.T) { // Fungsi test untuk encode map ke JSON
	product := map[string]interface{}{ // Membuat map dengan key string dan value fleksibel
		"id":    "1",           // Key id dengan value string
		"name":  "Macbook Pro", // Key name dengan value string
		"price": 20000000,      // Key price dengan value integer
	}

	bytes, err := json.Marshal(product) // Mengubah map menjadi JSON dalam bentuk byte
	if err != nil {                      // Mengecek apakah terjadi error saat encode
		panic(err) // Menghentikan program jika terjadi error
	}
	fmt.Println(string(bytes)) // Mengubah byte JSON ke string lalu menampilkannya
}

// Kesimpulan:
// Kode ini menunjukkan cara menggunakan package encoding/json di Golang untuk melakukan decode (Unmarshal) data JSON ke dalam map[string]interface{} dan encode (Marshal) map menjadi JSON. TestMapDecode berfungsi membaca JSON dan mengakses nilainya secara dinamis, sedangkan TestMapEncode menunjukkan bagaimana data map dapat diubah kembali menjadi format JSON. Pendekatan ini cocok digunakan ketika struktur data bersifat fleksibel atau belum didefinisikan menggunakan struct.

// Kelebihan penggunaan map terletak pada fleksibilitas tipe data setiap field yang dapat berubah, sehingga lebih adaptif dibandingkan penggunaan struct yang memiliki tipe data tetap.