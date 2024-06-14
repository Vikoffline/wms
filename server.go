package main

import (
	"fmt"
	"net/http"
)

func handleAuth(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("token"); err == nil {
		http.Redirect(w, r, "/", http.StatusOK)
		return
	}

	fmt.Println("sadasd")
	if r.Method == http.MethodPost {
		r.ParseForm()
		action := r.Form.Get("action")
		var token string
		var err error

		switch action {
		case "Login":
			token, err = userLogIn(r.Form.Get("login"), r.Form.Get("password"))
		case "SignUp":
			token, err = userSignUp(r.Form.Get("login"), r.Form.Get("password"), r.Form.Get("password2"))
		}

		if err != nil {
			w.Write([]byte(err.Error()))
			switch token {
			case "Forbidden":
				w.WriteHeader(http.StatusForbidden)
			case "InnerError":
				w.WriteHeader(http.StatusInternalServerError)
			}

		} else {
			cke := http.Cookie{
				Name:  "token",
				Value: token,
			}
			http.SetCookie(w, &cke)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func Serv() {
	mux := http.NewServeMux()
	mux.HandleFunc("/testLogin", testAuth)
	mux.HandleFunc("/auth", handleAuth)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
