package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse(time.DateOnly, "2025-07-11")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
}
