package pages

import (
	pkg "av/times/pkg"
	"av/times/pkg/auth"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)



func AuthHandler(c *echo.Echo, urlToTemplate pkg.URLToTemplate) {
    group:= c.Group(urlToTemplate.URL)
    group.Use(pkg.RendererMiddleware("views/" + urlToTemplate.Template))

    _get := func(c echo.Context) error {
        data := map[string]interface{}{
            "Title": "SignIn to The Something Post",
        }
        return c.Render(http.StatusOK, "main", data)
    }

    _post := func(c echo.Context) error {
        err := c.Request().ParseForm()
        if err != nil {
            return err
        }
        email := c.FormValue("email")
        data := map[string]interface{}{
            "Title": "Email to The Something Post",
            "email": email,
        }
        log.Printf("email: %v", email)
        time.Sleep(1 * time.Second)
        return c.Render(http.StatusOK, "code", data)
    }


    _googleCallback := func(c echo.Context) error {
        token := c.QueryParam("idToken")
        if token == "" {
            return c.String(http.StatusUnauthorized, "No token provided")
        }
        client, err := auth.App.Auth(context.Background())
        if err != nil {
            log.Printf("Error initializing Firebase Auth: %v", err)
            return c.String(http.StatusInternalServerError, "Error initializing Firebase Auth")
        }
        decodedToken, err := client.VerifyIDToken(context.Background(), token)
        if err != nil {
            log.Printf("Error verifying token: %v", err)
            return c.String(http.StatusUnauthorized, "Invalid Google token")
        }
        userID := decodedToken.UID
        userEmail := decodedToken.Claims["email"].(string)
        jwtToken, err := auth.SignUser(userID, userEmail)
        if err != nil {
            log.Printf("Error jwt generate: %v", err)
            return c.String(http.StatusInternalServerError, "Cannot generate token")
        }
        c.SetCookie(&http.Cookie{
            Name:     auth.AuthCookie,
            Value:    jwtToken,
            Path:     "/",
            HttpOnly: true,
            Secure:   true,
            SameSite: http.SameSiteLaxMode,
            Expires:  time.Now().Add(time.Hour * auth.TokenExpire),
        })
        return c.Redirect(http.StatusFound, "/")
    }

    _signOut := func(c echo.Context) error {
        c.SetCookie(auth.ExpiredCookie())
        return c.Redirect(http.StatusFound, "/")
    }

    group.GET("", _get)
    group.POST("/email", _post)
    group.GET("/google-callback", _googleCallback)
    group.GET("/sign-out", _signOut)
}

