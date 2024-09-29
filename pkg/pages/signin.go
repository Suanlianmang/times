
package pages

import (
	pkg "av/times/pkg"
	"net/http"

    "time"
	"github.com/labstack/echo/v4"
)


func SignIn(c *echo.Echo, urlToTemplate pkg.URLToTemplate) {
    group:= c.Group(urlToTemplate.URL)
    group.Use(pkg.RendererMiddleware("views/" + urlToTemplate.Template))
    _main := func(c echo.Context) error {
        data := map[string]interface{}{
            "Title": "SignIn to The Something Post",
        }
        return c.Render(http.StatusOK, "main", data)
    }
    _email := func(c echo.Context) error {
        err := c.Request().ParseForm()
        if err != nil {
            return err
        }
        email := c.FormValue("email")

        data := map[string]interface{}{
            "Title": "Email to The Something Post",
            "email": email,
        }

        time.Sleep(1 * time.Second)
        return c.Render(http.StatusOK, "code", data)
    }

    group.GET("", _main)
    group.POST("/email", _email)
}

