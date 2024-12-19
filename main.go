package main

import (
	"Github.com/Uallessonivo/RealWorld/cmd"
	"Github.com/Uallessonivo/RealWorld/config"
)

func main() {
	config.InitDatabase()
	cmd.Execute()
}
