package cacoo

const AccountURL = "https://cacoo.com/api/v1/account.xml"

type AccountResponse struct {
	Name     string `xml:"name"`
	Nickname string `xml:"nickname"`
	Type     string `xml:"type"`
	ImageURL string `xml:"imageUrl"`
}
