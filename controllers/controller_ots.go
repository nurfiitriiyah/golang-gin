package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

// to get one data with {id}
func (idb *InDB) GetOTS(c *gin.Context) {
	/*

	 */
	var (
		Ots    []structs.TbOutstanding
		result gin.H
	)

	idb.DB.Find(&Ots)

	if len(Ots) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": Ots,
			"count":  len(Ots),
		}
	}

	c.JSON(http.StatusOK, result)
}
