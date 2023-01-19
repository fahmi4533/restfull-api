package handler

import (
	"fmt"
	"net/http"
	"pustaka/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetQuery(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "hello word",
		"subtitel": "Jamaah masjid",
	})
}
func QueryHandler(c *gin.Context) {
	name := c.Query("name")
	harga := c.Query("harga")

	c.JSON(http.StatusOK, gin.H{"name": name, "harga": harga})
}
func BookInputData(c *gin.Context) {
	var bookInput book.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {

		errorMassages := []string{}
		for _, v := range err.(validator.ValidationErrors) {
			errorMassage := fmt.Sprintf("Error n field %s, condition : %s", v.Field(), v.ActualTag())
			errorMassages = append(errorMassages, errorMassage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMassages,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"name":  bookInput.Name,
		"harga": bookInput.Harga,
		// "sub_titel": bookInput.Subtitel,
	})
}
