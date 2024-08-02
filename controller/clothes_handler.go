package controller

import "github.com/ravenocx/clothes-store/services"

type ClothesController struct {
	clothesService services.ClothesService
}

func NewClothesController(clothesService services.ClothesService) *ClothesController {
	return &ClothesController{clothesService: clothesService}
}
