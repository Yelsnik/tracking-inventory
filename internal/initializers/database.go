package initializers

import (
	"fmt"
	"github.com/Yelsnik/tracking-inventory.git/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func ConnectToDatabase(user, password, dbname string) {

	var err error

	dsn := fmt.Sprintf("host=host.docker.internal user=%s password=%s dbname=%s port=5432 sslmode=disable", user, password, dbname)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	failOnError(err, "ERROR DB CONNECTION")

	Migrate(models.Inventory{})
	Migrate(models.User{})
}

type Migration interface {
	models.User | models.Inventory
}

func Migrate[M Migration](model M) error {

	err := DB.AutoMigrate(&model)

	if err != nil {
		return err
	}
	return nil
}
