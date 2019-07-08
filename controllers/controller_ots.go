package controllers

import (
	"../structs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
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
	errs := c.BindJSON(&createParams)
	if errs != nil {
		c.JSON(http.StatusUnauthorized, errs.Error())
		c.Abort()
	}
	dataNew := createParams.DataNew
	dataNewLength := createParams.DataNewLength

	fmt.Println("-----------------------------------------------------")
	fmt.Println(dataNew)
	fmt.Println("-----------------------------------------------------")
	fmt.Println(dataNewLength)
	fmt.Println("-----------------------------------------------------")

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
		retl, err := idb.DB.Raw("CALL getFindRetail(? , ?)", dataNew, dataNewLength).Find(&OtsArrayRet).Rows()
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
		pack, err := idb.DB.Raw("call getFindPack(? , ?)", dataNew, dataNewLength).Find(&OtsArrayPack).Rows()
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
		disp, err := idb.DB.Raw("CALL getFindDisp(? , ?)", dataNew, dataNewLength).Find(&OtsArray).Rows()
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
		rows, err := idb.DB.Raw("CALL getFindArea(? , ?)", dataNew, dataNewLength).Find(&OtsArrayArea).Rows()
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
		late, err := idb.DB.Raw("CALL getFindAging(? , ?)", dataNew, dataNewLength).Rows()
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
		trans, err := idb.DB.Raw("call getFindTrans(? , ?)", dataNew, dataNewLength).Find(&OtsArrayTrans).Rows()
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

func GetIOST(c *gin.Context) {

}
