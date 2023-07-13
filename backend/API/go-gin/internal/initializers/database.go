package initializers

import "fmt"

type db_data struct {
	Host string
	Port string
	User string
	Password string
	DbName string
	Ssl bool
}

// THESE ARE THE ACTUAL DATABASE SETTINGS - CHANGING THESE WILL IMPACT CONNECTION SETTINGS!
var DB_Data db_data = db_data{
	Host: "localhost",
	Port: "5432",
	User: "postgres",
	Password: "postgres",
	DbName: "go_gin_api_db",
	Ssl: false,
}
// =====================================================================================

func DataAsDSN(data db_data) string {
	sslmode := "disable"
	if (data.Ssl) {
		sslmode = "enable"
	}

	var dsn string = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", data.User, data.Password, data.DbName, data.Host, data.Port, sslmode)
	return dsn
}