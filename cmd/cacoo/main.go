package main

import (
	"context"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cacoo/pkg/cacoo"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("CACOO_API_KEY")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cacoo.DiagramsURL+"?apiKey="+apiKey, nil)
	if err != nil {
		log.Panicln(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		_, err = io.Copy(ioutil.Discard, resp.Body)
		if err != nil {
			log.Println(err)
		}
		err = resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var d cacoo.DiagramsResponse
	err = xml.NewDecoder(io.TeeReader(resp.Body, os.Stderr)).Decode(&d)
	if err != nil {
		log.Panicln(err)
	}
}
