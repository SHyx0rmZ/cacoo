package main

import (
	"context"
	"encoding/xml"
	"fmt"
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

	c := cacoo.NewClient(apiKey)

	//req, err := http.NewRequestWithContext(ctx, http.MethodGet, cacoo.DiagramsURL+"?apiKey="+apiKey, nil)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cacoo.AccountURL, nil)
	if err != nil {
		log.Panicln(err)
	}
	resp, err := c.Do(req)
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

	var d cacoo.AccountResponse
	err = xml.NewDecoder(io.TeeReader(resp.Body, os.Stderr)).Decode(&d)
	if err != nil {
		log.Panicln(err)
	}

	req, err = cacoo.NewUsersRequest(ctx, d.Name)
	if err != nil {
		log.Panicln(err)
	}
	resp2, err := c.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		_, err = io.Copy(ioutil.Discard, resp2.Body)
		if err != nil {
			log.Println(err)
		}
		err = resp2.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var u cacoo.UsersResponse
	err = xml.NewDecoder(io.TeeReader(resp2.Body, os.Stderr)).Decode(&u)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(u)

	req, err = cacoo.NewLicenseRequest(ctx)
	if err != nil {
		log.Panicln(err)
	}
	resp3, err := c.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		_, err = io.Copy(ioutil.Discard, resp3.Body)
		if err != nil {
			log.Println(err)
		}
		err = resp3.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var l cacoo.LicenseResponse
	err = xml.NewDecoder(io.TeeReader(resp3.Body, os.Stderr)).Decode(&l)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(u)
}
