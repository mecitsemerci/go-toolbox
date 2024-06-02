package resilience

import (
	"errors"
	"sync"
	"time"
)

// ErrCircuitOpen is returned when the circuit breaker is open and a request is attempted.
var ErrCircuitOpen = errors.New("circuit breaker is open")

// State represents the state of the circuit breaker.
type State int

const (
	// Closed is the initial state of the circuit breaker.
	Closed State = iota
	// Open means that the circuit breaker is open and will not allow requests.
	Open
	// HalfOpen means that the circuit breaker is half open and will allow a single request.
	HalfOpen
)

// CircuitBreakerPolicy is a simple implementation of a circuit breaker pattern.
type CircuitBreakerPolicy struct {
	failureThreshold int
	recoveryTimeout  time.Duration
	state            State
	failureCount     int
	lastAttempt      time.Time
	mutex            sync.Mutex
}

// NewCircuitBreakerPolicy creates a new CircuitBreakerPolicy with the given failure threshold and recovery timeout.
func NewCircuitBreakerPolicy(failureThreshold int, recoveryTimeout time.Duration) *CircuitBreakerPolicy {
	return &CircuitBreakerPolicy{
		failureThreshold: failureThreshold,
		recoveryTimeout:  recoveryTimeout,
		state:            Closed,
	}
}

// Execute executes the given action and manages the circuit breaker state.
// If the circuit breaker is open, it returns ErrCircuitOpen.
// If the circuit breaker is half open, it allows a single request and updates the state accordingly.
// If the action returns an error, it increments the failure count and updates the state if necessary.
// If the action returns no error, it resets the failure count.
func (cbp *CircuitBreakerPolicy) Execute(action func() error) error {
	cbp.mutex.Lock()
	defer cbp.mutex.Unlock()

	switch cbp.state {
	case Open:
		if time.Since(cbp.lastAttempt) > cbp.recoveryTimeout {
			cbp.state = HalfOpen
		} else {
			return ErrCircuitOpen
		}
	case HalfOpen:
		if time.Since(cbp.lastAttempt) > cbp.recoveryTimeout {
			cbp.state = Closed
		}
	}

	err := action()
	if err != nil {
		cbp.failureCount++
		if cbp.failureCount >= cbp.failureThreshold {
			cbp.state = Open
		}
		cbp.lastAttempt = time.Now()
	} else {
		cbp.failureCount = 0
	}

	return err
}
