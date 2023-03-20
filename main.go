package main

import (
	"GoLang/Modules/Item/Transport/GinItem"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello")

	dsn := "root:Phuongyen1234@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	db, errDb := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errDb != nil {
		log.Fatalln(errDb)
	}

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", GinItem.CreateItem(db))
			items.GET("", GinItem.GetListItem(db))
			items.GET("/:id", GinItem.GetItem(db))
			items.GET("/pagingItems", GinItem.GetPagingItem(db))
			items.PUT("/:id", GinItem.UpdateItem(db))
			items.DELETE("/:id", GinItem.DeleteItem(db))
		}
	}
	r.Run(":5002") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
