package main

import "GoLandPruebas/api"

func main() {
	if err := api.NewServer().Run(); err != nil {
		panic(err.Error())
	}
}
