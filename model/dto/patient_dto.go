package dto

type Patient_request struct {
	Id           int    `json:"id" gorm:"type:int;column:id"`
	Patient_name string `json:"patient_name" gorm:"type:varchar;column:patient_name"`
	Umur         string `json:"umur" gorm:"type:varchar;column:umur"`
	Alamat       string `json:"alamat" gorm:"type:varchar;column:alamat"`
}

type Patient_response struct {
	Patient_name string `json:"patient_name" gorm:"type:varchar;column:patient_name"`
	Umur         string `json:"umur" gorm:"type:varchar;column:umur"`
	Alamat       string `json:"alamat" gorm:"type:varchar;column:alamat"`
}
