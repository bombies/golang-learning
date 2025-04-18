package structs

import "fmt"

func F() {
	myCar := struct {
		make  string
		model string
	}{
		make:  "Toyota",
		model: "Corolla",
	}

	fmt.Printf("%s %s\n", myCar.make, myCar.model)
}
