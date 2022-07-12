package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Lukman",
		MiddleName: "Sanjaya",
		LastName:   "Ok",
		Hobbies:    []string{"Coding", "Gaming", "Sleeping"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Lukman","MiddleName":"Sanjaya","LastName":"Ok","Age":0,"Married":false,"Hobbies":["Coding","Gaming","Sleeping"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Lukman",
		Hobbies:   []string{"Coding", "Gaming", "Sleeping"},
		Addresses: []Address{
			{
				Street:     "Jalan Belum Ada",
				Country:    "Indonesia",
				PostalCode: "123",
			},
			{
				Street:     "Jalan Sudah Ada",
				Country:    "Indonesia",
				PostalCode: "456",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Lukman","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":["Coding","Gaming","Sleeping"],"Addresses":[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"123"},{"Street":"Jalan Sudah Ada","Country":"Indonesia","PostalCode":"456"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}

func TestOnlyeJsonArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"123"},{"Street":"Jalan Sudah Ada","Country":"Indonesia","PostalCode":"456"}]`
	jsonBytes := []byte(jsonString)

	address := &[]Address{}
	err := json.Unmarshal(jsonBytes, address)
	if err != nil {
		panic(err)
	}
	fmt.Println(address)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan Belum Ada",
			Country:    "Indonesia",
			PostalCode: "123",
		},
		{
			Street:     "Jalan Sudah Ada",
			Country:    "Indonesia",
			PostalCode: "456",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
