package gorm

import (
	"github.com/google/uuid"
	"github.com/varakornpz/models"
)


func GetUserByEmail(email string) (models.User , error){
	var user models.User
	result := DB.First(&user ,"email = ?" , email)

	return user , result.Error
}

func GetUserByUUID(uuid uuid.UUID) (models.User , error){
	var user models.User
	result := DB.First(&user , "uuid = ?" , uuid)

	return user , result.Error
}


func PutNewUser(user *models.User) error {
	result := DB.Create(user)
	return  result.Error
}