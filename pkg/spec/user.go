package spec

import "fmt"

type User struct {
	Name   string
	System bool
}

func (u User) Identifier() string {
	return u.Name
}

func (u User) EqualityHash() string {
	return fmt.Sprintf("%s|%v", u.Name, u.System)
}

type Users []User

type UsersDiff struct {
	Changed []User
	Removed []User
}
