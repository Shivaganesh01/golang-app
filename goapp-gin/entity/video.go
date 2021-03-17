package entity

import "time"

type Person struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"firstname" binding:"required" gorm:"type:varchar(50)"`
	LastName  string `json:"lastname" binding:"required" gorm:"type:varchar(50)"`
	Age       int8   `json:"age" binding:"gte=1,lte=120"`
	Email     string `json:"email" binding:"required,email" gorm:"type:varchar(200)"`
}

type Video struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title       string `json:"title" binding:"min=2,max=50" validate:"isGoodTitle" gorm:"type:varchar(50)"`
	Description string `json:"description" binding:"max=200" gorm:"type:varchar(200)"`
	URL         string `json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Creator     Person `json:"creator" binding:"required" gorm:"foreignkey:PersonID"`
	PersonID    uint64 `json:"-"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
