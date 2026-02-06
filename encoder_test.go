package belajar_golang_json // Menentukan nama package tempat file ini berada

import (
	"encoding/json" // Package untuk encoding & decoding data JSON
	"os"            // Package untuk operasi file dan sistem
	"testing"       // Package testing bawaan Go
)

func TestEncoder(t *testing.T) {
	// Membuat file baru bernama CustomerOut.json
	// Jika file sudah ada, maka isinya akan ditimpa
	writer, _ := os.Create("CustomerOut.json")

	// Membuat JSON encoder yang menulis output ke file (writer)
	encoder := json.NewEncoder(writer)

	// Membuat object Customer yang akan diubah menjadi JSON
	customer := Customer{
		FirstName: "Hilmi", // Mengisi field FirstName
		LastName:  "Yahya", // Mengisi field LastName
	}

	// Mengubah struct customer menjadi JSON
	// lalu langsung menuliskannya ke file CustomerOut.json
	encoder.Encode(customer)
}

// Kesimpulan:
// Kode ini digunakan untuk melakukan encoding data dari struct Customer menjadi format JSON menggunakan json.Encoder, kemudian menyimpannya langsung ke dalam file CustomerOut.json. Pendekatan encoder ini cocok untuk menulis data JSON ke file atau network stream secara langsung tanpa perlu menyimpan hasil JSON dalam bentuk string terlebih dahulu.
