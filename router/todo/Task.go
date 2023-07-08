package todo

import (
	"firstApp/conf"
	"firstApp/models"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	Task := models.Task{}
	err := c.BindJSON(&Task)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	res := conf.DB.Create(&Task)
	if res.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    Task,
	})

}

func Update(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Title       string
		Description string
		Status      string
		UserID      uint
	}
	c.Bind(&body)

	var Task models.Task
	err := conf.DB.First(&Task, id).Error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
		return
	}
	conf.DB.Model(&Task).Updates(body)
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    Task,
	})

}

func Delete(c *gin.Context) {
	Task := models.Task{}
	err := c.BindJSON(&Task)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	res := conf.DB.Delete(&Task)
	if res.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    &Task,
	})

}

func GetAll(c *gin.Context) {
	Task := []models.Task{}
	res := conf.DB.Find(&Task)
	if res.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    Task,
	})

}

func GetOne(c *gin.Context) {
	Task := models.Task{}
	res := conf.DB.First(&Task, c.Param("id"))
	if res.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    Task,
	})

}
