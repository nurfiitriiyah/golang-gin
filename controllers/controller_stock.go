package controllers

import (
	"gin/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (idb *InDB) GetIOS(c *gin.Context) {
	type ArrayData struct {
		DateStock  time.Time
		TotalStock [7]int
	}

	type Datas struct {
		Result []ArrayData
	}

	var (
		stock       structs.TbIostock
		singleStock [7]int
	)
	var data Datas

	queryStock, err := idb.DB.Table("tb_iostock").Select("iostok_date as stock_date,sum(iostok_in) as stock_in,sum(iostok_out) as stock_out,sum(iostok_stok) as stock_total,sum(iostok_gnt_in) as stock_change,sum(iostok_que) as stock_que,sum(iostok_otw) as stock_otw,sum(iostok_osdo) as stock_osdo").Group("iostok_date").Rows()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
	}
	for queryStock.Next() {
		err := queryStock.Scan(&stock.Iostock_Date, &stock.Iostock_in, &stock.Iostock_out, &stock.Iostock_stok, &stock.Iostock_gnt_in, &stock.Iostock_que, &stock.Iostock_otw, &stock.Iostock_osdo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		} else {
			singleStock[0] = stock.Iostock_in
			singleStock[1] = stock.Iostock_out
			singleStock[2] = stock.Iostock_stok
			singleStock[3] = stock.Iostock_gnt_in
			singleStock[4] = stock.Iostock_que
			singleStock[5] = stock.Iostock_otw
			singleStock[6] = stock.Iostock_osdo
			data.Result = append(data.Result, ArrayData{
				DateStock:  stock.Iostock_Date,
				TotalStock: singleStock,
			})
		}
	}
	c.JSON(http.StatusOK, data)

}
