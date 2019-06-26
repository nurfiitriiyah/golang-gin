package controllers

import (
	"gin/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (idb *InDB) GetIOS(c *gin.Context) {

	var (
		stock      structs.TbIostock
		resultDate []time.Time
		result     gin.H

		singleStockIn    []int
		singleStockOut   []int
		singleStockTotal []int
		singleStockQntIn []int
		singleStockQue   []int
		singleStockOtw   []int
		singleStockOsdo  []int
	)

	type ResultStruct struct {
		SingleStockIn    []int
		SingleStockOut   []int
		SingleStockTotal []int
		SingleStockQntIn []int
		SingleStockQue   []int
		SingleStockOtw   []int
		SingleStockOsdo  []int
	}

	var datas ResultStruct

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
			resultDate = append(resultDate, stock.Iostock_Date)
			singleStockIn = append(singleStockIn, stock.Iostock_in)
			singleStockOut = append(singleStockOut, stock.Iostock_out)
			singleStockTotal = append(singleStockTotal, stock.Iostock_stok)
			singleStockQntIn = append(singleStockQntIn, stock.Iostock_gnt_in)
			singleStockQue = append(singleStockQue, stock.Iostock_que)
			singleStockOtw = append(singleStockOtw, stock.Iostock_otw)
			singleStockOsdo = append(singleStockOsdo, stock.Iostock_osdo)
		}
	}

	result = gin.H{
		"label": resultDate,
		"data": gin.H{
			"StockIn":    singleStockIn,
			"StockOut":   singleStockOut,
			"StockTotal": singleStockTotal,
			"StockQntIn": singleStockQntIn,
			"StockQue":   singleStockQue,
			"StockOtw":   singleStockOtw,
			"StockOsdo":  singleStockOsdo,
		},
	}
	c.JSON(http.StatusOK, result)

}
