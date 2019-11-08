package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cacoo/pkg/cacoo"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("CACOO_API_KEY")

	c := cacoo.NewClient(apiKey)

	a, err := c.Account(ctx)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(a)

	u, err := c.User(ctx, a.Name)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(u)

	l, err := c.License(ctx)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(l)

	o, err := c.Organizations(ctx)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(o)

	ds, err := c.Diagrams(ctx)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(ds)

	d, err := c.Diagram(ctx, ds[0].DiagramID)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(d)
}
