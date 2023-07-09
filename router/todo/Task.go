package todo

import (
	"firstApp/conf"
	"firstApp/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(c *gin.Context) {
	task := models.Task{}

	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	task.UserID = uint(userID.(float64))
	task.Status = "pending"

	err := c.Bind(&task)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := conf.DB.Create(&task)
	if res.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    task,
	})

}

func Update(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	errBind := c.Bind(&task)
	if errBind != nil {
		c.JSON(400, gin.H{
			"message": errBind.Error(),
		})
		return
	}

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
	task := models.Task{}
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	res := conf.DB.Delete(&task)
	if res.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    &task,
	})

}

func GetAll(c *gin.Context) {
	task := []models.Task{}
	res := conf.DB.Find(&task)
	if res.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    task,
	})

}

func GetOne(c *gin.Context) {
	task := models.Task{}
	res := conf.DB.First(&task, c.Param("id")).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username")
	}).Find(&task)

	if res.Error != nil {
		// if error not found
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    task,
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
