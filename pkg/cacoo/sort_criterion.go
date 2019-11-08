package cacoo

type SortCriterion string

const (
	Updated SortCriterion = "updated"
	Title   SortCriterion = "title"
	Owner   SortCriterion = "owner"
	Folder  SortCriterion = "folder"
)

func (c SortCriterion) String() string {
	return string(c)
}
