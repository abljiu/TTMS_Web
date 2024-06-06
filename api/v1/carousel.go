package v1

import (
	"TTMS_Web/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCarousel(c *gin.Context) {
	var listCarousel service.CarouselService
	if err := c.ShouldBind(&listCarousel); err == nil {
		res := listCarousel.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func AddCarousel(c *gin.Context) {
	var addCarousel service.CarouselService
	if err := c.ShouldBind(&addCarousel); err == nil {
		res := addCarousel.Add(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
