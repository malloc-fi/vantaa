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

// Settings defines the Settings struct.
type Settings struct {
	HashCost           int
	JWTExpirationDelta int
	DbUrl              string
	PrivateKeyPath     string
	PublicKeyPath      string
	isset              bool
}

// Init get the environments value from flags, default fall back is development.
func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting development environment due to lack of GO_ENV value")
		env = "development"
	}
	LoadSettingsByEnv(env)
}

// LoadSettingsByEnv load the .toml setting files from
// ~/.config/vantaa/<environment>.toml.
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

// GetEnvironment returns the current environment
func GetEnvironment() string {
	return env
}

// Get gets and return Settings according to the configuration files
func Get() Settings {
	if !settings.isset {
		Init()
	}
	return settings
}

// IsTestEnvironment check if the current environment is test
func IsTestEnvironment() bool {
	return env == "test"
}
