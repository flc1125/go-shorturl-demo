package bootstrap

import (
	"os"

	"shorturl/app/providers/env"
	"shorturl/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Application struct {
	Echo *echo.Echo
}

func NewApplication() *Application {
	app := &Application{}

	app.Echo = echo.New()

	app.Env() // 注意顺序
	app.Routes()
	app.Middlewares()

	return app
}

// 载入 ENV 配置
func (app *Application) Env() {
	env.Env()
}

// 路由配置
func (app *Application) Routes() {
	routes.Web(app.Echo)
}

// 中间件配置
func (app *Application) Middlewares() {
	app.Echo.Use(middleware.Logger())
	app.Echo.Use(middleware.Recover())
}

// 启动服务
func (app *Application) Start() {
	app.Echo.Logger.Fatal(
		app.Echo.Start(
			os.Getenv("HTTP_HOST") + ":" + os.Getenv("HTTP_PORT")))
}
