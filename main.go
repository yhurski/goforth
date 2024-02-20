package main

import (
	"fmt"
	"github.com/yhurski/goforth/internal/forth"
)

func main() {
	fmt.Printf("Welcome to GoForth, version: %s\n", Version)

	forth.DoForth()
}
