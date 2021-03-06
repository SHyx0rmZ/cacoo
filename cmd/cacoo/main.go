package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"cacoo/pkg/cacoo"
)

var commands = map[string]func(context.Context, *cacoo.Client){
	"account":  Account,
	"diagrams": Diagrams,
	"user":     User,
}

func Account(ctx context.Context, client *cacoo.Client) {
	a, err := client.Account(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("    Name:", a.Name)
	fmt.Println("Nickname:", a.Nickname)
	fmt.Println("    Type:", a.Type)
	fmt.Println("ImageURL:", a.ImageURL)
}

func Diagrams(ctx context.Context, client *cacoo.Client) {
	fs := flag.NewFlagSet("diagrams", flag.ContinueOnError)
	fs.Usage = usageWithNewline

	var opts struct {
		Filter string
		Limit  int
	}

	fs.StringVar(&opts.Filter, "--filter", "owned", `one of "all", "owned", "shared", "stencil", "template", "recyclebin"`)
	fs.IntVar(&opts.Limit, "--limit", 3, "maximum number of diagrams to return")

	err := fs.Parse(flag.Args()[1:])
	if err != nil {
		log.Fatalln(err)
	}

	var ps []cacoo.DiagramsRequestParameter
	if opts.Filter != "" {
		ps = append(ps, cacoo.WithFilter(cacoo.DiagramFilter(opts.Filter)))
	}
	if opts.Limit != 0 {
		ps = append(ps, cacoo.WithLimit(opts.Limit))
	}

	ds, err := client.Diagrams(ctx, ps...)
	if err != nil {
		log.Fatalln(err)
	}

	for i, d := range ds {
		fmt.Printf("Diagram #%d", i)
		fmt.Println()
		fmt.Println("  Title:", d.Title)
		fmt.Println("    URL:", d.URL)

		if i != len(ds)-1 {
			fmt.Println()
		}
	}
}

func User(ctx context.Context, client *cacoo.Client) {
	if flag.NArg() != 2 {
		fmt.Fprintln(os.Stderr, "usage:", os.Args[0], "user", "<user-name>")
	}

	u, err := client.User(ctx, flag.Arg(1))
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("    Name:", u.Name)
	fmt.Println("Nickname:", u.Nickname)
	fmt.Println("    Type:", u.Type)
	fmt.Println("ImageURL:", u.ImageURL)
}

func available() {
	fmt.Fprintln(os.Stderr, "Available commands are:")
	fmt.Fprintln(os.Stderr, "   ", "account          - list information on account associated with the API key")
	fmt.Fprintln(os.Stderr, "   ", "diagrams         - list diagrams")
	fmt.Fprintln(os.Stderr, "   ", "user <user-name> - list information on specified user")
	os.Exit(2)
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage:", os.Args[0], "[--api-key=<api-key>] <command>")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Some commands require an API key, which can either be passed as a flag, or")
	fmt.Fprintln(os.Stderr, "through the CACOO_API_KEY environment variable.")
	fmt.Fprintln(os.Stderr)
	available()
}

func usageWithNewline() {
	fmt.Fprintln(os.Stderr)
	usage()
}

func main() {
	var apiKey string
	var debug bool

	flag.StringVar(&apiKey, "api-key", "", "API key used to access private diagrams")
	flag.BoolVar(&debug, "debug", false, "Print API responses")
	flag.Usage = usageWithNewline
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
	}

	if len(strings.TrimSpace(apiKey)) == 0 {
		apiKey = os.Getenv("CACOO_API_KEY")
	}

	ctx := context.Background()
	c := cacoo.NewClient(apiKey, debug)

	cmd, ok := commands[flag.Arg(0)]
	if !ok {
		fmt.Fprintln(os.Stderr, "unknown command:", flag.Arg(0))
		fmt.Fprintln(os.Stderr)
		available()
	}

	cmd(ctx, c)
	return

	//l, err := c.License(ctx)
	//if err != nil {
	//	log.Panicln(err)
	//}
	//fmt.Println(l)
	//
	//o, err := c.Organizations(ctx)
	//if err != nil {
	//	log.Panicln(err)
	//}
	//fmt.Println(o)
	//
	//ds, err := c.Diagrams(ctx, cacoo.WithFilter(cacoo.FilterOwnedDiagrams), cacoo.WithLimit(3))
	//if err != nil {
	//	log.Panicln(err)
	//}
	//fmt.Println(ds)
	//
	//d, err := c.Diagram(ctx, ds[0].DiagramID)
	//if err != nil {
	//	log.Panicln(err)
	//}
	//fmt.Println(d)
	//
	//co, err := c.DiagramContent(ctx, d.DiagramID)
	//if err != nil {
	//	log.Panicln(err)
	//}
	//fmt.Println(co)
}
