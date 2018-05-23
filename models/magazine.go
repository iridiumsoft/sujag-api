package models

type Magazine struct {
	Title     string  `json:"title,omitempty"`
	Is        string  `json:"is,omitempty"`
	CreatedOn float64 `json:"created_on,omitempty"`
}
