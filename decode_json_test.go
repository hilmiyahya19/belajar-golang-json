package belajar_golang_json // package contoh untuk belajar encoding/decoding JSON di Go

import (
	"encoding/json" // package untuk konversi JSON ↔ struct Go
	"fmt"           // untuk menampilkan output ke console
	"testing"       // package bawaan Go untuk unit test
)

func TestDecodeJSON(t *testing.T) { // fungsi unit test untuk decode JSON ke struct
	// string JSON sebagai sumber data
	jsonString := `{"FirstName":"Hilmi","LastName":"Yahya","Age":22,"Married":false}`

	jsonBytes := []byte(jsonString) // ubah string JSON menjadi slice byte (format input Unmarshal)

	customer := &Customer{} // buat pointer ke struct Customer kosong (akan diisi oleh Unmarshal)

	// decode (unmarshal) JSON → isi ke struct customer berdasarkan nama field
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil { // cek jika terjadi error saat proses decode
		panic(err) // hentikan program jika decode gagal
	}

	fmt.Println(customer)            // cetak seluruh isi struct customer
	fmt.Println(customer.FirstName)  // cetak field FirstName
	fmt.Println(customer.LastName)   // cetak field LastName
}

// Kesimpulan:
// Kode ini melakukan proses decode JSON menjadi struct Customer menggunakan json.Unmarshal. Data JSON diubah dulu menjadi []byte lalu dimasukkan ke struct melalui pointer agar bisa diisi. Field pada struct akan terisi otomatis jika nama field cocok dan bersifat exported (huruf besar). Test ini menunjukkan bahwa data JSON dapat langsung dipetakan ke struct Go tanpa pengisian manual.
