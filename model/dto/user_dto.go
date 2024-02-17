package dto

type User_request struct {
	Id       int    `json:"id" gorm:"type:int;column:id"`
	Email    string `json:"email" gorm:"type: varchar; column:email"`
	Password string `json:"password" gorm:"type: varchar; column:password"`
	Name     string `json:"name" gorm:"type: varchar; column:name"`
}

type User_response struct {
	Id    int    `json:"id" gorm:"type:int;column:id"`
	Email string `json:"email" gorm:"type: varchar; column:email"`
	Name  string `json:"name" gorm:"type: varchar; column:name"`
}
