package resilience

// FallbackPolicy is a struct that encapsulates a fallback action to be executed when a primary action fails.
type FallbackPolicy struct {
	fallbackAction func() error // The fallback action to be executed when the primary action fails.
}

// NewFallbackPolicy creates a new instance of FallbackPolicy with the provided fallback action.
//
// Parameters:
// - fallbackAction: A function that represents the fallback action to be executed when the primary action fails.
//
// Returns:
// - A pointer to a new instance of FallbackPolicy.
func NewFallbackPolicy(fallbackAction func() error) *FallbackPolicy {
	return &FallbackPolicy{
		fallbackAction: fallbackAction,
	}
}

// Execute executes the primary action and, if it fails, executes the fallback action.
//
// Parameters:
// - action: A function that represents the primary action to be executed.
//
// Returns:
// - An error if the primary action fails and the fallback action is executed, or nil if the primary action succeeds.
func (fp *FallbackPolicy) Execute(action func() error) error {
	err := action()
	if err != nil {
		return fp.fallbackAction()
	}
	return nil
}
