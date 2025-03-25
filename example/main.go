package main

import (
	"fmt"
	"json"
)

type Address struct {
	Street string
	City   string
	State  string
	Zip    string `json:"zip_code"`
}

type User struct {
	ID        int `json:"id"`
	Name      string
	Email     string
	Addresses []Address
}

func printUser(user User) {
	fmt.Printf("ID: %d\n", user.ID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Email: %s\n", user.Email)

	for i, address := range user.Addresses {
		fmt.Printf("Address %d: %s, %s, %s %s\n", i+1, address.Street, address.City, address.State, address.Zip)
	}
}

func main() {
	input := `{
        "id": 1,
        "name": "John Doe",
        "email": "johndoe@example.com",
        "addresses": [
            {
                "street": "123 Main St",
                "city": "Springfield",
                "state": "IL",
                "zip_code": "62701"
            },
            {
                "street": "456 Elm St",
                "city": "Springfield",
                "state": "IL",
                "zip_code": "62701"
            }
        ]
    }`

	var user User
	err := json.Decode(input, &user)
	if err != nil {
		panic(err)
	}

	printUser(user)

	encoded, err := json.Encode(user)
	if err != nil {
		panic(err)
	}
	println(encoded)
}
