package settings

import (
	"bytes"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

var homedir = os.Getenv("HOME")
var buffer bytes.Buffer
var env = "development"
var environments = map[string]string{
	"production":  ".config/vantaa/production.toml",
	"development": ".config/vantaa/development.toml",
	"test":        ".config/vantaa/test.toml",
}

var settings Settings = Settings{}

type Settings struct {
	HashCost           int
	JWTExpirationDelta int
	DbUrl              string
	PrivateKeyPath     string
	PublicKeyPath      string
	isset              bool
}

func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting development environment due to lack of GO_ENV value")
		env = "development"
	}
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	buffer.WriteString(homedir)
	buffer.WriteString("/")
	buffer.WriteString(environments[env])
	if _, err := toml.DecodeFile(buffer.String(), &settings); err != nil {
		fmt.Println("Failed to load configuration file in ~/config/vantaa/")
		panic(err)
	}
	settings.isset = true
}

func GetEnvironment() string {
	return env
}

func Get() Settings {
	if !settings.isset {
		Init()
	}
	return settings
}

func IsTestEnvironment() bool {
	return env == "test"
}
