package dto

type Destination struct {
	Id              int     `json:"id"`
	Img             *string `json:"image"`
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	DescriptiveText *string `json:"descriptive_text"`
}
