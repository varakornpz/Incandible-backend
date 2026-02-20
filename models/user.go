package models


import (
    "time"
	"gorm.io/gorm"
    "github.com/lib/pq"
    "github.com/google/uuid"
)


type User struct{
    CreatedAt      time.Time
    UpdatedAt      time.Time
    DeletedAt      gorm.DeletedAt `gorm:"index"`
    UUID            uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"` 
    Email          string
    Name            string
    ProfilePic     string   `gorm:"type:text"`
    RegisteredCand pq.StringArray `gorm:"type:text[]"`
}