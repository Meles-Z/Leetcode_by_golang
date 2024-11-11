package entities

type User struct{
	Model
	FirstName  string `gorm:"size:30;not null" json:"firstName"`
	MiddleName string `gorm:"size:30;not null" json:"middleName"`
	LastName   string `gorm:"size:30;not null" json:"lastName"`

	PhoneNumber string  `gorm:"size:13;not null" json:"phoneNumber"`
	Email       *string `gorm:"size:50" json:"email,omitempty"`
	Address     string  `gorm:"size:255;" json:"address"`
}