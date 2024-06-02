package resilience

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRetryPolicy(t *testing.T) {
	retryPolicy := NewRetryPolicy(3, time.Millisecond*10)

	var attempt int
	action := func() error {
		attempt++
		if attempt < 3 {
			return errors.New("failed")
		}
		return nil
	}

	err := retryPolicy.Execute(action)
	assert.NoError(t, err)
	assert.Equal(t, 3, attempt)
}

func TestRetryPolicyMaxRetries(t *testing.T) {
	retryPolicy := NewRetryPolicy(2, time.Millisecond*10)

	var attempt int
	action := func() error {
		attempt++
		return errors.New("failed")
	}

	err := retryPolicy.Execute(action)
	assert.Error(t, err)
	assert.Equal(t, 2, attempt)
}
