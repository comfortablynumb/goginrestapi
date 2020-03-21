package config

// Structs

type AppConfig struct {
	Port             int    `default:"8080"`
	LogLevel         string `default:"DEBUG"`
	DbUri            string `default:"file:test.db?cache=shared&mode=memory"`
	DbMigrationsPath string `default:"file://database/migrations"`
	DefaultLocale    string `default:"en"`
	DefaultLimit     int    `default:"50"`
}

// Static functions

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}
