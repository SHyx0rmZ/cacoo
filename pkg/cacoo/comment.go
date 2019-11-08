package cacoo

type Comment struct {
	User    User   `xml:"user"`
	Content string `xml:"content"`
	Created Date   `xml:"created"`
	Updated Date   `xml:"updated"`
}
