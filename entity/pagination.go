package entity

type Pagination struct {
	Next          int         `json:"next"`
	Previous      int         `json:"previous"`
	RecordPerPage int         `json:"recordPerPage"`
	CurrentPage   int         `json:"currentPage"`
	TotalPage     int         `json:"totalPage"`
	Items         interface{} `json:"items"`
}
