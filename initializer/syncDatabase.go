package initializers

import "github.com/hiroki-0815/go-jwt2/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
