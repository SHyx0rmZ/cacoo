package cacoo

import (
	"context"
	"net/http"
)

func NewOrganizationsRequest(ctx context.Context) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, http.MethodGet, "https://cacoo.com/api/v1/organizations.xml", nil)
}

type OrganizationsResponse struct {
	Result []Organization `xml:"result>organization"`
	Count  int            `xml:"count"`
}

type Organization struct {
	ID      int    `xml:"id"`
	Key     string `xml:"key"`
	Name    string `xml:"name"`
	Created Date   `xml:"created"`
	Updated Date   `xml:"updated"`
}
