package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rprajapati0067/golang-interview-test-main/api/domain"
	"github.com/rprajapati0067/golang-interview-test-main/api/domain/dto"
	"github.com/rprajapati0067/golang-interview-test-main/api/errors"
	"github.com/rprajapati0067/golang-interview-test-main/api/services"
)

func CreateSubscribers(c *gin.Context) {
	var request dto.Subscriber
	actualObject := domain.Subscriber{}
	if err := c.BindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr)
		return
	}
	//fmt.Println(request)

	if err := actualObject.SetName(request.Name); err != nil {
		c.JSON(400, err)
		return
	}
	if err := actualObject.SetUrl(request.Url); err != nil {
		c.JSON(400, err)
		return
	}

	result, err := services.SubscriberService.CreateSubscriber(actualObject)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	res := dto.Subscriber{
		Id:   result.Id(),
		Name: result.Name(),
		Url:  result.Url(),
	}
	c.JSON(http.StatusCreated, res)

}
