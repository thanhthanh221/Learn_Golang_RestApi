package GinItem

import (
	"GoLang/Common"
	"GoLang/Modules/Item/Business"
	"GoLang/Modules/Item/Storange"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetListItem(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		store := Storange.NewSQLSTore(db)
		business := Business.NewGetListItemBusiness(store)

		data, err := business.GetListItem(context)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		context.JSON(http.StatusOK, Common.SimpleSuccessResponse(data))
	}
}
