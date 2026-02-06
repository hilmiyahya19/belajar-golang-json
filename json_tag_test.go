package belajar_golang_json // Menentukan nama package tempat file ini berada

import (
	"encoding/json" // Package untuk melakukan encode (marshal) dan decode (unmarshal) JSON
	"fmt"           // Package untuk menampilkan output ke console
	"testing"       // Package untuk kebutuhan unit testing
)

// Struct Product merepresentasikan data produk
// JSON tag digunakan untuk menentukan nama field saat diubah ke JSON
type Product struct {
	Id       string `json:"id"`        // Field Id akan menjadi "id" di JSON
	Name     string `json:"name"`      // Field Name akan menjadi "name" di JSON
	ImageURL string `json:"image_url"` // Field ImageURL akan menjadi "image_url" di JSON
}

// Test untuk mengubah struct Product menjadi JSON menggunakan JSON tag
func TestJSONTag(t *testing.T) {
	// Membuat object Product
	product := Product{
		Id:       "1",                            // ID produk
		Name:     "Macbook Pro",                  // Nama produk
		ImageURL: "http://example.com/image.jpg", // URL gambar produk
	}

	// Mengubah struct Product menjadi JSON
	bytes, _ := json.Marshal(product)
	// Menampilkan hasil JSON ke console
	fmt.Println(string(bytes))
}

// Test untuk mengubah JSON menjadi struct Product menggunakan JSON tag
func TestJSONTagDecode(t *testing.T) {
	// JSON string dengan nama field sesuai JSON tag
	jsonString := `{"id":"1","name":"Macbook Pro","image_url":"http://example.com/image.jpg"}`
	// Mengubah JSON string menjadi byte slice
	jsonBytes := []byte(jsonString)

	// Membuat pointer ke struct Product sebagai tujuan decode
	product := &Product{}

	// Mengubah JSON menjadi struct Product
	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err) // Menghentikan program jika terjadi error saat decode
	}

	// Menampilkan hasil decode
	fmt.Println(product)
}

// Kesimpulan:
// Kode ini menunjukkan penggunaan JSON tag pada struct di Go untuk mengatur nama field saat proses encode dan decode JSON. Dengan JSON tag, nama field pada struct Go dapat berbeda dengan nama field di JSON, sehingga lebih fleksibel dan sesuai dengan standar API atau kebutuhan sistem lain tanpa harus mengubah nama field di kode Go.
