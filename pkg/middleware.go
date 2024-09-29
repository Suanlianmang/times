
package pkg

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Templates struct {
    templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
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


