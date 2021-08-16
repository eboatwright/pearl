package pearl


import (
	"math"
	"reflect"
)


// The Vector2 struct
type Vector2 struct {
	X float64
	Y float64
}

// Returns the magnitude of a Vector2 type
func (v *Vector2) Magnitude() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// Normalizes Vector2
func (v *Vector2) Normalize() {
	magnitude := v.Magnitude()
	if magnitude > 0 {
		v.Divide(magnitude)
	}
}

// Adds Vector2 and other (changes Vector2)
func (v *Vector2) Add(other interface {}) {
	t := reflect.TypeOf(other).String()
	if t == "pearl.Vector2" {
		otherVector := other.(Vector2)
		v.X += otherVector.X
		v.Y += otherVector.Y
	} else {
		val := reflect.Indirect(reflect.ValueOf(other))
		otherValue := val.Convert(reflect.TypeOf(float64(0))).Float()

		v.X += otherValue
		v.Y += otherValue
	}
}

// Subtracts Vector2 and other (changes Vector2)
func (v *Vector2) Subtract(other interface {}) {
	t := reflect.TypeOf(other).String()
	if t == "pearl.Vector2" {
		otherVector := other.(Vector2)
		v.X -= otherVector.X
		v.Y -= otherVector.Y
	} else {
		val := reflect.Indirect(reflect.ValueOf(other))
		otherValue := val.Convert(reflect.TypeOf(float64(0))).Float()

		v.X -= otherValue
		v.Y -= otherValue
	}
}

// Multiplies Vector2 and other (changes Vector2)
func (v *Vector2) Multiply(other interface {}) {
	t := reflect.TypeOf(other).String()
	if t == "pearl.Vector2" {
		otherVector := other.(Vector2)
		v.X *= otherVector.X
		v.Y *= otherVector.Y
	} else {
		val := reflect.Indirect(reflect.ValueOf(other))
		otherValue := val.Convert(reflect.TypeOf(float64(0))).Float()

		v.X *= otherValue
		v.Y *= otherValue
	}
}

// Divides Vector2 and other (changes Vector2)
func (v *Vector2) Divide(other interface {}) {
	t := reflect.TypeOf(other).String()
	if t == "pearl.Vector2" {
		otherVector := other.(Vector2)
		v.X /= otherVector.X
		v.Y /= otherVector.Y
	} else {
		val := reflect.Indirect(reflect.ValueOf(other))
		otherValue := val.Convert(reflect.TypeOf(float64(0))).Float()

		v.X /= otherValue
		v.Y /= otherValue
	}
}

// Floors Vector2 (changes Vector2)
func (v *Vector2) Vector2Floor() {
	v.X = float64(int64(v.X))
	v.Y = float64(int64(v.Y))
}


// Returns a copy of the Vector2 normalized (doesn't change Vector2)
func Normalized(v Vector2) Vector2 {
	v.Normalize()
	return v
}


// Adds b to vector. returns result
func Vector2Add(vector Vector2, other interface {}) Vector2 {
	result := Vector2 { vector.X, vector.Y }

	t := reflect.TypeOf(other).String()
	if t == "pearl.Vector2" {
		otherVector := other.(Vector2)
		result.X += otherVector.X
		result.Y += otherVector.Y
	} else {
		val := reflect.Indirect(reflect.ValueOf(other))
		otherValue := val.Convert(reflect.TypeOf(float64(0))).Float()

		result.X += otherValue
		result.Y += otherValue
	}

	return result
}

// Subtracts a and b (doesn't change variables)
func Vector2Subtract(vector Vector2, other interface {}) Vector2 {
	result := Vector2 { vector.X, vector.Y }

	t := reflect.TypeOf(other).String()
	if t == "pearl.Vector2" {
		otherVector := other.(Vector2)
		result.X -= otherVector.X
		result.Y -= otherVector.Y
	} else {
		val := reflect.Indirect(reflect.ValueOf(other))
		otherValue := val.Convert(reflect.TypeOf(float64(0))).Float()

		result.X -= otherValue
		result.Y -= otherValue
	}

	return result
}

// Multiplies a and b (doesn't change variables)
func Vector2Multiply(vector Vector2, other interface {}) Vector2 {
	result := Vector2 { vector.X, vector.Y }

	t := reflect.TypeOf(other).String()
	if t == "pearl.Vector2" {
		otherVector := other.(Vector2)
		result.X *= otherVector.X
		result.Y *= otherVector.Y
	} else {
		val := reflect.Indirect(reflect.ValueOf(other))
		otherValue := val.Convert(reflect.TypeOf(float64(0))).Float()

		result.X *= otherValue
		result.Y *= otherValue
	}

	return result
}

// Divides a and b (doesn't change variables)
func Vector2Divide(vector Vector2, other interface {}) Vector2 {
	result := Vector2 { vector.X, vector.Y }

	t := reflect.TypeOf(other).String()
	if t == "pearl.Vector2" {
		otherVector := other.(Vector2)
		result.X /= otherVector.X
		result.Y /= otherVector.Y
	} else {
		val := reflect.Indirect(reflect.ValueOf(other))
		otherValue := val.Convert(reflect.TypeOf(float64(0))).Float()

		result.X /= otherValue
		result.Y /= otherValue
	}

	return result
}

// Returns Vector2 floored (doesn't change Vector2)
func Vector2Floor(vector Vector2) Vector2 {
	return Vector2 { float64(int64(vector.X)), float64(int64(vector.Y)) }
}


var (
	// Shorthand for 0, 0
	VZERO = Vector2 { 0, 0 }
	// Shorthand for 1, 1
	VONE   = Vector2 { 1, 1 }
	// Shorthand for 0, -1
	VUP    = Vector2 { 0, -1 }
	// Shorthand for 0, 1
	VDOWN  = Vector2 { 0, 1 }
	// Shorthand for -1, 0
	VLEFT  = Vector2 { -1, 0 }
	// Shorthand for 1, 0
	VRIGHT = Vector2 { 1, 0 }
)