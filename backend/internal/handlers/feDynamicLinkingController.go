package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HtmlGET(c *gin.Context) {
	c.File(DIR_FRONT + "/html/" + c.Param("html"))
}

func CssGET(c *gin.Context) {
	log.Println(DIR_FRONT + "/css/" + c.Param("styling"))
	c.File(DIR_FRONT + "/css/" + c.Param("styling"))
}

func JsGET(c *gin.Context) {
	c.File(DIR_FRONT + "/js/" + c.Param("script"))
}

func ImageGet(c *gin.Context) {
	c.File(DIR_FRONT + "/images/" + c.Param("image"))
}
