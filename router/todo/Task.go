package todo

import (
	"firstApp/conf"
	"firstApp/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	var task models.Task

	c.BindJSON(&task)

	err := conf.DB.First(&task, id).Error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
		return
	}
	conf.DB.Model(&task).Updates(task)
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    task,
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
	res := conf.DB.First(&Task, c.Param("id")).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("username , id")
	}).Find(&Task)
	if res.Error != nil {
		// if error not found
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    Task,
	})

}

func List(c *gin.Context) {
	page, pageStrErr := strconv.Atoi(c.Query("page"))
	if pageStrErr != nil || page < 1 {
		page = 1
	}
	limit, limitStrErr := strconv.Atoi(c.Query("limit"))
	if limitStrErr != nil || limit < 10 {
		limit = 10
	}
	var count int64
	cntErr := conf.DB.Model(&models.Task{}).Count(&count).Error
	if cntErr != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
	}
	offset := (page - 1) * limit
	task := []models.Task{}
	err := conf.DB.Limit(limit).Offset(offset).Find(&task).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"list": map[string]interface{}{
			"count": count,
			"data":  task,
		},
	})

}
