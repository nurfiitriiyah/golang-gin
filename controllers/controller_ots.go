package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

// to get one data with {id}
func (idb *InDB) GetOTS(c *gin.Context) {
	var wg sync.WaitGroup
	/*Get OTS*/
	var (
		Ots    structs.TbOutstanding
		Retail structs.TbRetail
		result gin.H

		labelArea []string
		TotalArea []int

		labelDisp []string
		TotalDisp []int

		labelPack []string
		TotalPack []int

		labelRetl []string
		TotalRetl []int
	)

	rows, err := idb.DB.Table("tb_outstandings").Select("outstanding_area,sum(outstanding_quantity)").Group("outstanding_area").Rows()
	disp, err := idb.DB.Table("tb_outstandings").Select("outstanding_dispatcher,sum(outstanding_quantity)").Group("outstanding_dispatcher").Rows()
	pack, err := idb.DB.Table("tb_outstandings").Select("outstanding_package,sum(outstanding_quantity)").Group("outstanding_package").Rows()
	retl, err := idb.DB.Table("tb_outstandings").Select("retail_label,sum(outstanding_quantity)").Joins("JOIN tb_pairing_retail on outstanding_retail = retail_id").Group("outstanding_retail").Rows()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
	} else {
		wg.Add(4)

		go func() {
			for retl.Next() {
				err := retl.Scan(&Retail.Retail_label, &Ots.Outstanding_quantity)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err)
					c.Abort()
				} else {
					labelRetl = append(labelRetl, Retail.Retail_label)
					TotalRetl = append(TotalRetl, Ots.Outstanding_quantity)
				}
			}
			defer wg.Done()
		}()

		go func() {
			for pack.Next() {
				err := pack.Scan(&Ots.Outstanding_package, &Ots.Outstanding_quantity)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err)
					c.Abort()
				} else {
					labelPack = append(labelPack, Ots.Outstanding_package)
					TotalPack = append(TotalPack, Ots.Outstanding_quantity)
				}
			}
			defer wg.Done()
		}()

		go func() {
			for disp.Next() {
				err := disp.Scan(&Ots.Outstanding_dispatcher, &Ots.Outstanding_quantity)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err)
					c.Abort()
				} else {
					labelDisp = append(labelDisp, Ots.Outstanding_dispatcher)
					TotalDisp = append(TotalDisp, Ots.Outstanding_quantity)
				}
			}
			defer wg.Done()
		}()

		go func() {
			var i = 0
			for rows.Next() {
				err := rows.Scan(&Ots.Outstanding_area, &Ots.Outstanding_quantity)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err)
					c.Abort()
				} else {
					labelArea = append(labelArea, Ots.Outstanding_area)
					TotalArea = append(TotalArea, Ots.Outstanding_quantity)
				}
				i++
			}
			defer wg.Done()
		}()

	}
	wg.Wait()
	result = gin.H{
		"disp": gin.H{
			"label": labelDisp,
			"total": TotalDisp,
		},
		"area": gin.H{
			"label": labelArea,
			"total": TotalArea,
		},
		"pack": gin.H{
			"label": labelPack,
			"total": TotalPack,
		},
		"retail": gin.H{
			"label": labelRetl,
			"total": TotalRetl,
		},
	}

	c.JSON(http.StatusOK, result)
}
