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
