package main

import (
	"fmt"
)

func main() {

	var x float32

	_, _ = fmt.Scan(&x)

	y := int(x)

	fmt.Println(y)

}
