package api

import (
        "encoding/json"
	"github.com/valyala/fasthttp"
        models "../models"
        "fmt"
        uuid "github.com/satori/uuid"
)

func init() {
        models.Users = map[string]models.User{}
        models.Sessions = map[string]models.User{}
}

//CreateSession - create user
func CreateSession(ctx *fasthttp.RequestCtx) {
        var userStruct models.UserRegistration
        err := json.Unmarshal(ctx.PostBody(), &userStruct)
        if err != nil {
                return
        }
        fmt.Println(userStruct)
        login := userStruct.Name
        password := userStruct.Password
        email := userStruct.Email
        models.Users[login] = models.User {
                Name: login,
                Email: email,
                Password: password,
        }
        fmt.Println(models.Users)
        ctx.Write([]byte(login))
        ctx.SetStatusCode(fasthttp.StatusOK)
}

//GetSession - authorization
func GetSession(ctx *fasthttp.RequestCtx) {
        var userStruct models.UserLogin
        err := json.Unmarshal(ctx.PostBody(), &userStruct)
        if err != nil {
                return
        }
        fmt.Println(userStruct)
        login := userStruct.Name
        password := userStruct.Password
        user, ok := models.Users[login]
        if !ok || password != user.Password {
                ctx.SetStatusCode(fasthttp.StatusNotFound)
                return
        }
        sessionID, _ := uuid.NewV4()
        models.Sessions[sessionID.String()] = user
        cookie := fasthttp.Cookie{}
        cookie.SetKey("sessionid")
        cookie.SetValue(sessionID.String())
        cookie.SetPath("/")
        //cookie.SetHTTPOnly(true)
        cookie.SetMaxAge(60*60)
        ctx.Response.Header.SetCookie(&cookie)
        ctx.SetStatusCode(fasthttp.StatusOK)
}

//CheckSession - user authorization status
func CheckSession(ctx *fasthttp.RequestCtx) {
        cookie := ctx.Request.Header.Cookie("sessionid")
        if string(cookie) == "" {
                ctx.SetStatusCode(fasthttp.StatusNotFound)
                return
        }
        ctx.Write([]byte(models.Sessions[string(cookie)].Name))
	ctx.SetStatusCode(fasthttp.StatusOK)
}

//DestroySession - deauthorization
func DestroySession(ctx *fasthttp.RequestCtx) {
        cookieValue := ctx.Request.Header.Cookie("sessionid")
        if string(cookieValue) == "" {
                ctx.SetStatusCode(fasthttp.StatusNotFound)
                return
        }
        cookie := fasthttp.Cookie{}
        cookie.SetKey("sessionid")
        cookie.SetValue("")
        cookie.SetPath("/")
        cookie.SetMaxAge(-1)
        ctx.Response.Header.SetCookie(&cookie)
	delete(models.Sessions, string(cookieValue))
        ctx.SetStatusCode(fasthttp.StatusOK)
}
