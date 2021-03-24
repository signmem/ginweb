package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpRoute() {
	router := gin.Default()
<<<<<<< HEAD
	// httpApi := "/v1"

	v1 := router.Group("/v1")
	{
		v1.GET("/health_check", func(c *gin.Context) {
			info := http_health_check()
			result := gin.H{
				"msg":    info,
				"result": true,
			}
			c.JSON(http.StatusOK, result)
		})

		v1.GET("/plain", func(c *gin.Context) {

			// method  get
			// content-type  text/plain

			status := true
			info := GetPlainTextInfo(c)

			result := gin.H{
				"msg":    info,
				"result": status,
			}
			c.JSON(http.StatusOK, result)
		})

		v1.POST("/plain", func(c *gin.Context) {

			// method  post
			// content-type  text/plain

			status := true
			info := PostPlainTextInfo(c)
			result := gin.H{
				"msg":    info,
				"result": status,
			}
			c.JSON(http.StatusOK, result)
		})

		v1.GET("/bind", func(c *gin.Context) {

			// method  get
			// content-type  application/json

			status := true
			info := GetBindTextInfo(c)
			result := gin.H{
				"msg":    info,
				"result": status,
			}
			c.JSON(http.StatusOK, result)
		})

		v1.POST("/bind", func(c *gin.Context) {

			// method  get
			// content-type  application/json

			status := true
			info := GetBindTextInfo(c)
			result := gin.H{
				"msg":    info,
				"result": status,
			}
			c.JSON(http.StatusOK, result)
		})

		v1.POST("/config", func(c *gin.Context) {

			// raw data return
			// configdata genrate by post input

			configData, err := GetConfigTemplate(c)
			if err != nil {
				msg := fmt.Sprint(err)
				respo := gin.H{
					"msg":    msg,
					"result": false,
				}
				c.JSON(http.StatusOK, respo)
			}
			c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(configData))
		})

	}

	v2 := router.Group("/v2")
	{
		v2.GET("/health_check", func(c *gin.Context) {
			info := http_health_check()
			result := gin.H{
				"msg":    info,
				"result": true,
			}
			c.JSON(http.StatusOK, result)
		})
	}
=======
	httpApi := "/v1"

	router.GET("/health_check", func(c *gin.Context) {
		info := http_health_check()
		result := gin.H {
			"msg" : info,
			"result" : true,
		}
		c.JSON(http.StatusOK, result)
	})

	router.GET( httpApi + "/plain", func(c *gin.Context) {

		// method  get
		// content-type  text/plain

		status := true
		info := GetPlainTextInfo(c)

		result := gin.H {
			"msg": info,
			"result": status,
		}
		c.JSON(http.StatusOK, result)
	})

	router.POST( httpApi + "/plain", func(c *gin.Context) {

		// method  post
		// content-type  text/plain

		status := true
		info := PostPlainTextInfo(c)
		result := gin.H {
			"msg": info,
			"result": status,
		}
		c.JSON(http.StatusOK, result)
	})

	router.GET(httpApi + "/bind", func(c *gin.Context){

		// method  get
		// content-type  application/json

		status := true
		info := GetBindTextInfo(c)
		result := gin.H {
			"msg": info,
			"result": status,
		}
		c.JSON(http.StatusOK, result)
	})

	router.POST(httpApi + "/bind", func(c *gin.Context){

		// method  get
		// content-type  application/json

		status := true
		info := GetBindTextInfo(c)
		result := gin.H {
			"msg": info,
			"result": status,
		}
		c.JSON(http.StatusOK, result)
	})

	router.POST(httpApi + "/config", func(c *gin.Context){

		// raw data return
		// configdata genrate by post input

		configData, err := GetConfigTemplate(c)
		if err != nil {
			msg := fmt.Sprint(err)
			respo := gin.H {
				"msg" : msg,
				"result" : false,
			}
			c.JSON(http.StatusOK, respo)
		}
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(configData))
	} )
>>>>>>> 6ac36c6d2bb102e9da7e9fa1c13b1d5dba06755b

	router.Run(":80")
}
