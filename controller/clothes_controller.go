package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ravenocx/clothes-store/domain/dto"
	"github.com/ravenocx/clothes-store/domain/entities"
	"github.com/ravenocx/clothes-store/utils"
)

func (c *ClothesController) InsertCloth(ctx *gin.Context) {
	var clothPayload dto.ClothesPayload

	if err := ctx.ShouldBind(&clothPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := utils.NewValidator()
	if err := validate.Struct(&clothPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cloth := &entities.Clothes{
		Color: clothPayload.Color,
		Size:  clothPayload.Size,
		Price: clothPayload.Price,
		Stock: clothPayload.Stock,
	}

	cloth, err := c.clothesService.InsertCloth(cloth)
	if e, ok := err.(*utils.ErrorMessage); ok {
		ctx.JSON(e.ErrorCode(), gin.H{"error": e.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "successfully add cloth",
		"data":    cloth,
	})
}
