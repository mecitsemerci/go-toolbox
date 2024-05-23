# jsonconv

A simple Go package for converting Go values to JSON and vice versa, handling errors appropriately.

## Installation

To install the package, run the following command:

```sh
go get -u github.com/mecitsemerci/go-toolbox/jsonconv
```


## Usage

Here's an example of how to use the `jsonconv` package:

```go
package main

import (
	"log"
	"time"

	"github.com/mecitsemerci/go-toolbox/jsonconv"
)

// Person represents a person with a name, age, admin status, email, creation time, and roles.
type Person struct {
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	IsAdmin   bool      `json:"is_admin"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Roles     []string  `json:"roles,omitempty"`
}

func main() {
	// Create a Person struct with sample data.
	person := Person{
		Name:      "John Doe",
		Age:       30,
		IsAdmin:   true,
		Email:     "john.doe@example.com",
		CreatedAt: time.Now(),
		Roles:     []string{"admin", "user"},
	}

	// Serialize the Person struct to JSON using jsonconv.Serialize function.
	// Returns a JSON string and an error if any.
	jsonStr, err := jsonconv.Serialize(person)
	if err != nil {
		log.Fatalf("Error serializing: %v", err)
	}
	log.Printf("Serialized JSON: %s\n", jsonStr)

	// Deserialize the JSON string back to a Person struct using jsonconv.Deserialize function.
	// Returns a Person struct and an error if any.
	deserializedPerson, err := jsonconv.Deserialize[Person](jsonStr)
	if err != nil {
		log.Fatalf("Error deserializing: %v", err)
	}
	log.Printf("Deserialized Person: %+v\n", deserializedPerson)

	// Serialize the Person struct to JSON using the convenience function jsonconv.SerializeCool.
	// Returns a formatted JSON string.
	coolJsonStr := jsonconv.SerializeCool(person)
	log.Printf("Serialized JSON (cool version): %s\n", coolJsonStr)

	// DeserializeInto the JSON string back to a Person struct using jsonconv.DeserializeInto function.
	// Returns an error if any.
	var deserializedIntoPerson Person
	err = jsonconv.DeserializeInto(jsonStr, &deserializedIntoPerson)
	if err != nil {
		log.Fatalf("Error deserializing into: %v", err)
	}
	log.Printf("DeserializedInto Person: %+v\n", deserializedIntoPerson)
}

```

```bash

Serialized JSON: {"name":"John Doe","age":30,"is_admin":true,"email":"john.doe@example.com","created_at":"2024-05-23T10:29:31.612383+03:00","roles":["admin","user"]}
Deserialized Person: {Name:John Doe Age:30 IsAdmin:true Email:john.doe@example.com CreatedAt:2024-05-23 10:29:31.612383 +0300 +03 Roles:[admin user]}
Serialized JSON (cool version): {"name":"John Doe","age":30,"is_admin":true,"email":"john.doe@example.com","created_at":"2024-05-23T10:29:31.612383+03:00","roles":["admin","user"]}
DeserializedInto Person: {Name:John Doe Age:30 IsAdmin:true Email:john.doe@example.com CreatedAt:2024-05-23 10:29:31.612383 +0300 +03 Roles:[admin user]}

```

This example demonstrates how to serialize a `Person` struct to JSON using the `Serialize` function, deserialize the JSON string back to a `Person` struct using the `Deserialize` function, and use the convenience function `SerializeCool` to serialize the struct without checking for errors.

## Functions

The `jsonconv` package provides the following functions:

- `Serialize(value interface{}) (string, error)`: Converts a Go value `v` to JSON and returns the JSON string and an error if any.
- `SerializeCool(value interface{}) string`: Converts a Go value `v` to JSON and returns the JSON string. This function does not handle errors.
- `Deserialize[YourStruct](jsonStr string) (yourStruct, error)`: Converts a JSON string `jsonStr` to a Go struct `yourStruct` and returns an error if any.
- `DeserializeInto(jsonStr string, result *yourStruct) error`: Converts a JSON string `jsonStr` to a Go value pointed to by `result` and returns an error if any.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/mecitsemerci/go-toolbox/blob/main/LICENSE) file for more details.


