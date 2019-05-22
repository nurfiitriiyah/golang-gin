package controllers

import (
	"fmt"
	"gin/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

/**
Get default data ots, and create pie chart in frontend
*/
func (idb *InDB) GetOTS(c *gin.Context) {
	token, err := parseBearerToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()
	} else {
		decoded := token.Claims
		fmt.Println(decoded)

		var wg sync.WaitGroup

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

			labelLate []int
			TotalLate []int

			labelTransport []string
			TotalTransport []int

			labelTransportOther []string
			TotalTransportOther []int
		)

		wg.Add(6)

		go func() {
			retl, err := idb.DB.Table("tb_outstandings").Select("retail_label,sum(outstanding_quantity)").Joins("JOIN tb_pairing_retail on outstanding_retail = retail_id").Group("outstanding_retail").Rows()
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			}
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
			pack, err := idb.DB.Table("tb_outstandings").Select("outstanding_package,sum(outstanding_quantity)").Group("outstanding_package").Rows()

			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			}
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
			disp, err := idb.DB.Table("tb_outstandings").Select("outstanding_dispatcher,sum(outstanding_quantity)").Group("outstanding_dispatcher").Rows()
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			}
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
			rows, err := idb.DB.Table("tb_outstandings").Select("outstanding_area,sum(outstanding_quantity)").Group("outstanding_area").Rows()
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			}
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

		go func() {

			TotalLateM10 := 0
			TotalLateL1 := 0
			TotalLateB610 := 0
			late, err := idb.DB.Table("tb_outstandings").Select("outstanding_late,sum(outstanding_quantity)").Group("outstanding_late").Rows()
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			}

			for late.Next() {
				err := late.Scan(&Ots.Outstanding_late, &Ots.Outstanding_quantity)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err)
					c.Abort()
				} else {
					if Ots.Outstanding_late < 1 {
						TotalLateL1 = TotalLateL1 + Ots.Outstanding_quantity
					} else {
						if Ots.Outstanding_late > 5 && Ots.Outstanding_late < 10 {
							TotalLateB610 = TotalLateB610 + Ots.Outstanding_quantity
						} else {
							if Ots.Outstanding_late > 9 {
								TotalLateM10 = TotalLateM10 + Ots.Outstanding_quantity
							} else {
								labelLate = append(labelLate, Ots.Outstanding_late)
								TotalLate = append(TotalLate, Ots.Outstanding_quantity)
							}
						}
					}

				}

			}

			labelLate = append(labelLate, 6)
			TotalLate = append(TotalLate, TotalLateL1)

			labelLate = append(labelLate, 7)
			TotalLate = append(TotalLate, TotalLateB610)

			labelLate = append(labelLate, 8)
			TotalLate = append(TotalLate, TotalLateM10)

			defer wg.Done()
		}()

		go func() {
			trans, err := idb.DB.Table("tb_outstandings").Select("outstanding_transporter,sum(outstanding_quantity)").Group("outstanding_transporter").Rows()
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			}
			var i = 0
			for trans.Next() {
				err := trans.Scan(&Ots.Outstanding_transporter, &Ots.Outstanding_quantity)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err)
					c.Abort()
				} else {
					if i < 10 {
						labelTransport = append(labelTransport, Ots.Outstanding_transporter)
						TotalTransport = append(TotalTransport, Ots.Outstanding_quantity)
					} else {
						labelTransportOther = append(labelTransportOther, Ots.Outstanding_transporter)
						TotalTransportOther = append(TotalTransportOther, Ots.Outstanding_quantity)
					}

				}
				i++
			}

			defer wg.Done()
		}()

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
			"late": gin.H{
				"label": labelLate,
				"total": TotalLate,
			}, "transport": gin.H{
				"top10": gin.H{
					"label": labelTransport,
					"total": TotalTransport,
				},
				"other": gin.H{
					"label": labelTransportOther,
					"total": TotalTransportOther,
				},
			},
		}

		c.JSON(http.StatusOK, result)
	}

}

/**
Will show detail when pie chart is clicked
**/
func (idb *InDB) GetDetailOTS(c *gin.Context) {
	var data map[string][]structs.DetailData
	err := c.Bind(&data)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
		c.Abort()
	} else {
		fmt.Println(data)
	}

}
