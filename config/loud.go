package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type DatabaseConfig struct {
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_NAME string
}

func InitConfig() *ProgramConfig {
	godotenv.Load()

	var res = new(ProgramConfig)
	res = loadConfig()

	if res == nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

type ProgramConfig struct {
	SERVER_PORT string
	SECRET      string
	REFSECRET   string
}

func LoadDBConfig() *DatabaseConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	} else {
		log.Println(".env file loaded successfully")
	}

	var res = new(DatabaseConfig)

	if val, found := os.LookupEnv("DB_USER"); found {
		res.DB_USER = val
	}

	if val, found := os.LookupEnv("DB_PASS"); found {
		res.DB_PASS = val
	}

	if val, found := os.LookupEnv("DB_HOST"); found {
		res.DB_HOST = val
	}

	if val, found := os.LookupEnv("DB_PORT"); found {
		res.DB_PORT = val
	}

	if val, found := os.LookupEnv("DB_NAME"); found {
		res.DB_NAME = val
	}

	return res
}

func loadConfig() *ProgramConfig {
	var res = new(ProgramConfig)

	if val, found := os.LookupEnv("SERVER_PORT"); found {
		res.SERVER_PORT = val
	}
	if val, found := os.LookupEnv("SECRET"); found {
		res.SECRET = val
	}
	if val, found := os.LookupEnv("REFSECRET"); found {
		res.REFSECRET = val
	}
	return res
}

type AWSConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	Region          string
}

func LoadAwsConfig() *AWSConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	} else {
		log.Println(".env file loaded successfully")
	}

	var res = new(AWSConfig)
	if val, found := os.LookupEnv("AWS_ACCESS_KEY_ID"); found {
		res.AccessKeyID = val
	}
	if val, found := os.LookupEnv("AWS_ACCESS_KEY_SECRET"); found {
		res.AccessKeySecret = val
	}
	if val, found := os.LookupEnv("AWS_REGION"); found {
		res.Region = val
	}
	return res
}
