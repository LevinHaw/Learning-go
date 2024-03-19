package main

import (
	"sesi_7/pkg/database"
	"sesi_7/pkg/routers"
)

func main() {
	port := ":7000"

	db := database.SQLInit()

	gorm := database.GormInit(db)

	defer db.Close()

	start := routers.StartServer(db, gorm)
	start.Run(port)

}
