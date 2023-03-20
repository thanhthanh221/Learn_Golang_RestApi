package GinItem

import (
	"GoLang/Common"
	"GoLang/Modules/Item/Business"
	"GoLang/Modules/Item/Models"
	"GoLang/Modules/Item/Storange"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateItem(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		var data Models.TodoItemUpdate

		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := Storange.NewSQLSTore(db)
		bus := Business.NewGetUpdateItemBusiness(store)

		if err := bus.UpdateItemById(context.Request.Context(), id, &data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, Common.SimpleSuccessResponse(true))
	}
}
