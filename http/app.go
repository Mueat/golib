package http

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	Request  Request
	Response Response
}

func InitApp(c *gin.Context) App {
	req := Request{
		Ctx: c,
	}
	resp := Response{
		Ctx: c,
	}
	return App{Request: req, Response: resp}
}

// 获取body数据
func (a *App) GetBody() []byte {
	return a.Request.GetBody()
}

// 绑定
func (a *App) Bind(v interface{}) error {
	return a.Request.Bind(v)
}

// 获取body中的string
func (a *App) GetBodyStr(k string) string {
	return a.Request.GetBodyStr(k)
}

// 获取body中的int64
func (a *App) GetBodyInt64(k string) int64 {
	return a.Request.GetBodyInt64(k)
}

// 获取body中的bool值
func (a *App) GetBodyBool(k string) bool {
	return a.Request.GetBodyBool(k)
}

// 绑定body中特定的key值
func (a *App) BodyBind(k string, v interface{}) error {
	return a.Request.BodyBind(k, v)
}

//从url的query中获取指定key的内容，如果key不存在，则返回def内容
func (a *App) GetQuery(key string, def string) string {
	return a.Request.GetQuery(key, def)
}

//从url的params中获取指定key内容
func (a *App) GetParam(key string) string {
	return a.Request.GetParam(key)
}

//从form中获取值
func (a *App) GetForm(key, def string) string {
	return a.Request.GetForm(key, def)
}

//获取客户端IP
func (a *App) GetIP() string {
	return a.Request.GetIP()
}

//获取user-agent
func (a *App) GetUserAgent() string {
	return a.Request.GetUserAgent()
}

//设置值
func (a *App) Set(key string, v interface{}) {
	a.Request.Set(key, v)
}

//获取值
func (a *App) Get(key string) (value interface{}, exists bool) {
	return a.Request.Get(key)
}

func (a *App) GetString(key string) string {
	return a.Request.GetString(key)
}

func (a *App) GetStringMap(key string) map[string]interface{} {
	return a.Request.GetStringMap(key)
}

func (a *App) GetStringMapString(key string) map[string]string {
	return a.Request.GetStringMapString(key)
}

func (a *App) GetStringSlice(key string) []string {
	return a.Request.GetStringSlice(key)
}

func (a *App) GetStringMapStringSlice(key string) map[string][]string {
	return a.Request.GetStringMapStringSlice(key)
}

func (a *App) GetBool(key string) bool {
	return a.Request.GetBool(key)
}

func (a *App) GetInt(key string) int {
	return a.Request.GetInt(key)
}

func (a *App) GetInt64(key string) int64 {
	return a.Request.GetInt64(key)
}

func (a *App) GetUint(key string) uint {
	return a.Request.GetUint(key)
}

func (a *App) GetFloat64(key string) float64 {
	return a.Request.GetFloat64(key)
}

// Abort
func (a *App) Abort() {
	a.Response.Abort()
}

func (a *App) AbortWithStatus(code int) {
	a.Response.AbortWithStatus(code)
}

func (a *App) Status(code int) *App {
	a.Response.StatusCode = code
	return a
}

func (a *App) Send(str string) {
	a.Response.Send(str)
}

func (a *App) Json(v interface{}) {
	a.Response.Json(v)
}

func (a *App) Success(v interface{}) {
	a.Response.Success(v)
}

func (a *App) Error(code int, msg string) {
	a.Response.Error(code, msg)
}
