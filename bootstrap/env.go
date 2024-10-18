package bootstrap

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBType                 string `mapstructure:"DB_TYPE"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

// Helper function to handle writing environment variables and errors
func writeEnvVariableToFile(file *os.File, key, value string) {
	_, err := file.WriteString(key + "=" + value + "\n")
	if err != nil {
		log.Fatalf("Error writing %s to .env file: %v", key, err)
	}
}

func exportEnvToFile() {
	// List of environment variables
	envVars := map[string]string{
		"APP_ENV":                   os.Getenv("APP_ENV"),
		"SERVER_ADDRESS":            os.Getenv("SERVER_ADDRESS"),
		"CONTEXT_TIMEOUT":           os.Getenv("CONTEXT_TIMEOUT"),
		"DB_TYPE":                   os.Getenv("DB_TYPE"),
		"DB_HOST":                   os.Getenv("DB_HOST"),
		"DB_PORT":                   os.Getenv("DB_PORT"),
		"DB_USER":                   os.Getenv("DB_USER"),
		"DB_PASS":                   os.Getenv("DB_PASS"),
		"DB_NAME":                   os.Getenv("DB_NAME"),
		"ACCESS_TOKEN_EXPIRY_HOUR":  os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"),
		"REFRESH_TOKEN_EXPIRY_HOUR": os.Getenv("REFRESH_TOKEN_EXPIRY_HOUR"),
		"ACCESS_TOKEN_SECRET":       os.Getenv("ACCESS_TOKEN_SECRET"),
		"REFRESH_TOKEN_SECRET":      os.Getenv("REFRESH_TOKEN_SECRET"),
	}

	// Create the .env file
	file, err := os.Create(".env")
	if err != nil {
		log.Fatal("Error creating .env file: ", err)
	}
	defer file.Close() // Ensure the file is closed after writing

	// Write each environment variable to the file
	for key, value := range envVars {
		writeEnvVariableToFile(file, key, value)
	}
}

func NewEnv() *Env {
	value := os.Getenv("APP_ENV")
	log.Println("Environment variable: ", value)

	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		log.Println("No .env file found, creating one...")
		exportEnvToFile()
	}

	env := Env{}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No .env file found")
	}

	err = viper.Unmarshal(&env)
	log.Println("Environment data: ", env)

	if err != nil || env.AppEnv == "" {
		log.Fatal("Error upon loading can't be loaded: ")
	}

	if env.AppEnv == "development" {
		log.Println("Environment data: ", env)
		log.Println("The App is running in development environment")
	}

	return &env
}
