package main

import (
	"sesi_7/pkg/database"
	"sesi_7/pkg/routers"
)

func main() {
	port := ":6000"

	db := database.SQLInit()

	defer db.Close()

	start := routers.StartServer(db)
	start.Run(port)

}
