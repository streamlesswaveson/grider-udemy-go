package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

func main()  {

	resp, err := http.Get("https://google.com")

	if nil != err {
		fmt.Errorf("%v", err)
		os.Exit(1)
	}

	fmt.Println(resp)


	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	//io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

type logWriter struct {

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("wrote this many", len(bs))
	return len(bs), nil
}