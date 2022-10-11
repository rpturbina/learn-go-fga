package main

import (
	"github.com/rpturbina/learn-go-fga/belajar-jwt/database"
	"github.com/rpturbina/learn-go-fga/belajar-jwt/router"
)

func main() {
	database.StartBD()
	r := router.StartApp()
	r.Run(":8080")
}
