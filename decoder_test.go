package belajar_golang_json // Menentukan nama package tempat file ini berada

import (
	"encoding/json" // Package untuk encoding & decoding data JSON di Go
	"fmt"           // Package untuk output ke console (Println)
	"os"            // Package untuk operasi sistem, termasuk baca file
	"testing"       // Package testing bawaan Go untuk membuat unit test
)

func TestStreamDecoder(t *testing.T) {
	// Membuka file Customer.json dalam mode read-only
	// os.Open mengembalikan file (reader) dan error
	reader, _ := os.Open("Customer.json")

	// Membuat JSON decoder berbasis stream dari file
	// Decoder cocok untuk membaca JSON dari file atau network secara bertahap
	decoder := json.NewDecoder(reader)

	// Membuat pointer ke struct Customer
	// Data hasil decode JSON akan dimasukkan ke struct ini
	customer := &Customer{}

	// Melakukan decoding JSON dari file ke struct customer
	// Decode akan membaca JSON dan mengisi field sesuai tag/struktur
	decoder.Decode(customer)

	// Menampilkan hasil struct customer ke console
	fmt.Println(customer)
}

// Kesimpulan:
// Kode ini digunakan untuk melakukan decoding data JSON dari file Customer.json menggunakan json.Decoder (stream-based decoding) ke dalam struct Customer. Pendekatan ini cocok untuk membaca JSON dari sumber seperti file atau network karena lebih efisien dibandingkan membaca seluruh isi file ke memory terlebih dahulu. Fungsi ini dibungkus sebagai unit test agar bisa dijalankan menggunakan perintah `go test`.
