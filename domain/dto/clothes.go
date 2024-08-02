package dto

type ClothesPayload struct {
	Color string  `json:"color" validate:"required"`
	Size  string  `json:"size" validate:"required,oneof='S' 'M' 'L' 'XL' 'XXL'"`
	Price float64 `json:"price" validate:"required,min=1"`
	Stock int     `json:"stock" validate:"required,min=1"`
}

type GetClothesQuery struct {
	Color    string `json:"color"`
	Size     string `json:"size" validate:"oneof='S' 'M' 'L' 'XL' 'XXL'"`
	Stock    string
}
