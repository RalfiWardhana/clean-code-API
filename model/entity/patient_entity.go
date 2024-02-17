package entity

type Patient struct {
	Id           int    `json:"id" gorm:"type:int;column:id"`
	Patient_name string `json:"name" gorm:"type:varchar;column:nama"`
	Umur         string `json:"umur" gorm:"type:varchar;column:umur"`
	Alamat       string `json:"alamat" gorm:"type:varchar;column:alamat"`
}
