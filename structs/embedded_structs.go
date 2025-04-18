package structs

type sender struct {
	user      // Embedded struct - This allows top-level access to the fields of user
	rateLimit int
}

func UsingEmbeddedStructs() {
	aSender := sender{
		rateLimit: 1,
		user: user{
			name:   "Ajani",
			number: 1234567890,
		},
	}

	// Accessing the fields of the embedded struct directly
	println(aSender.name)   // Accessing the name field of user
	println(aSender.number) // Accessing the number field of user
}
