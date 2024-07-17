package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Initialize sets up the logger
func Initialize() (*logrus.Logger, error) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	logger.SetOutput(file)

	logger.SetLevel(logrus.InfoLevel)
	return logger, nil
}
