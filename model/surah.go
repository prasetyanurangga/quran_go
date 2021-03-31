package model


type Surah struct {
	IdSurah uint `json:"idSurah" gorm:"column:idSurah"`
	NumberOfAyah uint `json:"numberOfAyah" gorm:"column:numberOfAyah"`
	NameShort string `json:"nameShort" gorm:"column:nameShort"`
	NameLong string `json:"nameLong" gorm:"column:nameLong"`
	NameTransliterationEn string `json:"nameTransliterationEn" gorm:"column:nameTransliterationEn"`
	NameTransliterationId string `json:"nameTransliterationId" gorm:"column:nameTransliterationId"`
	NameTranslationEn string `json:"nameTranslationEn" gorm:"column:nameTranslationEn"`
	NameTranslationId string `json:"nameTranslationId" gorm:"column:nameTranslationId"`
	RevelationArab string `json:"revelationArab" gorm:"column:revelationArab"`
	RevelationEn string `json:"revelationEn" gorm:"column:revelationEn"`
	RevelationId string `json:"revelationId" gorm:"column:revelationId"`
	Tafsir string `json:"tafsir" gorm:"column:tafsir"`
	Ayah []Ayah `json:"ayah"`
}


func (Surah) TableName() string {
  return "surah"
}