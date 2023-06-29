package model

type User struct {
	Id       string `json:"id" bson:"_id"`
	UserName string `json:"userName" bson:"userName"`
	Password string `json:"password" bson:"password"`
}
