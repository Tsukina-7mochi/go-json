# go-json

A JSON decoder/encoder implementation for practicing Go reflection.
Supports Boolean values, strings, numbers, arrays, objects, conversion of field names, and Struct Tags.

## Example

See [./example/main.go](./example/main.go) for more details.

```go
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

    // decode (parse) JSON into struct
    var user User
    err := json.Decode(input, &user)
    if err != nil {
        panic(err)
    }

    // encode (stringify) struct into JSON
    encoded, err := json.Encode(user)
    if err != nil {
        panic(err)
    }

    println(encoded)
}
```
