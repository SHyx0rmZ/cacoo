package cacoo

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
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

func NewCreateDiagramRequest(ctx context.Context, diagram Diagram) (*http.Request, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://cacoo.com/api/v1/diagrams/create.xml", nil)
	if err != nil {
		return nil, err
	}
	q := r.URL.Query()
	if diagram.Title != "" {
		q.Add("title", diagram.Title)
	}
	if diagram.Description != "" {
		q.Add("description", diagram.Description)
	}
	if diagram.Security != "" {
		q.Add("security", string(diagram.Security))
	}
	if diagram.FolderID != 0 {
		q.Add("folderId", strconv.Itoa(diagram.FolderID))
	}
	r.URL.RawQuery = q.Encode()
	return r, nil
}

func (c *Client) CreateDiagram(ctx context.Context, diagram Diagram) (Diagram, error) {
	var r Diagram
	err := c.do(NewCreateDiagramRequest(ctx, diagram))(&r)
	if err != nil {
		return Diagram{}, err
	}
	return r, nil
}

func NewCopyDiagramRequest(ctx context.Context, diagramID string, diagram Diagram) (*http.Request, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://cacoo.com/api/v1/diagrams/"+diagramID+"/copy.xml", nil)
	if err != nil {
		return nil, err
	}
	q := r.URL.Query()
	if diagram.Title != "" {
		q.Add("title", diagram.Title)
	}
	if diagram.Description != "" {
		q.Add("description", diagram.Description)
	}
	if diagram.Security != "" {
		q.Add("security", string(diagram.Security))
	}
	if diagram.FolderID != 0 {
		q.Add("folderId", strconv.Itoa(diagram.FolderID))
	}
	r.URL.RawQuery = q.Encode()
	fmt.Println(r.URL)
	return r, nil
}

func (c *Client) CopyDiagram(ctx context.Context, diagramID string, diagram Diagram) (Diagram, error) {
	var r Diagram
	err := c.do(NewCopyDiagramRequest(ctx, diagramID, diagram))(&r)
	if err != nil {
		return Diagram{}, err
	}
	return r, nil
}

func NewArchiveDiagramRequest(ctx context.Context, diagramID string) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, http.MethodPost, "https://cacoo.com/api/v1/diagrams/"+diagramID+"/delete.xml", nil)
}

func (c *Client) ArchiveDiagram(ctx context.Context, diagramID string) error {
	err := c.do(NewArchiveDiagramRequest(ctx, diagramID))(nil)
	return err
}

//func NewChatMessagesRequest(ctx context.Context,diagramID string)
