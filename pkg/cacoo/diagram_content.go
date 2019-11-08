package cacoo

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
)

type DiagramContent struct {
	Sheets []struct {
		Name            string     `xml:"name,attr"`
		UID             string     `xml:"uid,attr"`
		Groups          []CGroup   `xml:"group"`
		Polygons        []CPolygon `xml:"polygon"`
		Lines           []CLine    `xml:"line"`
		Texts           []CText    `xml:"text"`
		Images          []CImage   `xml:"image"`
		BackgroundSheet string     `xml:"background-sheet,attr"`
	} `xml:"sheet"`
}

type CPosition struct {
	X      float64 `xml:"x,attr"`
	Y      float64 `xml:"y,attr"`
	Width  float64 `xml:"width,attr"`
	Height float64 `xml:"height,attr"`
	Angle  float64 `xml:"angle,attr"`
}

type CGroup struct {
	CPosition
	UID       string     `xml:"uid,attr"`
	StencilID string     `xml:"attr-stencil-id,attr"`
	Groups    []CGroup   `xml:"group"`
	Polygons  []CPolygon `xml:"polygon"`
	Lines     []CLine    `xml:"line"`
	Texts     []CText    `xml:"text"`
	Images    []CImage   `xml:"image"`
}

type CPolygon struct {
	CPosition
	UID             string `xml:"uid,attr"`
	Path            CPath  `xml:"path"`
	FillStyle       string `xml:"fill-style,attr"`
	FillColor       string `xml:"fill-color,attr"`
	FillOpacity     string `xml:"fill-opacity,attr"`
	BorderStyle     string `xml:"border-style,attr"`
	BorderColor     string `xml:"border-color,attr"`
	BorderThickness string `xml:"border-thickness,attr"`
}

type CLine struct {
	CPosition
	UID    string   `xml:"uid,attr"`
	Type   string   `xml:"type,attr"`
	Start  CEdge    `xml:"start"`
	End    CEdge    `xml:"end"`
	Points []CPoint `xml:"point"`
	Labels []CText  `xml:"labels>text"`
}

type CEdge struct {
	ConnectUID string `xml:"connect-uid,attr"`
	Style      string `xml:"style,attr"`
}

type CText struct {
	CPosition
	HorizontalAlignment string `xml:"h-align,attr"`
	VerticalAlignment   string `xml:"v-align,attr"`
	UID                 string `xml:"uid,attr"`
	// there can be multiple text styles per text, which we ignore here
	TextStyle CTextStyle
	Text      string
}

func (t *CText) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	sc := start.Copy()
	var r1 struct {
		CPosition
		HorizontalAlignment string `xml:"h-align,attr"`
		VerticalAlignment   string `xml:"v-align,attr"`
		UID                 string `xml:"uid"`
		// there can be multiple text styles per text, which we ignore here
		TextStyle CTextStyle `xml:"textStyle"`
		Text      string     `xml:",innerxml"`
	}
	err := dec.DecodeElement(&r1, &sc)
	if err != nil {
		return fmt.Errorf("un-marshaling CText: %w", err)
	}
	t.CPosition = r1.CPosition
	t.HorizontalAlignment = r1.HorizontalAlignment
	t.VerticalAlignment = r1.VerticalAlignment
	t.UID = r1.UID
	var z CTextStyle
	if r1.TextStyle != z {
		t.TextStyle = r1.TextStyle
		t.Text = ""
		return nil
	}
	t.TextStyle = CTextStyle{}
	t.Text = r1.Text
	return nil
}

type CTextStyle struct {
	Font      string `xml:"font,attr"`
	Color     string `xml:"color,attr"`
	Size      int    `xml:"size,attr"`
	Italic    bool   `xml:"italic,attr"`
	Underline bool   `xml:"underline,attr"`
	Weight    string `xml:"weight,attr"`
	Text      string `xml:",innerxml"`
}

type CImage struct {
	CPosition
	UID      string `xml:"uid"`
	SourceID string `xml:"source-id,attr"`
}

type CPath struct {
	Close  bool     `xml:"close,attr"`
	Points []CPoint `xml:"point"`
}

type CPoint struct {
	Type string  `xml:"type,attr"`
	X    float64 `xml:"x,attr"`
	Y    float64 `xml:"y,attr"`
}

func NewDiagramContentRequest(ctx context.Context, diagramID string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://cacoo.com/api/v1/diagrams/"+diagramID+"/contents.xml", nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Add("returnValues", "textStyle,shapeStyle,uid,position,point")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

func (c *Client) DiagramContent(ctx context.Context, diagramID string) (DiagramContent, error) {
	var r DiagramContent
	err := c.do(NewDiagramContentRequest(ctx, diagramID))(&r)
	if err != nil {
		return DiagramContent{}, err
	}
	return r, nil
}
