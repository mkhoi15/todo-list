package ginitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"to-list/common"
	"to-list/module/item/biz"
	"to-list/module/item/storage"
)

func GetItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		store := storage.NewSQLStorage(db)
		business := biz.NewGetItemBiz(store)

		data, err := business.GetItemById(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(data))
	}
}
