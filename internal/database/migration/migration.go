package migration

import (
	"fmt"
	"moonlay-test/internal/app/models"
	"moonlay-test/internal/database"
)

func RunMigration() {
	err := database.DB.AutoMigrate(
		&models.TodoList{},
		&models.SubList{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
