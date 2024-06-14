package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func userLogIn(login, password string) (string, error) {
	var err error

	Mn := NewManager()
	err = Mn.Find(login)
	if err != nil {
		return "Forbidden", errors.New("CError: User not found : " + err.Error())
	}
	err = bcrypt.CompareHashAndPassword([]byte(Mn.Password), []byte(password))

	if err != nil {
		return "Forbidden", errors.New("CError: Wrong password : " + err.Error())
	}

	Sn := NewSession()
	Sn.managerId = Mn.Id
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		return "InnerError", errors.New("CError: Error on generating session key : " + err.Error())
	}
	Sn.Token = base64.URLEncoding.EncodeToString(key)

	err = Sn.Create()
	if err != nil {
		return "InnerError", errors.New("CError: Error on creating session : " + err.Error())
	}

	return Sn.Token, nil
}

func userSignUp(login, password, password2 string) (string, error) {
	var err error

	var IsValidLogin = true
	var IsValidPass = true
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /

	if !IsValidLogin {
		return "Forbidden", errors.New("CError: Invalid login")
	}
	if !IsValidPass {
		return "Forbidden", errors.New("CError: Passwords are not valid or do not match")
	}

	Mn := NewManager()
	Mn.Login = login
	hashedPswd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "InnerError", errors.New("CError: Error on encrypting password : " + err.Error())
	}

	Mn.Password = string(hashedPswd)

	err = Mn.Create()

	if err != nil {
		return "InnerError", errors.New("CError: Error on creating user : " + err.Error())
	}

	Sn := NewSession()
	Sn.managerId = Mn.Id
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		return "InnerError", errors.New("CError: Error on generating session key : " + err.Error())
	}
	Sn.Token = base64.URLEncoding.EncodeToString(key)

	err = Sn.Create()
	if err != nil {
		return "InnerError", errors.New("CError: Error on creating session : " + err.Error())
	}

	return Sn.Token, nil
}
