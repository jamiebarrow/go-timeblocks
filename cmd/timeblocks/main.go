package main

import (
	"fmt"

	"github.com/jamiebarrow/go-timeblocks/pkg/domain"
)

func main() {
	fmt.Println("== timeblocks ==")


	fmt.Println(domain.NewTimeblock("some task"))
}
