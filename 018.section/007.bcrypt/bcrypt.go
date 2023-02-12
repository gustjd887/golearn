package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	fmt.Println([]byte(s))
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)

	loginPwrd := `passwod123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPwrd))
	if err != nil {
		fmt.Println("you can't login")
		// fmt.Println(err)
		return //여기서 프로그램 종료시키는 것인듯...? 아마 0 return될듯... 모르겠음
	}
	fmt.Println("you logged in")
}
