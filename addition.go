// Package addition is used for the addition of two numbers.
package addition

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds two numbers and returns the result.
//
// More information about addition can be found at [Math is fun].
//
// [Math is fun]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](x, y T) T {
	return x + y
}
