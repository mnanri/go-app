package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	// ローカル環境の場合、http://localhost:8080/ にGETアクセスされるとハンドラーを実行する
	e.GET("/", helloHandler)
	e.GET("/:path", urlHandler)
	e.GET("/query", queryHandler)
	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}

// ハンドラーを定義
func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func urlHandler(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("path"))
}

func queryHandler(c echo.Context) error {
	return c.String(http.StatusOK, c.QueryParam("key"))
}
