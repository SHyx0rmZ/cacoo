package cacoo

type Sheet struct {
	URL            string `xml:"url" json:"url"`
	ImageURL       string `xml:"imageUrl" json:"imageUrl"`
	ImageURLForAPI string `xml:"imageUrlForApi" json:"imageUrlForApi"`
	UniqueID       string `xml:"uid" json:"uid"`
	Name           string `xml:"name" json:"name"`
	Width          int    `xml:"width" json:"width"`
	Height         int    `xml:"height" json:"height"`
}
