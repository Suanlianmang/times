package pages

import (
	pkg "av/times/pkg"
	"net/http"

	"github.com/labstack/echo/v4"
)


func Index(c *echo.Echo, urlToTemplate pkg.URLToTemplate) {
    group:= c.Group(urlToTemplate.URL)
    group.Use(pkg.RendererMiddleware("views/" + urlToTemplate.Template))
    _main := func(c echo.Context) error {
        data := map[string]interface{}{
            "Title": "Home Page",
            "PageTitle": "This is a page title",
            "Intro": `
                Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                Lorem Ipsum has been the industry's standard dummy text ever since
                the 1500s, when an unknown printer took a galley of type and
                scrambled it to make a type specimen book. It has survived not only
                five centuries, but also the leap into electronic typesetting,
                remaining essentially unchanged. It was popularised in the 1960s
                with the release of Letraset sheets containing Lorem Ipsum
                passages, and more recently with desktop publishing software like
                Aldus PageMaker including versions of Lorem Ipsum.
            `,
        }
        return c.Render(http.StatusOK, "main", data)
    }
    group.GET("", _main)
}

