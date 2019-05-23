package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var name, addr string

	m := make(map[string]string)
	fmt.Println("Enter the name:")
	fmt.Scan(&name)

	fmt.Println("Enter the address:")
	fmt.Scan(&addr)

	m["name"] = name
	m["address"] = addr

	js, _ := json.Marshal(m)

	fmt.Println(string(js))
}
