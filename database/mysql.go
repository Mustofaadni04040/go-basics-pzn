package database

var connection string

func init() { // auto executed when package is imported
	connection = "mysql connection"
}

func GetDatabase() string {
	return connection
}