package routes

import (
	"fmt"
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {

	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"result": events,
	})
}
func CreateEvent(c *gin.Context) {
	var req models.Event
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("rrrrrrrrrrrrrrrrr")
	err = req.Save()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event",
		})
	}
	c.JSON(http.StatusCreated, req)

}
func GetEventById(c *gin.Context) {
	var req models.Event
	id := c.Param("id")
	resp, err := req.GetEventById(id)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": resp,
	})

}
