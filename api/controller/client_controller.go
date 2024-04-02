package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waltherx/honda-backend/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClientController struct {
	ClientUsecase domain.ClientUsecase
}

func (tc *ClientController) Create(c *gin.Context) {
	var client domain.Client

	err := c.ShouldBind(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")
	client.ID = primitive.NewObjectID()

	client.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.ClientUsecase.Create(c, &client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Client created successfully",
	})
}

func (u *ClientController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	clients, err := u.ClientUsecase.FetchByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, clients)
}
