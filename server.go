package main

import (
	"DompetKilat-SimpleWeb/routes"
)

func main() {
	e := routes.Init()
	defer e.Close()

	e.Logger.Fatal(e.Start(":8000"))
}
