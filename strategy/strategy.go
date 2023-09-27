package main

import "fmt"

type Context interface{
	setService(service LogInStrategy)
}

type User struct{
	name string
	logInBy LogInStrategy
}

func (u *User) setService(service LogInStrategy) {
	u.logInBy = service
	service.login(u)
}

type LogInStrategy interface{
	login(u *User)
}

type logInByGoogle struct{}

type logInByTwitter struct{}

type logInByFacebook struct{}

type logInByInstagram struct{}

func (i logInByGoogle) login(u *User) {
	fmt.Printf("%s, you successfully loged in by Google account\n", u.name);
}

func (i logInByTwitter) login(u *User) {
	fmt.Printf("%s, you successfully loged in by Twitter account\n", u.name);
}

func (i logInByFacebook) login(u *User) {
	fmt.Printf("%s, you successfully loged in by Facebook account\n", u.name);
}

func (i logInByInstagram) login(u *User) {
	fmt.Printf("%s, you successfully loged in by Instagram account\n", u.name);
}

func main() {
	google := logInByGoogle{}
	facebook := logInByFacebook{}
	twitter := logInByTwitter{}
	instagram := logInByInstagram{}



	user := User{name: "Poseidon"}

	user.setService(google)
	user.setService(facebook)
	user.setService(twitter)
	user.setService(instagram)
}