package main

import (
    "{{Root}}/models"
    "github.com/jinzhu/gorm"
)

func RunMigrations(db *gorm.DB) {
    db.AutoMigrate(&models.User{})
}
