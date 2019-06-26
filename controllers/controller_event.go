package controllers

import (
	"gin/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	//"net/http"
)

func (idb *InDB) GetProvid(c *gin.Context) {
	var (
		provid []structs.TbProvidEvent
		result gin.H
	)
	idb.DB.Table("tb_provids").Select("provid_code,provid_name").Find(&provid)
	result = gin.H{
		"result": provid,
	}
	c.JSON(http.StatusOK, result)

}

func (idb *InDB) GetBagCode(c *gin.Context) {
	var (
		bagCode []structs.TbBagcodeEvent
		result  gin.H
	)
	idb.DB.Table("tb_bagcodes").Select("bagcode_code,bagcode_name").Find(&bagCode)
	result = gin.H{
		"result": bagCode,
	}
	c.JSON(http.StatusOK, result)

}
