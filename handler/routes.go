package handler

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	r.GET("/person/:person_id", GetPersonById)
	r.POST("/person", CreatePerson)
}

func StartApp() {
	r := gin.Default()
	Routes(r)

	r.Run(PORT)
}
