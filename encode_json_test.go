package belajar_golang_json // deklarasi package untuk contoh encoding JSON di Go

import (
	"encoding/json" // package untuk encode/decode JSON di Go
	"fmt"           // untuk menampilkan output ke console
	"testing"       // package bawaan Go untuk unit test
)

// fungsi helper untuk mengubah data apapun menjadi JSON lalu mencetaknya
func logJson(data interface{}) { // interface{} → bisa menerima tipe data apa saja
	bytes, err := json.Marshal(data) // ubah data Go menjadi byte JSON (encode)
	if err != nil {                  // cek jika terjadi error saat encoding
		panic(err) // hentikan program jika gagal encode
	}
	fmt.Println(string(bytes)) // konversi byte ke string lalu cetak JSON
}

// fungsi unit test untuk mencoba proses encode JSON berbagai tipe data
func TestEncode(t *testing.T) {
	logJson("Hilmi")                    // encode string ke JSON
	logJson(1)                          // encode angka ke JSON
	logJson(true)                       // encode boolean ke JSON
	logJson([]string{"Hilmi", "Yahya"}) // encode slice/array string ke JSON array
}

// Kesimpulan:
// Kode ini mendemonstrasikan proses encoding (marshal) data Go menjadi format JSON menggunakan json.Marshal. Fungsi logJson dibuat fleksibel dengan parameter interface{}, sehingga bisa menerima berbagai tipe data seperti string, angka, boolean, dan slice. TestEncode menjalankan beberapa contoh pemanggilan untuk melihat hasil JSON yang dihasilkan.
