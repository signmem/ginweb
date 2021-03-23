package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpRoute() {
	router := gin.Default()
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

	router.Run(":80")
}
