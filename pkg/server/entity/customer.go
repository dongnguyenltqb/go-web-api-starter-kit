package entity

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	Name        string `json:"customer_name" bson:"customer_name" gorm:"column:customer_name"`
	Email       string `json:"customer_email" bson:"customer_email" gorm:"type:varchar(50);unique_index"`
	Address     string `json:"customer_address" bson:"customer_address"`
	Description string `json:"customer_description" bson:"customer_description"`
}
