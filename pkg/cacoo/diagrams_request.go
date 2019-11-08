package cacoo

import (
	"net/url"
	"strconv"
)

type diagramsRequestParameters struct {
	Offset          string
	Limit           string
	Type            string
	SortOn          string
	SortType        string
	FolderID        string
	Keyword         string
	OrganizationKey string
}

type DiagramsRequestParameter func(*diagramsRequestParameters)

func WithOffset(offset int) DiagramsRequestParameter {
	return func(p *diagramsRequestParameters) {
		p.Offset = strconv.Itoa(offset)
	}
}

func WithLimit(limit int) DiagramsRequestParameter {
	return func(p *diagramsRequestParameters) {
		p.Limit = strconv.Itoa(limit)
	}
}

func WithFilter(filter DiagramFilter) DiagramsRequestParameter {
	return func(p *diagramsRequestParameters) {
		p.Type = filter.String()
	}
}

func WithSortCriterion(sortCriterion SortCriterion) DiagramsRequestParameter {
	return func(p *diagramsRequestParameters) {
		p.SortOn = sortCriterion.String()
	}
}

func WithSortOrder(sortOrder SortOrder) DiagramsRequestParameter {
	return func(p *diagramsRequestParameters) {
		p.SortType = sortOrder.String()
	}
}

func WithFolderID(folderID int) DiagramsRequestParameter {
	return func(p *diagramsRequestParameters) {
		p.FolderID = strconv.Itoa(folderID)
	}
}

func WithKeyword(keyword string) DiagramsRequestParameter {
	return func(p *diagramsRequestParameters) {
		p.Keyword = keyword
	}
}

func WithOrganizationKey(organizationKey string) DiagramsRequestParameter {
	return func(p *diagramsRequestParameters) {
		p.OrganizationKey = organizationKey
	}
}

func (p diagramsRequestParameters) Encode() string {
	query := make(url.Values)
	if len(p.Offset) != 0 {
		query.Add("offset", p.Offset)
	}
	if len(p.Limit) != 0 {
		query.Add("limit", p.Limit)
	}
	if len(p.Type) != 0 && p.Type != "all" {
		query.Add("type", p.Type)
	}
	if len(p.SortOn) != 0 && p.SortOn != "updated" {
		query.Add("sortOn", p.SortOn)
	}
	if len(p.SortType) != 0 && p.SortType != "desc" {
		query.Add("sortType", p.SortType)
	}
	if len(p.FolderID) != 0 {
		query.Add("folderId", p.FolderID)
	}
	if len(p.Keyword) != 0 {
		query.Add("keyword", p.Keyword)
	}
	if len(p.OrganizationKey) != 0 {
		query.Add("organizationKey", p.OrganizationKey)
	}
	return query.Encode()
}
