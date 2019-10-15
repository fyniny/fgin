package main

import (
	"bytes"
	"fmt"
	"github.com/fyniny/fgin"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func main() {
	h := []fgin.Handler{
		{
			Prefix: "/v1",
			Descriptors: []fgin.Descriptor {
				{
					Path:   "/api",
					Method: "GET",
					Function: func(ctx *fgin.Context) {
						ctx.JSON(http.StatusOK, gin.H{
							"path": "/v1/api",
						})
					},
				},
			},
			SubHandler: &fgin.Handler {
				Prefix: "/project",
				Middleware: []fgin.HandlerFunc{
					func(ctx *fgin.Context) {
						id := ctx.Param("name")
						if id == "id" {
							ctx.AbortWithStatusJSON(http.StatusInternalServerError, fgin.H{
								"msg": "name should not be id",
							})
						}
					},
				},
				Descriptors: []fgin.Descriptor {
					{
						Path: "/:name",
						Method: "GET",
						Function: func(ctx *fgin.Context) {
							ctx.JSON(http.StatusOK, fgin.H{
								"/api/project/": ctx.Param("name"),
							})
						},
					},
				},
			},
		},
		{
			Prefix: "/v2",
			Descriptors: []fgin.Descriptor{
				{
					Path:   "/api",
					Method: "GET",
					Function: func(context *fgin.Context) {
						context.JSON(http.StatusOK, gin.H{
							"path": "/v2/api",
						})
					},
				},
			},
		},
	}

	router := fgin.SetMode(gin.ReleaseMode).New(nil)
	router.Build(h...)

	req := httptest.NewRequest("GET", "http://localhost:9090/v1/project/ids", bytes.NewBufferString(`{"lzf": "l123"}`))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	b, _ := ioutil.ReadAll(w.Body)
	fmt.Println(string(b))
}
