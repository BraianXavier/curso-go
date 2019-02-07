package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go concurrency Patterns

//<-chan - canal somente de leitura

func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	// t1 := titulo("http://www.google.com", "http://www.cod3r.com.br")
	t2 := titulo("http://www.amazon.com", "http://www.youtube.com")
	fmt.Println("Primeiros", <-t2)

	// fmt.Println("Segundos", <-t1, "i", <-t2)
}
