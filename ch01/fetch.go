package ch01

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Fetch prints the content found at a URL.
func Fetch(urls []string) {
	for _, url := range urls {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)

		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)
	}
}
