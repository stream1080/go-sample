package main

import "go-sample/router"

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://github.com/stream108

// @contact.name 一江溪水
// @contact.url https://github.com/stream1080
// @contact.email example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	r := router.Router()

	r.Run(":8080")
}