package models

type User struct {
        //ID uint64 `json:"id"`
        Name string `json:"name"`
        Email string `json:"email"`
        Password string `json:"password"`
        Score uint64 `json:"score"`
        Avatar string `json:"avatar"`
}

type UserProfile struct {
        Name string `json:"name"`
        Email string `json:"email"`
        Score uint64 `json:"score"`
        Avatar string `json:"avatar"`
}

type UserRegistration struct {
        Name string `json:"nickname"`
        Email string `json:"email"`
        Password string `json:"password"`
}

type UserLogin struct {
        Name string `json:"nickname"`
        Password string `json:"password"`
}

type UserChange struct {
        Email string `json:"email"`
        Password string `json:"password"`
}

var Users map[string]User
