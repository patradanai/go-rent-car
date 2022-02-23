package main

import (
	"car-booking/configs"
)

func main() {
	if _, err := configs.ConnectToDB(); err != nil {
		panic(err)
	}
}
