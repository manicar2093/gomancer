package main

import (
	"fmt"
	"github.com/manicar2093/gomancer/testfixtures"
)

func main() {
	fmt.Println("testfixtures attributes: ", len(testfixtures.ModelBinaryIdSuccess.Attributes))
	fmt.Println(testfixtures.ModelBinaryIdSuccess.ToStringCmd())
}
