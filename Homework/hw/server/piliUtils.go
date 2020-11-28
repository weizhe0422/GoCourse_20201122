package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	Model "github.com/weizhe0422/GoCourse_20201122/Homework/hw/PILI/Ｍodel"
	"net/http"
	"strconv"
)

// 取得全部資料
func Get(c *gin.Context) {
	if err := PILIUtil.GetAllRecords(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("failed to get records: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, PILIUtil.AllDramas)
	return
}

// 取得單一筆資料
func GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	data, err := PILIUtil.GetSpecificRecord(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, data)
	return
}

// 新增資料
func Post(c *gin.Context) {
	var r Model.Role
	if err := c.ShouldBind(&r); err!= nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := PILIUtil.InsertRecord([]Model.Role{r}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("failed to insert data: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}

type RoleVM struct {
	ID      uint   `json:"id"`      // Key
	Name    string `json:"name"`    // 角色名稱
	Summary string `json:"summary"` // 介紹
}

// 更新資料, 更新角色名稱與介紹
func Put(c *gin.Context) {
	var (
		id int
		err error
	)

	params := c.Param("id")
	if id, err = strconv.Atoi(params); err!=nil{
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var r Model.Role
	if err := c.ShouldBind(&r); err!= nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	for _, record := range PILIUtil.AllDramas {
		if record.ID == uint(id) {
			record.Summary = r.Summary
			record.Name = r.Name
			record.Skills = r.Skills

			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("ok to update %d data", id),
			})

			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": fmt.Sprintf("can't find %d data to update", id),
	})
}

// 刪除資料
func Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	if ok, Msg := PILIUtil.DeleteOneRecord(uint(id)); !ok{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("failed to delete: %v", Msg),
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"ok to delete",
	})
}
