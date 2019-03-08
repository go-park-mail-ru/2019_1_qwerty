package api

import (
        //"encoding/json"
	"github.com/valyala/fasthttp"
)

//CheckSession - проверка статуса авторизации пользователя
func CheckSession(ctx *fasthttp.RequestCtx) {
        cookie := ctx.Request.Header.Cookie("sessionid")
        if string(cookie) == "" {
                ctx.SetStatusCode(fasthttp.StatusNotFound)
                return
        }
	ctx.SetStatusCode(fasthttp.StatusOK)
}
