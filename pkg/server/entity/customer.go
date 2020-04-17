package entity

type Customer struct {
	Name        string  `json:"customer_name" bson:"customer_name"`
	Email       string  `json:"customer_email" bson:"customer_email"`
	Address     Address `json:"customer_address" bson:"customer_address"`
	Description string  `json:"customer_description" bson:"customer_description"`
}

type Address struct {
}
