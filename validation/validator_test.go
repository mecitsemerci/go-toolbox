package validation

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestUser struct {
	Name      string
	Age       int
	CreatedAt time.Time
	Roles     []string
	IsActive  bool

}

func TestNotEmpty_String(t *testing.T) {
	v := NewValidator[TestUser]().
		AddRule(NotEmpty(func(u TestUser) string { return u.Name }))

	tests := []struct {
		user TestUser
		want bool
	}{
		{TestUser{Name: "Alice"}, true},
		{TestUser{Name: ""}, false},
	}

	for _, tt := range tests {
		err := v.Validate(tt.user)
		if tt.want {
			assert.NoError(t, err, "Expected no error for user: %v", tt.user)
		} else {
			assert.Error(t, err, "Expected error for user: %v", tt.user)
		}
	}
}

func TestNotEmpty_Array(t *testing.T) {
	v := NewValidator[TestUser]().
		AddRule(NotEmpty(func(u TestUser) []string { return u.Roles }))

	tests := []struct {
		user TestUser
		want bool
	}{
		{TestUser{Name: "Alice", Roles: []string{"Admin", "User"}}, false},
		{TestUser{Name: "", Roles: []string{}}, true},
	}

	for _, tt := range tests {
		err := v.Validate(tt.user)
		if tt.want {
			assert.NoError(t, err, "Expected no error for user: %v", tt.user)
		} else {
			assert.Error(t, err, "Expected error for user: %v", tt.user)
		}
	}
}

func TestMaxLength(t *testing.T) {
	v := NewValidator[TestUser]().
		AddRule(MaxLength(func(u TestUser) string { return u.Name }, 5))

	tests := []struct {
		user TestUser
		want bool
	}{
		{TestUser{Name: "Alice"}, true},
		{TestUser{Name: "Alexander"}, false},
	}

	for _, tt := range tests {
		err := v.Validate(tt.user)
		if tt.want {
			assert.NoError(t, err, "Expected no error for user: %v", tt.user)
		} else {
			assert.Error(t, err, "Expected error for user: %v", tt.user)
		}
	}
}

func TestMinLength(t *testing.T) {
	v := NewValidator[TestUser]().
		AddRule(MinLength(func(u TestUser) string { return u.Name }, 3))

	tests := []struct {
		user TestUser
		want bool
	}{
		{TestUser{Name: "Ali"}, true},
		{TestUser{Name: "Al"}, false},
	}

	for _, tt := range tests {
		err := v.Validate(tt.user)
		if tt.want {
			assert.NoError(t, err, "Expected no error for user: %v", tt.user)
		} else {
			assert.Error(t, err, "Expected error for user: %v", tt.user)
		}
	}
}

func TestIsInRange(t *testing.T) {
	v := NewValidator[TestUser]().
		AddRule(IsInRange(func(u TestUser) int { return u.Age }, 18, 60))

	tests := []struct {
		user TestUser
		want bool
	}{
		{TestUser{Age: 30}, true},
		{TestUser{Age: 10}, false},
		{TestUser{Age: 70}, false},
	}

	for _, tt := range tests {
		err := v.Validate(tt.user)
		if tt.want {
			assert.NoError(t, err, "Expected no error for user: %v", tt.user)
		} else {
			assert.Error(t, err, "Expected error for user: %v", tt.user)
		}
	}
}

func TestCombinedRules(t *testing.T) {
	v := NewValidator[TestUser]().
		AddRule(NotEmpty(func(u TestUser) string { return u.Name })).
		AddRule(MaxLength(func(u TestUser) string { return u.Name }, 5)).
		AddRule(IsInRange(func(u TestUser) int { return u.Age }, 18, 60))

	tests := []struct {
		user TestUser
		want bool
	}{
		{TestUser{Name: "Alice", Age: 30}, true},
		{TestUser{Name: "", Age: 30}, false},
		{TestUser{Name: "Alexander", Age: 30}, false},
		{TestUser{Name: "Alice", Age: 70}, false},
	}

	for _, tt := range tests {
		err := v.Validate(tt.user)
		if tt.want {
			assert.NoError(t, err, "Expected no error for user: %v", tt.user)
		} else {
			assert.Error(t, err, "Expected error for user: %v", tt.user)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	var (
		val      *int
		emptyVal any
		emptyStr string
		emptyInt *int
	)
	number := 5
	val = &number
	tests := []struct {
		val  any
		want bool
	}{
		{emptyStr, true},
		{[]int{}, true},
		{map[string]any{}, true},
		{make([]int, 0), true},
		{"Foo", false},
		{[]int{1, 2, 3}, false},
		{map[string]string{"foo": "bar"}, false},
		{emptyInt, true},
		{emptyVal, true},
		{nil, true},
		{val, false},
	}
	for _, tt := range tests {
		result := isEmpty(tt.val)
		assert.Equal(t, tt.want, result)
	}
}


