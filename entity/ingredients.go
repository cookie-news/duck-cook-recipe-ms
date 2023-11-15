package entity

type Ingredients struct {
	Name    string  `json:"name"`
	Qty     float64 `json:"quantity"`
	Measure string  `json:"measure"`
}
