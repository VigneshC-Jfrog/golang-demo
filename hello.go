package main // import "github.com/you/hello"

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
    _, err := http.Get("https://golang.org/")
    if err != nil {
        fmt.Println(err)
    }
}
