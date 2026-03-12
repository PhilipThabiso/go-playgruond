package main

import (
	"fmt"

	"github.com/PhilipThabiso/go-playground/compare"
)

func main() {
	if err := compare.InitCompare(); err != nil {
		fmt.Printf("%w", err)
	}
}
