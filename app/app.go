package app

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ravenocx/clothes-store/db"
	"github.com/ravenocx/clothes-store/domain/repositories"
	"github.com/ravenocx/clothes-store/routes"
	"github.com/ravenocx/clothes-store/services"
)

func StartApp() {
	app := gin.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin, Content-Type, Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	db, err := db.OpenConnection()
	if err != nil {
		panic(err)
	}

	clothesRepo := repositories.NewClothesRepository(db)
	clothesService := services.NewClothesService(clothesRepo)

	routes.SetupClothesRoutes(app, clothesService)

	err = app.Run(os.Getenv(":" + "SERVER_PORT"))
	if err != nil {
		panic(err)
	}
}
