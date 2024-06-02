# Go Generic Validator

This package provides a generic `Validator` type that can be used to validate entities based on provided rules.

## Installation

To use this package, first install it using Go modules:

```bash
go get github.com/mecitsemerci/go-toolbox/validation
```

## Usage

Here are some examples of how to use the `Validator` type:

### Example 1: Validating a simple struct

```go
package main

import (
    "errors"
    "fmt"
    "github.com/mecitsemerci/go-toolbox/validation"
)

type User struct {
    Name string
    Age  int
}

func main() {
    validator := validation.NewValidator[User]()

    validator.AddRule(func(user User) error {
        if user.Name == "" {
            return errors.New("name is required")
        }
        return nil
    })

    validator.AddRule(func(user User) error {
        if user.Age < 18 {
            return errors.New("age must be at least 18")
        }
        return nil
    })

    user := User{Name: "John Doe", Age: 25}
    if err := validator.Validate(user); err!= nil {
        fmt.Println("Validation failed:", err)
    } else {
        fmt.Println("Validation passed")
    }
}
```

### Example 2: Validating a slice of structs

```go
package main

import (
    "errors"
    "fmt"
    "github.com/mecitsemerci/go-toolbox/validation"
)

type Product struct {
    Name  string
    Price float64
}

func main() {
    products := []Product{
        {Name: "Product 1", Price: 10.99},
        {Name: "", Price: 5.99},
        {Name: "Product 3", Price: -2.99},
    }

    validator := validation.NewValidator[Product]()

    validator.AddRule(func(product Product) error {
        if product.Name == "" {
            return errors.New("name is required")
        }
        return nil
    })

    validator.AddRule(func(product Product) error {
        if product.Price <= 0 {
            return errors.New("price must be greater than 0")
        }
        return nil
    })

    for _, product := range products {
        if err := validator.Validate(product); err!= nil {
            fmt.Printf("Validation failed for product %s: %v\n", product.Name, err)
        } else {
            fmt.Printf("Validation passed for product %s\n", product.Name)
        }
    }
}
```

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/mecitsemerci/go-toolbox/blob/main/LICENSE) file for more details.