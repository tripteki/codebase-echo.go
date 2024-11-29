package main

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func main () {

    app := echo.New ()

    app.Logger.Debug (
        app.Start (":1323")
    )
}
