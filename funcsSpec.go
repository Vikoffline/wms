package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func userSignIn(login, password string) (string, error) {
	var err error

	Mn := NewManager()
	err = Mn.Find(login)
	if err != nil {
		return "Forbidden", fmt.Errorf("CError: %s: %w", errUserNotFound, err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(Mn.Password), []byte(password))

	if err != nil {
		return "Forbidden", fmt.Errorf("CError: %s: %w", errWrongPassword, err)
	}

	Sn := NewSession()
	Sn.managerId = Mn.Id
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		return "InnerError", fmt.Errorf("CError: %s: %w", errSesKeyGen, err)
	}
	Sn.Token = base64.URLEncoding.EncodeToString(key)

	err = Sn.Create()
	if err != nil {
		return "InnerError", fmt.Errorf("CError: %s: %w", errSesCreation, err)
	}

	return Sn.Token, nil
}

func userSignUp(login, password, password2 string) (string, error) {
	var err error
	var IsValidLogin = strings.Trim(login, " ") != ""
	var IsValidPass = (password == password2) && strings.Trim(password, " ") != ""
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /

	if !IsValidLogin {
		return "Forbidden", fmt.Errorf("CError: %s: %w", errInvalidLogin, err)
	}
	if !IsValidPass {
		return "Forbidden", fmt.Errorf("CError: %s: %w", errInvalidPasssword, err)
	}

	Mn := NewManager()
	Mn.Login = login
	Mn.roleId = "Rl_2"
	hashedPswd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "InnerError", fmt.Errorf("CError: %s: %w", errPasEncryption, err)
	}

	Mn.Password = string(hashedPswd)

	err = Mn.Create()

	if err != nil {
		return "InnerError", fmt.Errorf("CError: %s: %w", errUserCreation, err)
	}

	Sn := NewSession()
	Sn.managerId = Mn.Id
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		return "InnerError", fmt.Errorf("CError: %s: %w", errSesKeyGen, err)
	}
	Sn.Token = base64.URLEncoding.EncodeToString(key)

	err = Sn.Create()
	if err != nil {
		return "InnerError", fmt.Errorf("CError: %s: %w", errSesCreation, err)
	}

	return Sn.Token, nil
}
