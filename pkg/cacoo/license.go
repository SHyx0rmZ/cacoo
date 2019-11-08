package cacoo

import (
	"context"
	"net/http"
)

func NewLicenseRequest(ctx context.Context) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, http.MethodGet, "https://cacoo.com/api/v1/account/license.xml", nil)
}

func (c *Client) License(ctx context.Context) (LicenseResponse, error) {
	var r LicenseResponse
	err := c.do(NewLicenseRequest(ctx))(&r)
	if err != nil {
		return LicenseResponse{}, err
	}
	return r, nil
}

type LicenseResponse struct {
	Plan                              string `xml:"plan"`
	RemainingSheets                   int    `xml:"remainingSheets"`
	RemainingSharedFolders            int    `xml:"remainingSharedFolders"`
	MaxNumberOfSharersPerDiagram      int    `xml:"maxNumberOfSharersPerDiagram"`
	MaxNumberOfSharersPerSharedFolder int    `xml:"maxNumberOfSharersPerSharedFolder"`
	CanCreateSheet                    bool   `xml:"canCreateSheet"`
	CanCreateSharedFolder             bool   `xml:"canCreateSharedFolder"`
	CanExportVectorFormat             bool   `xml:"canExportVectorFormat"`
}
