package cacoo

type DiagramFilter string

const (
	FilterAllDiagrams    DiagramFilter = "all"
	FilterOwnedDiagrams  DiagramFilter = "owned"
	FilterSharedDiagrams DiagramFilter = "shared"
	FilterStencil        DiagramFilter = "stencil"
	FilterTemplate       DiagramFilter = "template"
	FilterRecycleBin     DiagramFilter = "recyclebin"
)

func (f DiagramFilter) String() string {
	return string(f)
}
