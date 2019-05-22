package main

import (
	"fmt"
	"strings"
)

func main() {

	var x string

	_, _ = fmt.Scan(&x)

	x = strings.ToLower(x)

	if strings.HasPrefix(x, "i") && strings.HasSuffix(x, "n") && strings.Contains(x, "a") {

		fmt.Println("Found!")

	} else {

		fmt.Println("Not Found!")

	}

}
