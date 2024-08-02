package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ravenocx/clothes-store/controller"
	"github.com/ravenocx/clothes-store/services"
)

func SetupClothesRoutes(r *gin.Engine, clothesService services.ClothesService) {
	clothesController := controller.NewClothesController(clothesService)

	router := r.Group("/api/v1/clothes")

	router.POST("", clothesController.InsertCloth)
	router.GET("", clothesController.GetClothes)
	router.PUT("/:id", clothesController.UpdateCloth)
	router.DELETE("/:id", clothesController.DeleteCloth)

	router.PUT("/:id/stock/add", clothesController.IncreaseStock)
	router.PUT("/:id/stock/reduce", clothesController.DecreaseStock)
}
