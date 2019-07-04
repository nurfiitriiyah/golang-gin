package controllers

import (
	//"fmt"
	"../structs"
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
		"label":            resultDate,
		"singleStockIn":    singleStockIn,
		"singleStockOut":   singleStockOut,
		"singleStockTotal": singleStockTotal,
		"singleStockQntIn": singleStockQntIn,
		"singleStockQue":   singleStockQue,
		"singleStockOtw":   singleStockOtw,
		"singleStockOsdo":  singleStockOsdo,
	}
	c.JSON(http.StatusOK, result)

}
func (idb *InDB) GetDetailIOS(c *gin.Context) {

	var createParams structs.CreateParams
	errs := c.BindJSON(&createParams)
	if errs != nil {
		c.JSON(http.StatusUnauthorized, errs.Error())
		c.Abort()
	}

	var (
		stock      structs.TbIostock
		resultDate []time.Time
		result     gin.H

		secondCharBagCode string
		firstCharBagCode  string

		findFirstProvid  string
		singleStockIn    []int
		singleStockOut   []int
		singleStockTotal []int
		singleStockQntIn []int
		singleStockQue   []int
		singleStockOtw   []int
		singleStockOsdo  []int
	)

	nums := createParams.Data
	if nums[2] != "" {
		firstCharBagCode = "%" + nums[2][:1] + "%"
		secondCharBagCode = "%" + trimFirstRune(nums[2]) + "%"
	} else {
		firstCharBagCode = "%%"
		secondCharBagCode = "%%"
	}

	if nums[3] != "" {
		findFirstProvid = "%" + nums[3] + "%"
	} else {
		findFirstProvid = "%%"
	}

	queryStock, err := idb.DB.Table("tb_iostock").Select("iostok_date as stock_date,sum(iostok_in) as stock_in,sum(iostok_out) as stock_out,sum(iostok_stok) as stock_total,sum(iostok_gnt_in) as stock_change,sum(iostok_que) as stock_que,sum(iostok_otw) as stock_otw,sum(iostok_osdo) as stock_osdo").Group("iostok_date").Where("iostok_date BETWEEN ? AND ? AND iostok_type like ? AND iostok_packg like ? AND iostok_dispatch like ?", nums[0], nums[1], firstCharBagCode, secondCharBagCode, findFirstProvid).Rows()

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
		"label":            resultDate,
		"singleStockIn":    singleStockIn,
		"singleStockOut":   singleStockOut,
		"singleStockTotal": singleStockTotal,
		"singleStockQntIn": singleStockQntIn,
		"singleStockQue":   singleStockQue,
		"singleStockOtw":   singleStockOtw,
		"singleStockOsdo":  singleStockOsdo,
	}
	c.JSON(http.StatusOK, result)
}
