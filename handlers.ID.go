package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckID godoc
// @Summary check id number
// @Description check id validation and return city information
// @Tags ID
// @ID check-id
// @Accept  json
// @Produce  html
// @Param id path string true "id"
// @Header 200 {string} Token "qwerty"
// @Router /ID/validation/{id} [get]
func CheckID(c *gin.Context) {
	if id := c.Param("id"); id != "" {
		// Check if the id is valid and return city info
		if id, err := idValidation(id); err == nil && id != nil {
			c.JSON(http.StatusOK, id)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
