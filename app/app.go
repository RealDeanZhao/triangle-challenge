package app

// TriangleType is the alias of type int that represents the type of the triangles
type TriangleType int

/*
* None: There is an error such as the float64 overflow or it is not a triangle
* Scalene: No sides of the triangle are equal
* Isosceles: Any two sides of the triangle are equal
* Equilateral: All sides of the triangle are equal
 */
const (
	None = TriangleType(iota)
	Scalene
	Isosceles
	Equilateral
)

// DetermineTriangleType is used to determine the type by the input
// It will first check if it is a triangle
// Then it will check the type.
func DetermineTriangleType(a float64, b float64, c float64) TriangleType {
	if !IsTriangle(a, b, c) {
		return None
	}
	if AllSidesAreEqual(a, b, c) {
		return Equilateral
	} else if TwoSidesAreEqual(a, b, c) {
		return Isosceles
	}
	return Scalene
}

// IsTriangle is used to determine if it is a triangle by the input.
// It will first check if all the sides are positive numbers.
// Then it will check if the sum of any two sides is greater than the third side.
func IsTriangle(a float64, b float64, c float64) bool {
	if AllSidesArePositive(a, b, c) {
		return AnyTwoSidesAreGreaterThanTheThrid(a, b, c)
	}
	return false
}

// AllSidesArePositive is used to check if all the sides are positive numbers
func AllSidesArePositive(a float64, b float64, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	return true
}

// AnyTwoSidesAreGreaterThanTheThrid is used to check if the sum of any two sides is greater than the third side.
func AnyTwoSidesAreGreaterThanTheThrid(a float64, b float64, c float64) bool {
	// We may need to consider about the float64 overflow if a + b > the max value of float64
	if a+b <= c || a+c <= b || b+c <= a {
		return false
	}
	return true
}

// AllSidesAreEqual is used to check if all the sides are equal or not.
func AllSidesAreEqual(a float64, b float64, c float64) bool {
	if a == b && b == c {
		return true
	}
	return false
}

// TwoSidesAreEqual is used to check if any two sides are equal or not.
func TwoSidesAreEqual(a float64, b float64, c float64) bool {
	if a == b || a == c || b == c {
		return true
	}
	return false
}
