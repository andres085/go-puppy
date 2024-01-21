package puppy

import (
	"fmt"

	"github.com/andres085/go-dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark());
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks());
}

func FromV1() {
	fmt.Println("Hello from v1.1.0")	
}

