package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Renn-Amm/Project-Edge/internal/form"
    "github.com/Renn-Amm/Project-Edge/internal/models"
    "github.com/Renn-Amm/Project-Edge/internal/services"
)

type UserHandler struct {
    userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) SignUp(c *gin.Context) {
    var user models.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.userService.SignUp(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func (h *UserHandler) Login(c *gin.Context) {
    var loginRequest form.LoginRequest

    if err := c.ShouldBindJSON(&loginRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := h.userService.Login(&loginRequest)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}