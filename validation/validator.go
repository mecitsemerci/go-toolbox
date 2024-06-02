package validation

// Validator is a generic type that validates entities based on provided rules.
// The type parameter T represents the type of the entity to be validated.
type Validator[T any] struct {
	rules []func(T) error // rules to be applied on the entity
}

// NewValidator creates a new instance of Validator.
// It returns a pointer to the newly created Validator instance.
func NewValidator[T any]() *Validator[T] {
	return &Validator[T]{}
}

// AddRule adds a new validation rule to the Validator.
// The rule is a function that takes an entity of type T and returns an error if the rule fails.
// The method returns a pointer to the Validator instance to allow chaining of method calls.
func (v *Validator[T]) AddRule(rule func(T) error) *Validator[T] {
	v.rules = append(v.rules, rule)
	return v
}

// Validate applies all the added rules to the given entity.
// It iterates over each rule and calls it with the entity as the argument.
// If any rule fails (returns a non-nil error), the method returns that error immediately.
// If all rules pass, the method returns nil.
func (v *Validator[T]) Validate(entity T) error {
	for _, rule := range v.rules {
		if err := rule(entity); err != nil {
			return err
		}
	}
	return nil
}
