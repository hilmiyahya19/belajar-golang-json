package belajar_golang_json // Menentukan nama package

import (
	"encoding/json" // Package untuk encode dan decode JSON
	"fmt"           // Package untuk menampilkan output ke console
	"testing"       // Package untuk keperluan unit testing
)

// Test untuk encode struct yang memiliki slice (array) sederhana
func TestJSONArray(t *testing.T) {
	// Membuat object Customer dengan data hobi berupa slice string
	customer := Customer{
		FirstName: "Hilmi",                              // Nama depan
		LastName:  "Yahya",                              // Nama belakang
		Hobbies:   []string{"Gaming", "Watching", "Coding"}, // Slice hobi
	}

	// Mengubah struct Customer menjadi JSON
	bytes, _ := json.Marshal(customer)
	// Menampilkan hasil JSON
	fmt.Println(string(bytes))
}

// Test untuk decode JSON array sederhana ke struct
func TestJSONArrayDecode(t *testing.T) {
	// JSON string yang berisi data customer dengan array hobbies
	jsonString := `{"FirstName":"Hilmi","LastName":"Yahya","Hobbies":["Gaming","Watching","Coding"]}`
	// Mengubah string JSON menjadi byte slice
	jsonBytes := []byte(jsonString)

	// Membuat pointer ke struct Customer sebagai tujuan decode
	customer := &Customer{}
	// Mengubah JSON menjadi struct Customer
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err) // Menghentikan program jika terjadi error
	}

	// Menampilkan hasil decode
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

// Test untuk encode JSON dengan array of object (nested / complex array)
func TestJSONArrayComplex(t *testing.T) {
	// Membuat Customer dengan field Addresses berupa slice Address
	customer := Customer{
		FirstName: "Hilmi", // Nama depan
		Addresses: []Address{
			{
				Street:     "Jalan A",     // Nama jalan
				Country:    "Indonesia",   // Negara
				PostalCode: "111111",      // Kode pos
			},
			{
				Street:     "Jalan B",
				Country:    "Indonesia",
				PostalCode: "222222",
			},
		},
	}

	// Mengubah struct Customer menjadi JSON
	bytes, _ := json.Marshal(customer)
	// Menampilkan hasil JSON
	fmt.Println(string(bytes))
}

// Test untuk decode JSON kompleks (array di dalam object)
func TestJSONArrayComplexDecode(t *testing.T) {
	// JSON string dengan field Addresses berupa array object
	jsonString := `{"FirstName":"Hilmi","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan A","Country":"Indonesia","PostalCode":"111111"},{"Street":"Jalan B","Country":"Indonesia","PostalCode":"222222"}]}`
	// Mengubah JSON string menjadi byte slice
	jsonBytes := []byte(jsonString)

	// Membuat pointer ke Customer untuk menampung hasil decode
	customer := &Customer{}
	// Decode JSON ke struct Customer
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	// Menampilkan hasil decode
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

// Test untuk decode JSON yang isinya hanya array object (tanpa wrapper struct)
func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	// JSON string berupa array Address
	jsonString := `[{"Street":"Jalan A","Country":"Indonesia","PostalCode":"111111"},{"Street":"Jalan B","Country":"Indonesia","PostalCode":"222222"}]`
	// Mengubah JSON string menjadi byte slice
	jsonBytes := []byte(jsonString)

	// Membuat pointer ke slice Address sebagai tujuan decode
	addresses := &[]Address{}
	// Decode JSON array ke slice Address
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	// Menampilkan hasil decode
	fmt.Println(addresses)
}

// Test untuk encode slice Address langsung menjadi JSON array
func TestOnlyJSONArrayComplex(t *testing.T) {
	// Membuat slice Address
	addresses := []Address{
		{
			Street:     "Jalan A",
			Country:    "Indonesia",
			PostalCode: "111111",
		},
		{
			Street:     "Jalan B",
			Country:    "Indonesia",
			PostalCode: "222222",
		},
	}

	// Mengubah slice Address menjadi JSON array
	bytes, _ := json.Marshal(addresses)
	// Menampilkan hasil JSON
	fmt.Println(string(bytes))
}

// Kesimpulan:
// Kode ini menunjukkan berbagai skenario penggunaan JSON Array di Go, mulai dari encode dan decode array sederhana, array kompleks (nested object), hingga JSON yang hanya berisi array tanpa pembungkus struct. Melalui json.Marshal dan json.Unmarshal, data Go dapat dengan mudah dikonversi ke format JSON dan sebaliknya, yang sangat umum digunakan dalam komunikasi API, penyimpanan data, dan pertukaran data antar sistem.
