package utils

import (
	"log"

	"github.com/EricOgie/ope-be/logger"
	"github.com/spf13/viper"
)

type Config struct {
	AppName       string `mapstructure:"APP_NAME"`
	AppEnv        string `mapstructure:"APP_ENV"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`

	SigningKey string `mapstructure:"JWT_SECRETE"`

	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBUser     string `mapstructure:"DB_USER"`
	DBAddress  string `mapstructure:"DB_ADDR"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	Mailer          string `mapstructure:"MAIL_MAILER"`
	MailHost        string `mapstructure:"MAIL_HOST"`
	MailPort        string `mapstructure:"MAIL_PORT"`
	MailUserName    string `mapstructure:"MAIL_USERNAME"`
	MailPassword    string `mapstructure:"MAIL_PASSWORD"`
	MailEncryption  string `mapstructure:"MAIL_ENCRYPTION"`
	MailFromAddress string `mapstructure:"MAIL_FROM_ADDRESS"`
	MailFromName    string `mapstructure:"MAIL_FROM_NAME"`
}

// LoadConfig reads/loads environment variables into config struct.
// It returns a Config struct with all loaded envs as attributes.
// Each env can then be accessed by the DOT notation on the Config struct like so: config.DBAddress
func LoadConfig(path string) (config Config, err error) {
	// Set config file path to env fle
	viper.AddConfigPath(path)
	// define what file to be looked with config name
	viper.SetConfigName("")
	// Define the type of file to ve looked
	viper.SetConfigType("env")
	// configure auto override config variables with set environment variables
	viper.AutomaticEnv()

	// Initiate read config value
	err = viper.ReadInConfig()
	if err != nil {
		logger.Error("ReadConfigError " + err.Error())
		return
	} else {
		err = viper.Unmarshal(&config)
		return
	}

}

// RunSanityCheck runs a check on the system to ensure all essential variables are properly read from env.
// It will KILL the system if it can not read set variables from config
func RunSanityCheck(err error) {
	if err != nil {
		logger.Debug("Error While loading Config. ErrorMsg: " + err.Error())
		log.Fatal("cannot load config:", err)
	}
}
