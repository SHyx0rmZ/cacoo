package cacoo

const DiagramsURL = "https://cacoo.com/api/v1/diagrams.xml"

type DiagramType string

const (
	AllDiagrams    DiagramType = "all"
	OwnedDiagrams  DiagramType = "owned"
	SharedDiagrams DiagramType = "shared"
	Stencil        DiagramType = "stencil"
	Template       DiagramType = "template"
	RecycleBin     DiagramType = "recyclebin"
)

const (
	Updated = "updated"
	Title   = "title"
	Owner   = "owner"
	Folder  = "folder"
)

const (
	Ascending  = "asc"
	Descending = "descending"
)

type DiagramsRequest struct {
	Offset          int         `xml:"offset"`
	Limit           int         `xml:"limit"`
	Type            DiagramType `xml:"type"`
	SortOn          string      `xml:"sortOn"`
	SortType        string      `xml:"sortType"`
	FolderID        int         `xml:"folderId"`
	Keyword         string      `xml:"keyword"`
	OrganizationKey string      `xml:"organizationKey,omitempty"`
}

const (
	Private = "private"
	URL     = "url"
	Public  = "public"
)

const (
	Normal = "normal"
	//Stencil  = "stencil"
	//Template = "template"
)

const (
	Cacoo = "cacoo"
	Other = "other"
)

type DiagramsResponse struct {
	Result []struct {
		URL            string `xml:"url"`
		ImageURL       string `xml:"imageUrl"`
		ImageURLForAPI string `xml:"imageUrlForApi"`
		DiagramID      string `xml:"diagramId"`
		Title          string `xml:"title"`
		Description    string `xml:"description"`
		Security       string `xml:"security"`
		Type           string `xml:"type"`
		Owner          struct {
			Name     string `xml:"name"`
			Nickname string `xml:"nickname"`
			Type     string `xml:"type"`
			ImageURL string `xml:"imageUrl"`
		} `xml:"owner"`
		Editing    bool   `xml:"editing"`
		Own        bool   `xml:"own"`
		Shared     bool   `xml:"shared"`
		FolderID   int    `xml:"folderId"`
		FolderName string `xml:"folderName"`
		SheetCount int    `xml:"sheetCount"`
		Created    string/* RFC 2822 date */ `xml:"created"`
		Updated    string/* " */ `xml:"updated"`
	} `xml:"result>diagram"`
	Count int `xml:"count"`
}
