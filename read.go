package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Name struct {
	fname,lname string
}

func main()  {

	var tfile string

	fmt.Println("Enter the file name with extension:")
	_, err := fmt.Scan(&tfile)

	if err != nil{
		fmt.Println("Unable to read the file")
	}

	f,_:=ioutil.ReadFile(tfile)

	var names []Name
	var name Name

	data := strings.Split(string(f), "\n")

	for _,d := range data{

		line := strings.Split(string(d)," ")

		name.fname = line[0]
		name.lname = line[1]

		names = append(names,name)
	}


	fmt.Println(names)


}
