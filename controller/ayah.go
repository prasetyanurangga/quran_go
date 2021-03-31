package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prasetyanurangga/quran_go/model"
	"github.com/prasetyanurangga/quran_go/config" 
	// "github.com/swaggo/swag/example/celler/httputil"

)

// GetAyahByIdSurah godoc
// @Summary Get Ayah By Id Surah
// @Description Get Ayah By Id Surah
// @ID get-ayah-by-id-surah
// @Accept  json
// @Produce  json
// @Tags Ayah
// @Param id path int true "IdSurah"
// @Success 200 {object} config.JSONSuccessResult{data=[]model.Ayah,success=bool,count=int}
// @Failure 500 {object} config.JSONErrorResult{success=bool,count=int}
// @Router /ayah/get_by_id_surah/{id} [get]
func GetAyahByIdSurah(c *gin.Context) {
  var ayah []model.Ayah

  if err := config.DB.Where("idSurah = ?", c.Param("id")).Find(&ayah).Error; err != nil {
  	config.ErrorResponse(c, 0)
  } else {
  	config.SuccessResponse(c, ayah, len(ayah))
  }
}


// GetAyahByJuz godoc
// @Summary Get Ayah By Juz
// @Description Get Ayah By Juz
// @ID get-ayah-by-juz
// @Accept  json
// @Tags Ayah
// @Produce  json
// @Param juz path int true "Juz"
// @Success 200 {object} config.JSONSuccessResult{data=[]model.Ayah,success=bool,count=int}
// @Failure 500 {object} config.JSONErrorResult{success=bool,count=int}
// @Router /ayah/get_by_juz/{juz} [get]
func GetAyahByJuz(c *gin.Context) {
  var ayah []model.Ayah

  if err := config.DB.Where("juz = ?", c.Param("juz")).Find(&ayah).Error; err != nil {
  	config.ErrorResponse(c, 0)
  } else {
  	config.SuccessResponse(c, ayah, len(ayah))
  }
}


// GetAyahByIdAyah godoc
// @Summary Get Ayah By Id Ayah
// @Description Get Ayah By Id Ayah
// @ID get-ayah-by-id-ayah
// @Accept  json
// @Produce  json
// @Tags Ayah
// @Param id path int true "IdAyah"
// @Success 200 {object} config.JSONSuccessResult{data=model.Ayah,success=bool,count=int}
// @Failure 500 {object} config.JSONErrorResult{success=bool,count=int}
// @Router /ayah/get_by_id_ayah/{id} [get]
func GetAyahByIdAyah(c *gin.Context) {
  var ayah model.Ayah

  if err := config.DB.Where("IdAyahInSurah = ?", c.Param("id")).First(&ayah).Error; err != nil {
  	config.ErrorResponse(c, 0)
  } else {
  	config.SuccessResponse(c, ayah, 1)
  }
}


// SearchAyah godoc
// @Summary Search Ayah
// @Description Search Ayah
// @ID search-ayah
// @Accept  json
// @Tags Ayah
// @Produce  json
// @Param keyword path string true "Keyword"
// @Success 200 {object} config.JSONSuccessResult{data=[]model.Ayah,success=bool,count=int}
// @Failure 500 {object} config.JSONErrorResult{success=bool,count=int}
// @Router /ayah/search/{keyword} [get]
func SearchAyah(c *gin.Context) {
  var ayah []model.Ayah

  var keyword = c.Param("keyword")
  if err := config.DB.Where(
  	"indoText LIKE ? OR "+ 
  	"tafsirLong LIKE ? OR "+
  	"tafsirShort LIKE ? OR "+ 
  	"arabText LIKE ?", 
  	"%"+keyword+"%", 
  	"%"+keyword+"%", 
  	"%"+keyword+"%", 
  	"%"+keyword+"%",
  ).Find(&ayah).Error; err != nil {
  	config.ErrorResponse(c, 0)
  } else {
  	config.SuccessResponse(c, ayah, len(ayah))
  }
}