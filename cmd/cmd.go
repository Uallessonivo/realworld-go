package cmd

import (
	"Github.com/Uallessonivo/RealWorld/pkg/database"
	"Github.com/Uallessonivo/RealWorld/pkg/server"
)

func Execute() {
	database.InitDb()
	server.InitServerApp()
}
