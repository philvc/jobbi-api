package test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/nightborn-be/invoice-backend/database"
	"github.com/nightborn-be/invoice-backend/database/model"
	"github.com/nightborn-be/invoice-backend/repository"
	"gorm.io/gorm"
)

type TestEnvironment struct {
	Database   *gorm.DB
	Repository repository.Repository
}

// Retrieves a test database using in memory sqlite database
func Default() TestEnvironment {
	database := database.TestDefault()

	return TestEnvironment{
		Database:   database,
		Repository: repository.Default(database),
	}
}

// Fills database with mock data
func (testEnvironment TestEnvironment) Initialise() {
	createAddresses(testEnvironment.Database)
	createUsers(testEnvironment.Database)
	createClients(testEnvironment.Database)
}

// Fills database with Addresses
func createAddresses(database *gorm.DB) {

}

// Fills database with Users
func createUsers(database *gorm.DB) {

}

// Fills database with Clients
func createClients(database *gorm.DB) {
	client := model.Client{
		Name: gofakeit.Person().FirstName,
	}
	database.Create(&client)
}
