package fgin

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
)

// Handler a group of router
type Handler struct {
	Prefix      string
	Descriptors []Descriptor
	Middleware  []HandlerFunc
	SubHandler  *Handler
}

// Descriptor is a data structure of api interface
type Descriptor struct {
	Path     string
	Method   string
	Function HandlerFunc
}

type H gin.H

func (h H) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return gin.H(h).MarshalXML(e, start)
}

type Context struct {
	*gin.Context
}

type HandlerFunc func(*Context)
