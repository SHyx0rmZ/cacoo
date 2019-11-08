package cacoo

type Diagram struct {
	URL            string       `xml:"url"`
	ImageURL       string       `xml:"imageUrl"`
	ImageURLForAPI string       `xml:"imageUrlForApi"`
	DiagramID      string       `xml:"diagramId"`
	Title          string       `xml:"title"`
	Description    string       `xml:"description"`
	Security       SecurityType `xml:"security"`
	Type           DiagramType  `xml:"type"`
	Owner          User         `xml:"owner"`
	Editing        bool         `xml:"editing"`
	Own            bool         `xml:"own"`
	Shared         bool         `xml:"shared"`
	FolderID       int          `xml:"folderId"`
	FolderName     string       `xml:"folderName"`
	SheetCount     int          `xml:"sheetCount"`
	Created        Date         `xml:"created"`
	Updated        Date         `xml:"updated"`
	Sheets         []Sheet      `xml:"sheets>sheet"`
	Comments       []Comment    `xml:"comments>comment"`
}
