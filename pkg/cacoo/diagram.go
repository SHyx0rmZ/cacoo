package cacoo

type Diagram struct {
	URL            string       `xml:"url" json:"url"`
	ImageURL       string       `xml:"imageUrl" json:"imageUrl"`
	ImageURLForAPI string       `xml:"imageUrlForApi" json:"imageUrlForApi"`
	DiagramID      string       `xml:"diagramId" json:"diagramId"`
	Title          string       `xml:"title" json:"title"`
	Description    string       `xml:"description" json:"description"`
	Security       SecurityType `xml:"security" json:"security"`
	Type           DiagramType  `xml:"type" json:"type"`
	Owner          User         `xml:"owner" json:"owner"`
	Editing        bool         `xml:"editing" json:"editing"`
	Own            bool         `xml:"own" json:"own"`
	Shared         bool         `xml:"shared" json:"shared"`
	FolderID       int          `xml:"folderId" json:"folderId"`
	FolderName     string       `xml:"folderName" json:"folderName"`
	SheetCount     int          `xml:"sheetCount" json:"sheetCount"`
	Created        Date         `xml:"created" json:"created"`
	Updated        Date         `xml:"updated" json:"updated"`
	Sheets         []Sheet      `xml:"sheets>sheet" json:"sheets"`
	Comments       []Comment    `xml:"comments>comment" json:"comments"`
}
