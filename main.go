package main

import (
	"os"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/extractors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/varakornpz/auth"
	"github.com/varakornpz/gorm"
	"github.com/varakornpz/providers"
	// "github.com/golang-jwt/jwt/v5"
)



func main(){
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	providers.InitAppConf()
	auth.InitGoogleAuthConf()

	app := fiber.New()
	InitCORSConf(app)
	app.Get("/" , func (c fiber.Ctx) error  {
		return c.SendString("huh , wtf is this place.")
	})

	authRoute := app.Group("/auth")

	authRoute.Get("/google/callback" , auth.GoogleAuthCallBack)
	authRoute.Get("/google/signin" , auth.GoogleAuthSignin)

	mainAppRoute := app.Group("/app")
	mainAppRoute.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(providers.AppConf.JWTSecret)},
		Extractor:  extractors.FromAuthHeader("Bearer"),
		ErrorHandler: func (c fiber.Ctx , err error) error {
			return c.SendStatus(fiber.ErrUnauthorized.Code)
		},
	}))

	mainAppRoute.Get("/hi" , func (c fiber.Ctx) error  {
		return c.SendString("You are wowza")
	})


	gorm.InitDB()

	

	

	app.Listen(":3334")
}