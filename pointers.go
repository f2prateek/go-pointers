// Package pointers provides factories for creating pointers to primitives.
package pointers

// Bool returns a pointer to a bool primitive.
func Bool(x bool) *bool {
	return &x
}

// Byte returns a pointer to a byte primitive.
func Byte(x byte) *byte {
	return &x
}

// Rune returns a pointer to a rune primitive.
func Rune(x rune) *rune {
	return &x
}

// String returns a pointer to a string primitive.
func String(x string) *string {
	return &x
}

// Int returns a pointer to an int primitive.
func Int(x int) *int {
	return &x
}

// Int8 returns a pointer to an int8 primitive.
func Int8(x int8) *int8 {
	return &x
}

// Int32 returns a pointer to an int32 primitive.
func Int32(x int32) *int32 {
	return &x
}

// Int64 returns a pointer to an int64 primitive.
func Int64(x int64) *int64 {
	return &x
}

// UInt returns a pointer to an uint primitive.
func UInt(x uint) *uint {
	return &x
}

// UInt16 returns a pointer to an uint16 primitive.
func UInt16(x uint16) *uint16 {
	return &x
}

// UInt32 returns a pointer to an uint32 primitive.
func UInt32(x uint32) *uint32 {
	return &x
}

// UInt64 returns a pointer to an uint64 primitive.
func UInt64(x uint64) *uint64 {
	return &x
}

// Float32 returns a pointer to a float32 primitive.
func Float32(x float32) *float32 {
	return &x
}

// Float64 returns a pointer to a float64 primitive.
func Float64(x float64) *float64 {
	return &x
}
