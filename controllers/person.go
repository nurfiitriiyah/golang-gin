package controllers

import (
	"../structs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

// to get one data with {id}
func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// get all data in person
func (idb *InDB) GetPersons(c *gin.Context) {
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

		labelTimeUpdate time.Time
	)

	var wg sync.WaitGroup

	wg.Add(6)
	/**
	  Retail
	  **/
	go func() {
		retl, err := idb.DB.Table("tb_outstandings as ots").Select("floor(((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000)), ret.retail_label ").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id").Group("retail_id").Rows()
		if err != nil {
			fmt.Println("----------------ERROR ret------------------")
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		for retl.Next() {
			err := retl.Scan(&Ots.Outstanding_quantity, &Retail.Retail_label)
			if err != nil {
				fmt.Println("----------------ERROR ret------------------")
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			} else {
				labelRetl = append(labelRetl, Retail.Retail_label)
				TotalRetl = append(TotalRetl, Ots.Outstanding_quantity)
			}
		}
		defer wg.Done()
	}()
	/**
	  Pack
	  **/
	go func() {
		pack, err := idb.DB.Table("tb_outstandings as ots").Select("floor(((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000)), bag.bagcode_name ").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id").Group("bagcode_code").Rows()

		if err != nil {
			fmt.Println("----------------ERROR pack------------------")
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		for pack.Next() {
			err := pack.Scan(&Ots.Outstanding_quantity, &Ots.Outstanding_package)
			if err != nil {
				fmt.Println("----------------ERROR pack------------------")
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			} else {
				labelPack = append(labelPack, Ots.Outstanding_package)
				TotalPack = append(TotalPack, Ots.Outstanding_quantity)
			}
		}
		defer wg.Done()
	}()
	/**
	  Disp
	  **/
	go func() {
		disp, err := idb.DB.Table("tb_outstandings as ots").Select("outstanding_updt,floor(((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000)),  provid_ktgr ").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id").Group("provid_ktgr").Rows()
		if err != nil {
			fmt.Println("----------------ERROR dispatch------------------")
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		for disp.Next() {
			err := disp.Scan(&Ots.Outstanding_update, &Ots.Outstanding_quantity, &Ots.Outstanding_dispatcher)
			if err != nil {
				fmt.Println("----------------ERROR dispatch------------------")
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			} else {
				labelTimeUpdate = Ots.Outstanding_update
				labelDisp = append(labelDisp, Ots.Outstanding_dispatcher)
				TotalDisp = append(TotalDisp, Ots.Outstanding_quantity)
			}
		}
		defer wg.Done()
	}()
	/**
	  Area
	  **/
	go func() {
		rows, err := idb.DB.Table("tb_outstandings as ots").Select(" floor(((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000)), area.tp_area_alias2 ").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id").Group("area.tp_area_alias2 ").Rows()
		if err != nil {
			fmt.Println("----------------ERROR area------------------")
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		var i = 0
		for rows.Next() {
			err := rows.Scan(&Ots.Outstanding_quantity, &Ots.Outstanding_area)
			if err != nil {
				fmt.Println("----------------ERROR area------------------")
				fmt.Println(err)
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
	/**
	  Late
	  **/
	go func() {

		TotalLateM10 := 0
		TotalLateL1 := 0
		TotalLateB610 := 0
		late, err := idb.DB.Table("tb_outstandings as ots").Select("floor(((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000)), ots.outstanding_late").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id").Group("ots.outstanding_late").Rows()
		if err != nil {
			fmt.Println("----------------ERROR late------------------")
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}

		for late.Next() {
			err := late.Scan(&Ots.Outstanding_quantity, &Ots.Outstanding_late)
			if err != nil {
				fmt.Println("----------------ERROR late------------------")
				fmt.Println(err)
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
	/**
	  Trans
	  **/
	go func() {
		trans, err := idb.DB.Table("tb_outstandings as ots").Select("floor(((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000)) as total, trs.transporter_name ").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id").Group(" transporter_name").Order("total desc").Rows()
		if err != nil {
			fmt.Println("----------------ERROR trans------------------")
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		var i = 0
		for trans.Next() {
			err := trans.Scan(&Ots.Outstanding_quantity, &Ots.Outstanding_transporter)
			if err != nil {
				fmt.Println("----------------ERROR trans------------------")
				fmt.Println(err)
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
		"lastUpdate": labelTimeUpdate,
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

// create new data to the database
func (idb *InDB) CreatePerson(c *gin.Context) {
	token, err := parseBearerToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()
	} else {
		decoded := token.Claims
		fmt.Println(decoded)
		var (
			person structs.Person
			result gin.H
		)
		first_name := c.PostForm("first_name")
		last_name := c.PostForm("last_name")
		person.First_Name = first_name
		person.Last_Name = last_name
		idb.DB.Create(&person)
		result = gin.H{
			"result": person,
		}
		fmt.Println(result)
		c.JSON(http.StatusOK, result)
	}
}

// update data with {id} as query
func (idb *InDB) UpdatePerson(c *gin.Context) {
	id := c.Query("id")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	var (
		person    structs.Person
		newPerson structs.Person
		result    gin.H
	)

	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	newPerson.First_Name = first_name
	newPerson.Last_Name = last_name
	err = idb.DB.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

// delete data with {id}
func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
