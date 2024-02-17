package entity

type Order struct {
	Id               int    `json:"id" gorm:"type:integer;column:id"`
	Id_patient       int    `json:"id_patient" gorm:"type:integer;column:id_patient"`
	Id_doctor        int    `json:"id_doctor" gorm:"type:integer;column:id_doctor"`
	Waktu            string `json:"waktu" gorm:"type:varchar;column:waktu"`
	Harga_konsultasi int    `json:"harga_konsultasi" gorm:"type:integer;column:harga_konsultasi"`
	Harga_obat       int    `json:"harga_obat" gorm:"type:integer;column:harga_obat"`
	Total_harga      int    `json:"total_harga" gorm:"type:integer;column:total_harga"`
	Doctor_name      string `json:"Doctor_name" gorm:"type:varchar;column:Doctor_name"`
	Patient_name     string `json:"Patient_name" gorm:"type:varchar;column:Patient_name"`
}
