package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"
)

/**
Get default data ots, and create pie chart in frontend
*/
func (idb *InDB) GetOTS(c *gin.Context) {
	var (
		OtsArray      []structs.TbOutstandingStruct
		OtsArrayArea  []structs.TbOutstandingStruct
		OtsArrayRet   []structs.TbOutstandingStruct
		OtsArrayPack  []structs.TbOutstandingStruct
		OtsArrayTrans []structs.TbOutstandingStruct

		Ots    structs.TbOutstandingStruct
		Retail structs.TbRetail
		result gin.H

		labelArea []string
		TotalArea []int
		//
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
	)

	var wg sync.WaitGroup

	wg.Add(6)
	/**
	  Retail
	  **/
	go func() {
		retl, err := idb.DB.Raw("CALL getDefaultRetail()").Find(&OtsArrayRet).Rows()
		lengthRet := len(OtsArrayRet)
		prepIncRet := 0
		prepTempLabelRet := ""
		prepTempTotalRet := 0
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		defer retl.Close()
		for retl.Next() {
			err := retl.Scan(&Ots.Outstanding_quantitys, &Retail.Retail_label)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			}
			if prepIncRet < lengthRet {
				if prepIncRet == 0 || (Retail.Retail_label == prepTempLabelRet) {
					prepTempTotalRet = prepTempTotalRet + Ots.Outstanding_quantitys
				} else {
					labelRetl = append(labelRetl, prepTempLabelRet)
					TotalRetl = append(TotalRetl, prepTempTotalRet)
					prepTempTotalRet = 0
					prepTempTotalRet = prepTempTotalRet + Ots.Outstanding_quantitys
				}

				if prepIncRet == (lengthRet - 1) {
					labelRetl = append(labelRetl, prepTempLabelRet)
					TotalRetl = append(TotalRetl, prepTempTotalRet)
				}

			}
			prepTempLabelRet = Retail.Retail_label
			prepIncRet++
		}
		defer wg.Done()
	}()
	/**
	  Pack
	  **/
	go func() {
		pack, err := idb.DB.Raw("call getDefaultPack()").Find(&OtsArrayPack).Rows()
		lengthPack := len(OtsArrayRet)
		prepIncPack := 0
		prepTempLabelPack := ""
		prepTempTotalPack := 0
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		defer pack.Close()
		for pack.Next() {
			err := pack.Scan(&Ots.Outstanding_quantitys, &Ots.Outstanding_package)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			} else {
				if prepIncPack < lengthPack {
					if prepIncPack == 0 || (Ots.Outstanding_package == prepTempLabelPack) {
						prepTempTotalPack = prepTempTotalPack + Ots.Outstanding_quantitys
					} else {
						labelPack = append(labelPack, prepTempLabelPack)
						TotalPack = append(TotalPack, prepTempTotalPack)
						prepTempTotalPack = 0
						prepTempTotalPack = prepTempTotalPack + Ots.Outstanding_quantitys

					}

					if prepIncPack == (lengthPack - 1) {
						labelPack = append(labelPack, prepTempLabelPack)
						TotalPack = append(TotalPack, prepTempTotalPack)
					}
				}
				prepTempLabelPack = Ots.Outstanding_package
			}
			prepIncPack++
		}
		defer wg.Done()
	}()
	/**
	  Disp
	  **/
	go func() {
		disp, err := idb.DB.Raw("CALL getDefaultDispatch()").Find(&OtsArray).Rows()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		lengthDisp := len(OtsArray)
		prepTempLabelDisp := ""
		prepTempTotalDisp := 0
		prepIncDisp := 0
		defer disp.Close()
		for disp.Next() {
			err := disp.Scan(&Ots.Outstanding_quantitys, &Ots.Outstanding_dispatcher)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			} else {
				if prepIncDisp < lengthDisp {
					if prepIncDisp == 0 || (Ots.Outstanding_dispatcher == prepTempLabelDisp) {
						prepTempTotalDisp = prepTempTotalDisp + Ots.Outstanding_quantitys
					} else {
						labelDisp = append(labelDisp, prepTempLabelDisp)
						TotalDisp = append(TotalDisp, prepTempTotalDisp)
						prepTempTotalDisp = 0
						prepTempTotalDisp = prepTempTotalDisp + Ots.Outstanding_quantitys
					}
					if prepIncDisp == (lengthDisp - 1) {
						labelDisp = append(labelDisp, prepTempLabelDisp)
						TotalDisp = append(TotalDisp, prepTempTotalDisp)
					}
				}

				prepTempLabelDisp = Ots.Outstanding_dispatcher
				prepIncDisp = prepIncDisp + 1
			}
		}
		defer wg.Done()
	}()
	/**
	  Area
	  **/
	go func() {
		rows, err := idb.DB.Raw("CALL getDefaultArea()").Find(&OtsArrayArea).Rows()
		lengthArea := len(OtsArrayArea)
		prepTempLabelArea := ""
		prepTempTotalArea := 0
		prepIncArea := 0
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&Ots.Outstanding_quantitys, &Ots.Outstanding_area)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			} else {

				if prepIncArea < lengthArea {
					if prepIncArea == 0 || (Ots.Outstanding_area == prepTempLabelArea) {
						prepTempTotalArea = prepTempTotalArea + Ots.Outstanding_quantitys
					} else {
						labelArea = append(labelArea, prepTempLabelArea)
						TotalArea = append(TotalArea, prepTempTotalArea)
						prepTempTotalArea = 0
						prepTempTotalArea = prepTempTotalArea + Ots.Outstanding_quantitys
					}
					if prepIncArea == (lengthArea - 1) {
						labelArea = append(labelArea, prepTempLabelArea)
						TotalArea = append(TotalArea, prepTempTotalArea)
					}

				}
			}
			prepTempLabelArea = Ots.Outstanding_area
			prepIncArea++
		}
		defer wg.Done()
	}()
	/**
	  Late
	  **/
	go func() {
		lateLess0 := 0
		lateIs0 := 0
		late1 := 0
		late2 := 0
		late3 := 0
		late4 := 0
		late5 := 0
		lateB6T10 := 0
		lateM10 := 0
		late, err := idb.DB.Raw("CALL getDefaultLate()").Rows()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		defer late.Close()
		for late.Next() {
			err := late.Scan(&Ots.Outstanding_quantitys, &Ots.Outstanding_late)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			} else {
				if Ots.Outstanding_late < 0 {
					lateLess0 = lateLess0 + Ots.Outstanding_quantitys
				}
				if Ots.Outstanding_late == 0 {
					lateIs0 = lateIs0 + Ots.Outstanding_quantitys
				}
				if Ots.Outstanding_late == 1 {
					late1 = late1 + Ots.Outstanding_quantitys
				}
				if Ots.Outstanding_late == 2 {
					late2 = late2 + Ots.Outstanding_quantitys
				}
				if Ots.Outstanding_late == 3 {
					late3 = late3 + Ots.Outstanding_quantitys
				}
				if Ots.Outstanding_late == 4 {
					late4 = late4 + Ots.Outstanding_quantitys
				}
				if Ots.Outstanding_late == 5 {
					late5 = late5 + Ots.Outstanding_quantitys
				}
				if Ots.Outstanding_late > 5 && Ots.Outstanding_late < 11 {
					lateB6T10 = lateB6T10 + Ots.Outstanding_quantitys
				}
				if Ots.Outstanding_late > 10 {
					lateM10 = lateM10 + Ots.Outstanding_quantitys
				}
			}
		}
		labelLate = append(labelLate, 6)
		TotalLate = append(TotalLate, lateLess0)

		labelLate = append(labelLate, 9)
		TotalLate = append(TotalLate, lateIs0)

		labelLate = append(labelLate, 1)
		TotalLate = append(TotalLate, late1)

		labelLate = append(labelLate, 2)
		TotalLate = append(TotalLate, late2)

		labelLate = append(labelLate, 3)
		TotalLate = append(TotalLate, late3)

		labelLate = append(labelLate, 4)
		TotalLate = append(TotalLate, late4)

		labelLate = append(labelLate, 5)
		TotalLate = append(TotalLate, late5)

		labelLate = append(labelLate, 7)
		TotalLate = append(TotalLate, lateB6T10)

		labelLate = append(labelLate, 8)
		TotalLate = append(TotalLate, lateM10)
		defer wg.Done()
	}()
	/**
	  Trans
	  **/
	go func() {
		trans, err := idb.DB.Raw("call getDefaultTrans()").Find(&OtsArrayTrans).Rows()
		lengthTrans := len(OtsArrayTrans)
		prepTempLabelTrans := ""
		prepTempTotalTrans := 0
		prepIncTrans := 0
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
		defer trans.Close()

		for trans.Next() {
			err := trans.Scan(&Ots.Outstanding_quantitys, &Ots.Outstanding_transporter)
			if err != nil {

				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			} else {
				if prepIncTrans < lengthTrans {
					if prepIncTrans == 0 || (Ots.Outstanding_transporter == prepTempLabelTrans) {
						prepTempTotalTrans = prepTempTotalTrans + Ots.Outstanding_quantitys
					} else {
						labelTransport = append(labelTransport, prepTempLabelTrans)
						TotalTransport = append(TotalTransport, prepTempTotalTrans)
						prepTempTotalTrans = 0
						prepTempTotalTrans = prepTempTotalTrans + Ots.Outstanding_quantitys

					}
					if prepIncTrans == (lengthTrans - 1) {
						labelTransport = append(labelTransport, prepTempLabelTrans)
						TotalTransport = append(TotalTransport, prepTempTotalTrans)
					}
				}
				prepTempLabelTrans = Ots.Outstanding_transporter
				prepIncTrans++
			}
		}
		defer wg.Done()
	}()

	wg.Wait()

	result = gin.H{
		"lastUpdate": "2019-05-02 19:00:00",
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
		},
		"transport": gin.H{
			"label": labelTransport,
			"total": TotalTransport,
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
		Ots    structs.TbOutstandingStruct
		result gin.H
		Retail structs.TbRetail

		labelDisp []string
		TotalDisp []string

		labelRitl []string
		TotalRitl []string

		labelPack []string
		TotalPack []string

		labelArea []string
		TotalArea []string

		labelLate []int
		TotalLate []string

		labelTransport []string
		TotalTransport []string

		labelTransportOther []string
		TotalTransportOther []string

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
	disp := idb.DB.Table("tb_outstandings as ots").Select("outstanding_updt,((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000),  provid_ktgr ").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id join tb_pairing_order as ord ON bag.bagcode_name = ord.order_label").Where("order_pages = 'OTS'").Group("provid_ktgr").Order("order_data ASC")
	retl := idb.DB.Table("tb_outstandings as ots").Select("((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000), ret.retail_label ").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id").Group("retail_id")
	area := idb.DB.Table("tb_outstandings as ots").Select(" ((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000), area.tp_area_alias2 ").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id").Group("area.tp_area_alias2 ").Order("tp_area_code")
	late := idb.DB.Table("tb_outstandings as ots").Select("((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000), ots.outstanding_late").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id").Group("ots.outstanding_late")
	trans := idb.DB.Table("tb_outstandings as ots").Select("((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000) as total, trs.transporter_name").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id").Group(" transporter_name").Order("total desc")
	pack := idb.DB.Table("tb_outstandings as ots").Select("((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000), bag.bagcode_name ").Joins("join tb_provids as prov on ots.outstanding_loct = prov.provid_code join tb_bagcodes as bag on ots.outstanding_pack = bag.bagcode_code join tb_transporters as trs on ots.outstanding_trans = trs.transporter_code join tb_destinations as dest on ots.outstanding_dest = dest.destination_code join tb_tp_areas as area on ots.outstanding_area = area.tp_area_code join tb_retails as ret on ots.outstanding_ret = ret.retail_id join tb_pairing_order as ord ON bag.bagcode_name = ord.order_label").Where("order_pages = 'OTS'").Group("bagcode_code").Order("order_data ASC")

	for _, num := range nums {
		subStrn := string([]rune((re.FindAllString(num, -1))[0])[0:1])
		value := trimFirstRune(num)

		switch subStrn {
		case "1":
			disp = disp.Select("outstanding_updt,floor(((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000)),   SUBSTR(provid_name,1,12)").Group("provid_name").Where("provid_ktgr = ?", value)
			retl = retl.Where("provid_ktgr = ?", value)
			area = area.Where("provid_ktgr = ?", value)
			late = late.Where("provid_ktgr = ?", value)
			trans = trans.Where("provid_ktgr = ?", value)
			pack = pack.Where("provid_ktgr = ?", value)
			break

		case "2":
			disp = disp.Where("tp_area_alias2 = ?", value)
			retl = retl.Where("tp_area_alias2 = ?", value)
			area = area.Select("floor(((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000)) AS total,SUBSTR(dest.destination_zona,1,12)").Where("tp_area_alias2 = ?", value).Group("destination_zona")
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
		case "9":
			prepFindWil := value + "%"
			disp = disp.Where("destination_zona LIKE ?", prepFindWil)
			retl = retl.Where("destination_zona LIKE ?", prepFindWil)
			area = area.Select("((Sum(ots.outstanding_zak) * bag.bagcode_kg) / 1000),SUBSTR(dest.destination_wilayah, 1, 12)").Where("destination_zona LIKE ?", prepFindWil).Group("dest.destination_wilayah")
			late = late.Where("destination_zona LIKE ?", prepFindWil)
			trans = trans.Where("destination_zona LIKE ?", prepFindWil)
			pack = pack.Where("destination_zona LIKE ?", prepFindWil)
			break
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

		var TotalLateM10, TotalLateL1, TotalLateB610 float64
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
				i, errs := strconv.ParseFloat(Ots.Outstanding_quantity, 64)
				if errs != nil {
				} else {
					if Ots.Outstanding_late < 1 {
						TotalLateL1 = TotalLateL1 + i
					} else {
						if Ots.Outstanding_late > 5 && Ots.Outstanding_late < 10 {
							TotalLateB610 = TotalLateB610 + i
						} else {
							if Ots.Outstanding_late > 9 {
								TotalLateM10 = TotalLateM10 + i
							} else {
								labelLate = append(labelLate, Ots.Outstanding_late)
								TotalLate = append(TotalLate, Ots.Outstanding_quantity)
							}
						}
					}
				}

			}

		}
		late6Days := strconv.FormatFloat(TotalLateL1, 'f', 2, 64)
		lateB6T10 := strconv.FormatFloat(TotalLateB610, 'f', 2, 64)
		lateM10 := strconv.FormatFloat(TotalLateM10, 'f', 2, 64)

		if TotalLateL1 > 0.0 {

			labelLate = append(labelLate, 6)
			TotalLate = append(TotalLate, late6Days)
		}

		if TotalLateB610 > 0.0 {
			labelLate = append(labelLate, 7)
			TotalLate = append(TotalLate, lateB6T10)
		}

		if TotalLateM10 > 0.0 {
			labelLate = append(labelLate, 8)
			TotalLate = append(TotalLate, lateM10)

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
			"label": labelTransport,
			"total": TotalTransport,
		},
	}

	c.JSON(http.StatusOK, result)

}

func GetIOST(c *gin.Context) {

}
