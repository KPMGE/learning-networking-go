// This is just a simple example of how to time-out a request using the go's strandart net/http package.
// Author: Kevin Carvalho de Jesus

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {
	// create a context to time it out 5 seconds. You can simulate a timeout setting this to a smaller value
	fiveSeconds := time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), fiveSeconds)
	// cancel context
	defer cancel()

	// create a new request to google's page
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// make request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Fatalln("Time out exceeded!")
		}
		log.Fatalln(err)
	}
	// we should always close the response body!
	defer resp.Body.Close()

	// read response body and print it
	body, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}
