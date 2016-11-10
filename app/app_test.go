package app

import (
	"testing"
)

func TestAllSidesArePositive(t *testing.T) {
	if AllSidesArePositive(-1, 4, 5) {
		t.Error("AllSidesArePositive(-1, 4, 5) should return false")
	}
	if AllSidesArePositive(1, -1, 5) {
		t.Error("AllSidesArePositive(1, -1, 5) should return false")
	}
	if AllSidesArePositive(1, 1, -5) {
		t.Error("AllSidesArePositive(1, 1, -5) should return false")
	}
	if !AllSidesArePositive(1, 1, 1) {
		t.Error("AllSidesArePositive(1, 1, 1) should return true")
	}
}

func TestAnyTwoSidesAreGreaterThanTheThrid(t *testing.T) {
	if !AnyTwoSidesAreGreaterThanTheThrid(3, 4, 5) {
		t.Error("AnyTwoSidesAreGreaterThanTheThrid(3, 4, 5) should return true")
	}
	if AnyTwoSidesAreGreaterThanTheThrid(1, 4, 5) {
		t.Error("AnyTwoSidesAreGreaterThanTheThrid(1, 4, 5) should return false")
	}
	if AnyTwoSidesAreGreaterThanTheThrid(0, 4, 5) {
		t.Error("AnyTwoSidesAreGreaterThanTheThrid(0, 4, 5) should return false")
	}
}

func TestIsTriangle(t *testing.T) {
	if !IsTriangle(3, 4, 5) {
		t.Error("IsTriangle(3, 4, 5) should return true")
	}
	if IsTriangle(1, 2, 5) {
		t.Error("IsTriangle(1, 2, 5) should return false")
	}
	if IsTriangle(-1, 2, 5) {
		t.Error("IsTriangle(-1, 2, 5) should return false")
	}
}

func TestAllSidesAreEqual(t *testing.T) {
	if !AllSidesAreEqual(3, 3, 3) {
		t.Error("AllSidesAreEqual(3, 3, 3) should return true")
	}
	if AllSidesAreEqual(3, 4, 5) {
		t.Error("AllSidesAreEqual(3, 4, 5) should return false")
	}
}

func TestTwoSidesAreEqual(t *testing.T) {
	if !TwoSidesAreEqual(3, 3, 4) {
		t.Error("TwoSidesAreEqual(3, 3, 4) should return true")
	}

	if TwoSidesAreEqual(3, 4, 5) {
		t.Error("TwoSidesAreEqual(3, 4, 5) should return false")
	}
}

func TestDetermineTriangleType(t *testing.T) {
	if DetermineTriangleType(1, 1, 2) != None {
		t.Error("DetermineTriangleType(1, 1, 2) should return None")
	}
	if DetermineTriangleType(3, 4, 5) != Scalene {
		t.Error("DetermineTriangleType(3, 4, 5) should return Scalene")
	}
	if DetermineTriangleType(4, 4, 5) != Isosceles {
		t.Error("DetermineTriangleType(3, 4, 5) should return Isosceles")
	}
	if DetermineTriangleType(4, 4, 4) != Equilateral {
		t.Error("DetermineTriangleType(3, 4, 5) should return Equilateral")
	}
}
