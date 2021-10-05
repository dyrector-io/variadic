package variadic_test

import (
	"strings"
	"testing"

	"github.com/dyrector-io/variadic"
	"github.com/stretchr/testify/assert"
)

func TestJoins(t *testing.T) {
	expected := "this-is-a-test"

	// normally you do this when concatenating strings, with a separator
	res := strings.Join([]string{"this", "is", "a", "test"}, "-")

	assert.Equal(t, expected, res)

	// i dont know about you, but i don't like creating arrays on the fly
	res = variadic.JoinV("-", "this", "is", "a", "test")
	assert.Equal(t, expected, res)

	// variadic join also works with an array
	res = variadic.JoinV("-", []string{"this", "is", "a", "test"}...)
	assert.Equal(t, expected, res)
}
