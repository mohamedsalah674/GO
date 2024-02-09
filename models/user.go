package models

type Address struct {
	State   string
	City    string
	Pincode string
}

type User struct {
	Name    string  `json: "name" bson:"user_name"`
	Age     int     `json: "age" bson:"user_age"`
	Address Address `json: "address" bson:"user_address"`
}
