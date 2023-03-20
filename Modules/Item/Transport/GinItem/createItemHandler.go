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

func CreateItem(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		var data Models.TodoItemCreation
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := Storange.NewSQLSTore(db)
		business := Business.NewCreateItemBusiness(store)

		if err := business.CreateNewItem(context.Request.Context(), &data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, Common.SimpleSuccessResponse(data.Id))
	}
}
