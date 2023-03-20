package GinItem

import (
	"GoLang/Common"
	"GoLang/Modules/Item/Business"
	"GoLang/Modules/Item/Storange"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DeleteItem(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {

		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := Storange.NewSQLSTore(db)
		biz := Business.NewDeleteItemBusiness(store)

		if err := biz.DeleteItemById(context.Request.Context(), id); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, Common.SimpleSuccessResponse(true))
	}
}
