package sanbox

import (
	"fmt"
	"gin/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

package controllers

import (
"../structs"
"fmt"
"github.com/gin-gonic/gin"
"net/http"
)

// to get one data with {id}
func (idb *InDB) GetOTS(c *gin.Context) {
	/*
		Get OTS
	*/
	var (
		OtsDispatch    structs.TbOutstanding
		OtsArea    []structs.TbOutstanding
		result gin.H
	)

	rows, err := idb.DB.Table("tb_outstandings").Select("*").Group("outstanding_dispatcher").Scan(&OtsDispatch).Rows()
	if err!= nil{
		fmt.Println(err)
	}else{
		fmt.Println(rows)
	}

	idb.DB.Find(&OtsArea).Group("outstanding_area")

	fmt.Println(rows)
	result = gin.H{
		"dispatch": OtsDispatch,
		//"area": OtsArea,
	}

	c.JSON(http.StatusOK, result)
}

