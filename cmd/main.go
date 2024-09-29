package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"av/times/db"
	"av/times/pkg"
	"av/times/pkg/pages"
)




func main() {

    var err error
    // Get a greeting message and print it.
    err = godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("PORT is missing")
    }

    db_URL := os.Getenv("DB_URL")
    if port == "" {
        log.Fatal("DB_URL is missing")
    }
    ctx := context.Background()
    conn, err := pgx.Connect(ctx, db_URL)
    if err != nil {
        log.Fatal("Cannot connect to database: ", err)
    }
    defer conn.Close(ctx)
    queries := db.New(conn)
    var user db.User
    user, err = queries.GetUserByEmail(ctx, "test@email.com")
    if err != nil {
        newUser, err := queries.CreateUser(ctx, db.CreateUserParams{
            Email: "test@email.com",
            IsStaff: false,
            Name: pgtype.Text{String: "User 1", Valid: true},
        })
        if err != nil {
            log.Fatal("Cannot create user: ", err)
        }
        user = newUser
    }

	if err != nil {
        log.Fatal("Cannot connect to database: ", err)
	}
    log.Print(user)

    e := echo.New()
    e.Use(middleware.Logger())
    e.Static("/static", "assets")

    // Middleware

    // Main Page
    pages.Index(e, pkg.IndexUrl)
    pages.Post(e, pkg.PostUrl)
    pages.SignIn(e, pkg.SignInUrl)

    e.Logger.Fatal(e.Start(":" + port))
}
