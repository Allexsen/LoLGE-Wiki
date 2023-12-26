package router

import (
	"github.com/Allexsen/LoLGE-Wiki/backend/internal/handlers"
)

func setupFeDynamicLinkingRouter() {
	r.GET("/html/:html/", handlers.HtmlGET)

	r.GET("/css/:styling/", handlers.CssGET)

	r.GET("/js/:script/", handlers.JsGET)

	r.GET("/images/:image/", handlers.ImageGet)
}
