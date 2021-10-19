package test

import (
	"github.com/philvc/jobbi-api/database"
	"github.com/philvc/jobbi-api/repository"
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
}

// Fills database with Addresses
func createAddresses(database *gorm.DB) {

}

// Fills database with Users
func createUsers(database *gorm.DB) {

}


