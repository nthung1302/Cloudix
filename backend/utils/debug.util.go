package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Response(c *gin.Context, status int, code int, message string, data any) {
	resp := gin.H{
		"status":  code,
		"message": message,
		"data":    data,
	}

	c.JSON(status, resp)
}

var Logger = logrus.New()

func InitLogger() {
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	Logger.SetLevel(logrus.ErrorLevel)

	Logger.SetOutput(&lumberjack.Logger{
		Filename:   "logs/error.log",
		MaxSize:    50,
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   true,
		LocalTime:  true,
	})
}
