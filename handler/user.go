package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"startup-funding/helper"
	"startup-funding/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)
		errorMassage := gin.H{"errors": errors}
		response := helper.APIResponse("Register Account failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register Account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentokentoken")

	response := helper.APIResponse("Account has been Registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
