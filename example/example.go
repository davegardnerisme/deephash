package main

import (
	"fmt"

	"github.com/davegardnerisme/deephash"
)

type example struct {
	Foo string
	Bar float64
}

func main() {
	eg1 := "foo"
	eg2 := example{
		Foo: "foo",
		Bar: 43.0,
	}
	eg3 := &example{
		Foo: "foo",
		Bar: 43.0,
	}

	fmt.Printf("String\t%x\n", deephash.Hash(eg1))
	fmt.Printf("Struct\t%x\n", deephash.Hash(eg2))
	fmt.Printf("Pointer\t%x\n", deephash.Hash(eg3))
}
