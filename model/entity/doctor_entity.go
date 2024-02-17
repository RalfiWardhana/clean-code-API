package entity

type Doctor struct {
	Doctor_name      string `json:"doctor_name" gorm:"type:varchar;column:doctor_name"`
	Jadwal_praktek   string `json:"jadwal_praktek" gorm:"type:varchar;column:jadwal_praktek"`
	Harga_konsultasi int    `json:"harga_konsultasi" gorm:"type:integer; column:harga_konsultasi"`
}
