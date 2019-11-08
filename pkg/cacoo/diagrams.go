package cacoo

import (
	"context"
	"net/http"
)

const DiagramsURL = "https://cacoo.com/api/v1/diagrams.xml"

type DiagramType string

const (
	AllDiagrams    DiagramType = "all"
	OwnedDiagrams  DiagramType = "owned"
	SharedDiagrams DiagramType = "shared"
	Stencil        DiagramType = "stencil"
	Template       DiagramType = "template"
	RecycleBin     DiagramType = "recyclebin"
)

const (
	Updated = "updated"
	Title   = "title"
	Owner   = "owner"
	Folder  = "folder"
)

const (
	Ascending  = "asc"
	Descending = "descending"
)

type DiagramsRequest struct {
	Offset          int         `xml:"offset"`
	Limit           int         `xml:"limit"`
	Type            DiagramType `xml:"type"`
	SortOn          string      `xml:"sortOn"`
	SortType        string      `xml:"sortType"`
	FolderID        int         `xml:"folderId"`
	Keyword         string      `xml:"keyword"`
	OrganizationKey string      `xml:"organizationKey,omitempty"`
}

const (
	Private = "private"
	URL     = "url"
	Public  = "public"
)

const (
	Normal = "normal"
	//Stencil  = "stencil"
	//Template = "template"
)

const (
	Cacoo = "cacoo"
	Other = "other"
)

type DiagramsResponse struct {
	Result []Diagram `xml:"result>diagram"`
	Count  int       `xml:"count"`
}

func NewDiagramsRequest(ctx context.Context) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://cacoo.com/api/v1/diagrams.xml", nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Add("type", "all")
	return req, nil
}

func (c *Client) Diagrams(ctx context.Context) ([]Diagram, error) {
	var r DiagramsResponse
	err := c.do(NewDiagramsRequest(ctx))(&r)
	if err != nil {
		return nil, err
	}
	return r.Result, nil
}

func NewDiagramRequest(ctx context.Context, diagramID string) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, http.MethodGet, "https://cacoo.com/api/v1/diagrams/"+diagramID+".xml", nil)
}

func (c *Client) Diagram(ctx context.Context, diagramID string) (DiagramResponse, error) {
	var r DiagramResponse
	err := c.do(NewDiagramRequest(ctx, diagramID))(&r)
	if err != nil {
		return DiagramResponse{}, err
	}
	return r, nil
}

type DiagramResponse Diagram

type Diagram struct {
	URL            string    `xml:"url"`
	ImageURL       string    `xml:"imageUrl"`
	ImageURLForAPI string    `xml:"imageUrlForApi"`
	DiagramID      string    `xml:"diagramId"`
	Title          string    `xml:"title"`
	Description    string    `xml:"description"`
	Security       string    `xml:"security"`
	Type           string    `xml:"type"`
	Owner          User      `xml:"owner"`
	Editing        bool      `xml:"editing"`
	Own            bool      `xml:"own"`
	Shared         bool      `xml:"shared"`
	FolderID       int       `xml:"folderId"`
	FolderName     string    `xml:"folderName"`
	SheetCount     int       `xml:"sheetCount"`
	Created        Date      `xml:"created"`
	Updated        Date      `xml:"updated"`
	Sheets         []Sheet   `xml:"sheets>sheet"`
	Comments       []Comment `xml:"comments>comment"`
}

type Sheet struct {
	URL            string `xml:"url"`
	ImageURL       string `xml:"imageUrl"`
	ImageURLForAPI string `xml:"imageUrlForApi"`
	UniqueID       string `xml:"uid"`
	Name           string `xml:"name"`
	Width          int    `xml:"width"`
	Height         int    `xml:"height"`
}

type Comment struct {
	User    User   `xml:"user"`
	Content string `xml:"content"`
	Created Date   `xml:"created"`
	Updated Date   `xml:"updated"`
}

//func NewChatMessagesRequest(ctx context.Context,diagramID string)
