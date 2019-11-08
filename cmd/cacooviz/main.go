package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"cacoo/pkg/cacoo"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	apiKey := os.Getenv("CACOO_API_KEY")

	c := cacoo.NewClient(apiKey)
	co, err := c.DiagramContent(ctx, "VjBWOBNheo4AV46C")
	if err != nil {
		log.Panicln(err)
	}

	f, err := os.Create("cacoo.dot")
	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	sheet := co.Sheets[1] // in the example diagram linked, we want to export (at least parts of) the second sheet

	fmt.Fprintln(f, `digraph "`+sheet.Name+`"`, "{")
	fmt.Fprintln(f, "node [style=filled];")
	for _, group := range sheet.Groups {
		if len(group.Polygons) != len(group.Texts) {
			continue
		}
		var label struct {
			Text   string
			Index  int
			HAlign string
		}
		for i, text := range group.Texts {
			if text.Text != "" {
				label.Text = text.Text
				label.Index = i
				label.HAlign = text.HorizontalAlignment
				break
			}
			if text.TextStyle.Text != "" {
				label.Text = text.TextStyle.Text
				label.Index = i
				label.HAlign = text.HorizontalAlignment
				break
			}
		}
		if label.Text == "" {
			continue
		}
		polygon := group.Polygons[label.Index]
		// this is a very stupid (and bad) check to guess if we're dealing with a rectangle
		if label.HAlign == "left" {
			label.Text = strings.TrimRight(`\l`+strings.ReplaceAll(label.Text, "\n", "\n\\l"), `\l`)
		}
		label.Text = strings.ReplaceAll(label.Text, "\n", `\n`)
		fmt.Fprint(f, `"`+polygon.UID+`" [label="`+label.Text+`"`)
		if polygon.FillColor != "" {
			fmt.Fprint(f, `,fillcolor="`+polygon.FillColor+`"`)
		}
		if polygon.BorderColor != "" {
			fmt.Fprint(f, `,bordercolor="`+polygon.BorderColor+`"`)
		}
		if polygon.Path.Close && len(polygon.Path.Points) == 4 {
			fmt.Fprint(f, `,shape="rectangle"`)
		}
		fmt.Fprintln(f, `];`)
	}
	for _, line := range sheet.Lines {
		var label struct {
			Text   string
			Index  int
			HAlign string
		}
		for i, text := range line.Labels {
			if text.Text != "" {
				label.Text = text.Text
				label.Index = i
				label.HAlign = text.HorizontalAlignment
				break
			}
			if text.TextStyle.Text != "" {
				label.Text = text.TextStyle.Text
				label.Index = i
				label.HAlign = text.HorizontalAlignment
				break
			}
		}
		fmt.Fprint(f, `"`+line.Start.ConnectUID+`"`, " -> ", `"`+line.End.ConnectUID+`"`)
		if label.Text != "" {
			label.Text = strings.ReplaceAll(label.Text, "\n", `\n`)
			fmt.Fprint(f, ` [label="`+label.Text+`"]`)
		}
		fmt.Fprintln(f, ";")
	}
	fmt.Fprintln(f, "}")
}
