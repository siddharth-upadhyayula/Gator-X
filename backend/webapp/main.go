package main

import (
	model "webapp/model"
	view "webapp/views"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.BaseReview{})
	db.AutoMigrate(&model.Places{})

	server := gin.Default()

	server.Use(gin.Logger())

	server.Use(gin.Recovery())

	server.GET("/getallplaces", view.GetallplacesView(db))
	server.POST("/postplace", view.PostplaceView(db))
	server.GET("/getallreviews", view.GetallreviewsView(db))
	server.POST("/postreview", view.PostreviewView(db))
	server.POST("/editreview/", view.EditreviewView(db))
	server.DELETE("/delete/", view.DeletereviewView(db))

	server.Run()

}