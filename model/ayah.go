package model

type Ayah struct {
	IdAyahInQuran uint `json:"idAyahInQuran" gorm:"column:idAyahInQuran"` 
  	IdAyahInSurah uint `json:"idAyahInSurah" gorm:"column:idAyahInSurah"` 
  	IdSurah uint `json:"idSurah" gorm:"column:idSurah"` 
  	ArabText string `json:"arabText" gorm:"column:arabText"` 
  	IndoText string `json:"indoText" gorm:"column:indoText"` 
  	TafsirLong string `json:"tafsirLong" gorm:"column:tafsirLong"`
  	TafsirShort string `json:"tafsirShort" gorm:"column:tafsirShort"` 
  	TransliterationEn string `json:"transliterationEn" gorm:"column:transliterationEn"` 
  	EnText string `json:"enText" gorm:"column:enText"`
  	Juz uint `json:"juz" gorm:"column:juz"`
  	Page uint `json:"page" gorm:"column:page"`
  	Manzil uint `json:"manzil" gorm:"column:manzil"`
  	Ruku uint `json:"ruku" gorm:"column:ruku"`
  	HizbQuarter uint `json:"hizbQuarter" gorm:"column:hizbQuarter"`
  	SajdaRecommended bool `json:"sajda_recommended" gorm:"column:sajda_recommended"`
  	SajdaObligatory bool `json:"sajda_obligatory" gorm:"column:sajda_obligatory"`
  	Audio1 string `json:"audio_1" gorm:"column:audio_1"`
  	Audio2 string`json:"audio_2" gorm:"column:audio_2"`
  	Audio3 string `json:"audio_3" gorm:"column:audio_3"`
}

func (Ayah) TableName() string {
  return "ayah"
}