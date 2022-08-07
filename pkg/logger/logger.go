package logger

import (
	"fmt"
	"github.com/fd239/gopher_keeper/config"
	"log"
	"os"

	zap "go.uber.org/zap"
)

const (
	currentLogsPath = "_logs"
	serviceName     = "ocs"
)

//NewLogger create new logger
func NewLogger(cfg *config.Config) *zap.SugaredLogger {
	var zapConfig zap.Config

	// Create directory _logs if not exist
	if _, err := os.Stat(currentLogsPath); os.IsNotExist(err) {
		err = os.Mkdir(currentLogsPath, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Open log file by path _logs/{serviceName}.log
	logFilePath := fmt.Sprintf("%s/%s.log", currentLogsPath, serviceName)
	_, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}

	debugMode := cfg.Logger.Debug
	if debugMode {
		zapConfig = zap.NewDevelopmentConfig()
	} else {
		zapConfig = zap.NewProductionConfig()
	}

	zapConfig.Encoding = "json"
	zapConfig.OutputPaths = []string{"stdout", logFilePath}
	zapConfig.ErrorOutputPaths = []string{"stdout"}
	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
	return logger.Sugar()
}
