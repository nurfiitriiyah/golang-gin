package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
	//"net/http"
)

func (idb *InDB) GetProvid(c *gin.Context) {
	var (
		provid []structs.TbProvidEvent
		result gin.H
	)
	idb.DB.Table("tb_provids").Select("provid_code,provid_name").Where("provid_ktgr = 'warehouse'").Order("provid_name asc").Find(&provid)
	result = gin.H{
		"result": provid,
	}
	c.JSON(http.StatusOK, result)

}

func (idb *InDB) GetBagCode(c *gin.Context) {
	var (
		bagCode []structs.TbBagcodeEvent
		result  gin.H
		params  structs.TbProvidEvent
	)
	c.Bind(&params)
	parameter := params.Provid_Code
	idb.DB.Table("tb_iostock as stock").Select("bagcode_code,bagcode_name").Joins("join tb_provids as provid on provid.provid_code = stock.iostok_dispatch join tb_bagcodes as bags on bags.bagcode_code = CONCAT(stock.iostok_packg,stock.iostok_type)").Where("provid.provid_code = ?", parameter).Group("bagcode_code").Find(&bagCode)
	result = gin.H{
		"result": bagCode,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetRole(c *gin.Context) {
	var (
		roles  []structs.TbRoles
		result gin.H
	)
	idb.DB.Table("tb_roles").Select("*").Order("role_id").Find(&roles)
	result = gin.H{
		"result": roles,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetPlan(c *gin.Context) {
	var (
		plans  []structs.TbPlan
		result gin.H
	)
	idb.DB.Table("tb_plans").Select("*").Order("plan_id").Find(&plans)
	result = gin.H{
		"result": plans,
	}
	c.JSON(http.StatusOK, result)
}
