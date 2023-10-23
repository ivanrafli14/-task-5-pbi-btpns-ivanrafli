package main

import (
	"github.com/ivanrafli14/API-BTPN/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load();

	if err != nil {
		panic(err)
	}

	r := router.SetupRouter()
	r.Run(":3000")
}