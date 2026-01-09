package router

import (
	postgres "github.com/fnxr21/hw-live-music/backend/internal/db"
	"github.com/fnxr21/hw-live-music/backend/internal/handlers"
	repositories "github.com/fnxr21/hw-live-music/backend/internal/repositories"
	"github.com/labstack/echo/v4"
)

func User(e *echo.Group) {
	// Initialize repository
	repo := repositories.Repository(postgres.DB)

	// Initialize handler
	h := handlers.HandlerUser(repo)

	// Define routes
	e.POST("/user/register", h.CreateUser)
	e.GET("/user/:id", h.GetUserByID)
	e.GET("/users", h.ListUsers)
	e.PUT("/user/:id", h.UpdateUser)
	e.DELETE("/user/:id", h.DeleteUser)
}
