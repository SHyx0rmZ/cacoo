package cacoo

type SortOrder string

const (
	Ascending  SortOrder = "asc"
	Descending SortOrder = "descending"
)

func (o SortOrder) String() string {
	return string(o)
}
