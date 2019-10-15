package fgin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router interface {
	gin.IRouter
	http.Handler
	Build(handler ...Handler) *EngineRouter
}

func New(engine *gin.Engine) Router {
	if engine == nil {
		engine = gin.New()
	}
	return &EngineRouter{Engine: engine}
}

func SetMode(mode string) *struct{New func(*gin.Engine)Router} {
	gin.SetMode(mode)
	return &struct{
		New func(*gin.Engine)Router
	}{
		New: New,
	}
}
