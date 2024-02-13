package Add

import "golang.org/x/exp/constraints"

// Add function takes two integer parameters and return sum of two integers
//
// https://www.mathsisfun.com/numbers/addition.html

type Number interface {
	constraints.Integer | constraints.Float
}

func Add[T Number](v1, v2 T) T {
	return v1 + v2
}
