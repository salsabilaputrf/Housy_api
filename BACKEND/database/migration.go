package database

import (
  "housy/models"
  "housy/pkg/mysql"
  "fmt"
)

// Automatic Migration if Running App
func RunMigration() {
  err := mysql.DB.AutoMigrate(
    &models.User{},
    &models.ListAs{},
    &models.City{},
    &models.House{},
  )

  if err != nil {
    fmt.Println(err)
    panic("Migration Failed")
  }

  fmt.Println("Migration Success")
}