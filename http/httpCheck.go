package http

import (
	"github.com/gin-gonic/gin"
	"github.com/signmem/ginweb/template"
)

func http_health_check() string {
	return "ok"
}

func GetPlainTextInfo(c *gin.Context) UserInfo  {

	var user UserInfo
	var domainInfo []string
	user.Name   = c.DefaultQuery("name", "none")
	domainInfo, status := c.GetQueryArray("domain")
	if status != true {
		domainInfo = append(domainInfo, "none")
	}
	user.Domain = domainInfo
	user.Token  = c.DefaultQuery("token", "none")

	return user

}

func PostPlainTextInfo(c *gin.Context) UserInfo {
	var user UserInfo
	var domainInfo []string
	user.Name   = c.DefaultPostForm("name", "none")
	domainArray := c.PostFormArray("domain")
	if len(domainArray) > 0 {
		for _, domain := range domainArray {
			domainInfo = append(domainInfo, domain)
		}
	} else {
		domainInfo = append(domainInfo, "none")
	}
	user.Domain = domainInfo
	user.Token  = c.DefaultPostForm("token", "none")

	return user
}

func GetBindTextInfo(c *gin.Context) string {
	var user UserInfo
	c.ShouldBindJSON(&user)

	return user.String()
}

func GetConfigTemplate(c *gin.Context)  ([]byte, error) {
	var user UserInfo
	c.ShouldBindJSON(&user)

	templateVars := make(map[string]interface{})
	templateVars["Name"] = user.Name
	templateVars["Domain"] = user.Domain
	templateVars["Token"] = user.Token

	configDetail := template.ProcessFile("./templates/config.tmpl",templateVars)
	return []byte(configDetail), nil
}
