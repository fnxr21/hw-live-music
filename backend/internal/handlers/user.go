package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/fnxr21/hw-live-music/backend/internal/models"
	"github.com/fnxr21/hw-live-music/backend/internal/repositories"
	"github.com/labstack/echo/v4"
)

type handlerUser struct {
	UserRepo repositories.User
}

func HandlerUser(userRepo repositories.User) *handlerUser {
	return &handlerUser{userRepo}
}

// CreateUser handles creating a new user
func (h *handlerUser) CreateUser(c echo.Context) error {
	var user models.RefUser
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	createdUser, err := h.UserRepo.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, createdUser)
}

// ListUsers returns all users
func (h *handlerUser) ListUsers(c echo.Context) error {
	users, err := h.UserRepo.ListUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}

// GetUserByID returns a single user by UUID
func (h *handlerUser) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user, err := h.UserRepo.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "User not found",
		})
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser updates an existing user
func (h *handlerUser) UpdateUser(c echo.Context) error {
	  idParam := c.Param("id")

    // Parse UUID
    userID, err := uuid.Parse(idParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid user ID",
        })
    }

    var user models.RefUser
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid request payload",
        })
    }
	user.UserID = userID 


	updatedUser, err := h.UserRepo.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser deletes a user by UUID
func (h *handlerUser) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if err := h.UserRepo.DeleteUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User deleted successfully",
	})
}
