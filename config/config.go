package config

type Config struct {
	Database Database
}

type Database struct {
	ConnectionString string
}
