package routes

import (
	"receipt-processor-challenge/models"
	"receipt-processor-challenge/pkg"

	"github.com/gin-gonic/gin"
)

type ReciptHandler struct {
	reciptStore *pkg.ReciptStore
}

func NewReciptHandler(rs *pkg.ReciptStore) *ReciptHandler {
	return &ReciptHandler{
		reciptStore: rs,
	}
}

func (rh *ReciptHandler) AddNewRecipt(c *gin.Context) {
	var r models.Recipt
	err := c.BindJSON(&r)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid JSON",
		})
		return
	}
	res := rh.reciptStore.AddNewRecipt(r)
	c.JSON(200, gin.H{
		"ID": res,
	})
}

func (rh *ReciptHandler) GetRecipt(c *gin.Context) {
	date := c.Param("ID")
	r := rh.reciptStore.GetRecipt(date)
	if r == -1 {
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
		return
	}
	c.JSON(200, gin.H{
		"points": r,
	})
}
