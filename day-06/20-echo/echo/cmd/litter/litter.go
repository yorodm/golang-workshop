package main

import (
	"os"
	"sample20/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
)

func main() {

	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	key := os.Getenv("LITTER_SECRET")
	if key == "" {
		e.Logger.Fatal("Cannot start without secret key")
	}
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for signup and login requests
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	// Database connection
	db, err := mgo.Dial("localhost")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Create indices
	if err = db.Copy().DB("twitter").C("users").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	// Initialize handler
	h := &handler.Handler{DB: db, Key: key}

	// Routes
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
	e.POST("/follow/:id", h.Follow)
	e.POST("/posts", h.CreatePost)
	e.GET("/feed", h.FetchPost)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
