package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/mecitsemerci/go-toolbox/resilience"
)

func main() {
	retryPolicy := resilience.NewRetryPolicy(3, 2*time.Second)
	cb := resilience.NewCircuitBreakerPolicy(2, 5*time.Second)
	fallbackPolicy := resilience.NewFallbackPolicy(func() error {
		fmt.Println("Executing fallback action")
		return nil
	})

	res := resilience.NewResilience().
		WithRetry(retryPolicy).
		WithCircuitBreaker(cb).
		WithFallback(fallbackPolicy)

	err := res.Execute(func() error {
		fmt.Println("Executing main action")
		return errors.New("main action failed")
	})

	if err != nil {
		fmt.Printf("Final error: %v\n", err)
	}
}
