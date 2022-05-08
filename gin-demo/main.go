package main

import "gin-demo/router"

func main() {
	r := router.Router()

	r.Run("127.0.0.1:8080")
}
