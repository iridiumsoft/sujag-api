package models

type Magazine struct {
	Title     string  `json:"title,omitempty"`
	Is        string  `json:"is,omitempty"`
	Img       string  `json:"img,omitempty"`
	CreatedOn float64 `json:"created_on,omitempty"`
}
