package main

import (
	"github.com/Nilay1999/gin-gonic-server/initializers"
	"github.com/Nilay1999/gin-gonic-server/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Init()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}