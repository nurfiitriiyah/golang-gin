package controllers

import (
	"../structs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"sync"
	"time"
)

/**
Get default data ots, and create pie chart in frontend
*/
func (idb *InDB) GetOTS(c *gin.Context) {
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

/**
Will show detail when pie chart is clicked
**/

func (idb *InDB) GetDetailOTS(c *gin.Context) {
	var createParams structs.CreateParams

	var (
		Ots    structs.TbOutstanding
		result gin.H
		Retail structs.TbRetail

		labelDisp []string
		TotalDisp []int

		labelRitl []string
		TotalRitl []int

		labelPack []string
		TotalPack []int

		labelArea []string
		TotalArea []int

		labelLate []int
		TotalLate []int

		labelTransport []string
		TotalTransport []int

		labelTransportOther []string
		TotalTransportOther []int

		prepFindTransport string
		labelTimeUpdate   time.Time
	)

	re := regexp.MustCompile("[0-9]+")
	errs := c.BindJSON(&createParams)
	if errs != nil {
		c.JSON(http.StatusUnauthorized, errs.Error())
		c.Abort()
	}
	nums := createParams.Data
	disp := idb.DB.Table("tb_outstandings as ots").Select("outstanding_updt,floor(((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000)),  provid_ktgr ").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id").Group("provid_ktgr")
	retl := idb.DB.Table("tb_outstandings as ots").Select("floor(((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000)), ret.retail_label ").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id").Group("retail_id")
	area := idb.DB.Table("tb_outstandings as ots").Select(" floor(((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000)), area.tp_area_alias2 ").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id").Group("area.tp_area_alias2 ")
	late := idb.DB.Table("tb_outstandings as ots").Select("floor(((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000)), ots.outstanding_late").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id").Group("ots.outstanding_late")
	trans := idb.DB.Table("tb_outstandings as ots").Select("floor(((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000)) as total, trs.transporter_name ").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id").Group(" transporter_name").Order("total desc")
	pack := idb.DB.Table("tb_outstandings as ots").Select("floor(((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000)), bag.bagcode_name ").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id").Group("bagcode_code")

	for _, num := range nums {
		subStrn := string([]rune((re.FindAllString(num, -1))[0])[0:1])
		value := trimFirstRune(num)

		switch subStrn {
		case "1":
			disp = disp.Select("outstanding_updt,floor(((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000)),  provid_name").Group("provid_name").Where("provid_ktgr = ?", value)
			retl = retl.Where("provid_ktgr = ?", value)
			area = area.Where("provid_ktgr = ?", value)
			late = late.Where("provid_ktgr = ?", value)
			trans = trans.Where("provid_ktgr = ?", value)
			pack = pack.Where("provid_ktgr = ?", value)
			break

		case "2":
			disp = disp.Where("tp_area_alias2 = ?", value)
			retl = retl.Where("tp_area_alias2 = ?", value)
			area = area.Where("tp_area_alias2 = ?", value)
			late = late.Where("tp_area_alias2 = ?", value)
			trans = trans.Where("tp_area_alias2 = ?", value)
			pack = pack.Where("tp_area_alias2 = ?", value)
			break

		case "3":
			subStrnDel := string(num[1:2])
			switch subStrnDel {
			case "6":
				disp = disp.Where("outstanding_late < ?", 1)
				retl = retl.Where("outstanding_late < ?", 1)
				area = area.Where("outstanding_late < ?", 1)
				late = late.Where("outstanding_late < ?", 1)
				trans = trans.Where("outstanding_late < ?", 1)
				pack = pack.Where("outstanding_late < ?", 1)
			case "7":
				disp = disp.Where("outstanding_late BETWEEN  ? AND ?", 6, 10)
				retl = retl.Where("outstanding_late BETWEEN  ? AND ?", 6, 10)
				area = area.Where("outstanding_late BETWEEN  ? AND ?", 6, 10)
				late = late.Where("outstanding_late BETWEEN  ? AND ?", 6, 10)
				trans = trans.Where("outstanding_late BETWEEN  ? AND ?", 6, 10)
				pack = pack.Where("outstanding_late BETWEEN  ? AND ?", 6, 10)
			case "8":
				disp = disp.Where("outstanding_late > ?", 10)
				retl = retl.Where("outstanding_late > ?", 10)
				area = area.Where("outstanding_late > ?", 10)
				late = late.Where("outstanding_late > ?", 10)
				trans = trans.Where("outstanding_late > ?", 10)
				pack = pack.Where("outstanding_late > ?", 10)
			default:
				disp = disp.Where("outstanding_late = ?", value)
				retl = retl.Where("outstanding_late = ?", value)
				area = area.Where("outstanding_late = ?", value)
				late = late.Where("outstanding_late = ?", value)
				trans = trans.Where("outstanding_late = ?", value)
				pack = pack.Where("outstanding_late = ?", value)
			}
			break

		case "4":
			prepFindTransport = value + "%"
			disp = disp.Where("transporter_name LIKE ?", prepFindTransport)
			retl = retl.Where("transporter_name LIKE ?", prepFindTransport)
			area = area.Where("transporter_name LIKE ?", prepFindTransport)
			late = late.Where("transporter_name LIKE ?", prepFindTransport)
			trans = trans.Where("transporter_name LIKE ?", prepFindTransport)
			pack = pack.Where("transporter_name LIKE ?", prepFindTransport)
			break

		case "5":

			disp = disp.Where("bagcode_name = ?", value)
			retl = retl.Where("bagcode_name = ?", value)
			area = area.Where("bagcode_name = ?", value)
			late = late.Where("bagcode_name = ?", value)
			trans = trans.Where("bagcode_name = ?", value)
			pack = pack.Where("bagcode_name = ?", value)
			break

		case "6":

			disp = disp.Where("retail_label = ?", value)
			retl = retl.Where("retail_label = ?", value)
			area = area.Where("retail_label = ?", value)
			late = late.Where("retail_label = ?", value)
			trans = trans.Where("retail_label = ?", value)
			pack = pack.Where("retail_label = ?", value)
			break
		case "7":

			disp = disp.Where("provid_name = ?", value)
			retl = retl.Where("provid_name = ?", value)
			area = area.Where("provid_name = ?", value)
			late = late.Where("provid_name = ?", value)
			trans = trans.Where("provid_name = ?", value)
			pack = pack.Where("provid_name = ?", value)
			break
		case "8":
			subStrnDelete := value[:1]
			valueDel := trimFirstRune(value)

			switch subStrnDelete {
			case "1":
				disp = disp.Where("provid_ktgr != ?", valueDel)
				retl = retl.Where("provid_ktgr != ?", valueDel)
				area = area.Where("provid_ktgr != ?", valueDel)
				late = late.Where("provid_ktgr != ?", valueDel)
				trans = trans.Where("provid_ktgr != ?", valueDel)
				pack = pack.Where("provid_ktgr != ?", valueDel)
			case "2":
				disp = disp.Where("tp_area_alias2 != ?", valueDel)
				retl = retl.Where("tp_area_alias2 != ?", valueDel)
				area = area.Where("tp_area_alias2 != ?", valueDel)
				late = late.Where("tp_area_alias2 != ?", valueDel)
				trans = trans.Where("tp_area_alias2 != ?", valueDel)
				pack = pack.Where("tp_area_alias2 != ?", valueDel)
			case "3":
				subStrnDelLate := valueDel[:1]
				switch subStrnDelLate {
				case "6":
					disp = disp.Where("outstanding_late > ?", 0)
					retl = retl.Where("outstanding_late > ?", 0)
					area = area.Where("outstanding_late > ?", 0)
					late = late.Where("outstanding_late > ?", 0)
					trans = trans.Where("outstanding_late > ?", 0)
					pack = pack.Where("outstanding_late > ?", 0)
				case "7":
					disp = disp.Where("outstanding_late NOT BETWEEN  ? AND ?", 6, 10)
					retl = retl.Where("outstanding_late NOT BETWEEN  ? AND ?", 6, 10)
					area = area.Where("outstanding_late NOT  BETWEEN  ? AND ?", 6, 10)
					late = late.Where("outstanding_late NOT BETWEEN  ? AND ?", 6, 10)
					trans = trans.Where("outstanding_late NOT BETWEEN  ? AND ?", 6, 10)
					pack = pack.Where("outstanding_late NOT BETWEEN  ? AND ?", 6, 10)
				case "8":
					disp = disp.Where("outstanding_late < ?", 10)
					retl = retl.Where("outstanding_late < ?", 10)
					area = area.Where("outstanding_late < ?", 10)
					late = late.Where("outstanding_late < ?", 10)
					trans = trans.Where("outstanding_late < ?", 10)
					pack = pack.Where("outstanding_late < ?", 10)
				default:
					disp = disp.Where("outstanding_late = ?", valueDel)
					retl = retl.Where("outstanding_late = ?", valueDel)
					area = area.Where("outstanding_late = ?", valueDel)
					late = late.Where("outstanding_late = ?", valueDel)
					trans = trans.Where("outstanding_late = ?", valueDel)
					pack = pack.Where("outstanding_late = ?", valueDel)
				}

			case "4":
				disp = disp.Where("transporter_name != ?", valueDel)
				retl = retl.Where("transporter_name != ?", valueDel)
				area = area.Where("transporter_name != ?", valueDel)
				late = late.Where("transporter_name != ?", valueDel)
				trans = trans.Where("transporter_name != ?", valueDel)
				pack = pack.Where("transporter_name != ?", valueDel)
			case "5":

				disp = disp.Where("bagcode_name != ?", valueDel)
				retl = retl.Where("bagcode_name != ?", valueDel)
				area = area.Where("bagcode_name != ?", valueDel)
				late = late.Where("bagcode_name != ?", valueDel)
				trans = trans.Where("bagcode_name != ?", valueDel)
				pack = pack.Where("bagcode_name != ?", valueDel)
			case "6":

				disp = disp.Where("retail_label != ?", valueDel)
				retl = retl.Where("retail_label != ?", valueDel)
				area = area.Where("retail_label != ?", valueDel)
				late = late.Where("retail_label != ?", valueDel)
				trans = trans.Where("retail_label != ?", valueDel)
				pack = pack.Where("retail_label != ?", valueDel)
			}
		}
	}

	var wg sync.WaitGroup

	wg.Add(6)
	go func() {
		resDisp, err := disp.Rows()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		var i = 0
		for resDisp.Next() {
			err := resDisp.Scan(&Ots.Outstanding_update, &Ots.Outstanding_quantity, &Ots.Outstanding_location)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			} else {
				labelTimeUpdate = Ots.Outstanding_update
				labelDisp = append(labelDisp, Ots.Outstanding_location)
				TotalDisp = append(TotalDisp, Ots.Outstanding_quantity)
			}
			i++
		}
		defer wg.Done()
	}()
	/**
	  Retail
	  **/
	go func() {
		resRet, err := retl.Rows()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		var i = 0
		for resRet.Next() {
			err := resRet.Scan(&Ots.Outstanding_quantity, &Retail.Retail_label)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			} else {
				labelRitl = append(labelRitl, Retail.Retail_label)
				TotalRitl = append(TotalRitl, Ots.Outstanding_quantity)
			}
			i++
		}
		defer wg.Done()
	}()
	/**
	  Pack
	  **/
	go func() {
		resPack, err := pack.Rows()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		var i = 0
		for resPack.Next() {
			err := resPack.Scan(&Ots.Outstanding_quantity, &Ots.Outstanding_package)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			} else {
				labelPack = append(labelPack, Ots.Outstanding_package)
				TotalPack = append(TotalPack, Ots.Outstanding_quantity)
			}
			i++
		}
		defer wg.Done()
	}()
	/**
	  Area
	  **/
	go func() {
		resArea, err := area.Rows()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		var i = 0
		for resArea.Next() {
			err := resArea.Scan(&Ots.Outstanding_quantity, &Ots.Outstanding_area)
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
	/**
	  Late
	  **/
	go func() {

		TotalLateM10 := 0
		TotalLateL1 := 0
		TotalLateB610 := 0
		resLate, err := late.Rows()

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		for resLate.Next() {
			err := resLate.Scan(&Ots.Outstanding_quantity, &Ots.Outstanding_late)
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

		if TotalLateL1 > 0 {
			labelLate = append(labelLate, 6)
			TotalLate = append(TotalLate, TotalLateL1)
		}
		if TotalLateB610 > 0 {
			labelLate = append(labelLate, 7)
			TotalLate = append(TotalLate, TotalLateB610)
		}
		if TotalLateM10 > 0 {
			labelLate = append(labelLate, 8)
			TotalLate = append(TotalLate, TotalLateM10)
		}

		defer wg.Done()
	}()
	/**
	  Trans
	  **/
	go func() {
		resTrans, err := trans.Rows()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		var i = 0
		for resTrans.Next() {
			err := resTrans.Scan(&Ots.Outstanding_quantity, &Ots.Outstanding_transporter)
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
		"lastUpdate": labelTimeUpdate,

		"disp": gin.H{
			"label": labelDisp,
			"total": TotalDisp,
		},
		"retail": gin.H{
			"label": labelRitl,
			"total": TotalRitl,
		},
		"pack": gin.H{
			"label": labelPack,
			"total": TotalPack,
		},
		"area": gin.H{
			"label": labelArea,
			"total": TotalArea,
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
