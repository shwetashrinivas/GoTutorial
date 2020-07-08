package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//Custom Writer Interface
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//1
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	//2 func Copy(dst Writer , src Reader) ( writen int 64, err error)
	io.Copy(os.Stdout, resp.Body)

	//3
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

//logWriter is implementing Writer Interfaces
func (logWriter) Write(bs []byte) (int, error) {
	//return 1 , nil
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
