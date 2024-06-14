package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var cl = http.Client{}

func clAuth(action, login, password, password2 string) {
	vals := url.Values{}
	vals.Add("action", action)
	vals.Add("login", login)
	vals.Add("password", password)
	vals.Add("password2", password2)

	res, err := cl.PostForm("http://localhost:8080/auth", vals)
	if err != nil {
		panic(err)
	}

	cks := res.Cookies()
	fmt.Println("Client: ", cks)
	body, _ := io.ReadAll(res.Body)
	fmt.Println("Client: ", string(body))

}

func testAuth(w http.ResponseWriter, r *http.Request) {
	clAuth("Login", "root", "DevPassword", "")
}
