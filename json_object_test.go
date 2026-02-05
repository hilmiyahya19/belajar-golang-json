package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName string
	LastName string
	Age int
	Married bool
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Hilmi",
		LastName: "Yahya",
		Age: 22,
		Married: false,
	}
	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}