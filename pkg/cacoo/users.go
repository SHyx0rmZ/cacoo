package cacoo

import (
	"context"
	"fmt"
	"net/http"
)

func NewUsersRequest(ctx context.Context, name string) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://cacoo.com/api/v1/users/%s.xml", name), nil)
}

func (c *Client) User(ctx context.Context, name string) (User, error) {
	var r User
	err := c.do(NewUsersRequest(ctx, name))(&r)
	if err != nil {
		return User{}, err
	}
	return r, nil
}

type UsersResponse User

type User struct {
	Name     string `xml:"name"`
	Nickname string `xml:"nickname"`
	Type     string `xml:"type"`
	ImageURL string `xml:"imageUrl"`
}
