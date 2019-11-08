package cacoo

import (
	"context"
	"net/http"
)

const AccountURL = "https://cacoo.com/api/v1/account.xml"

type AccountResponse User

func NewAccountRequest(ctx context.Context) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, http.MethodGet, AccountURL, nil)
}

func (c *Client) Account(ctx context.Context) (User, error) {
	var r User
	err := c.do(NewAccountRequest(ctx))(&r)
	if err != nil {
		return User{}, err
	}
	return r, nil
}
