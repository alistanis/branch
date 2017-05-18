package main

import (
	"fmt"
	"os"

	"strings"

	"github.com/alistanis/branch"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	pathSplit := strings.Split(path, string(os.PathSeparator))
	r, err := branch.Open(pathSplit)
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
