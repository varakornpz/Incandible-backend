package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)


func InitCORSConf(app *fiber.App){
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000", "https://incanedible.varakorn.net" ,
		},
		AllowCredentials: true,
		AllowHeaders: []string{
			"Origin","Content-Type","Accept","Authorization",
		},
		AllowMethods: []string{
			"GET","POST","PUT","DELETE" , "OPTIONS",
		},
	}))	
}