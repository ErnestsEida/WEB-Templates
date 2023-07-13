package initializers

type db_data struct {
	Host string
	Port string
	User string
	Password string
	Ssl bool
	DbName string
}

var DB_Data db_data = db_data{
	Host: "localhost",
	Port: "5432",
	User: "postgres",
	Password: "postgres",
	Ssl: false,
	DbName: "go_gin_api_db",
}

// func DataAsDSN() string {

// 	return value
// }