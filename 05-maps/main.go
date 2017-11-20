package main

import "fmt"

func main()  {

	colors := map[string]string{
		"red":"ff0000",
		"green": "008000",
		"blue": "0000ff",
	}

	for k,v := range colors {
		fmt.Println(k,"=",v)
	}

	delete(colors,"green")

	fmt.Println(colors)

}