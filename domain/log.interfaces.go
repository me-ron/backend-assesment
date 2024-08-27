package domain

import "github.com/gin-gonic/gin"

type LogCreate interface {
	GetAllLogs(gin.Context)
}

type LogUsecase interface {
	GetAllLogs() ([]Log, error)
}

type LogRepository interface {
	GetAllLogs() ([]Log, error)
}