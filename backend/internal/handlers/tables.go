package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/fnxr21/hw-live-music/backend/internal/models"
	"github.com/fnxr21/hw-live-music/backend/internal/repositories"
	"github.com/labstack/echo/v4"
)

type handlerTable struct {
	TableRepo repositories.Table
}

func HandlerTable(tableRepo repositories.Table) *handlerTable {
	return &handlerTable{tableRepo}
}

// CreateTable handles creating a new table
func (h *handlerTable) CreateTable(c echo.Context) error {
	var table models.RefTable
	if err := c.Bind(&table); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	createdTable, err := h.TableRepo.CreateTable(table)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, createdTable)
}

// ListTables returns all active tables
func (h *handlerTable) ListTables(c echo.Context) error {
	tables, err := h.TableRepo.ListTables()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, tables)
}

// GetTableByID returns a single table by UUID
func (h *handlerTable) GetTableByID(c echo.Context) error {
	id := c.Param("id")
	table, err := h.TableRepo.GetTableByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if table == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Table not found"})
	}
	return c.JSON(http.StatusOK, table)
}

// UpdateTable updates an existing table
func (h *handlerTable) UpdateTable(c echo.Context) error {
	idParam := c.Param("id")
	tableID, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid table ID"})
	}

	var table models.RefTable
	if err := c.Bind(&table); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	table.TableID = tableID
	table.UpdatedAt = time.Now()

	updatedTable, err := h.TableRepo.UpdateTable(table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updatedTable)
}

// DeleteTable performs a soft delete
func (h *handlerTable) DeleteTable(c echo.Context) error {
	id := c.Param("id")
	if err := h.TableRepo.DeleteTable(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Table deleted successfully"})
}
