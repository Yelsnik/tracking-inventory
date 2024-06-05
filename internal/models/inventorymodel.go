package models

import (
	"gorm.io/gorm"
	"time"
	//_ "github.com/lib/pq"
	//"github.com/gin-gonic/gin"
)

type Gormmodel struct {
	ID        string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Inventory struct {
	Gormmodel
	Item         string `json:"item"`
	SerialNumber string `json:"serialNumber"`
	Value        string `json:"value"`
	User         string `gorm:"foreignKey:ID"`
}

type User struct {
	Gormmodel
	Name     string `gorm:"size:255;not null;" json:"name"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"-"`
}
