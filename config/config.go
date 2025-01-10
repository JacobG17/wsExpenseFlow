package config

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func GetConfig() Config {
	return Config{
		DBHost: "localhost",
		DBPort: "5432",
		DBUser: "postgres",
		DBPass: "hermano1",
		DBName: "expenseflowdb",
	}
}