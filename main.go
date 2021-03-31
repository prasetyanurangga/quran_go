package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/prasetyanurangga/quran_go/config"
	"github.com/prasetyanurangga/quran_go/controller"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/prasetyanurangga/quran_go/docs"
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
)


// @title Quran Go
// @description This is quran rest api
// @version 1.0
// @host quran-go.herokuapp.com
// @BasePath /
// @contact.name API Support
// @contact.email angganurprasetya4@gmail.com
func main() {
	r := gin.Default()

	config.ConnectDB()

	surah := r.Group("/surah")
	{
	  	surah.GET("get_all", controller.GetAllSurah)
		surah.GET("search/:keyword", controller.SearchSurah)
		surah.GET("get_by_id/:id", controller.GetSurahById)
		surah.GET("get_by_id_with_ayah/:id", controller.GetSurahByIdWithAyah)
	}

	ayah := r.Group("/ayah")
	{
		ayah.GET("get_by_id_surah/:id", controller.GetAyahByIdSurah)
		ayah.GET("get_by_juz/:juz", controller.GetAyahByJuz)
		ayah.GET("get_by_id_ayah/:id", controller.GetAyahByIdAyah)
		ayah.GET("search/:keyword", controller.SearchAyah)
	}

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/swagger/index.html")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}