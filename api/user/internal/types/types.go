// Code generated by goctl. DO NOT EDIT.
package types

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
