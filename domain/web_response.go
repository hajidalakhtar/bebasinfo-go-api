package domain

type WebResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type WebRespPaginate struct {
	Code     int               `json:"code"`
	Status   string            `json:"status"`
	Message  string            `json:"message"`
	Data     interface{}       `json:"data"`
	Paginate PaginatedResponse `json:"paginate"`
}

type PaginatedResponse struct {
	TotalItems  int64 `json:"total_items"`
	TotalPages  int   `json:"total_pages"`
	CurrentPage int   `json:"current_page"`
	NextPage    int   `json:"next_page"`
	PrevPage    int   `json:"prev_page"`
}
