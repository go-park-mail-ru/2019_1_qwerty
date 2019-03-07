package api

import (
        "encoding/json"
	"github.com/valyala/fasthttp"
)
//Credintials 1
type Credintials struct {
	Session string `json:"session"`
        Testing string `json:"testing"`
}

//CheckLogin 1
func CheckLogin(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
        ctx.SetContentType("application/json")
        ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:8000")
        ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
        cookie := ctx.Request.Header.Cookie("sessionid")
        json.NewEncoder(ctx).Encode(string(cookie))
}
