package main

import (
	"fmt"
	"os"

	"github.com/alistanis/branch"
)

func main() {
	r, err := branch.Open(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	h, err := r.Head()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(h.Name().Short())
}
