package router

import (
	postgres "github.com/fnxr21/hw-live-music/backend/internal/db"
	"github.com/fnxr21/hw-live-music/backend/internal/handlers"
	repositories "github.com/fnxr21/hw-live-music/backend/internal/repositories"
	"github.com/labstack/echo/v4"
)

func Table(e *echo.Group) {
	// Initialize repository
	repo := repositories.Repository(postgres.DB)

	// Initialize handler
	h := handlers.HandlerTable(repo)

	// Define routes
	e.POST("/table", h.CreateTable)         // Create a new table
	e.GET("/table/:id", h.GetTableByID)    // Get table by UUID
	e.GET("/tables", h.ListTables)         // List all tables
	e.PUT("/table/:id", h.UpdateTable)     // Update table
	e.DELETE("/table/:id", h.DeleteTable)  // Soft delete table
}
