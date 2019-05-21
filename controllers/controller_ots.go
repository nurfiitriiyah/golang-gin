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
		Ots    structs.TbOutstanding
		result gin.H
	)
	areaOts := make(map[string][]interface{})

	rows, err := idb.DB.Table("tb_outstandings").Select("outstanding_area,sum(outstanding_quantity)").Group("outstanding_dispatcher").Rows()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
	} else {
		i := 1
		for rows.Next() {
			err := rows.Scan(&Ots.Outstanding_area, &Ots.Outstanding_quantity)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			} else {
				areaOts[Ots.Outstanding_area] = append(areaOts[Ots.Outstanding_area], Ots.Outstanding_area, Ots.Outstanding_quantity)
				fmt.Println(Ots)
			}
			i++
		}
		fmt.Println(areaOts)

	}
	result = gin.H{
		"area": areaOts,
	}

	c.JSON(http.StatusOK, result)

}
