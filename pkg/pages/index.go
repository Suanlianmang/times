package pages

import (
	pkg_middleware "av/times/pkg"
	"net/http"

	"github.com/labstack/echo/v4"
)


func Index(c *echo.Echo, url string, template string) {
    group:= c.Group(url)
    group.Use(pkg_middleware.RendererMiddleware("views/" + template + ".html"))


    _main := func(c echo.Context) error {
        data := map[string]interface{}{
            "Title": "Home Page",
        }
        return c.Render(http.StatusOK, "main", data)
    }

    group.GET("", _main)
}

