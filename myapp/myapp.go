package myapp

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/varakornpz/gorm"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)


func GetUserData(c fiber.Ctx) error{
	userToken := jwtware.FromContext(c)

	if userToken == nil {
            return c.Status(401).JSON(fiber.Map{"msg" : "Token not found"})
    }

	claims , claimsOk := userToken.Claims.(jwt.MapClaims)
	if !claimsOk {
		log.Error().Msg("Claims error in myapp.go")
		return  c.SendStatus(fiber.ErrInternalServerError.Code)
	}

	uuidStr , uuidOk := claims["uuid"].(string)
	if !uuidOk {
		log.Error().Msg("uuid error in myapp.go")
		return  c.SendStatus(fiber.ErrInternalServerError.Code)
	}

	parsedUUID, uuidErr := uuid.Parse(uuidStr)
	if uuidErr != nil {
		log.Error().Msg("uuid convert error in myapp.go")
		return  c.SendStatus(fiber.ErrInternalServerError.Code)
	}

	user , getErr := gorm.GetUserByUUID(parsedUUID)
	if getErr != nil {
		log.Error().Msg("get user error in myapp.go")
		return  c.SendStatus(fiber.ErrInternalServerError.Code)
	}


	return c.JSON(fiber.Map{
		"email" : user.Email ,
		"name" : user.Name ,
		"profile_pic" : user.ProfilePic ,
		"canes" : user.RegisteredCand ,
	})
}