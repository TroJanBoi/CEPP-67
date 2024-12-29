package main

import (
	"github.com/TroJanBoi/CEPP-67/services"
)

func main() {

	db, err := services.SetUpDB()

	if err != nil {
		panic(err)
	}
	app := services.InitializeRouter(db)
	services.SetUpServer(app)
}