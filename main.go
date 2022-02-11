package main

import (
	"fmt"
	"github.com/peterhellberg/swapi"
	"strings"
)

func main() {
	c := swapi.DefaultClient

	if atst, err := c.Vehicle(19); err == nil {
		fmt.Println("name: ", atst.Name)
		fmt.Println("model:", atst.Model)
	}

	ab:=strings.Builder{}
	ab.WriteString("abc")
	ab.WriteString("def")
	fmt.Println(ab)
}