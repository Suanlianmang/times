package pkg

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Templates struct {
    templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    if data == nil {
        data = map[string]interface{}{}
    }
    dataMap, ok := data.(map[string]interface{})
    if !ok {
        return fmt.Errorf("data must be a map")
    }
    userId := c.Get("userID")
    dataMap["IsAuth"] = userId != nil
    return t.templates.ExecuteTemplate(w, name, dataMap)
}

func newTemplates(pattern string) *Templates {
    tmpl := template.Must(template.ParseGlob("views/partial/*.html"))
    return &Templates{
        templates: template.Must(tmpl.ParseGlob(pattern)),
    }
}


func RendererMiddleware(pattern string) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            renderer := newTemplates(pattern)
            c.Echo().Renderer = renderer
            return next(c)
        }
    }
}

