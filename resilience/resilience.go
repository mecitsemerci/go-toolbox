package resilience

// Resilience struct holds the policies for retry, circuit breaker and fallback.
type Resilience struct {
	RetryPolicy          *RetryPolicy
	CircuitBreakerPolicy *CircuitBreakerPolicy
	FallbackPolicy       *FallbackPolicy
}

// NewResilience returns a new instance of Resilience with default policies set to nil.
func NewResilience() *Resilience {
	return &Resilience{}
}

// WithRetry sets the retry policy for the Resilience instance.
// Returns the Resilience instance for method chaining.
func (r *Resilience) WithRetry(retryPolicy *RetryPolicy) *Resilience {
	r.RetryPolicy = retryPolicy
	return r
}

// WithCircuitBreaker sets the circuit breaker policy for the Resilience instance.
// Returns the Resilience instance for method chaining.
func (r *Resilience) WithCircuitBreaker(circuitBreakerPolicy *CircuitBreakerPolicy) *Resilience {
	r.CircuitBreakerPolicy = circuitBreakerPolicy
	return r
}

// WithFallback sets the fallback policy for the Resilience instance.
// Returns the Resilience instance for method chaining.
func (r *Resilience) WithFallback(fallbackPolicy *FallbackPolicy) *Resilience {
	r.FallbackPolicy = fallbackPolicy
	return r
}

// Execute executes the provided action with the configured policies.
// It first applies the retry policy, then the circuit breaker policy, and finally the fallback policy.
// Returns the error returned by the action or the fallback policy if applicable.
func (r *Resilience) Execute(action func() error) error {
	var err error
	if r.RetryPolicy != nil {
		err = r.RetryPolicy.Execute(action)
	} else {
		err = action()
	}

	if err != nil && r.CircuitBreakerPolicy != nil {
		err = r.CircuitBreakerPolicy.Execute(action)
	}

	if err != nil && r.FallbackPolicy != nil {
		return r.FallbackPolicy.Execute(action)
	}

	return err
}
