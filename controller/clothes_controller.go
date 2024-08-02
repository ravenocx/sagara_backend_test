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

func (c *ClothesController) GetClothes(ctx *gin.Context) {
	getClothesQuery := dto.GetClothesQuery{}

	if color := ctx.Query("color"); color != "" {
		getClothesQuery.Color = color
	}

	if size := ctx.Query("size"); size != "" {
		getClothesQuery.Size = size

		// validate the size query
		validate := utils.NewValidator()
		if err := validate.Struct(&getClothesQuery); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// log.Printf("Query Params : %+v" , getClothesQuery)
	clothes, err := c.clothesService.GetClothes(getClothesQuery)
	if e, ok := err.(*utils.ErrorMessage); ok {
		ctx.JSON(e.ErrorCode(), gin.H{"error": e.Error()})
		return
	}

	resp := []entities.Clothes{}
	resp = append(resp, clothes...)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully get clothes",
		"data":    resp,
	})

}

func (c *ClothesController) UpdateCloth(ctx *gin.Context) {
	id := ctx.Param("id")

	var clothPayload dto.ClothesPayload

	if err := ctx.ShouldBind(&clothPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := utils.NewValidator()

	if clothPayload.Size != "" {
		if err := validate.Struct(&clothPayload); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	cloth, err := c.clothesService.GetClothByID(id)
	if e, ok := err.(*utils.ErrorMessage); ok {
		ctx.JSON(e.ErrorCode(), gin.H{"error": e.Error()})
		return
	}

	cloth.Color = clothPayload.Color
	cloth.Size = clothPayload.Size
	cloth.Price = clothPayload.Price
	cloth.Stock = clothPayload.Stock

	cloth, err = c.clothesService.UpdateCloth(cloth)
	if e, ok := err.(*utils.ErrorMessage); ok {
		ctx.JSON(e.ErrorCode(), gin.H{"error": e.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully updated cloth",
		"data":    cloth,
	})

}

func (c *ClothesController) DeleteCloth(ctx *gin.Context) {
	id := ctx.Param("id")

	if err, ok := c.clothesService.DeleteCloth(id).(*utils.ErrorMessage); ok {
		ctx.JSON(err.ErrorCode(), gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted cloth data",
	})
}
