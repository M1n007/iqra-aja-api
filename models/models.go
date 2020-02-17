package model

//Surah models
type Surah struct {
	Id               int    `form:"id" json:"id"`
	Surat_name       string `form:"surat_name" json:"surat_name"`
	Surat_text       string `form:"surat_text" json:"surat_text"`
	Surat_terjemahan string `form:"surat_terjemahan" json:"surat_terjemahan"`
	Count_ayat       int    `form:"count_ayat" json:"count_ayat"`
}

type Meta struct {
	Page       int `form:"page" json:"page"`
	Size       int `form:"size" json:"size"`
	Total_page int `form:"total_page" json:"total_page"`
	Total_data int `form:"total_data" json:"total_data"`
}

//Ayat models
type Ayat struct {
	Aya_id               int    `form:"aya_id" json:"aya_id"`
	Sura_id              int    `form:"sura_id" json:"sura_id"`
	Aya_number           int    `form:"aya_number" json:"aya_number"`
	Juz_id               int    `form:"juz_id" json:"juz_id"`
	Aya_text             string `form:"aya_text" json:"aya_text"`
	Translation_aya_text string `form:"translation_aya_text" json:"translation_aya_text"`
}

//Pagination Response Get Surah ...
type PaginationResponseSurah struct {
	Err     bool    `json:"err"`
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Surah `form:"data" json:"data"`
	Meta    []Meta  `form:"meta" json:"meta"`
}

//Pagination Response Get Ayat ...
type PaginationResponseAyat struct {
	Err     bool   `json:"err"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Ayat `form:"data" json:"data"`
}
