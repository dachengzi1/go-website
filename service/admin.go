package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-website/model"
	"net/http"
)

func SaveAdmin(c *gin.Context) {
	var admin model.Admin
	err := c.ShouldBindJSON(&admin)
	if err != nil {
		fmt.Println("user err")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "User information is not complete",
		})
		return
	}
	err = admin.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0})
}

func FindAdmin(c *gin.Context) {
	var admin model.Admin
	id := c.Param("id")
	err := admin.FindOne(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"admin": admin})
}

func FindAdmins(c *gin.Context) {
	var admin model.Admin
	admins, err := admin.Find()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"admins": admins})
}

func UpdateAdmin(c *gin.Context) {
	var admin model.Admin
	err := admin.UpdateOne("root45", "1111")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"admin": admin})
}

func UpdateAdmins(c *gin.Context) {
	var admin model.Admin
	err := admin.Update("root45", "222")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"admin": admin})
}

func DeleteAdmin(c *gin.Context) {
	var admin model.Admin
	err := admin.Delete("root45")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
