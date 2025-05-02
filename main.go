// Hands-on exercise #15 (was #10)
//
// ● use the terminal to make a go workspace
//  ○ mkdir <name>
// 	○ cd <name>
//  ○ go mod init <somename>
//
// ● write a hello world program
// 	○ nano main.go
// 		○ (if this doesn't work on your machine, use an IDE)
// 	○ write go code
//
// ● run your program
// 	○ go run main.go
//
// curriculum item # 053-hands-on-exercise-10

package main

import (
	"fmt"

	"github.com/GoesToEleven/dog"
	"github.com/GoesToEleven/puppy"
)

func main() {
	fmt.Println("Main Parent Bark")
	fmt.Println(dog.WhenGrownUp("Woof"))
	fmt.Println(puppy.BigBark())
}
