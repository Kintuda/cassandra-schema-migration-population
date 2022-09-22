package main

import (
	"github.com/spf13/viper"
)

type MigratorConfig struct {
	ContactPoints string `mapstructure:"CONTACT_POINTS"`
}

func LoadConfigFromEnv(m *MigratorConfig) error {
	v := viper.New()

	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&m); err != nil {
		return err
	}

	return nil
}
