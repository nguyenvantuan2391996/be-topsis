package handler

import (
	"net/http"

	"topsis/handler/constant"
	"topsis/handler/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateUser(c *gin.Context) {
	logrus.Info("Start api create user...")

	// Parse request
	var userRequest *model.UserRequest
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		logrus.Errorf("Parse request create user fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.ParseRequestFail,
		})
		return
	}

	// Validate request
	listErr := h.ValidateRequest(userRequest)
	if len(listErr) > 0 {
		logrus.Errorf("Request is invalid: %v", listErr)
		c.JSON(http.StatusBadRequest, listErr)
		return
	}

	userCreated, err := h.userDomain.CreateUser(c, userRequest.Name)
	if err != nil {
		logrus.Errorf("Create new user fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.CreateNewUserFail,
		})
		return
	}

	logrus.Info("Create new user success")
	c.JSON(http.StatusCreated, &model.UserResponse{
		ID:   userCreated.ID,
		Name: userCreated.Name,
	})
	return
}
