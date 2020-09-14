package main

type User struct {
	UserName       string `bson:"user_name" json:"user_name"`
	Age            int    `bson:"age" json:"age"`
	SomeOneISSubTo string `bson:"some_one_issub_to" json:"some_one_issub_to"`
}

// User builder pattern code
type UserBuilder struct {
	user *User
}

func NewUserBuilder() *UserBuilder {
	user := &User{}
	b := &UserBuilder{user: user}
	return b
}

func (b *UserBuilder) UserName(userName string) *UserBuilder {
	b.user.UserName = userName
	return b
}

func (b *UserBuilder) Age(age int) *UserBuilder {
	b.user.Age = age
	return b
}

func (b *UserBuilder) SomeOneISSubTo(someOneISSubTo string) *UserBuilder {
	b.user.SomeOneISSubTo = someOneISSubTo
	return b
}

func (b *UserBuilder) Build() (*User, error) {
	return b.user, nil
}

// User handler pattern code
type IUser interface {
	//todo

}

func NewUser() IUser {
	return &User{
		//todo
	}

}
