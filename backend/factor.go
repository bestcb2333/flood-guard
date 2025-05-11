package main

import (
	"math"
	"time"

	"github.com/gin-gonic/gin"
)

type RiskFactor struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func AddFactorRoute(r *gin.Engine, pbc *PreloaderBaseConfig) {

	r.GET("/factors", func(c *gin.Context) {

		var regions []Region
		if err := pbc.DB.Find(&regions).Error; err != nil {
			c.JSON(500, Resp("查询区域数据失败", err, nil))
			return
		}

		var factors []RiskFactor
		for _, region := range regions {

			var factor RiskFactor
			factor.Name = region.Name
			var times uint

			if region.Altitude != nil {
				altitude := float64(*region.Altitude)
				factor.Value += 1 / (altitude + 1)
				times++
			}

			if region.Drainage != nil {
				drainage := float64(*region.Drainage)
				factor.Value += 1 / (drainage + 1)
				times++
			}

			if region.Forecast != nil {
				if *region.Forecast < 0 || *region.Forecast > 10 {
					c.JSON(400, Resp("该区域的预报指数不合法", nil, nil))
					return
				}
				factor.Value += float64(*region.Forecast) / 10
				times++
			}

			var history History
			if err := pbc.DB.Last(
				&history,
				"created_at > ? AND region_id = ?",
				time.Now().AddDate(0, 0, -1), region.ID,
			).Error; err == nil {
				if history.Rainfall != nil {
					factor.Value += *history.Rainfall / (*history.Rainfall + 1)
					times++
				}
				if history.Waterlevel != nil {
					factor.Value += *history.Waterlevel / (*history.Waterlevel + 1)
					times++
				}
			}

			if times != 0 {
				factor.Value /= float64(times)
				factor.Value = math.Round(factor.Value*100) / 100
				factors = append(factors, factor)
			}
		}

		c.JSON(200, Resp("查询成功", nil, factors))
	})
}
