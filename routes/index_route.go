package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexRoute struct {
}

func NewIndexRoute() *IndexRoute {
	return &IndexRoute{}
}

func (r *IndexRoute) indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (r *IndexRoute) IndexRoute(rg *gin.RouterGroup) {
	rg.GET("/", r.indexHandler)
}
