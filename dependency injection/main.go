package main

import "fmt"

type User struct {
	name string
}

func newUser(name string) User {
	return User{name: name}
}

func uName() string {
	return "Rohit"
}

func (u *User) Get(message string) string {
	return fmt.Sprintf("Hi %s, %s", u.name, message)
}

func run(user User){
	result:= user.Get("How are you doing ?")
	fmt.Println(result)
}

func main(){
	user:= Initialize()
	run(user)
}