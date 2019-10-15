# handler 说明文档

```go
package fgin
type Handler struct {
	Prefix      string
	Descriptors []Descriptor
	Middleware  []gin.HandlerFunc
	SubHandler  *Handler
}

// Descriptor is a data structure of api interface
type Descriptor struct {
	Path     string
	Method   string
	Function gin.HandlerFunc
}
```

## Handler

`Handler` 是一个router组，相当于`gin.RouterGroup`

`Prefix` 为一组路由的前缀

`Middleware` 声明的中间件

`SubHandler` 子`router`，相当于`router.router`，支持无限嵌套。

## Descriptor

一个`Descriptor`实例就是一个api对象
