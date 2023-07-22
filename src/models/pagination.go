package models

type JsonResponse struct {
	Page      int         `json:"curr_page"`
	TotalPage int         `json:"total_page"`
	TotalObjs int64       `json:"total_objs"`
	PerPage   int         `json:"per_page"`
	Objs      interface{} `json:"objs"`
}

type ListRequest struct {
	Search   []Filter `json:"search,omitempty"`
	Sort     Filter   `json:"sort,omitempty"`
	Page     int      `json:"page,omitempty"`
	Limit    int      `json:"limit,omitempty"`
	Download bool     `json:"download"`
}

type Filter struct {
	Field string `json:"field,omitempty"`
	Value string `json:"value,omitempty"`
}
