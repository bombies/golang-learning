package structs

// Memory layout of a struct is important.
// The order of fields in a struct can affect memory alignment and padding.
// This can lead to performance issues, especially in large structs.
// Usually, we order the fields from largest to smallest size.

// Poorly designed struct
type poorUser struct {
	a int8  // 8 bytes
	b int16 // 16 bytes
	c int8  // 8 bytes
	// The order of fields in the struct can affect memory alignment and padding.
	// In this case, the struct is poorly designed because the fields are not ordered by size.
	// This can lead to performance issues, especially in large structs.
}

// Well designed struct
type wellUser struct {
	a int16 // 16 bytes
	b int8  // 8 bytes
	c int8  // 8 bytes
	// The order of fields in the struct can affect memory alignment and padding.
	// In this case, the struct is well designed because the fields are ordered by size.
	// This can lead to performance issues, especially in large structs.
}
