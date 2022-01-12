package database

import "fmt"

//Config to maintain DB configuration properties
type Config struct {
	Server       string
	User         string
	Password     string
	DataBaseName string
	Port         int
}

var GetConnectionString = func(config Config) string {

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		config.Server, config.User, config.Password, config.Port, config.DataBaseName)
	return connectionString

}
