package main

import "BluebellSpace/router"

func main() {
	route := router.New()
	_ = route.Run("127.0.0.1:7749")
}
