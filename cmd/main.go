package main

import "github.com/dkgv/dislikes/internal/api"

func main() {
	api := api.NewAPI()
	err := api.Start()
	if err != nil {
		panic(err)
	}
}
