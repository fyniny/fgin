package fgin

import "github.com/gin-gonic/gin"

type EngineRouter struct {
	*gin.Engine
}

// Build build router
func (router *EngineRouter) Build(handler ...Handler) *EngineRouter {
	if handler == nil || len(handler) <= 0 {
		return nil
	}

	for _, handler := range handler {
		var group *gin.RouterGroup = &router.RouterGroup
		register(group, &handler)
	}

	return router
}

func register(group *gin.RouterGroup, handler *Handler) {
	if handler == nil {
		return
	}

	if handler.Prefix != "" {
		group = group.Group(handler.Prefix)
	}

	middleware := make([]gin.HandlerFunc, 0, len(handler.Middleware))

	for _, mid := range handler.Middleware {
		f := func(ctx *gin.Context) {
			Ctx := Context{Context: ctx}
			mid(&Ctx)
		}
		middleware = append(middleware, f)
	}

	group.Use(middleware...)
	for _, descriptor := range handler.Descriptors {
		group.Handle(descriptor.Method, descriptor.Path, func(context *gin.Context) {
			descriptor.Function(&Context{Context: context})
		})
	}

	register(group, handler.SubHandler)
}
