package main

import (
	"os"
	"io"
)

func main()  {
	if len(os.Args) != 2 {
		panic("expected filename")
	}

	filename := os.Args[1]

	f,err := os.Open(filename)
	assert(err)

	io.Copy(os.Stdout, f)


}

func assert(e error)  {
	if nil != e {
		panic(e)
	}

}