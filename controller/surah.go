package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prasetyanurangga/quran_go/model"
	"github.com/prasetyanurangga/quran_go/config"

)

// GetAllSurah godoc
// @Summary Get All Surah
// @Description Get All Surah
// @ID get-all-surah
// @Accept  json
// @Tags Surah
// @Produce  json
// @Success 200 {object} config.JSONSuccessResult{data=[]model.Surah,success=bool,count=int}
// @Failure 500 {object} config.JSONErrorResult{success=bool,count=int}
// @Router /surah/get_all/ [get]
func GetAllSurah(c *gin.Context) {
  var surah []model.Surah
  if err := config.DB.Find(&surah).Error; err != nil {
	config.ErrorResponse(c, 0)
  } else {
  	config.SuccessResponse(c, surah, len(surah))
  }
}

// GetSurahById godoc
// @Summary Get Surah By Id Surah
// @Description Get Surah By Id Surah
// @ID get-surah-by-id-surah
// @Accept  json
// @Produce  json
// @Tags Surah
// @Param id path string true "Id Surah"
// @Success 200 {object} config.JSONSuccessResult{data=model.Surah,success=bool,count=int}
// @Failure 500 {object} config.JSONErrorResult{success=bool,count=int}
// @Router /surah/get_by_id/{id} [get]
func GetSurahById(c *gin.Context) {
  var surah model.Surah
  if err := config.DB.Where("idSurah = ?", c.Param("id")).First(&surah).Error; err != nil {
  	config.ErrorResponse(c, 0)
  } else {
  	config.SuccessResponse(c, surah, 1)
  }
}

// GetSurahByIdWithAyah godoc
// @Summary Get Surah By Id Surah With Ayah
// @Description Get Surah By Id Surah With Ayah
// @ID get-surah-by-id-surah-with-ayah
// @Accept  json
// @Tags Surah
// @Produce  json
// @Param id path string true "Id Surah"
// @Success 200 {object} config.JSONSuccessResult{data=model.Surah,success=bool,count=int}
// @Failure 500 {object} config.JSONErrorResult{success=bool,count=int}
// @Router /surah/get_by_id_with_ayah/{id} [get]
func GetSurahByIdWithAyah(c *gin.Context) {
  var surah model.Surah
  if err := config.DB.Where("idSurah = ?", c.Param("id")).First(&surah).Error; err != nil {

  	config.ErrorResponse(c, 0)
  } else {

  	var ayah []model.Ayah
  	config.DB.Where("idSurah = ?", c.Param("id")).Find(&ayah)

  	surah.Ayah = ayah

  	config.SuccessResponse(c, surah, 1)
  }
}

// SearchSurah godoc
// @Summary Get Search Surah
// @Description Get Search Surah
// @ID search-surah
// @Accept  json
// @Tags Surah
// @Produce  json
// @Param keyword path string true "Keyword"
// @Success 200 {object} config.JSONSuccessResult{data=[]model.Surah,success=bool,count=int}
// @Failure 500 {object} config.JSONErrorResult{success=bool,count=int}
// @Router /surah/search/{keyword} [get]
func SearchSurah(c *gin.Context) {
  var surah []model.Surah

  var keyword = c.Param("keyword")
  if err := config.DB.Where(
  	"numberOfAyah LIKE ? OR "+ 
  	"nameShort LIKE ? OR "+
  	"nameLong LIKE ? OR "+ 
  	"nameTransliterationEn LIKE ? OR "+
  	"nameTransliterationId LIKE ? OR "+ 
  	"nameTranslationEn LIKE ? OR "+ 
  	"nameTranslationId LIKE ? OR "+ 
  	"revelationArab LIKE ? OR "+
  	"revelationEn LIKE ? OR "+
  	"revelationId LIKE ? OR "+ 
  	"tafsir LIKE ?",  
  	"%"+keyword+"%", 
  	"%"+keyword+"%", 
  	"%"+keyword+"%", 
  	"%"+keyword+"%",  
  	"%"+keyword+"%", 
  	"%"+keyword+"%", 
  	"%"+keyword+"%", 
  	"%"+keyword+"%", 
  	"%"+keyword+"%", 
  	"%"+keyword+"%", 
  	"%"+keyword+"%",
  ).Find(&surah).Error; err != nil {
  	config.ErrorResponse(c, 0)
  } else {
  	config.SuccessResponse(c, surah, len(surah))
  }
}