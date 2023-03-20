package GinItem

import (
	"GoLang/Common"
	"GoLang/Modules/Item/Business"
	"GoLang/Modules/Item/Storange"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItem(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := Storange.NewSQLSTore(db)
		business := Business.NewGetItemBusiness(store)

		data, err := business.GetItemById(context.Request.Context(), id)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		context.JSON(http.StatusOK, Common.SimpleSuccessResponse(data))
	}
}
