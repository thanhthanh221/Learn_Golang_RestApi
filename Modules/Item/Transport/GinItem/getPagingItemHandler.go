package GinItem

import (
	"GoLang/Common"
	"GoLang/Modules/Item/Business"
	"GoLang/Modules/Item/Models"
	"GoLang/Modules/Item/Storange"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPagingItem(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		var paging Common.Paging

		if err := context.ShouldBind(&paging); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		paging.Process()

		var filter Models.Filter

		if err := context.ShouldBind(&filter); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := Storange.NewSQLSTore(db)
		biz := Business.NewGetPagingItemBusiness(store)

		result, err := biz.GetPagingItem(context, &filter, &paging)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		context.JSON(http.StatusOK, Common.NewSuccessResponse(result, paging, nil))
	}
}
