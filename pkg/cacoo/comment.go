package cacoo

type Comment struct {
	User    User   `xml:"user" json:"user"`
	Content string `xml:"content" json:"content"`
	Created Date   `xml:"created" json:"created"`
	Updated Date   `xml:"updated" json:"updated"`
}
