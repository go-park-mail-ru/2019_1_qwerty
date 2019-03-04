package router

import (
	"../api"
	"github.com/buaazp/fasthttprouter"
)

// Instance - Экспортируемый экземпляр роутера
var Instance = fasthttprouter.New()

func init() {
	Instance.GET("/api/score/:index", api.GetNextAfter)
}
