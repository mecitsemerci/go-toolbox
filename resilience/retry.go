package resilience

import (
	"time"
)

// RetryPolicy represents a retry policy with a maximum number of retries and a delay between retries.
type RetryPolicy struct {
	MaxRetries int
	Delay      time.Duration
}

// NewRetryPolicy creates a new RetryPolicy with the given maximum number of retries and delay between retries.
func NewRetryPolicy(maxRetries int, delay time.Duration) *RetryPolicy {
	return &RetryPolicy{
		MaxRetries: maxRetries,
		Delay:      delay,
	}
}

// Execute executes the given action with the retry policy.
// It will retry the action up to MaxRetries times, waiting for Delay between each retry.
// If the action returns nil, it will return nil.
// If all retries fail, it will return the last error from the action.
func (r *RetryPolicy) Execute(action func() error) error {
	var err error
	for i := 0; i < r.MaxRetries; i++ {
		err = action()
		if err == nil {
			return nil
		}
		time.Sleep(r.Delay)
	}
	return err
}
