package cacoo

type Sheet struct {
	URL            string `xml:"url"`
	ImageURL       string `xml:"imageUrl"`
	ImageURLForAPI string `xml:"imageUrlForApi"`
	UniqueID       string `xml:"uid"`
	Name           string `xml:"name"`
	Width          int    `xml:"width"`
	Height         int    `xml:"height"`
}
