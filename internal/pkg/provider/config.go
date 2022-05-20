package provider

import (
	"github.com/spf13/viper"
	"strings"
)

func NewConfig() *viper.Viper {
	conf := viper.New()
	conf.AutomaticEnv()

	replacer := strings.NewReplacer(".", "_")
	conf.SetEnvKeyReplacer(replacer)

	return conf
}
