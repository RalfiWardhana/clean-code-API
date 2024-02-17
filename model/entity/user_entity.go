package entity

type User struct {
	Email    string `json:"email" gorm:"type: varchar; column:email"`
	Password string `json:"password" gorm:"type: varchar; column:password"`
	Name     string `json:"name" gorm:"type: varchar; column:name"`
}