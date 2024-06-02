package resilience

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCircuitBreakerPolicyExecute_WhenFailureThresholdIsNotReachedItAllowsRequests(t *testing.T) {
	cbp := NewCircuitBreakerPolicy(3, time.Minute)

	for i := 0; i < 2; i++ {
		err := cbp.Execute(func() error { return nil })
		require.NoError(t, err)
	}

	err := cbp.Execute(func() error { return errors.New("failed") })
	assert.Error(t, err)
}

func TestCircuitBreakerPolicyExecute_WhenFailureThresholdIsReachedItOpensTheCircuitBreaker(t *testing.T) {
	cbp := NewCircuitBreakerPolicy(3, time.Minute)

	for i := 0; i < 3; i++ {
		err := cbp.Execute(func() error { return errors.New("failed") })
		require.Error(t, err)
	}

	err := cbp.Execute(func() error { return nil })
	assert.EqualError(t, err, ErrCircuitOpen.Error())
}

func TestCircuitBreakerPolicyExecute_WhenCircuitBreakerIsOpenItDoesNotAllowRequestsUntilRecoveryTimeout(t *testing.T) {
	cbp := NewCircuitBreakerPolicy(3, 500*time.Millisecond)

	for i := 0; i < 3; i++ {
		err := cbp.Execute(func() error { return errors.New("failed") })
		require.Error(t, err)
	}

	time.Sleep(1 * time.Second)

	err := cbp.Execute(func() error { return nil })
	assert.NoError(t, err)

}

func TestCircuitBreakerPolicy_Execute_Closed(t *testing.T) {
	// Arrange
	cbp := NewCircuitBreakerPolicy(3, 10*time.Second)
	action := func() error { return nil }

	// Act
	err := cbp.Execute(action)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, Closed, cbp.state)
}

func TestCircuitBreakerPolicy_Execute_Open(t *testing.T) {
	// Arrange
	cbp := NewCircuitBreakerPolicy(3, 10*time.Second)
	action := func() error { return errors.New("error") }

	// Act
	for i := 0; i < 3; i++ {
		err := cbp.Execute(action)
		assert.Error(t, err)
	}

	// Assert
	assert.Equal(t, Open, cbp.state)
}

func TestCircuitBreakerPolicy_Execute_HalfOpen(t *testing.T) {
	// Arrange
	cbp := NewCircuitBreakerPolicy(3, 10*time.Second)
	action := func() error { return errors.New("error") }

	// Act
	for i := 0; i < 3; i++ {
		err := cbp.Execute(action)
		assert.Error(t, err)
	}

	// Assert
	assert.Equal(t, Open, cbp.state)

	// Wait for recovery timeout
	time.Sleep(11 * time.Second)

	// Act
	err := cbp.Execute(func() error { return nil })

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, HalfOpen, cbp.state)
}

