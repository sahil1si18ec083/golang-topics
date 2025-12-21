package routes

import (
	"fmt"
	"net/http"
	"rest-api/models"
	"strconv"

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
	fmt.Println(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userId := c.GetInt64("userId")
	req.UserId = userId

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
func UpdateById(c *gin.Context) {
	id := c.Param("id")
	idval, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	var req models.Event
	event, err := req.GetEventById(id)
	fmt.Println(event)
	userid, ok := c.Get("userId")
	fmt.Println(ok)
	fmt.Println(event.UserId, "a")
	fmt.Println(userid, "b")
	useridval, ok := userid.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid user",
		})
		return
	}

	if event.UserId != useridval {

		c.JSON(http.StatusForbidden, gin.H{
			"message": "Not allowed",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot find id",
		})
		return
	}

	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	var e models.Event
	_, err = e.UpdateById(idval, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server Error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Updated ",
	})

}

func DeleteById(c *gin.Context) {
	id := c.Param("id")
	idval, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	var E models.Event
	event, err := E.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot find id",
		})
		return
	}
	userid, ok := c.Get("userId")
	fmt.Println(ok)
	fmt.Println(event.UserId, "a")
	fmt.Println(userid, "b")
	useridval, ok := userid.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid user",
		})
		return
	}

	if event.UserId != useridval {

		c.JSON(http.StatusForbidden, gin.H{
			"message": "Not allowed",
		})
		return
	}
	err = E.Delete(idval)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot delete",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted Successfully",
	})

}
