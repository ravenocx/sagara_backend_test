package main

import (
	"github.com/joho/godotenv"
	"github.com/ravenocx/clothes-store/app"
	"github.com/ravenocx/clothes-store/db"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	_, err = db.OpenConnection()

	if err != nil {
		panic(err)
	}

	app.StartApp()

}
