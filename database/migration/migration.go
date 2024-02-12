package migration

import (
	"api/database"
	"api/model/entity"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.Users{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated!")
}
