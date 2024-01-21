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

func FromV2() {
	fmt.Println("Hello from v1.2.0")	
}

func FromV3() {
	fmt.Println("Hello from v1.3.0")	
}

func FromV4() {
	fmt.Println("Hello from v1.4.0")	
}
