package configs

import (
	"errors"
	"os"

	"github.com/kweusuf/url-shortner/pkg/constants"
	"github.com/kweusuf/url-shortner/pkg/model"
	"github.com/kweusuf/url-shortner/pkg/utils/log"
)

type AppConfig struct {
	DB      *model.DBConfig
	URI     *model.AppURI
	BaseURI string
	// Typically we will get list of users from an auth server but since we are bundling everything in a single service for POC, we save an array of users in the config
	Users          []string
	ConfigFilePath string
	ConfigFileName string
}

// mongodb+srv://kweusuf:<db_password>@job-scheduler-app.asftf.mongodb.net/?retryWrites=true&w=majority&appName=job-scheduler-app
func GetConfig() *AppConfig {
	return &AppConfig{
		DB: &model.DBConfig{
			Dialect:  "mysql",
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "password",
			Name:     "mysql",
			Charset:  "utf8",
		},
		URI: &model.AppURI{
			Host:       "localhost",
			Port:       ":3000",
			HttpScheme: "http",
		},
		BaseURI:        "/api/v1",
		Users:          []string{"admin", "eusuf"},
		ConfigFilePath: GetCWD() + string(os.PathSeparator) + "data" + string(os.PathSeparator),
		ConfigFileName: "conf.yml",
	}
}

func InitializeConfig(config *AppConfig) error {
	// TODO: We want to implement different ways to connect to some source and load the config.
	// 			The options are:
	// 				1. Default config hard coded in the code
	// 				2. Config saved in a json/xml/yml/properties file saved on a default location/location from a env variable
	// 				3. Config saved in a database. We can either pass the credentials at startup or load from env variable
	return errors.New(constants.ErrEmptyConfig)
}

func GetCWD() string {
	mydir, err := os.Getwd()
	if err != nil {
		log.Error(err.Error())
	}
	return mydir
}
