package belajar_golang_json // Menentukan nama package tempat file ini berada

import (
	"encoding/json" // Package untuk encode (marshal) dan decode (unmarshal) data JSON
	"fmt"           // Package untuk menampilkan output ke console
	"testing"       // Package bawaan Go untuk membuat dan menjalankan unit test
)

// Struct Address merepresentasikan data alamat
type Address struct {
	Street     string // Nama jalan
	Country    string // Nama negara
	PostalCode string // Kode pos
}

// Struct Customer merepresentasikan data pelanggan
type Customer struct {
	FirstName string    // Nama depan customer
	LastName  string    // Nama belakang customer
	Age       int       // Umur customer
	Married   bool      // Status pernikahan customer
	Hobbies   []string  // Daftar hobi customer dalam bentuk slice
	Addresses   []Address // Daftar alamat customer (relasi ke struct Address)
}

// Fungsi test untuk menguji konversi object Go ke JSON
func TestJSONObject(t *testing.T) {
	// Membuat instance Customer dan mengisi sebagian field-nya
	customer := Customer{
		FirstName: "Hilmi", // Mengisi nama depan
		LastName:  "Yahya", // Mengisi nama belakang
		Age:       22,      // Mengisi umur
		Married:   false,   // Mengisi status menikah
	}

	// Mengubah struct Customer menjadi JSON dalam bentuk byte slice
	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err) // Menghentikan program jika terjadi error saat proses marshal
	}

	// Menampilkan hasil JSON ke console dalam bentuk string
	fmt.Println(string(bytes))
}

// Kesimpulan:
// Kode ini mendemonstrasikan cara mengubah (marshal) sebuah struct di Go menjadi format JSON menggunakan package encoding/json. Struct Customer dibuat sebagai representasi data, kemudian di dalam fungsi test, data tersebut dikonversi menjadi JSON dan ditampilkan ke console. Pendekatan ini umum digunakan saat mengirim data ke API, menyimpan data, atau berkomunikasi antar sistem berbasis JSON.
