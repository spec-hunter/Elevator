package initialization

import (
	"github.com/ahlemarg/Elevator/src/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig() {
	v := viper.New()

	v.SetConfigFile("src/limit.yaml")

	if err := v.ReadInConfig(); err != nil {
		zap.S().Infof("Failed to read config file: %s\n", err.Error())
	}

	if err := v.Unmarshal(global.Config); err != nil {
		zap.S().Infof("failed to unmarshal struct: %s\n", err.Error())
	}

	v.WatchConfig()

	v.OnConfigChange(func(in fsnotify.Event) {
		zap.S().Debugf("config file update")

		v.ReadInConfig()
		if err := v.Unmarshal(global.Config); err != nil {
			zap.S().Infof("failed to unmarshal struct: %s\n", err.Error())
		}
	})
}
