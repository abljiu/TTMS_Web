package v1

//func AddTheater(c *gin.Context) {
//	var addTheater service.TheaterService
//	if err := c.ShouldBind(&addTheater); err == nil {
//		res := addTheater.Add(c.Request.Context())
//		c.JSON(http.StatusOK, res)
//	} else {
//		c.JSON(http.StatusBadRequest, ErrorResponse(err))
//		util.LogrusObj.Infoln("SubmitOrder", err)
//	}
//}
