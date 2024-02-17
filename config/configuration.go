package config

import "os"

const (
	LOCAL       = "local"
	DEVELOPMENT = "development"
)

// ENVIRONMENT
const ENVIRONMENT string = DEVELOPMENT

// MAP LIST ENVIRONMENT LOCAL AND DEVELOPMENT
var env = map[string]map[string]string{
	"local": {

		"MS_PORT": "9000",

		"MYSQL_HOST":   "127.0.0.1",
		"MYSQL_PORT":   "3306",
		"MYSQL_USER":   "root",
		"MYSQL_PASS":   "N#@98wrft45",
		"MYSQL_SCHEMA": "rumah_sakit",
	},
	"development": {

		"MS_PORT": "11000",

		"MYSQL_HOST":   "127.0.0.1",
		"MYSQL_PORT":   "3306",
		"MYSQL_USER":   "root",
		"MYSQL_PASS":   "N#@98wrft45",
		"MYSQL_SCHEMA": "rumah_sakit",
	},
}

var CONFIG = env[ENVIRONMENT]

// CHECK ENVIRONMENT AND GET ENVIRONMENT
func Getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

// INITIALIZE CONFIGURATION
func InitConfig() {
	for key := range CONFIG {
		CONFIG[key] = Getenv(key, CONFIG[key])

		os.Setenv(key, CONFIG[key])
	}
}
