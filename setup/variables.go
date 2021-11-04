package setup

import (
	"fmt"
	"os"

	"github.com/EricOgie/ope-be/konstants"
)

type ENV struct {
	ServerAddress string
	DBUser        string
	DBAddress     string
	DBPort        string
	DBName        string
}

// Helper function to fetch Set ENVIRONMENT Variables
func GetSetENVs() ENV {
	return ENV{
		ServerAddress: getServerAdd(),
		DBUser:        getDBUser(),
		DBAddress:     getDBAdd(),
		DBPort:        getDBPort(),
		DBName:        getDBName(),
	}
}

// ---------------------  PRIVATE HELPER FUNCTIONS --------------------------- //
func isEnvSet() bool {
	_, isSet := os.LookupEnv("SERVER_ADDRESS")

	if isSet {
		return true
	} else {
		return false
	}
}

func getServerAdd() string {
	if isEnvSet() {
		serverAdd := os.Getenv("SERVER_ADDRESS")
		serverPort := os.Getenv("SERVER_PORT")
		address := fmt.Sprintf("%s:%s", serverAdd, serverPort)
		return address
	} else {
		return konstants.LOCAL_ADD
	}
}

func getDBUser() string {
	_, isSet := os.LookupEnv("DB_USER")
	if isSet {
		return os.Getenv("DB_USER")
	} else {
		return konstants.LOCAL_DB_USER
	}
}

func getDBName() string {
	_, isSet := os.LookupEnv("DB_NAME")
	if isSet {
		return os.Getenv("DB_NAME")
	} else {
		return konstants.LOCAL_DB_NAME
	}
}

func getDBPassword() string {
	_, isSet := os.LookupEnv("DB_PASSWD")
	if isSet {
		return os.Getenv("DB_PASSWD")
	} else {
		return ""
	}
}

func getDBAdd() string {
	_, isSet := os.LookupEnv("DB_ADDR")
	if isSet {
		return os.Getenv("DB_ADDR")
	} else {
		return konstants.LOCAL_DB_ADD
	}
}

func getDBPort() string {
	_, isSet := os.LookupEnv("DB_PORT")
	if isSet {
		return os.Getenv("DB_PORT")
	} else {
		return ""
	}
}
