package sna

import (
	"encoding/json"
	"errors"
	"net/http"
)

// GetTenant - Returns ID of tenant
func (c *Client) GetTenant(tenantName string) (*int, error) {
	req, err := http.NewRequest("GET", getTenantUrl(c), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	tenants := TenantData{}
	err = json.Unmarshal(body, &tenants)
	if err != nil {
		return nil, err
	}

	for _, tenant := range tenants.Data {
		if tenant.DisplayName == tenantName {
			return &tenant.Id, nil
		}
	}

	return nil, errors.New("tenant does not exist")
}
