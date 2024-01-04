package sna

// Order -
type Order struct {
	ID    int         `json:"id,omitempty"`
	Items []OrderItem `json:"items,omitempty"`
}

// OrderItem -
type OrderItem struct {
	Tag      Tag `json:"tag"`
	Quantity int `json:"quantity"`
}

// TenantData -
type TenantData struct {
	Data []Tenant `json:"data"`
}

// Tenant -
type Tenant struct {
	Id          int    `json:"id"`
	DisplayName string `json:"displayName"`
}

// TagListData -
type TagListData struct {
	Data []TagList `json:"data"`
}

// TagList -
type TagList struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// TagData -
type TagData struct {
	Data Tag `json:"data"`
}

// Tag -
type Tag struct {
	ID                       int        `json:"id"`
	Name                     string     `json:"name"`
	Location                 string     `json:"location"`
	Ranges                   []string   `json:"ranges"`
	Description              string     `json:"description"`
	HostBaselines            bool       `json:"hostBaselines"`
	SuppressExcludedServices bool       `json:"suppressExcludedServices"`
	InverseSuppression       bool       `json:"inverseSuppression"`
	HostTrap                 bool       `json:"hostTrap"`
	SendToCta                bool       `json:"sendToCta"`
	DomainId                 int        `json:"domainId"`
	ParentId                 int        `json:"parentId"`
	Display                  TagDisplay `json:"display"`
	ParentDisplay            TagDisplay `json:"parentDisplay"`
}

// TagDisplay Tag Display Reference -
type TagDisplay struct {
	DomainID int      `json:"domain_id"`
	Editable bool     `json:"editable"`
	ID       int      `json:"id"`
	Location string   `json:"location"`
	Name     string   `json:"name"`
	Path     []string `json:"path"`
	IdPath   []int    `json:"id_path"`
}
