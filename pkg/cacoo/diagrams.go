package cacoo

import (
	"context"
	"net/http"
)

const DiagramsURL = "https://cacoo.com/api/v1/diagrams.xml"

type DiagramType2 string

type DiagramsResponse struct {
	Result []Diagram `xml:"result>diagram"`
	Count  int       `xml:"count"`
}

func NewDiagramsRequest(ctx context.Context, parameters ...DiagramsRequestParameter) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://cacoo.com/api/v1/diagrams.xml", nil)
	if err != nil {
		return nil, err
	}
	var ps diagramsRequestParameters
	for _, p := range parameters {
		p(&ps)
	}
	req.URL.RawQuery = ps.Encode()
	return req, nil
}

func (c *Client) Diagrams(ctx context.Context, parameters ...DiagramsRequestParameter) ([]Diagram, error) {
	var r DiagramsResponse
	err := c.do(NewDiagramsRequest(ctx, parameters...))(&r)
	if err != nil {
		return nil, err
	}
	return r.Result, nil
}

func NewDiagramRequest(ctx context.Context, diagramID string) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, http.MethodGet, "https://cacoo.com/api/v1/diagrams/"+diagramID+".xml", nil)
}

func (c *Client) Diagram(ctx context.Context, diagramID string) (Diagram, error) {
	var r Diagram
	err := c.do(NewDiagramRequest(ctx, diagramID))(&r)
	if err != nil {
		return Diagram{}, err
	}
	return r, nil
}

type DiagramContent struct {
	Sheet struct {
		Name     string     `xml:"name,attr"`
		Groups   []CGroup   `xml:"group"`
		Polygons []CPolygon `xml:"polygon"`
		Lines    []CLine    `xml:"line"`
		Texts    []CText    `xml:"text"`
		Images   []CImage   `xml:"image"`
	} `xml:"sheet"`
}

type CGroup struct {
	StencilID string     `xml:"attr-stencil-id,attr"`
	Groups    []CGroup   `xml:"group"`
	Polygons  []CPolygon `xml:"polygon"`
	Lines     []CLine    `xml:"line"`
	Texts     []CText    `xml:"text"`
	Images    []CImage   `xml:"image"`
}

type CPolygon struct {
}

type CLine struct {
}

type CText struct {
	Text string `xml:",innerxml"`
}

type CImage struct {
	SourceID string `xml:"source-id,attr"`
}

func NewDiagramContentRequest(ctx context.Context, diagramID string) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, http.MethodGet, "https://cacoo.com/api/v1/diagrams/"+diagramID+"/contents.xml", nil)
}

func (c *Client) DiagramContent(ctx context.Context, diagramID string) (DiagramContent, error) {
	var r DiagramContent
	err := c.do(NewDiagramContentRequest(ctx, diagramID))(&r)
	if err != nil {
		return DiagramContent{}, err
	}
	return r, nil
}

//func NewChatMessagesRequest(ctx context.Context,diagramID string)
