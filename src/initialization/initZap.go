package initialization

import (
	"go.uber.org/zap"
)

func InitLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("日志组件无法初始化.")
	}

	zap.ReplaceGlobals(logger)
}
