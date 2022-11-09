package main

import (
	"fmt"

	"github.com/sufu777/go-github-action/core"
)

func main() {
	fmt.Println(core.Greeting())
	fmt.Printf("1 + 2 = %d", core.Add(1, 2))
}
