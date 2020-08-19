package main

type User struct {
	UserName       string `bson:"user_name" json:"user_name"`
	Age            int    `bson:"age" json:"age"`
	SomeOneISSubTo string `bson:"some_one_issub_to" json:"some_one_is_sub_to"`
}
