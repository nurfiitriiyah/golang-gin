package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

// to get one data with {id}
func (idb *InDB) GetOTS(c *gin.Context) {
	/*Get OTS*/
	var (
		Ots    structs.TbOutstanding
		result gin.H
	)
	areaOts := make(map[string][]interface{})
	dispOts := make(map[string][]interface{})

	rows, err := idb.DB.Table("tb_outstandings").Select("outstanding_area,sum(outstanding_quantity)").Group("outstanding_area").Rows()
	disp, err := idb.DB.Table("tb_outstandings").Select("outstanding_dispatcher,sum(outstanding_quantity)").Group("outstanding_dispatcher").Rows()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
	} else {
		for disp.Next() {
			err := disp.Scan(&Ots.Outstanding_dispatcher, &Ots.Outstanding_quantity)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			} else {
				dispOts[Ots.Outstanding_dispatcher] = append(areaOts[Ots.Outstanding_dispatcher], Ots.Outstanding_quantity)
			}
		}
		for rows.Next() {
			err := rows.Scan(&Ots.Outstanding_area, &Ots.Outstanding_quantity)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			} else {
				areaOts[Ots.Outstanding_area] = append(areaOts[Ots.Outstanding_area], Ots.Outstanding_quantity)
			}
		}
	}
	result = gin.H{
		"area": areaOts,
		"disp": dispOts,
	}

	c.JSON(http.StatusOK, result)

}