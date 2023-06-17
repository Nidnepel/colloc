package config

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDb       string
	PostgresDbHost   string
}

func NewConfig() Config {
	return Config{
		PostgresUser:     "postgres",
		PostgresPassword: "password",
		PostgresDb:       "hotel-api-db",
		PostgresDbHost:   "localhost",
	}
}
