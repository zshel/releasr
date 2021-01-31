package config

import (
	"encoding/base64"
	"fmt"
	"github.com/packagrio/go-common/errors"
	"github.com/spf13/viper"
)

// When initializing this class the following methods must be called:
// Config.New
// Config.Init
// This is done automatically when created via the Factory.
type configuration struct {
	*viper.Viper
}

//Viper uses the following precedence order. Each item takes precedence over the item below it:
// explicit call to Set
// flag
// env
// config
// key/value store
// default

func (c *configuration) Init() error {
	c.Viper = viper.New()
	//set defaults
	c.SetDefault(PACKAGR_PACKAGE_TYPE, "generic")
	c.SetDefault(PACKAGR_SCM, "default")
	c.SetDefault(PACKAGR_VERSION_BUMP_TYPE, "patch")

	//set the default system config file search path.
	//if you want to load a non-standard location system config file (~/capsule.yml), use ReadConfig
	//if you want to load a repo specific config file, use ReadConfig
	c.SetConfigType("yaml")
	c.SetConfigName("packagr")
	c.AddConfigPath("$HOME/")

	//configure env variable parsing.
	c.SetEnvPrefix("PACKAGR")
	c.AutomaticEnv()
	//CLI options will be added via the `Set()` function

	return nil
}

func (c *configuration) GetBase64Decoded(key string) (string, error) {
	if len(c.GetString(key)) > 0 {
		key, err := base64.StdEncoding.DecodeString(c.GetString(key))
		if err != nil {
			return "", errors.ScmUnspecifiedError(fmt.Sprintf("Could not decode base64 key (%s): %s", key, err))
		}
		return string(key), nil
	} else {
		return "", nil
	}
}
