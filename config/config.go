package config

import (
	"os"

	"github.com/spf13/viper"
)

func Init(cfgFile string) error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".termfi")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	// Fill with defaults if needed
	for _, value := range configKeys {
		if viper.GetString(value) == "" {
			viper.Set(value, "")
		}
	}

	Write()

	return nil
}

func Write() error {
	return viper.WriteConfig()
}

func GetPwd() string {
	return viper.ConfigFileUsed()
}
