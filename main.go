package main

import "github.com/nozzlium/heymat_backend/app"

func main() {
	inst, err := app.InitApp()
	if err != nil {
		panic(err)
	}

	err = inst.Listen(":4343")
	if err != nil {
		panic(err)
	}
}
