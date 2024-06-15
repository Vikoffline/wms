package main

// var jar, _ = cookiejar.New(nil)
// var cl = http.Client{
// 	Jar: jar,
// }

// func testAuth(w http.ResponseWriter, r *http.Request) {
// 	query := r.URL.Query()
// 	vals := url.Values{}
// 	vals.Add("action", query.Get("Action"))
// 	vals.Add("login", query.Get("Login"))
// 	vals.Add("password", query.Get("Password"))
// 	vals.Add("password2", query.Get("Password2"))

// 	res, err := cl.PostForm("http://localhost:8080/auth", vals)
// 	if err != nil {
// 		panic(err)
// 	}

// 	cks := res.Cookies()
// 	fmt.Println("Client: ", cks)
// 	body, _ := io.ReadAll(res.Body)
// 	fmt.Println("Client: ", string(body))

// }
