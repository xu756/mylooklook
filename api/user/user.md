### 1. "用户登录"

1. route definition

- Url: /userapi/v1/login
- Method: POST
- Request: `userLogin`
- Response: `Res`

2. request definition

```golang
type UserLogin struct {
Username string `json:"username" form:"username"`
Password string `json:"password" form:"password"`
}

```

3. response definition

```golang
type Res struct {
Code int `json:"code"`
Msg string `json:"msg"`
Data interface{} `json:"data"`
}
```

