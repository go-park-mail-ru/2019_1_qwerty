package router

import (
	"../api"

	"github.com/buaazp/fasthttprouter"
        "github.com/valyala/fasthttp"
)

type handlerFunc func(ctx *fasthttp.RequestCtx)

// CORSDecorator - исполнение CORS
func CORSDecorator(CORSFunction handlerFunc) fasthttp.RequestHandler {
        return func(ctx* fasthttp.RequestCtx) {
                CORSFunction(ctx)
                ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:8000")
                ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
                ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
        }
}

// Instance - router
var Instance = fasthttprouter.New()

func init() {
	Instance.GET("/api/score/:index", api.GetNextAfter)
        Instance.GET("/api/user/check", CORSDecorator(api.CheckSession))
        Instance.POST("/api/user/login", CORSDecorator(api.GetSession))
        Instance.POST("/api/user/signup", CORSDecorator(api.CreateSession))
        Instance.GET("/api/user/logout", CORSDecorator(api.DestroySession))
}
