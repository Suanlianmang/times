package pages

import (
	pkg "av/times/pkg"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/russross/blackfriday/v2"
)


func Post(c *echo.Echo, urlToTemplate pkg.URLToTemplate) {
    group:= c.Group(urlToTemplate.URL)
    group.Use(pkg.RendererMiddleware("views/" + urlToTemplate.Template))
    _main := func(c echo.Context) error {

        mdContent, err := os.ReadFile("views/test.md")
        if err != nil {
            log.Fatalf("Failed to read file: %v", err)
        }

        htmlContent := blackfriday.Run(
            mdContent,
            blackfriday.WithExtensions(blackfriday.CommonExtensions),
            )
        data := map[string]interface{}{
            "Title": "Post Page",
            "htmlContent": template.HTML(htmlContent),
        }
        return c.Render(http.StatusOK, "main", data)
    }

    group.GET("", _main)
}

